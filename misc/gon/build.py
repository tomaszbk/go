#!/usr/bin/env python3
"""Build maintained Gon tools with checksum-verified, pinned dependencies."""
import json
import os
from pathlib import Path
import shutil
import subprocess

ROOT = Path(__file__).resolve().parents[2]
HERE = Path(__file__).resolve().parent
WORK = ROOT / "pkg" / "gon-tools"
GONPLS = ROOT / "tools" / "gonpls"
GO = ROOT / "bin" / ("go.exe" if os.name == "nt" else "go")
ENV = dict(os.environ, GOROOT=str(ROOT), GOTOOLCHAIN="local", GOWORK="off", GOFLAGS="")
ENV["PATH"] = str(GO.parent) + os.pathsep + ENV.get("PATH", "")
SOURCES = json.loads((HERE / "sources.json").read_text())


def run(*args, cwd=ROOT, capture=False):
    return subprocess.run([str(a) for a in args], cwd=cwd, env=ENV,
                          check=True, text=True, stdout=subprocess.PIPE if capture else None)


def main():
    if not GO.exists():
        raise SystemExit("Build Gon first: cd src && ./make.bash")
    WORK.mkdir(parents=True, exist_ok=True)
    for name, source in SOURCES.items():
        info = json.loads(run(GO, "mod", "download", "-json", source["module"],
                              cwd=HERE / "launcher", capture=True).stdout)
        if info.get("Sum") != source["sum"] or info.get("GoModSum") != source["goModSum"]:
            raise SystemExit(f"Source checksum mismatch for {source['module']}")
        target = WORK / name
        if target.exists():
            shutil.rmtree(target)
        shutil.copytree(info["Dir"], target)
        target.chmod(target.stat().st_mode | 0o700)
        for path in target.rglob("*"):
            path.chmod(path.stat().st_mode | (0o700 if path.is_dir() else 0o200))
        patch = HERE / "patches" / (name + ".patch")
        if patch.exists():
            # git apply expects a slash-separated directory, also on Windows.
            run("git", "apply", "--directory=" + target.relative_to(ROOT).as_posix(), patch)
    run(GO, "mod", "edit", "-replace=golang.org/x/tools=../tools", cwd=WORK / "staticcheck")
    platform = json.loads(run(GO, "env", "-json", "GOOS", "GOARCH", capture=True).stdout)
    suffix = ".exe" if platform["GOOS"] == "windows" else ""
    private = ROOT / "pkg" / "tool" / (platform["GOOS"] + "_" + platform["GOARCH"])
    public = ROOT / "gon" / "bin"
    public.mkdir(parents=True, exist_ok=True)
    # Keep the maintained module immutable during builds. Its relative replace
    # directives select the prepared dependencies without a generated workspace.
    run(GO, "build", "-mod=readonly", "-trimpath", "-buildvcs=false",
        "-o", private / ("gonpls" + suffix),
        "-ldflags=-X main.version=gonpls-v0.1.0+gopls.v0.23.0", ".", cwd=GONPLS)
    for name in ("gon", "gonpls"):
        run(GO, "build", "-o", public / (name + suffix), ".", cwd=HERE / "launcher")
    print(f"Built {public / ('gon' + suffix)} and {public / ('gonpls' + suffix)}")


if __name__ == "__main__":
    main()
