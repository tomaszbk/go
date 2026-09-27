#!/usr/bin/env python3
"""Real LSP regressions for compiler test inputs and ordinary package loading."""
from pathlib import Path
import tempfile

from test import Client, ROOT, position


def open_file(client, path, source=None):
    client.send("textDocument/didOpen", {"textDocument": {
        "uri": path.as_uri(), "languageId": "gon", "version": 1,
        "text": path.read_text() if source is None else source,
    }})


def check_package(folder):
    # Even a '// run' comment must not split an ordinary package or a *.dir
    # fixture. Definition must resolve to the second file, and an actual
    # duplicate declaration must still produce an error.
    a, b = folder / "a.go", folder / "b.go"
    a.write_text("// run\npackage main\nfunc main() { shared() }\n")
    b.write_text("package main\nfunc shared() {}\n")
    with tempfile.TemporaryFile(mode="w+") as log:
        client = Client(folder, log)
        try:
            open_file(client, a)
            open_file(client, b)
            client.diagnostics(a.as_uri(), 1, lambda ds: not ds)
            result = client.request("textDocument/definition", {
                "textDocument": {"uri": a.as_uri()},
                "position": position(a.read_text(), "shared"),
            })
            assert result and result[0]["uri"] == b.as_uri(), result
            client.send("textDocument/didChange", {
                "textDocument": {"uri": a.as_uri(), "version": 2},
                "contentChanges": [{"text": a.read_text() + "func shared() {}\n"}],
            })
            client.diagnostics(a.as_uri(), 2, lambda ds: any(d.get("code") == "DuplicateDecl" for d in ds))
        finally:
            client.close()


def main():
    with tempfile.TemporaryFile(mode="w+") as log:
        client = Client(ROOT, log)
        try:
            files = [ROOT / "test" / name for name in ("235.go", "noinit.go")]
            for path in files:
                open_file(client, path)
            for path in files:
                client.diagnostics(path.as_uri(), 1, lambda ds: not ds)
            path = files[0]
            client.send("textDocument/didChange", {
                "textDocument": {"uri": path.as_uri(), "version": 2},
                "contentChanges": [{"text": path.read_text() + "\nvar broken = missingGonName\n"}],
            })
            findings = client.diagnostics(path.as_uri(), 2, lambda ds: any("missingGonName" in d["message"] for d in ds))
            assert len(findings) == 1 and findings[0]["code"] == "UndeclaredName", findings
            # Recovery must retain isolation from the other open compiler test.
            client.send("textDocument/didChange", {
                "textDocument": {"uri": path.as_uri(), "version": 3},
                "contentChanges": [{"text": path.read_text()}],
            })
            client.diagnostics(path.as_uri(), 3, lambda ds: not ds)
        finally:
            client.close()

    with tempfile.TemporaryDirectory(prefix="gon-package-") as tmp:
        folder = Path(tmp).resolve()
        (folder / "go.mod").write_text("module example.com/workspacecheck\n\ngo 1.26\n")
        check_package(folder)
    with tempfile.TemporaryDirectory(prefix="gon-workspace-", suffix=".dir", dir=ROOT / "test") as tmp:
        check_package(Path(tmp).resolve())
    print("PASS: compiler test isolation, real errors and recovery, ordinary multi-file packages, multi-file compiler fixtures")


if __name__ == "__main__":
    main()
