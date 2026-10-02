#!/usr/bin/env python3
"""Run focused Gon integration gates. A passing subset is not feature completion."""
import argparse
import json
import os
from pathlib import Path
import shlex
import subprocess
import sys
import time

ROOT = Path(__file__).resolve().parents[2]
GON = ROOT / 'gon/bin' / ('gon.exe' if os.name == 'nt' else 'gon')


def checks(feature):
    def test(name, cwd, packages, pattern=None):
        args = [str(GON), 'test', '-json', *packages]
        if pattern:
            args += ['-run', pattern]
        return name, cwd, args + ['-count=1']
    common = [
        ('vendor', '.', [sys.executable, 'misc/gon/vendor.py', '--check']),
        ('install-tools', '.', [str(GON), 'install', 'cmd/vet', 'cmd/gofmt', 'cmd/cover', 'cmd/cgo']),
        ('build-tooling', '.', [sys.executable, 'misc/gon/build.py']),
        test('ast', '.', ['go/ast']),
        test('structural-tools', 'tools/x-tools', ['./go/ast/inspector', './go/ast/astutil', './go/ast/edge', './go/cfg', './refactor/satisfy', './internal/typesinternal'], 'TestGon|TestCond|TestError|TestInspectAllNodes'),
        test('analyzers', 'tools/gonpls', ['./internal/settings'], '^TestGonAnalyzers$'),
        test('refactor-safety', 'tools/x-tools', ['./internal/refactor/inline'], '^(TestGon|TestCalleeEffects|TestBasics|TestPrecedenceParens)'),
        test('staticcheck-safety', 'tools/staticcheck', ['./analysis/code', './go/ast/astutil'], '^TestGon'),
    ]
    pairs = []
    features = ['errorhandling', 'conditional'] if feature == 'tooling' else [feature]
    for name in features:
        pairs.append(test(name+'-execution', '.', ['cmd/internal/testdir'], 'Test/'+name+r'.go$'))
    if feature == 'tooling':
        return common + pairs + [
            test('ssa', 'tools/x-tools', ['./go/ssa'], '^TestGon'),
            test('staticcheck-ir', 'tools/staticcheck', ['./go/ir'], '^TestGon'),
            test('tooling-api', 'tools/gonpls', ['./internal/cmd'], '^TestGon'),
            test('typerefs', 'tools/gonpls', ['./internal/cache/typerefs'], '^TestRefs$'),
            test('unusedfunc', 'tools/gonpls', ['./internal/analysis/unusedfunc']),
            ('lsp', '.', [sys.executable, 'misc/gon/test.py']),
            ('cli', '.', [sys.executable, 'misc/gon/test_cli.py']),
        ]
    pattern = 'CondExpr|CondParen' if feature == 'conditional' else 'ErrorHandling|ErrorExpr'
    extra = []
    if feature == 'errorhandling':
        extra = [
            test('cgo', '.', ['cmd/cgo/internal/testerrorhandling'], '^Test(PairedCgoErrorHandling|CgoErrorHandlingDiagnostics)$'),
            test('cover', '.', ['cmd/cover'], '^Test(ErrorFlowCoverage|ErrorHandlingRanges|LegacyInstrumentationUnchanged)$'),
            ('lsp', '.', [sys.executable, 'misc/gon/test.py']),
            ('cli', '.', [sys.executable, 'misc/gon/test_cli.py']),
        ]
    return common + pairs + extra + [
        test('syntax', '.', ['cmd/compile/internal/syntax', 'go/parser', 'go/printer', 'go/format', 'cmd/gofmt'], pattern),
        test('types', '.', ['cmd/compile/internal/types2', 'go/types'], pattern+'|TestGenerate'),
        test('ssa', 'tools/x-tools', ['./go/ssa'], '^TestGon'),
        test('staticcheck-ir', 'tools/staticcheck', ['./go/ir'], '^TestGon'),
        test('vet', '.', ['cmd/vet'], '^TestCondExpr$' if feature == 'conditional' else '^TestVet$'),
    ]


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('feature', choices=['tooling', 'errorhandling', 'conditional'])
    parser.add_argument('--list', action='store_true', help='show commands without executing')
    parser.add_argument('--only', action='append', help='run selected check IDs; reports a partial run')
    args = parser.parse_args()
    plan = checks(args.feature)
    if args.only:
        unknown = set(args.only) - {name for name, _, _ in plan}
        if unknown:
            parser.error('unknown checks: ' + ', '.join(sorted(unknown)))
        plan = [step for step in plan if step[0] in args.only]
    if args.list:
        for name, cwd, cmd in plan:
            print(f'{name}: (cd {cwd} && {shlex.join(cmd)})')
        return 0
    baseline = os.environ.get('GON_BASELINE_GO') or os.environ.get('GO_ERROR_HANDLING_BASELINE')
    if not baseline or not Path(baseline).is_absolute() or not os.access(baseline, os.X_OK):
        parser.error('set GON_BASELINE_GO to an unmodified compatible Go executable (absolute path)')
    env = dict({k: v for k, v in os.environ.items() if k not in ('GOROOT', 'GOTOOLDIR')}, GON_ROOT=str(ROOT), GOWORK='off',
               GOTOOLCHAIN='local', GOFLAGS='', GON_BASELINE_GO=baseline,
               GO_ERROR_HANDLING_BASELINE=baseline, GO_CONDITIONAL_EXPRESSION_BASELINE=baseline)
    # Analysis loaders invoke "go"; select the private tool only in child processes.
    env['PATH'] = str(ROOT / 'bin') + os.pathsep + env.get('PATH', '')
    out = ROOT / 'pkg/gon-validation' / args.feature
    out.mkdir(parents=True, exist_ok=True)
    results = []
    for index, (name, cwd, cmd) in enumerate(plan):
        print(f'RUN {name}: {shlex.join(cmd)}', flush=True)
        start = time.monotonic()
        log = out / (name+'.log')
        with log.open('w') as stream:
            proc = subprocess.run(cmd, cwd=ROOT/cwd, env=env, stdout=stream, stderr=subprocess.STDOUT)
        result = dict(check=name, command=cmd, cwd=cwd, status='pass' if proc.returncode == 0 else 'fail',
                      exitCode=proc.returncode, seconds=round(time.monotonic()-start, 3), log=str(log))
        if len(cmd) > 1 and cmd[1] == 'test':
            events = []
            for line in log.read_text().splitlines():
                try:
                    events.append(json.loads(line))
                except ValueError:
                    pass
            result['testsRun'] = sum(e.get('Action') == 'run' for e in events)
            result['skippedTests'] = [e.get('Package', '') + '/' + e['Test'] for e in events
                                      if e.get('Action') == 'skip' and e.get('Test')]
            if not result['testsRun']:
                result['status'] = 'fail'
                result['reason'] = 'no matching tests ran'
        results.append(result)
        print(f'{result["status"].upper()} {name} ({result["seconds"]}s): {log}', flush=True)
        if result['status'] == 'fail':
            print(log.read_text()[-10000:], flush=True)
            if name in ('vendor', 'install-tools', 'build-tooling'):
                results.extend(dict(check=n, command=c, cwd=d, status='not-run', reason='setup failed')
                               for n, d, c in plan[index+1:])
                break
    pending = json.loads((ROOT/'misc/gon/features.json').read_text())[args.feature]['pending']
    summary = dict(feature=args.feature, partial=bool(args.only), pending=pending, results=results)
    (out/'summary.json').write_text(json.dumps(summary, indent=2)+'\n')
    if pending:
        print('FEATURE STILL OPEN: ' + '; '.join(pending))
    if args.only:
        print('PARTIAL RUN: unselected checks were not executed')
    if any(r['status']=='fail' for r in results):
        return 1
    print('PASS: selected checks' if args.only or pending else 'PASS: all integration gates')
    return 2 if pending or args.only else 0


if __name__ == '__main__':
    sys.exit(main())
