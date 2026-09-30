#!/usr/bin/env python3
"""Vendor the std and cmd dependencies and reapply Gon's changes to them.

'go mod vendor' restores the upstream contents of src/vendor and
src/cmd/vendor, which drops Gon's changes to vendored packages. This script
vendors both modules with this toolchain and applies
misc/gon/patches/cmd-vendor.patch again.

With --save, it records the current differences between the vendored files
and their module versions as that patch instead. Run it after changing a
vendored file, and before vendoring new dependency versions.
"""
import argparse
import difflib
import json
import os
from pathlib import Path
import subprocess

ROOT = Path(__file__).resolve().parents[2]
PATCH = ROOT / "misc" / "gon" / "patches" / "cmd-vendor.patch"
GO = ROOT / "bin" / ("go.exe" if os.name == "nt" else "go")
MODULES = ("src", "src/cmd")
ENV = {k: v for k, v in os.environ.items() if k not in ("GOROOT", "GOFLAGS")}
ENV.update(GOTOOLCHAIN="local", GOWORK="off")


def run(*args, cwd=ROOT, capture=False):
    return subprocess.run([str(a) for a in args], cwd=cwd, env=ENV, check=True, text=True,
                          stdout=subprocess.PIPE if capture else None)


def vendored_modules(module):
    """Return the module versions listed in module's vendor/modules.txt."""
    versions = {}
    for line in (ROOT / module / "vendor" / "modules.txt").read_text().splitlines():
        fields = line.split()
        if len(fields) >= 3 and fields[0] == "#" and not fields[1].startswith("explicit"):
            path, version = fields[1], fields[2]
            if "=>" in fields:
                raise SystemExit(f"{module}: replaced module {path} is not supported")
            info = json.loads(run(GO, "mod", "download", "-json", f"{path}@{version}",
                                  cwd=ROOT / module, capture=True).stdout)
            versions[path] = Path(info["Dir"])
    return versions


def save():
    chunks = []
    for module in MODULES:
        vendor = ROOT / module / "vendor"
        dirs = vendored_modules(module)
        for file in sorted(p for p in vendor.rglob("*") if p.is_file() and p.name != "modules.txt"):
            rel = file.relative_to(vendor).as_posix()
            owner = max((m for m in dirs if rel.startswith(m + "/")), key=len, default=None)
            if owner is None:
                raise SystemExit(f"{file}: not part of a vendored module")
            pristine = dirs[owner] / rel[len(owner) + 1:]
            new = file.read_text().splitlines(keepends=True)
            old = pristine.read_text().splitlines(keepends=True) if pristine.exists() else []
            if old == new:
                continue
            name = file.relative_to(ROOT).as_posix()
            before = f"a/{name}" if pristine.exists() else "/dev/null"
            chunks.append(f"diff --git a/{name} b/{name}\n")
            if not pristine.exists():
                chunks.append("new file mode 100644\n")
            chunks.extend(difflib.unified_diff(old, new, before, f"b/{name}"))
    PATCH.write_text("".join(chunks))
    print(f"Wrote {PATCH.relative_to(ROOT)} with {sum(c.startswith('diff --git') for c in chunks)} files")


def vendor():
    for module in MODULES:
        run(GO, "mod", "vendor", cwd=ROOT / module)
    if PATCH.read_text():
        try:
            run("git", "apply", PATCH)
        except subprocess.CalledProcessError:
            raise SystemExit(f"{PATCH.relative_to(ROOT)} no longer applies: port Gon's changes "
                             "to the new versions, then run this script with --save")
    print("Vendored std and cmd, and applied Gon's changes")


def main():
    parser = argparse.ArgumentParser(description=__doc__, formatter_class=argparse.RawDescriptionHelpFormatter)
    parser.add_argument("--save", action="store_true", help="record the current vendored changes as the patch")
    args = parser.parse_args()
    if not GO.exists():
        raise SystemExit("Build Gon first: cd src && ./make.bash")
    save() if args.save else vendor()


if __name__ == "__main__":
    main()
