#!/usr/bin/env python3
"""Generate cmd's vendor tree from maintained sources; --check detects drift.

Edit tools/x-tools, never src/cmd/vendor/golang.org/x/tools. Other modules
retain the versions selected by src/cmd/go.mod. No Gon patches are applied.
"""
import argparse
import os
from pathlib import Path
import subprocess
import tempfile

ROOT = Path(__file__).resolve().parents[2]
GO = ROOT / 'gon/bin' / ('gon.exe' if os.name == 'nt' else 'gon')
ENV = dict(os.environ, GOROOT=str(ROOT), GOTOOLCHAIN='local', GOWORK='off', GOFLAGS='')


def files(root):
    return {p.relative_to(root): p.read_bytes() for p in root.rglob('*') if p.is_file()}


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--check', action='store_true')
    args = parser.parse_args()
    target = ROOT / 'src/cmd/vendor'
    with tempfile.TemporaryDirectory(prefix='gon-vendor-') as tmp:
        generated = Path(tmp) / 'vendor'
        subprocess.run([str(GO), 'mod', 'vendor', '-o', str(generated)],
                       cwd=ROOT / 'src/cmd', env=ENV, check=True)
        before, after = files(target), files(generated)
        changed = sorted(str(p) for p in before.keys() | after.keys() if before.get(p) != after.get(p))
        if args.check:
            if changed:
                raise SystemExit('Vendor differs; run python3 misc/gon/vendor.py:\n' + '\n'.join(changed))
            print('PASS: vendor matches maintained sources')
        else:
            # Write only differences; do not touch unrelated source trees.
            for rel in before.keys() - after.keys():
                (target / rel).unlink()
            for rel, data in after.items():
                if before.get(rel) != data:
                    path = target / rel
                    path.parent.mkdir(parents=True, exist_ok=True)
                    path.write_bytes(data)
            print(f'Generated cmd vendor ({len(changed)} changed files)')


if __name__ == '__main__':
    main()
