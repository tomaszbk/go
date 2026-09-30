# Maintained Gon language server

This directory is the maintained Gon adaptation of gopls. Edit these sources
directly. It is a separate Go module within the Gon repository, built with Gon's
compiler, public parser, AST, types, and formatter. Its upstream module path is
retained for internal imports; the public executable remains `gonpls`.

The initial source import is gopls v0.23.0 from `golang.org/x/tools/gopls`.
[`UPSTREAM.json`](UPSTREAM.json) records the module and go.mod checksums, tag,
repository, and commit. Both checksums were verified using `go mod download`
before copying source. The complete existing `misc/gon/patches/gopls.patch`,
including the pending compiler-test workspace fix, was applied during import.
That patch was then removed; future changes belong here. The upstream license,
tests, documentation, module requirements, and checksums are retained.

From the Gon repository root, after building the compiler:

```sh
python3 misc/gon/build.py
./gon/bin/gonpls version
cd tools/gonpls
../../gon/bin/gon test ./internal/protocol
../../gon/bin/gon test ./internal/cache -run 'TestGonCompilerTest|TestStandalone'
../../gon/bin/gon test ./internal/cmd
```

The build downloads only the pinned x/tools and Staticcheck sources in
`misc/gon/sources.json`, checks both module sums before using them, and applies
their maintained Gon patches into disposable `pkg/gon-tools` directories.
Relative `replace` directives in this module select those adapted dependencies.
They must not be removed or replaced with unmodified upstream dependencies:
Gon's AST and export data require the adaptations. The compiler's separately
vendored analyzers remain in `src/cmd/vendor`.

The server is built directly from this directory with `-mod=readonly`,
`-trimpath`, and `-buildvcs=false`. Builds do not rewrite its source, `go.mod`, or
`go.sum`, and do not download another gopls tree. Selected module versions and
content sums fix the dependency inputs; the exact Gon compiler and target
platform also determine the output. The generated dependency trees and the old
`pkg/gon-tools/gopls` tree can be deleted without losing maintained source.

## Gon adaptations

Carried over from the former patch:

- The `gonpls` name, `gonpls.*` command namespace and `gon` LSP language ID.
- Semantic tokens for postfix `!` and `or` handlers.
- LSP cancellation codes instead of generic failures for cancelled requests.
- Standalone loading of compiler test inputs under the toolchain's `test` tree.
- No upstream telemetry uploads or crash reports.

Added for the tooling commands of the public `gon` launcher:

- `internal/cmd/gon*.go`: `gon query`, `refactor`, `check`, `explain` and
  `capabilities`, reached as `gonpls gon ...` from `main.go` (see
  [misc/gon/CLI.md](../../misc/gon/CLI.md)). Each command drives a gonpls
  session in its own process through its LSP server and snapshots.
- `settings.InternalOptions.OnDemandDiagnostics`, which the command engine sets
  so that the server skips background diagnostics; editors are unaffected.
- No pkg.go.dev links for type-checker codes that upstream x/tools does not
  document, including `InvalidErrorHandling`; the x/tools patch names codes
  148–151 and Gon's own range starting at 10000.
- Upstream command-line test expectations and `internal/cmd/usage` help files
  updated for the `gonpls` name and command namespace.

When updating the upstream baseline, start from the module identified by
`UPSTREAM.json`, compare upstream changes against this maintained tree, carry
forward the Gon adaptations, update provenance and pins, and run the real LSP
and executable compatibility regressions in `misc/gon`. Do not reintroduce a
second gopls patch containing the changes maintained here.
