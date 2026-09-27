#!/usr/bin/env python3
"""Executable and real stdio-LSP regression tests for Gon public commands."""
import json
import os
from pathlib import Path
import queue
import sys
import subprocess
import tempfile
import threading
import time

ROOT = Path(__file__).resolve().parents[2]
BIN = ROOT / "gon" / "bin"
SUFFIX = ".exe" if os.name == "nt" else ""
GON, LSP = (BIN / (n + SUFFIX) for n in ("gon", "gonpls"))

MODERN = '''package main

import (
    "errors"
    "fmt"
)

func read(fail bool) (int, error) {
    if fail { return 7, errors.New("failure") }
    return 21, nil
}

func twice(fail bool) (int, error) {
    value := read(fail) or problem {
        return 0, fmt.Errorf("wrapped: %w", problem)
    }
    next := read(false)!
    return value + next, nil
}

func main() {
    a, e := twice(false)
    b, f := twice(true)
    fmt.Println(a, e, b, f)
}
'''
LEGACY = MODERN.replace('value := read(fail) or problem {', 'value, problem := read(fail)\n    if problem != nil {').replace(
    'next := read(false)!', 'next, err := read(false)\n    if err != nil { return 0, err }')


def command(*args, cwd=None, env=None):
    return subprocess.run([str(a) for a in args], cwd=cwd, env=env,
                          check=True, text=True, capture_output=True).stdout


class Client:
    def __init__(self, folder, log):
        self.process = subprocess.Popen([str(LSP), "serve"], cwd=folder,
                                        stdin=subprocess.PIPE, stdout=subprocess.PIPE, stderr=log)
        self.messages = queue.Queue()
        self.notifications = []
        self.serial = 0
        threading.Thread(target=self.read, daemon=True).start()
        self.result = self.request("initialize", {
            "processId": os.getpid(), "rootUri": folder.as_uri(),
            "workspaceFolders": [{"uri": folder.as_uri(), "name": "gon-test"}],
            "capabilities": {"textDocument": {
                "publishDiagnostics": {"versionSupport": True},
                "semanticTokens": {"requests": {"full": True}, "formats": ["relative"],
                    "tokenTypes": ["namespace", "type", "class", "enum", "interface", "struct", "typeParameter", "parameter", "variable", "property", "enumMember", "event", "function", "method", "macro", "keyword", "modifier", "comment", "string", "number", "regexp", "operator"],
                    "tokenModifiers": ["declaration", "definition", "readonly", "static", "deprecated", "abstract", "async", "modification", "documentation", "defaultLibrary"]}}},
            "initializationOptions": {"semanticTokens": True, "staticcheck": True, "analyses": {"unusedwrite": True}, "diagnosticsDelay": "10ms"},
        })
        self.send("initialized", {})

    def read(self):
        try:
            while True:
                headers = {}
                while True:
                    line = self.process.stdout.readline()
                    if not line:
                        raise EOFError("gonpls stdout closed")
                    if line == b"\r\n":
                        break
                    key, value = line.decode().split(":", 1)
                    headers[key.lower()] = value.strip()
                self.messages.put(json.loads(self.process.stdout.read(int(headers["content-length"]))))
        except Exception as error:
            self.messages.put(error)

    def send(self, method, params, serial=None):
        msg = {"jsonrpc": "2.0", "method": method, "params": params}
        if serial is not None:
            msg["id"] = serial
        self.write(msg)

    def write(self, msg):
        body = json.dumps(msg).encode()
        self.process.stdin.write(f"Content-Length: {len(body)}\r\n\r\n".encode() + body)
        self.process.stdin.flush()

    def receive(self, timeout):
        msg = self.messages.get(timeout=timeout)
        if isinstance(msg, Exception):
            raise msg
        if "method" in msg and "id" in msg:
            # Acknowledge optional server requests; configuration is supplied
            # at initialization, so workspace/configuration isn't advertised.
            self.write({"jsonrpc": "2.0", "id": msg["id"], "result": None})
        elif "method" in msg:
            self.notifications.append(msg)
        return msg

    def request(self, method, params):
        self.serial += 1
        serial = self.serial
        self.send(method, params, serial)
        deadline = time.monotonic() + 90
        while time.monotonic() < deadline:
            msg = self.receive(max(0.01, deadline - time.monotonic()))
            if msg.get("id") == serial and "method" not in msg:
                assert "error" not in msg, (method, msg)
                return msg.get("result")
        raise TimeoutError(method)

    def diagnostics(self, uri, version, predicate):
        deadline = time.monotonic() + 90
        while time.monotonic() < deadline:
            for msg in self.notifications:
                p = msg.get("params", {})
                if (msg.get("method") == "textDocument/publishDiagnostics"
                    and p.get("uri") == uri and p.get("version") == version
                    and predicate(p["diagnostics"])):
                    return p["diagnostics"]
            try:
                self.receive(max(0.01, deadline - time.monotonic()))
            except queue.Empty:
                break
        raise TimeoutError((uri, version, self.notifications))

    def close(self):
        try:
            if self.process.poll() is not None:
                return
            self.request("shutdown", None)
            self.send("exit", None)
            self.process.wait(timeout=10)
            assert self.process.returncode == 0
        finally:
            if self.process.poll() is None:
                self.process.kill()
                self.process.wait()


def position(source, needle, offset=0):
    index = source.index(needle) + offset
    before = source[:index]
    return {"line": before.count("\n"), "character": len(before.rsplit("\n", 1)[-1])}


def apply_edits(source, edits):
    lines = source.splitlines(keepends=True)
    def offset(pos):
        return sum(map(len, lines[:pos["line"]])) + pos["character"]
    for edit in sorted(edits, key=lambda e: offset(e["range"]["start"]), reverse=True):
        start, end = (offset(edit["range"][p]) for p in ("start", "end"))
        source = source[:start] + edit["newText"] + source[end:]
    return source


def main():
    baseline = os.environ.get("GO_ERROR_HANDLING_BASELINE")
    if not baseline:
        raise SystemExit("Set GO_ERROR_HANDLING_BASELINE to an unmodified Go executable")
    before = command(baseline, "version")
    with tempfile.TemporaryDirectory(prefix="gon-tools-test-") as temp:
        folder = Path(temp).resolve()
        # Installer is idempotent, resolves symlinks, and refuses collisions.
        public = folder / "public"
        install = ROOT / "misc" / "gon" / "install.py"
        command(sys.executable, install, "--bin-dir", public)
        command(sys.executable, install, "--bin-dir", public)
        assert command(public / ("gon" + SUFFIX), "env", "GOROOT").strip() == str(ROOT)
        other = folder / "occupied"
        other.mkdir()
        (other / ("gon" + SUFFIX)).write_text("keep me")
        refused = subprocess.run([sys.executable, str(install), "--bin-dir", str(other)], capture_output=True)
        assert refused.returncode != 0
        assert (other / ("gon" + SUFFIX)).read_text() == "keep me"
        assert not (other / ("gonpls" + SUFFIX)).exists()
        (folder / "go.mod").write_text("module example.com/gon-test\n\ngo 1.26\n")
        file = folder / "main.go"
        file.write_text(LEGACY)
        expected = "42 <nil> 0 wrapped: failure\n"
        assert command(baseline, "run", file, cwd=folder) == expected
        assert command(GON, "run", file, cwd=folder) == expected
        file.write_text(MODERN)
        assert command(GON, "run", file, cwd=folder) == expected
        assert command(GON, "test", "./...", cwd=folder)
        command(GON, "vet", "./...", cwd=folder)
        # An inherited upstream GOROOT/toolchain must not override Gon.
        env = dict(os.environ, GOROOT=command(baseline, "env", "GOROOT").strip(), GOTOOLCHAIN="auto")
        assert command(GON, "run", file, cwd=folder, env=env) == expected
        assert command(GON, "env", "GOTOOLCHAIN", cwd=folder, env=env).strip() == "local"
        failed = subprocess.run([str(GON), "build", "./does-not-exist"], cwd=folder, capture_output=True)
        assert failed.returncode != 0
        uri = file.as_uri()
        doc = {"textDocument": {"uri": uri}}
        with (folder / "lsp.log").open("w+") as log:
            client = Client(folder, log)
            try:
                assert client.result["serverInfo"]["name"] == "gonpls", client.result
                assert all(c.startswith("gonpls.") for c in client.result["capabilities"]["executeCommandProvider"]["commands"])
                client.send("textDocument/didOpen", {"textDocument": {
                    "uri": uri, "languageId": "go", "version": 1, "text": MODERN}})
                diagnostics = client.diagnostics(uri, 1, lambda d: True)
                assert not [d for d in diagnostics if d.get("severity") == 1], diagnostics
                hover = client.request("textDocument/hover", dict(doc, position=position(MODERN, "problem)")))
                assert "error" in json.dumps(hover), hover
                imports = client.request("workspace/executeCommand", {"command": "gonpls.list_imports", "arguments": [{"URI": uri}]})
                assert "fmt" in json.dumps(imports), imports
                definitions = client.request("textDocument/definition", dict(doc, position=position(MODERN, "problem)")))
                assert definitions and definitions[0]["range"]["start"] == position(MODERN, "problem {"), definitions
                rename = client.request("textDocument/rename", dict(doc, position=position(MODERN, "problem)"), newName="failure"))
                assert json.dumps(rename).count('"newText": "failure"') == 2, rename
                completion = client.request("textDocument/completion", dict(doc, position=position(MODERN, "problem)", 3)))
                assert any(item["label"] == "problem" for item in completion["items"]), completion
                edits = client.request("textDocument/formatting", dict(doc, options={"tabSize": 4, "insertSpaces": False}))
                assert edits, "expected formatting edits"
                formatted = apply_edits(MODERN, edits)
                assert "or problem {" in formatted and "read(false)!" in formatted
                file.write_text(formatted)
                assert command(GON, "run", file, cwd=folder) == expected
                file.write_text(MODERN)
                tokens = client.request("textDocument/semanticTokens/full", doc)
                assert tokens and tokens["data"], tokens
                legend = client.result["capabilities"]["semanticTokensProvider"]["legend"]["tokenTypes"]
                classified = {}
                line = column = 0
                for delta, char, length, kind, _ in zip(*[iter(tokens["data"])] * 5):
                    line += delta
                    column = char if delta else column + char
                    classified[(line, column, length)] = legend[kind]
                for needle, size, kind in [("or problem", 2, "keyword"), ("!", 1, "operator")]:
                    pos = position(MODERN, needle)
                    assert classified[(pos["line"], pos["character"], size)] == kind, classified
                # Unsaved real errors must still be diagnosed, then cleared.
                broken = MODERN.replace("return value + next, nil", 'return "wrong" + next, nil')
                client.send("textDocument/didChange", dict(doc, textDocument={"uri": uri, "version": 2}, contentChanges=[{"text": broken}]))
                errors = client.diagnostics(uri, 2, lambda ds: any(d.get("severity") == 1 for d in ds))
                assert any("mismatch" in d["message"] or "invalid operation" in d["message"] for d in errors), errors
                client.send("textDocument/didChange", dict(doc, textDocument={"uri": uri, "version": 3}, contentChanges=[{"text": MODERN}]))
                client.diagnostics(uri, 3, lambda ds: not any(d.get("severity") == 1 for d in ds))
                unused = MODERN.replace('"errors"', '"errors"\n    "strings"')
                client.send("textDocument/didChange", dict(doc, textDocument={"uri": uri, "version": 4}, contentChanges=[{"text": unused}]))
                actions = client.request("textDocument/codeAction", dict(doc, range={"start": {"line": 0, "character": 0}, "end": {"line": 0, "character": 0}}, context={"diagnostics": [], "only": ["source.organizeImports"]}))
                action = next(a for a in actions if a.get("kind") == "source.organizeImports")
                changes = action["edit"].get("documentChanges", [])
                import_edits = [e for change in changes for e in change.get("edits", [])]
                import_edits += action["edit"].get("changes", {}).get(uri, [])
                organized = apply_edits(unused, import_edits)
                assert '"strings"' not in organized, action
                file.write_text(organized)
                assert command(GON, "run", file, cwd=folder) == expected
                # The dedicated VS Code extension uses a distinct language ID
                # to coexist with official Go providers in the same window.
                client.send("textDocument/didClose", doc)
                client.send("textDocument/didOpen", {"textDocument": {
                    "uri": uri, "languageId": "gon", "version": 5, "text": MODERN}})
                client.diagnostics(uri, 5, lambda ds: not any(d.get("severity") == 1 for d in ds))
                hover = client.request("textDocument/hover", dict(doc, position=position(MODERN, "problem)")))
                assert "error" in json.dumps(hover), hover
                # Both independent flow builders must produce real diagnostics
                # on a package that uses Gon syntax, not silently skip analysis.
                analyzed = MODERN + """
func analysisValues() (int, string, error) { return 1, "value", nil }
func analysisProbe() error {
    x := read(false)!
    x = 7
    fmt.Println(x)
    s := struct{ value int }{}
    s.value = read(false)!
    a, b := analysisValues()!
    a, b = 2, "replaced"
    fmt.Println(a, b)
    return nil
}
"""
                client.send("textDocument/didChange", dict(doc, textDocument={"uri": uri, "version": 6}, contentChanges=[{"text": analyzed}]))
                findings = client.diagnostics(uri, 6, lambda ds: {"SA4006", "unusedwrite"}.issubset({str(d.get("source")) for d in ds}))
                assert not any(d.get("severity") == 1 for d in findings), findings
                messages = {d["message"] for d in findings if d.get("source") == "SA4006"}
                assert {"this value of a is never used", "this value of b is never used"}.issubset(messages), findings
                # Cancel real code actions over the wire. A race may complete
                # successfully, but a cancellation must never use error code 0.
                for _ in range(3):
                    client.serial += 1
                    request_id = client.serial
                    client.send("textDocument/codeAction", dict(doc,
                        range={"start": {"line": 0, "character": 0}, "end": {"line": 0, "character": 0}},
                        context={"diagnostics": [], "only": ["source.organizeImports"]}), request_id)
                    client.send("$/cancelRequest", {"id": request_id})
                    while True:
                        reply = client.receive(30)
                        if reply.get("id") == request_id and "method" not in reply:
                            if "error" in reply:
                                assert reply["error"]["code"] in (-32800, -32802), reply
                            break

            finally:
                try:
                    client.close()
                finally:
                    log.seek(0)
                    logs = log.read()
                    if client.process.returncode != 0 or "panic" in logs.lower():
                        print(logs, flush=True)
            log.seek(0)
            logs = log.read()
            assert "panic" not in logs.lower() and "failed to implement" not in logs, logs
    assert command(baseline, "version") == before
    print("PASS: executable legacy/modern pair, baseline, isolation, CLI errors, LSP diagnostics, hover, definition, rename, completion, formatting, semantic tokens, imports, unsaved edits, SSA/Staticcheck diagnostics")


if __name__ == "__main__":
    main()
