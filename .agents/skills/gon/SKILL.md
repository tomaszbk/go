---
name: gon
description: Use in repositories that contain this skill, which selected the Gon toolchain (a Go fork adding postfix ! error propagation and "or err { }" handlers). Covers choosing gon instead of go/gopls, the implemented syntax and its rules, and the gon query, check, refactor and explain commands for understanding, editing and validating code.
---

# Gon projects

This repository uses **Gon**, a fork of the Go toolchain. Source files are
still `.go`, tests `_test.go`, and modules use `go.mod`/`go.work`. Upstream
`go`, `gofmt` and `gopls` do not understand Gon syntax: in this repository use
`gon` for every Go command and gonpls-based tooling for analysis.

## Select the toolchain

- Run `gon` from `PATH`, or the path in `gon.compilerPath` of this project's
  `.vscode/settings.json` if it is set. In the Gon toolchain repository
  itself, use `./gon/bin/gon` after `python3 misc/gon/build.py`.
- Check it with `gon capabilities --json`: it reports the toolchain root,
  the gonpls version and the supported commands. If a command below is missing,
  the toolchain is older; say so instead of falling back to upstream tools.
- `gon build|test|vet|fmt|run|mod|doc|env` behave like the Go commands. Never
  replace them with `go` or install tools globally.

## Implemented language additions

Postfix `!` propagates the error of a call whose **last result has exactly type
`error`**, from a function whose own last result is `error`:

```go
data := os.ReadFile(path)!          // (data, err) → data, or return on error
name, size := stat()!              // several success results
flush()!                           // error-only call used as a statement
```

On error, the enclosing function returns zero values for its other results
(named results are reset too, before defers run) and the original error. It
never wraps the error and never uses panic/recover.

`call() or err { ... }` handles the error locally. The binding name is
required, scoped to the block and of type `error`:

```go
data := os.ReadFile(path) or err {
    return Config{}, fmt.Errorf("read config %q: %w", path, err)
}
```

Rules that produce compile errors (code `InvalidErrorHandling`):

- The operand must be a function or method call (not a conversion or value),
  and its last result must be exactly `error` (aliases are fine; named error
  interfaces, concrete error types and type parameters are not).
- `!` needs the nearest enclosing function literal or declaration to return
  `error` last; `or` handlers do not.
- A handler for a call with success results must end in a terminating
  statement (`return`, `panic`, ...). Error-only handlers may fall through.
- No labels, `goto` or labeled `break`/`continue` in handlers; `defer f()!` and
  `go f()!` are invalid (use a function body).
- `!` drops partial results. When a call returns useful values together with
  an error (for example `io.Reader.Read`), keep the explicit
  `n, err := r.Read(p)` form.

Compatibility: a standalone `!` followed by a newline ends the statement.
Never split prefix negation across lines (`valid := !` + newline + `ok` is a
syntax error); write `!ok` or `! ok` on one line. `!=` is unchanged, `or` is
still a valid identifier, and ordinary Go error handling keeps working.

Not implemented yet (do not write them): Result or Option types, `?`
propagation, sum types, pattern matching, new lambda syntax, safe navigation or
coalescing operators, conditional expressions. The full, current specification
is `$(gon env GOROOT)/design/error-handling/README.md`.

## Work with the tooling

Semantic commands analyze the **saved** files through the gonpls engine, each
in its own process (usually under a second; the first call after rebuilding
the toolchain can take several seconds). Save files before querying them. Add
`--json` for machine-readable output
(schema version 1). Exit status: 0 success, 1 findings, 2 usage, 3 toolchain or
workspace failure.

1. **Understand** before editing:
   - `gon query symbols Name` finds declarations; results are fully qualified
     and usable as targets.
   - `gon query def <target> --doc`, `gon query refs <target>` (always before
     changing a declaration), `gon query impls <target>`,
     `gon query type file.go:line:col` (also shows Gon constructs).
   - Targets: `Name`, `Type.Method`, `pkg.Name`, `./dir.Name`,
     `import/path.Type.Method`, `file.go:line:col` (1-based line, byte column,
     as in compiler errors) or `file.go:#offset`.
   - `gon doc pkg.Symbol` shows documentation for the dependency versions the
     module selects.
2. **Edit**, preferring small consistent changes. For renames use
   `gon refactor rename <target> NewName --dry-run` to preview, then run it
   without `--dry-run`. It refuses to write if files changed since the
   analysis, and it only reformats files that were already gofmt-clean.
3. **Check** each batch of edits: `gon check ./path/... --json`. Fix
   `category: "language"` errors first; `analysis` findings are advisory.
   `gon explain <code>` explains a code such as `InvalidErrorHandling`.
   `gon check` does not build or run tests; its `notVerified` field says what
   remains.
4. **Validate**: `gon fmt` the changed packages, then `gon vet`, `gon test`
   for the affected packages, and `gon build ./...` before finishing.

`gon help tooling` and `gon help <command>` list all flags.
