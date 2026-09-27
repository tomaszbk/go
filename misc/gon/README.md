# Gon commands and language server

`gon` runs this fork's Go-compatible command. `gonpls` runs a Gon-aware build of
gopls. Both pin their subprocesses to this toolchain without replacing the user's
Go installation. Source files still use `.go`.

## Build

First build the compiler from the repository's `src` directory with `./make.bash`.
Then, from the repository root:

```sh
python3 misc/gon/build.py
./gon/bin/gon run a.go
./gon/bin/gonpls version
```

The build requires Python 3, Git, and network access for dependencies not already
cached. It builds the maintained server source directly from `tools/gonpls`.
It downloads x/tools and Staticcheck pinned in `sources.json`, verifies their
sums, copies them into `pkg/gon-tools`, and applies their checked-in patches.
Transitive dependency versions/checksums come from `tools/gonpls/go.mod` and
`go.sum`. The shared module cache is never patched. `pkg/gon-tools` is disposable
build output; edit the dependency patches for persistent changes. Edit gonpls
directly in `tools/gonpls`; its [GON.md](../../tools/gonpls/GON.md) records source
provenance and build details. The old `pkg/gon-tools/gopls` copy is no longer used
and may be removed.

Public executables live in `gon/bin`. The private server lives in
`pkg/tool/<os>_<arch>/gonpls`. Keep that layout intact when moving the toolchain.
The launchers resolve installation symlinks to find their owning toolchain.

`gon` retains upstream command behavior, including `gon version` reporting the
underlying Go version, so tools that query the compiler can still parse it.
`gonpls version` identifies the Gon server and its gopls baseline. Both set
`GOROOT`, prepend the private toolchain directory to their **child process**
`PATH`, and force `GOTOOLCHAIN=local`. A module requiring a newer toolchain must
be handled by upgrading Gon; silently downloading an upstream compiler would
lose Gon syntax support.

## Tooling commands for agents, scripts and CI

`gon query`, `gon refactor`, `gon check`, `gon explain` and `gon capabilities`
are additive commands implemented by the gonpls engine. They need no LSP or MCP
configuration, print compact text or versioned JSON (`--json`), and use stable
exit statuses. Each command analyzes the saved files in its own process and
keeps no state afterwards; editors keep using their own long-lived gonpls.
See [CLI.md](CLI.md) for the contract.

```sh
gon query refs ./store.Sum
gon query type store/store.go:29:16 --json
gon check ./... --json
gon refactor rename store.Sum Add --dry-run
gon explain InvalidErrorHandling
```

Agents learn these conventions from the project skill in
[`.agents/skills/gon`](../../.agents/skills/gon/SKILL.md). It is not installed
globally; to opt another repository into Gon, copy it with
`python3 misc/gon/install.py --project-skill /path/to/repo`.

## Install without replacing Go

```sh
python3 misc/gon/install.py
gon run a.go
gon fmt ./path/to/package
gonpls check a.go
```

The installer creates only `gon` and `gonpls` symlinks in `~/.local/bin`, or in
`--bin-dir /your/chosen/path`. It refuses to overwrite unrelated files. It does
not edit shell profiles, global Go settings, or upstream `go`/`gopls` commands.
If needed, add that **public** directory to `PATH`, not this fork's `bin` directory.
Remove the two installed symlinks to uninstall the public commands.

## VS Code

The dedicated **Gon** extension (`gon-lang.gon`) is maintained in the separate
sibling `vscode-gon` repository, cloned from `golang/vscode-go`. Install its local
VSIX, disable the official Go extension in this workspace, and reload VS Code.
The Gon extension supports both normal Go projects (official `go` + `gopls`) and
Gon projects (`gon` + `gonpls`). This repository's `.vscode/settings.json` selects
Gon and its local tools. For another Gon project, use **Gon: Enable for This
Project**, or set workspace settings explicitly:

```json
{
  "gon.enabled": true,
  "gon.compilerPath": "/absolute/path/to/toolchain/gon/bin/gon",
  "gon.languageServerPath": "/absolute/path/to/toolchain/gon/bin/gonpls",
  "gon.serverSettings": { "semanticTokens": true }
}
```

The extension switches `.go` documents to the `Gon` editor language only in
opted-in projects, with a separate language client per folder. It does not claim
`.go` globally. Ordinary Go projects retain the Go language and official tools;
Go and Gon folders have been tested together in a multi-root window with only
the Gon extension active. `gonpls` accepts both `go` and `gon` LSP language
identifiers and uses its own `gonpls.*` command namespace. The extension routes
server commands to the correct project even with multiple Gon servers running.

The editor-title **▶ Run Go/Gon File** button saves the active `.go` file and
runs it with the appropriate compiler. Use `gon.run.mode: "package"` for programs
with multiple source files and `gon.run.args` for program arguments. Output goes
to a task terminal. Test files use `go test` or `gon test` instead.

The extension neither installs tools automatically nor changes the terminal's
Go environment. Run **Gon: Restart Language Server** after rebuilding gonpls.
An older setup can still use the official Go extension's `go.alternateTools`
setting to point at Gon tools, but remove that override when switching to the
Gon extension. The extension includes native Delve DAP debugging and a test
explorer; both select Go or Gon per project.

## Repository layout

This repository owns the compiler, public launcher sources, the maintained
`tools/gonpls` module (including the tooling commands in
`internal/cmd/gon*.go`), pinned x/tools/Staticcheck patches, the agent skill,
and their build/test scripts. The dependencies copied into `pkg/gon-tools` are recreated by the build
and are not separate maintained Git repositories. The
VS Code extension has its own Git repository at `../vscode-gon`; its build emits
`gon-0.1.0.vsix`. Neither the local clone nor the VSIX implies publication to
GitHub or the VS Code marketplace.

## Support and current limits

The server uses Gon's parser, type checker and formatter, with adapted AST
traversal, control-flow analysis, interface constraint discovery, and semantic
tokens for `!` and `or`. Diagnostics, hover, definitions, local rename, completion,
formatting, and import organization are covered by the stdio LSP regression.
Normal parse/type errors remain errors; Gon syntax diagnostics are not hidden.

The SSA and Staticcheck IR builders lower Gon error expressions to ordinary
branches and returns, including named result resets, `defer`, typed-nil errors,
multiple successful values and returns inside range-over-function loops.
The earlier unsupported-analysis guard has been removed. SSA-based `unusedwrite`
and Staticcheck `SA4006` diagnostics are verified through actual LSP requests.
Enable optional Staticcheck checks with `gon.serverSettings.staticcheck: true`.
The pinned x/tools export reader also supports the compiler's V5 package format.
Upstream crash uploads are not started by Gon.

Other upstream refactorings and optional integrations have not all been validated
with the new syntax. The server is an initial Gon adaptation, not a claim of
complete compatibility with every gopls or third-party analysis feature.

## Regression checks

Compiler regression inputs under the selected toolchain's `GOROOT/test` are
loaded as standalone files when they begin with a test-harness recipe such as
`// run` or `// errorcheck`. These inputs are independent programs, not one Go
package; upstream gopls otherwise reports spurious duplicate declarations.
Multi-file `*.dir` fixtures and ordinary projects retain normal package loading.
Intentional errors in error-checking fixtures remain visible as diagnostics.
Run `python3 misc/gon/test_workspace.py` to check this behavior over real LSP.

```sh
GO_ERROR_HANDLING_BASELINE=/path/to/official/go python3 misc/gon/test.py
GO_ERROR_HANDLING_BASELINE=/path/to/official/go python3 misc/gon/test_cli.py
GO_ERROR_HANDLING_BASELINE=/path/to/official/go python3 misc/gon/test_analysis.py
GO_ERROR_HANDLING_BASELINE=/path/to/official/go \
  ./gon/bin/gon test cmd/internal/testdir -run='Test/errorhandling.go$' -count=1
```

The first test executes equivalent legacy and modern programs with Gon and the
legacy program with official Go. It tests inherited-environment isolation and
CLI failure status, then launches `gonpls serve` and makes actual LSP requests.
It checks valid syntax, handler binding types and navigation, rename, completion,
formatting, semantic tokens, import actions, unsaved type errors and their repair.
The analysis test executes equivalent Go/Gon programs with both toolchains,
executes both versions in the SSA interpreter and builds Staticcheck IR with
sanity checks. The final command is the comprehensive language regression.

`test_cli.py` drives the tooling commands through the public launcher: it runs
a legacy/modern pair with Gon and the legacy program with official Go, checks
and renames both, verifies that Go commands keep their behavior and that an
unsaved editor buffer does not reach command-line results. The Go tests of the
commands run from `tools/gonpls` with `gon test ./internal/cmd -run TestGon`;
they cover target resolution, JSON contracts, checks, stale and atomic plans,
formatting and explanations.
