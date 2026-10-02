#!/usr/bin/env python3
"""Build maintained Gon tools with pinned transitive dependencies."""
import json
import os
from pathlib import Path
import subprocess

ROOT = Path(__file__).resolve().parents[2]
HERE = Path(__file__).resolve().parent
GONPLS = ROOT / "tools" / "gonpls"
GO = ROOT / "bin" / ("go.exe" if os.name == "nt" else "go")
ENV = dict(os.environ, GOROOT=str(ROOT), GOTOOLCHAIN="local", GOWORK="off", GOFLAGS="")
ENV["PATH"] = str(GO.parent) + os.pathsep + ENV.get("PATH", "")


def run(*args, cwd=ROOT, capture=False):
    return subprocess.run([str(a) for a in args], cwd=cwd, env=ENV,
                          check=True, text=True, stdout=subprocess.PIPE if capture else None)


def main():
    if not GO.exists():
        raise SystemExit("Build Gon first: cd src && ./make.bash")
    platform = json.loads(run(GO, "env", "-json", "GOOS", "GOARCH", capture=True).stdout)
    suffix = ".exe" if platform["GOOS"] == "windows" else ""
    private = ROOT / "pkg" / "tool" / (platform["GOOS"] + "_" + platform["GOARCH"])
    public = ROOT / "gon" / "bin"
    public.mkdir(parents=True, exist_ok=True)
    # Keep the maintained module immutable during builds. Its relative replace
    # directives select the maintained dependencies without a generated workspace.
    run(GO, "build", "-mod=readonly", "-trimpath", "-buildvcs=false",
        "-o", private / ("gonpls" + suffix),
        "-ldflags=-X main.version=gonpls-v0.1.0+gopls.v0.23.0", ".", cwd=GONPLS)
    for name in ("gon", "gonpls"):
        run(GO, "build", "-o", public / (name + suffix), ".", cwd=HERE / "launcher")
    print(f"Built {public / ('gon' + suffix)} and {public / ('gonpls' + suffix)}")


if __name__ == "__main__":
    main()
