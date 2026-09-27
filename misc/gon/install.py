#!/usr/bin/env python3
"""Expose only gon and gonpls; never overwrite an existing command.

With --project-skill, copy the gon agent skill into a repository that adopts
Gon instead. The skill is project-scoped and never installed globally.
"""
import argparse
import filecmp
import os
from pathlib import Path
import shutil

ROOT = Path(__file__).resolve().parents[2]
SKILL = ROOT / ".agents" / "skills" / "gon"


def same_tree(a, b):
    compare = filecmp.dircmp(a, b)
    if compare.left_only or compare.right_only or compare.funny_files:
        return False
    _, mismatch, errors = filecmp.cmpfiles(a, b, compare.common_files, shallow=False)
    return not mismatch and not errors and all(
        same_tree(Path(a) / d, Path(b) / d) for d in compare.common_dirs)


def install_skill(repo):
    repo = repo.expanduser().absolute()
    if not repo.is_dir():
        raise SystemExit(f"{repo} is not a directory")
    target = repo / ".agents" / "skills" / "gon"
    if target.exists() or target.is_symlink():
        if target.is_dir() and not target.is_symlink() and same_tree(SKILL, target):
            print(f"{target} is up to date")
            return
        raise SystemExit(f"Refusing to overwrite {target}; remove it to install this toolchain's skill")
    target.parent.mkdir(parents=True, exist_ok=True)
    shutil.copytree(SKILL, target)
    print(f"Copied the gon skill to {target}; commit it to opt this repository into Gon")


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--bin-dir", type=Path, default=Path.home() / ".local" / "bin")
    parser.add_argument("--project-skill", type=Path, metavar="REPO",
                        help="copy the agent skill into REPO/.agents/skills/gon and exit")
    args = parser.parse_args()
    if args.project_skill:
        install_skill(args.project_skill)
        return
    destination = args.bin_dir.expanduser().absolute()
    suffix = ".exe" if os.name == "nt" else ""
    links = []
    for name in ("gon", "gonpls"):
        source = ROOT / "gon" / "bin" / (name + suffix)
        target = destination / (name + suffix)
        if not source.is_file():
            raise SystemExit("Build first: python3 misc/gon/build.py")
        if target.exists() or target.is_symlink():
            if target.is_symlink() and target.resolve() == source.resolve():
                continue
            raise SystemExit(f"Refusing to overwrite {target}; choose another --bin-dir")
        links.append((source, target))
    destination.mkdir(parents=True, exist_ok=True)
    for source, target in links:
        target.symlink_to(source)
        print(f"Installed {target} -> {source}")
    print(f"Public commands are in {destination}; keep this toolchain directory in place.")
    if str(destination) not in os.environ.get("PATH", "").split(os.pathsep):
        print(f"Add {destination} to PATH. Do not add the toolchain's private bin directory.")


if __name__ == "__main__":
    main()
