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
`InvalidErrorHandling` instead of `ErrorCode(152)`. `gonpls` and the launcher
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
