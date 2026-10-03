# Gon tooling validation

Validated on 2026-09-26 on macOS (`darwin/arm64`) with this fork's
`go1.28-devel_59b9e2dc57` and unmodified `/opt/homebrew/bin/go` as the baseline.

Commands run from the repository root unless specified otherwise:

```sh
python3 misc/gon/build.py
python3 misc/gon/install.py
GO_ERROR_HANDLING_BASELINE=/opt/homebrew/bin/go python3 misc/gon/test.py
GO_ERROR_HANDLING_BASELINE=/opt/homebrew/bin/go \
  ./gon/bin/gon test cmd/internal/testdir -run='Test/errorhandling.go$' -count=1
```

All passed. The build was repeated from fresh copied module sources, with pinned
checksums and patches applied. The tooling regression executes the legacy/modern
pair, verifies upstream Go isolation, tests installer collision refusal and
idempotence, and exchanges real JSON-RPC messages with `gonpls serve`.
Formatting and import-organization edits are applied and the resulting programs
are executed; semantic-token checks verify `or` and postfix `!` specifically.
The existing language harness passed in 3.872 seconds and includes the official
Go baseline, invalid constructs, export/import, generic instantiation and vet.

Adapted upstream package tests also passed, run from `pkg/gon-tools/tools`:

```sh
/Users/tzbk/Documents/gon/gon/bin/gon test \
  ./go/ast/astutil ./go/ast/inspector ./go/cfg \
  ./refactor/satisfy ./go/analysis/passes/lostcancel
```

The user's `a.go` was separately executed with `gon run a.go` and checked with
`gonpls check a.go`. A stdio LSP session rooted at the actual repository opened
that file, received no error diagnostics, and resolved its handler binding's
hover type to `error`.

Installation resolved to:

```text
go      /opt/homebrew/bin/go
gon     /Users/tzbk/.local/bin/gon
gonpls  /Users/tzbk/.local/bin/gonpls
```

VS Code workspace configuration selects the two Gon launchers. Variable-path
resolution was checked against the installed Go extension (0.56.1). The running
VS Code UI itself was not used as a test; reload its language server after build.
Windows/Linux launchers and a mixed multi-root VS Code window were not executed
in this validation. This is focused tooling validation, not a fresh run of the
entire Go distribution suite. See the follow-up validation below for native SSA/IR and editor coverage.

## Dedicated VS Code extension follow-up

The later dedicated extension was tested inside VS Code 1.139.1 with only the
Gon extension active and three workspace folders: ordinary Go plus two Gon
projects. Go used official gopls; Gon used gonpls. Hover, diagnostics, per-project
server-command dispatch, and actual Play-command execution of the paired
programs passed. The server now accepts `gon` as well as `go` language IDs and
advertises `gonpls.*` commands to distinguish them from official gopls commands.
`GO_ERROR_HANDLING_BASELINE=/opt/homebrew/bin/go python3 misc/gon/test.py` passed
again with explicit coverage for both language IDs and the command namespace.
Detailed editor validation is in the sibling `vscode-gon/VALIDATION.md`.

## Native analysis and debugger/test-explorer port

The unsupported-analysis guard has been removed. The build pins and patches
`honnef.co/go/tools v0.8.0-rc.1` in addition to gopls and x/tools. Both builders
lower `ast.ErrorExpr` directly to ordinary IR control flow. The x/tools importer
also reads V5 export data, preserving generic/non-generic method order.

Passed:

```sh
GO_ERROR_HANDLING_BASELINE=/opt/homebrew/bin/go python3 misc/gon/test_analysis.py
GO_ERROR_HANDLING_BASELINE=/opt/homebrew/bin/go python3 misc/gon/test.py
```

The executable analysis pair covers success/failure, partial results, typed nil,
named-result resets before defer, aggregate/generic zeros, short-circuiting,
multiple-value assignments/call arguments/returns, local handlers, error-only
calls, and propagation from a range-over-function loop. Both programs execute
with Gon, legacy executes with official Go, and both execute in the SSA
interpreter. Staticcheck builds both with IR sanity checks. Actual gonpls LSP
requests assert `SA4006` and `unusedwrite` diagnostics on Gon expressions.

Upstream regressions passed from `pkg/gon-tools/tools` (SSA: 128.564 s;
gcimporter: 139.135 s):

```sh
GON_ROOT=/Users/tzbk/Documents/gon /Users/tzbk/Documents/gon/gon/bin/gon test \
  ./go/ssa ./go/ssa/ssautil ./go/ast/astutil ./go/ast/inspector ./go/cfg \
  ./refactor/satisfy ./go/analysis/passes/lostcancel \
  ./go/analysis/passes/unusedwrite ./internal/pkgbits ./internal/gcimporter
```

From `pkg/gon-tools/staticcheck` (IR: 136.579 s):

```sh
GON_ROOT=/Users/tzbk/Documents/gon /Users/tzbk/Documents/gon/gon/bin/gon test -mod=mod \
  ./go/ir/... ./analysis/facts/purity ./staticcheck/sa4006
```

An exploratory run of the separate upstream interpreter fixture suite exposed
two fixtures that use integer-to-string conversions rejected by upstream
Go 1.28 language mode (`zeros.go`, `fixedbugs/issue55115.go`). Those fixtures were
not changed. The dedicated Go/Gon SSA executable pairs above pass. This is not
a claim that every upstream interpreter fixture or the entire Go distribution
suite passes on this development version.

The sibling extension now registers Delve DAP debugging and a compiler-backed
test explorer. Real VS Code integration tests cover Go/Gon breakpoints, locals,
evaluation, stepping and continuation, plus test discovery, run/debug, failure,
skip, dynamic subtests and cancellation. See `../vscode-gon/VALIDATION.md` relative
to the repository root for the complete editor command and tested platform.

## Code-action cancellation regression

`gon test ./internal/protocol -count=1` from `pkg/gon-tools/gopls` passed.
The generated `textDocument/codeAction` dispatcher is tested on both JSON-RPC
transports: successful actions, client cancellation, snapshot cancellation,
wrapped cancellation, and genuine failures (including failures racing with a
cancellation). Cancellations return LSP -32800/-32802 instead of generic code 0;
genuine failures remain failures. The stdio `misc/gon/test.py` regression also
passed with real code-action requests followed by `$/cancelRequest`.

## Compiler test workspace loading

The official gopls v0.23.0 and the previous gonpls both reproduced the erroneous
duplicate declarations and shadowed builtins in `test/235.go`. The directory
contains separate compiler-harness inputs, not a single package.

After the fix, these commands passed:

```sh
python3 misc/gon/test_workspace.py
GO_ERROR_HANDLING_BASELINE=/opt/homebrew/bin/go python3 misc/gon/test.py
gonpls check test/235.go
# From pkg/gon-tools/gopls:
gon test ./internal/cache -run 'TestGonCompilerTest|TestStandalone' -count=1
```

The workspace regression opens the actual `235.go` and `noinit.go` together,
asserts empty diagnostics, introduces an undefined name in an unsaved overlay,
asserts that real error, and checks recovery. Separate cases verify cross-file
definitions and real duplicate-declaration diagnostics in an ordinary module
and a multi-file compiler `*.dir` fixture. Detection is restricted to the
selected toolchain's test tree and recognized harness recipes. This does not
suppress intentional errors in negative compiler tests.

## Tooling commands

Validated on 2026-09-29 on macOS (`darwin/arm64`) with this fork's
`go1.28-devel_59b9e2dc57` and unmodified Go 1.27.1 at `/opt/homebrew/bin/go`
as the baseline. From the repository root, after `python3 misc/gon/build.py`:

```sh
GO_ERROR_HANDLING_BASELINE=/opt/homebrew/bin/go python3 misc/gon/test_cli.py
GO_ERROR_HANDLING_BASELINE=/opt/homebrew/bin/go python3 misc/gon/test.py
GO_ERROR_HANDLING_BASELINE=/opt/homebrew/bin/go python3 misc/gon/test_workspace.py
GO_ERROR_HANDLING_BASELINE=/opt/homebrew/bin/go python3 misc/gon/test_analysis.py
GO_ERROR_HANDLING_BASELINE=/opt/homebrew/bin/go \
  ./gon/bin/gon test cmd/internal/testdir -run='Test/errorhandling.go$' -count=1
(cd tools/gonpls && ../../gon/bin/gon test ./internal/cmd ./internal/settings ./internal/server -count=1)
(cd tools/gonpls && ../../gon/bin/gon test ./internal/cache/... ./internal/protocol/... \
  ./internal/doc/... ./internal/lsprpc/... ./internal/test/integration/diagnostics/... -count=1)
(cd pkg/gon-tools/tools && ../../../gon/bin/gon test ./internal/typesinternal -count=1)
```

All passed. `test_cli.py` ran a legacy/modern pair with Gon and the legacy
program with official Go, checked both with `gon check` (no errors), compared
their references, renamed a function in both, ran both programs with the
expected output and ran the renamed legacy program with the baseline again. It
verified that `gon build`, `vet`, `test`, `fmt` and `run` keep their behavior
and that an unsaved buffer in a real `gonpls serve` session was visible to the
editor but not to `gon query` until saved. The Go tests cover every target
spelling, UTF-8 byte columns, revisions, paging, Gon constructs in
`query type`, check categories, filters and load errors, canonical paths when
the working directory is reached through a symbolic link, stale and atomic
plans, conservative formatting, rejected renames, and explanations
synchronized with both type checkers.

The upstream `internal/cmd` tests expected the `gopls` name and `gopls.*`
commands; their expectations and help files now match the `gonpls` identity.
gopls diagnostics integration tests confirm that editor sessions still publish
diagnostics normally. `TestGonErrorCodes` checks that x/tools has the same
value for every type-checker code of this toolchain, so gonpls reports
`InvalidErrorHandling` instead of `ErrorCode(10000)`. `gonpls` and the launcher
also built for linux/amd64, linux/arm64, linux/riscv64, windows/amd64 and
windows/arm64.

A persistent service for these commands was implemented, measured and then
removed by decision of the project: with gonpls's on-disk cache, a new process
answered references queries in 0.2–0.9 s, against 0.04–0.2 s for the service,
which kept 0.2–1.5 GB resident. The measurements are recorded in
[design/agent-tooling.md](../../design/agent-tooling.md). The editor's
language server remains a long-lived process.

The VS Code extension checks are recorded in `../vscode-gon/VALIDATION.md`.
A rerun of that suite exposed that `go list` reported directories under an
inherited, symlinked `$PWD` (macOS `/var` instead of `/private/var`), which made
`gon check` print absolute paths that the problem matcher could not resolve.
`gon check` now runs `go list` with the canonical directory; the Go regression
reproduces the failure without the fix, and the suite passed again.

## Maintained modules and feature integration gates (2026-10-02)

Validated on darwin/arm64 with Gon `go1.28-devel_981e974870` and unmodified
`/opt/homebrew/bin/go` (`go1.27.1`). Other supported architectures were not
executed in this session. No broad distribution test suite was run.

Sources now live in `tools/x-tools` and `tools/staticcheck`. Gonpls and cmd select
that same x/tools source through relative replacements; `go list -m -json`
confirmed both resolve to `tools/x-tools`. There are no nested Git repositories.
The old tools, staticcheck and cmd-vendor patches were imported and removed;
source provenance is in `UPSTREAM.json`. Old `pkg/gon-tools` output was moved out
of the workspace before the final build, which passed without recreating it.

Commands from the repository root:

```sh
GON_BASELINE_GO=/opt/homebrew/bin/go python3 misc/gon/validate.py tooling
GON_BASELINE_GO=/opt/homebrew/bin/go python3 misc/gon/validate.py errorhandling
GON_BASELINE_GO=/opt/homebrew/bin/go python3 misc/gon/validate.py conditional
python3 misc/gon/vendor.py --check
```

Results:

- Tooling: all 17 gates passed (exit 0), covering build, AST, structural helpers,
  semantic helpers, conservative inliner refusal, both language executable pairs,
  SSA execution, Staticcheck IR, CLI/LSP, typerefs and unusedfunc.
- Error handling: all 18 gates passed (exit 0), additionally covering the existing
  cgo executable pair/diagnostics, cover flow profiles/legacy instrumentation,
  syntax, generated type checker consistency and vet.
- Conditional: all 14 implemented gates passed; exit 2 deliberately reports the
  unfinished feature integration items in `features.json`. This is not a claim
  that conditional expressions are complete.
- The analyzer corpus ran all 247 analyzers from the actual gonpls registry on
  five packages, including legacy/modern error and conditional fixtures, with
  no panics or failed analysis actions. New public Gon nodes must appear in the
  corpus; the test checks the inventory in `api/fork.txt`.
- A deliberate temporary change to a generated vendor file made `--check` fail
  and name the changed file. The probe restored the exact bytes immediately;
  the final vendor check passed.
- `git diff --check` and Python compilation of the build/vendor/validation scripts
  passed. Ignore rules were adjusted so imported source `internal/event/core`
  and upstream fixtures resembling compiler outputs are included in Git.

The runners record exact commands, test counts, skips and output under
`pkg/gon-validation/<profile>/summary.json` and adjacent logs. Upstream skips
were `typesinternal.TestErrorCodes` (replaced by the passing Gon-specific code
inventory test) and vet's stringintconv/stdversion/loopclosure subtests. No Gon
executable pair was skipped; each legacy pair also ran with the official baseline.

`ast.Children` and Walk share one structural inventory. Inspector cursor metadata
and mutable AST slots still require explicit registration. Semantic walkers are
not indiscriminately replaced: effects are conservative, unsupported copy/equality
is declined, and inlining Gon control-flow callee bodies returns an explicit
error. cgo/cover/editor completion for conditional expressions remains the next
feature work, together with its specification clarifications.

## Go master integration (2026-10-02)

Go master through `67c1d421161d3d1ae9f5fd005e84c29fd0d9f896` and x/tools
`98444708d405` were integrated and tested on darwin/arm64. The original checkout
was rebuilt from unmodified Go 1.27.1 after an isolated validation. Tooling
(17 gates) and error handling (18 gates) pass; conditional passes its 14 existing
gates but still exits 2 for pending integration. Both syntax/type-checker
families, upstream regressions, the new export format and selected receiver
operations in gonpls were checked. No whole-distribution suite was run.

See [the full record](../../handover/upstream-2026-10-02.md) for exact revisions,
commands, logs, conflict resolutions and skipped/dormant upstream tests. In
particular, two upstream gcimporter TestMain functions execute zero tests; their
package-level successes are not counted as importer validation.

## Conditional integration closure (2026-10-02)

Conditional integration is complete with explicit conservative limits: source
inlining declines Gon control-flow bodies and affected call sites; variable
extraction declines lazy branches and whole conditional expressions; cgo's
existing restriction on error handling in pointer-check rewritten C arguments
still applies. The user confirmed all four target-type clarifications, now
recorded in the local specification. `features.json` separates these supported
limits from pending work, which is empty for conditional expressions.

Validated on **darwin/arm64**, with unmodified **Go 1.27.1** at
`/opt/homebrew/bin/go`. Only checks affected by this integration ran; compiler,
SSA/IR and other unchanged gates were not repeated. No full profile, whole
distribution suite, full bootstrap, or other-architecture execution is claimed.

From the repository root:

```sh
GON_BASELINE_GO=/opt/homebrew/bin/go ./gon/bin/gon test cmd/cgo/internal/testconditional -run '^Test(PairedCgoConditional|CgoConditionalDiagnostics|CgoConditionalBootstrap)$' -count=1 -v
GO_ERROR_HANDLING_BASELINE=/opt/homebrew/bin/go ./gon/bin/gon test cmd/cgo/internal/testerrorhandling -run '^Test(PairedCgoErrorHandling|CgoErrorHandlingDiagnostics)$' -count=1
GON_BASELINE_GO=/opt/homebrew/bin/go ./gon/bin/gon test cmd/cover -run '^(TestCondFlowCoverage|TestErrorFlowCoverage|TestLegacyInstrumentationUnchanged|TestErrorHandlingRanges)$' -count=1 -v
GON_BASELINE_GO=/opt/homebrew/bin/go python3 misc/gon/validate.py conditional --only refactor-safety --only vet
GON_BASELINE_GO=/opt/homebrew/bin/go python3 misc/gon/test.py --conditional-only
```

All selected tests passed. cgo conditional took 6.033s and existing cgo error
handling 4.663s, with no skips. Its executable pairs check C and Go call order,
lazy branch selection, contextual pointer/nil types, actual pointer checks,
defer, errno propagation and handlers. Nine diagnostic/control cases cover
rejected source and nested function boundaries. The bootstrap adapter test
builds maintained cmd/cgo against upstream go/ast and processes ordinary source;
the baseline supplies other internal cmd dependencies.

Cover passed in 4.252s. The new conditional fixtures compare 45 profiles across
legacy Gon, modern Gon, baseline legacy and set/count/atomic modes. Executable
assertions cover both success/failure branches, named results, defer, handlers,
closure boundaries and nesting. Existing error-flow and unchanged legacy
instrumentation regressions passed as well.

The runner subset passed refactor-safety (53 tests/subtests, 3.439s) and vet
(1 test, 1.948s), without skips. Exit 2 indicates a partial selection; the
metadata still listed pending items when it ran, before this closure update.
Exact JSON events and the historical summary remain in
`pkg/gon-validation/conditional/`. The new inliner executable pair runs both
original and transformed legacy code under Gon and baseline Go, plus modern
code under Gon after verifying refusal without edits. It asserts lazy branch
selection, evaluation order and early error returns. Legacy inlining remains
available even in functions with unrelated conditional expressions.

From `tools/gonpls`:

```sh
../../gon/bin/gon test ./internal/cmd -run '^TestGonConditionalQuery$' -count=1
../../gon/bin/gon test ./internal/golang -run '^TestConditionalExtraction$' -count=1
```

Both passed. The LSP command above ran after rebuilding gonpls and passed its
legacy/baseline/modern executable pair, keyword tokens, hover/navigation,
completion for incomplete then/else/condition expressions and inferred sibling
types, and extraction-action refusal. Unit checks include package-level query
types/constants and extract-all across eager and lazy occurrences.

Distribution commands also passed:

```sh
./gon/bin/gon install cmd/cgo
python3 misc/gon/vendor.py
python3 misc/gon/vendor.py --check
python3 misc/gon/build.py
./gon/bin/gon install cmd/fix
git diff --check
```

The validation runner now exposes focused cgo, cover, editor-query,
editor-extraction and editor-lsp checks in the conditional profile. Its setup
also installs cmd/fix so regenerated inliner changes reach `gon fix`.
`GON_BASELINE_GO` is the current public setting; older aliases remain supported.

## Lambdas and null safety (2026-10-02)

Implemented `(params) => expression/block`, `?.`, `?(`, `??`, guarded
dereferencing and `??=` in the compiler and public syntax/type checker families.
Integration includes CFG/SSA/Staticcheck IR, analyzers, cgo/bootstrap adapters,
coverage, gonpls and the semantic CLI. The lambda design was already accepted;
the user authorized both implementations, with the recommended nil-safety
decisions, and requested only necessary focused tests.

Validation ran on **darwin/arm64**, using `/opt/homebrew/bin/go` (unmodified
Go 1.27.1) for legacy baselines. No other architecture execution, full bootstrap
or whole-distribution test run is claimed. All source changes remain uncommitted.

From the repository root:

```sh
./gon/bin/gon test go/token go/scanner go/ast go/parser go/printer go/format cmd/gofmt -count=1
./gon/bin/gon test cmd/compile/internal/syntax cmd/compile/internal/types2 go/types -run='Lambda|NilSafety|NullSafety|TestGenerate' -count=1
GON_BASELINE_GO=/opt/homebrew/bin/go ./gon/bin/gon test cmd/internal/testdir -run='^Test/(lambda|nullsafety|errorhandling|conditional)\.go$' -count=1 -json
GON_BASELINE_GO=/opt/homebrew/bin/go ./gon/bin/gon test cmd/cgo cmd/cgo/internal/testconditional -run='^Test(GonErrorBoundary|GonTargetType|GonWalk|PairedCgoGonFeatures|PairedCgoConditional|CgoConditionalDiagnostics|CgoConditionalBootstrap)$' -count=1
GON_BASELINE_GO=/opt/homebrew/bin/go ./gon/bin/gon test cmd/cover -run='^TestGonFlowCoverage$' -count=1
GON_BASELINE_GO=/opt/homebrew/bin/go ./gon/bin/gon test cmd/cover -run='^Test(ErrorHandlingRanges|LegacyInstrumentationUnchanged|ErrorFlowCoverage|CondFlowCoverage|GonFlowCoverage)$' -short -count=1
```

All passed. The final four-feature execution took 14.597s, without skips. The
lambda/nil-safety harnesses execute legacy and modern with Gon, legacy with the
baseline, modern without inlining and under the race detector, exported generic
and inlined cross-package bodies, plus vet on both variants. They reject 23
invalid lambda and 28 invalid nil-safety programs. Cases cover closure capture,
function boundaries, deferred work, partial errors, nil/typed-nil interfaces,
single evaluation, skipped arguments, nested guards, zero versus absence,
nil-map stores and `??=` without a store on the present path.

Cgo's final combined regression passed in 0.244s and 20.897s for the two
packages, without skips, including coverage, vet and the baseline.
`TestCgoConditionalBootstrap` checks the cgo bootstrap adapters;
`cmd/cgo`'s `TestGonErrorBoundary`, `TestGonTargetType`, `TestGonWalk` and cover's
`TestGonFunctionBoundary` also passed. The complete feature coverage test passed
in 6.31s across set/count/atomic and both language variants plus baseline. The
focused cover regression command with `-short` passed in 7.34s. Expression
lambda bodies count with their enclosing statement; block lambda bodies have
ordinary function counters. Package-initializer lambdas with `or` handlers are
covered, including handler entry counters.

From `tools/x-tools`:

```sh
GON_ROOT=/Users/tzbk/Documents/gon GON_BASELINE_GO=/opt/homebrew/bin/go ../../gon/bin/gon test ./go/ssa -run='^TestGonFeatures$' -count=1 -v
../../gon/bin/gon test ./go/analysis/passes/nilfunc ./go/analysis/passes/defers ./go/analysis/passes/waitgroup ./go/analysis/passes/unusedresult ./go/analysis/passes/printf ./go/analysis/passes/testinggoroutine ./go/analysis/passes/unreachable -run='^(Test|TestGon)$' -count=1
../../gon/bin/gon test ./go/analysis/passes/copylock ./go/analysis/passes/lostcancel -run='^TestGon' -count=1
```

From `tools/staticcheck`:

```sh
GON_ROOT=/Users/tzbk/Documents/gon ../../gon/bin/gon test ./go/ir -run='^TestGonFeatures$' -count=1
../../gon/bin/gon test ./internal/sharedcheck ./simple/s1023 ./staticcheck/sa4004 ./staticcheck/sa4009 ./staticcheck/sa5003 ./staticcheck/sa9001 -run='^TestGon' -count=1
```

These passed. SSA's fixture runs legacy on baseline and Gon, modern on Gon, then
interprets both SSA programs (1.798s). Staticcheck builds both variants with IR
sanity checks (0.207s). Extended cases distinguish absent chains, present typed
nil results and interface boxing, generic nilable unions, nested dereferences,
skipped propagation/handlers, lambda defaults and no-store assignments. Nil-map
panic recovery is tested by the native harness; it is not duplicated in the
interpreter fixture because the existing interpreter cannot represent recovered
runtime errors without loading runtime, even for the legacy program.

The analyzer tests assert legacy/modern diagnostic parity. They also cover
closure scopes, copylock interface conversions, lostcancel paths and preserving
necessary contextual types when suggesting removal of redundant declarations.
The old Staticcheck `TestTestdata` suites for the five changed analyzers could
not run: the imported checkout lacks their `testdata/go1.x/go.mod` fixtures.
This is recorded as unavailable evidence, not a passing regression; their new
`TestGon` fixtures and the full analyzer registry corpus passed.

The final maintained-tooling subset ran from the root:

```sh
GON_BASELINE_GO=/opt/homebrew/bin/go python3 misc/gon/validate.py lambda --only structural-tools --only analyzers --only refactor-safety --only staticcheck-safety --only typerefs
```

All five selected checks passed: 42, 1730, 58, 2 and 25 test events respectively.
The registry corpus exercises all 247 registered analyzers on seven packages.
The only skip was upstream `typesinternal.TestErrorCodes`, replaced by passing
`TestGonErrorCodes`. The runner exited 2 because this was a partial selection
and metadata was still open at execution time. It does not represent a full
profile run. Logs and the exact commands are in `pkg/gon-validation/lambda/`.

Structural checks assert cursor edges and mutable AST slots; CFG checks assert
that absent receivers bypass arguments and go directly to coalescing fallbacks.
Interface constraint discovery covers lambda results, navigation, `??` and
`??=`. Typerefs covers generic result inference, parameter shadowing and guards.
The inliner declines affected code; extraction declines movement across lazy
paths or contextual lambda boundaries. Conservative refactoring limits are
listed separately from unfinished work in `features.json`.

Final editor checks passed after rebuilding public gonpls (3.31s):

```sh
GON_BASELINE_GO=/opt/homebrew/bin/go python3 misc/gon/test.py --features-only
# From tools/gonpls:
../../gon/bin/gon test ./internal/golang ./internal/cmd -run='Test(GonFeatureExtraction|GonFeatureQuery|GonFeatureExplain)$' -count=1
```

The real LSP script passed in 4.75s, including baseline/modern execution,
semantic tokens, hover, parameter/return navigation, inferred parameter hints,
guarded-call signature help, incomplete-source completion, extraction refusal,
formatting and executable lambda/function-literal conversion roundtrips. It
asserts that conversions are declined for unsafe contexts, partially inferred
generic calls and comments that would be lost. Focused extraction/query/explain
tests passed in 0.344s and 1.039s for the two packages. `gon query type` reports
the new constructs and `gon explain` documents both new diagnostic codes.

Distribution checks passed:

```sh
./gon/bin/gon install cmd/compile
./gon/bin/gon install cmd/vet cmd/fix cmd/cover cmd/cgo cmd/gofmt
python3 misc/gon/vendor.py
python3 misc/gon/vendor.py --check
python3 misc/gon/build.py
python3 -m py_compile misc/gon/validate.py
git diff --check
```

The public tools were rebuilt after the final source changes. The `lambda` and
`nullsafety` validation profiles now list all focused feature gates. These
records combine the checks actually needed during implementation rather than
rerunning unchanged gates merely to obtain a full-profile status.

## Gon syntax modernization (2026-10-02)

`gon fix` now shares four Gon syntax analyzers with gonpls and `gon check`:
`gonerrors`, `gonconditional`, `gonnil`, and `gonlambda`. Maintained source is
`tools/x-tools/go/analysis/passes/gonmodernize`; registration in `fix.Suite`
also supplies optional editor hints and quick fixes. No compiler or language
semantics changed. `misc/gon/CLI.md` records commands and conservative limits.

Executed on darwin/arm64 with the unmodified Go 1.27.1 baseline:

```sh
python3 misc/gon/vendor.py
./gon/bin/gon install cmd/fix
python3 misc/gon/build.py
# From tools/gonpls:
../../gon/bin/gon run ./internal/doc/generate
../../gon/bin/gon test ./internal/doc/generate -run '^TestGenerated$' -count=1
# From the repository root:
GON_BASELINE_GO=/opt/homebrew/bin/go python3 misc/gon/validate.py tooling --only vendor --only analyzers --only cli --only syntax-fixes --only fix-execution
python3 -m py_compile misc/gon/test_fix.py misc/gon/validate.py
git diff --check
```

All five selected gates passed without skipped tests. `syntax-fixes` ran 29
test events, including diagnostics/goldens, post-fix type checking, conservative
refusals and all 24 analyzer orders through the actual fix edit merger. The
registry corpus ran 251 analyzers on seven packages (1,758 test events).
The existing semantic CLI regression passed, as did generated documentation.
The validator exited 2 because the selection was partial; no full profile or
execution on another architecture is claimed.

`test_fix.py` copies `fixfixtures/legacy.go` into a temporary module, executes it
with the baseline and Gon, modernizes it through the public `gon fix`, and
executes the resulting modern program with normal and disabled inlining. Both
programs assert exact values, error identity/wrapping and side-effect order.
Coverage includes success/failure, partial results, still-live error bindings,
named return values observed by defers, fall-through local handlers, lazy nil
defaults/callback arguments, typed-nil interfaces, mixed interface branch types,
zero-valued present fields and closure captures. It verifies `-diff` leaves
files intact and exits 1, application reaches a stable result, vet accepts the
modern program, and CLI/LSP diagnostics expose usable fixes for all four
analyzers. The modern program is generated by the tested command, not a
documentation-only example.

Nested overlapping fixes use the existing driver conflict handling and may
require rerunning the command. The analyzers intentionally decline unsupported
source shapes and transformations that would alter contextual types or lose
bindings, comments or imports. All language additions remain optional.
