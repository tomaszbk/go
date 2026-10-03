---
name: gon
description: Use in repositories that contain this skill, which selected the Gon toolchain (a Go fork adding postfix ! error propagation and "or err { }" handlers). Covers choosing gon instead of go/gopls, the implemented syntax and its rules, and the gon query, check, refactor and explain commands for understanding, editing and validating code.
---

# Gon projects

This repository uses **Gon**, a fork of the Go toolchain. Source files are
still `.go`, tests `_test.go`, and modules use `go.mod`/`go.work`. Upstream
`go`, `gofmt` and `gopls` do not understand Gon syntax: in this repository use
`gon` for every Go command and gonpls-based tooling for analysis.

Gon 1.0 preserves compatibility with Go 1.27+ (user decision, 2026-10-02),
subject to the accepted newline-after-prefix-`!` exception. The current
unmodified validation baseline is stable Go 1.27.1. This support floor does
not require rewriting module language directives or historical test records.

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

Conditional expressions select one value and evaluate only the chosen branch:

```go
label := if count == 1 { "item" } else { "items" }
data := if cached { readCache()! } else { fetch()! }
```

Both branches are required and must contain one single-valued expression.
No init statements, else-if chains, or direct nested conditional expressions
are allowed. A conditional inside a separate operand (such as a call argument)
is allowed. At statement start, `if` remains an ordinary Go statement.

A contextual target type applies separately to each branch. Without a target,
typed branches must agree; an untyped branch takes the other branch's type.
For interface targets, untyped non-nil branches first acquire their no-target
types; `nil` instead converts directly to the target, preserving nil interfaces.
Explicit conversions distribute over branches. `gon query type` reports the
construct as `conditional-expression`.

Gonpls declines variable extraction from lazy branches and extraction of a
whole conditional, and the source inliner declines Gon control-flow bodies or
affected call sites. These restrictions preserve evaluation and target types.

Lambdas take their parameter and result types from context:

```go
var twice func(int) int = (x) => x * 2
slices.SortFunc(users, (a, b) => cmp.Compare(a.Name, b.Name))
var load func() ([]byte, error) = () => {
    data := os.ReadFile(path)!
    return data, nil
}
```

Parameters must be parenthesized identifiers, without written types. Bodies
may be expressions or ordinary blocks. A lambda has ordinary Go closure and
capture semantics; `return`, `defer`, `!` and `or` inside it belong to that
lambda. A standalone `f := (x) => x` has no target and is invalid. Parentheses
around the whole lambda do not pass a target. Generic calls can infer results
from expression bodies once parameter types are known; block bodies require
known result types. Lambda errors use `InvalidLambda`.
Gonpls offers conversions between lambdas and function literals when signature
identity and source context are safe; it declines generic inference and types
that cannot be written at the current position. Conversion to a lambda requires
a direct typed declaration or assignment without named results.

Null safety operators perform lazy nil checks:

```go
name := user?.Name ?? "guest"
value := callback?(arg()) ?? 0
config ??= defaults()
timeout := *config?.Timeout ?? 30
```

`?.` guards a pointer or interface; `?(` guards a function. A nil guard skips
the remaining primary-expression chain, including arguments and indexes.
Parentheses end that chain. A value-producing chain without `??` must have a
type that can be nil; `??` also accepts guarded non-nilable results and guarded
dereferences. Defaults run only on absence or nil, not on zero, false or empty
values. `??=` evaluates the location once and stores only when its current
value is nil. Ordinary interface nil semantics remain: an interface containing
a typed nil is non-nil. `??` is right associative and cannot mix with other
binary operators without parentheses. Safe chains cannot be assignment targets
or direct `go`/`defer` calls. These are conveniences, not static non-null
guarantees. Diagnostics use `InvalidNilSafety`.

Not implemented yet (do not write them): Result or Option types, lone `?`
propagation, sum types and pattern matching. Local specifications live in
`$(gon env GOROOT)/design/{error-handling,conditional-expression,lambda,null-safety}/README.md`.

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
4. **Validate**: format the changed packages and run the relevant focused
   `gon test`/`gon vet` checks. In this toolchain repository follow
   `misc/gon/INTEGRATION.md`, use the maintained module that owns the change,
   and do not run `gon build ./...` from the repository root.

`gon help tooling` and `gon help <command>` list all flags.

## Modernize existing syntax

Use `gon fix -diff ./...` to preview safe Gon syntax conversions and `gon fix
./...` to apply them. `gon check --severity=hint ./...` reports the same optional
suggestions without writes, including edits with `--json`. Gonpls exposes them
as editor hints and quick fixes. The analyzers are `gonerrors` (`!`/`or`),
`gonconditional`, `gonnil` (`??=`, `??`, `?.`, `?(`) and `gonlambda`.
For a focused preview, use `gon fix -gonerrors -diff ./...`; run `gon tool fix
help <analyzer>` for details. `gon fix -diff` exits 1 when it has a diff.

Only recognized, semantics-preserving patterns are converted. Partial-result
handlers, still-used error bindings, unavailable lambda target types and edits
that would discard comments or required imports are conservatively declined.
Ordinary Go constructs remain valid; these hints are not mandatory diagnostics.
