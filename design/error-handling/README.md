# Error handling on existing Go returns

Status: implemented for conventional Go error returns in this fork. The current
semantics and intentionally unsupported handler forms are specified below.
Result, Option, and recovery expressions that supply replacement values remain
future work.

## Objective and compatibility

Reduce repetitive `if err != nil` checks while retaining conventional Go
signatures, explicit error handling, and interoperability with existing packages.
All existing valid Go code must continue to compile and preserve its
language-defined behavior after this and the other planned language changes,
except for the explicitly accepted newline-after-prefix-`!` exception below.

## Accepted decisions

### Postfix `!` propagates errors

Use postfix `!` on a call with a conventional Go error return to obtain its
successful results or propagate the error from the enclosing function.

```go
func loadConfig(path string) (Config, error) {
    data := os.ReadFile(path)!
    return parseConfig(data)
}
```

For this example, the intended legacy equivalent is:

```go
func loadConfig(path string) (Config, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return Config{}, err
    }
    return parseConfig(data)
}
```

Keep same-line prefix Boolean negation: `!valid` and `! valid` remain Boolean
negation, while `read()!` is the new postfix propagation form. Propagation must
use normal branch-and-return control flow, not panic/recover.

### Accepted newline rule for `!`

The user explicitly accepted breaking compatibility for prefix negation with a
newline immediately after `!`. This previously valid Go spelling will be rejected:

```go
valid := !
    disabled
```

Add the standalone `!` token to the tokens eligible for automatic semicolon
insertion at a newline or EOF. This rule applies regardless of prefix/postfix
use. Follow existing Go handling of trailing whitespace and comments, including
newlines in block comments. The token `!=` is unaffected.

```go
valid := !disabled       // Supported prefix negation.
valid := ! disabled      // Supported prefix negation with whitespace.
data := read()!          // Newline terminates the propagation statement.
data := read()! + suffix // No semicolon after !: another token follows.
```

A trailing comment does not restore support for multiline prefix negation:
`valid := ! // comment` followed by the operand on the next line is also rejected.
Only this specific compatibility exception has been accepted; it does not relax
the compatibility requirement for unrelated constructs.

### `or err { ... }` handles errors locally

Use an `or` block with an explicit error binding when the caller wants to handle
the error instead of using automatic propagation. The block runs only on error.

```go
func loadConfig(path string) (Config, error) {
    data := os.ReadFile(path) or err {
        return Config{}, fmt.Errorf("read config %q: %w", path, err)
    }
    return parseConfig(data)
}
```

The equivalent legacy implementation is:

```go
func loadConfig(path string) (Config, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return Config{}, fmt.Errorf("read config %q: %w", path, err)
    }
    return parseConfig(data)
}
```

The binding is explicit rather than an implicitly injected variable. The syntax
must permit another binding name when needed. `or` must be contextual: existing
Go programs may use `or` as an identifier and must keep working.

Local recovery with a replacement value remains a future extension. The current
implementation does not introduce a yield construct or implicit final-expression
result. For a call with successful-result positions, every handler path must
terminate. For an error-only call, a handler may finish normally and execution
continues after the statement.

Existing explicit error handling and intentional discards remain valid:

```go
value, err := operation()
value, _ := operation()
_ = saveMetrics() // For an error-only return.
```

These lines illustrate separate alternatives, not a complete program. Discarding
an error retains the actual accompanying values; it does not compute a fallback.

### Future Option support uses `?` for absence

Keep error propagation and absence propagation distinct, following V's division:

| Construct | Meaning | Phase |
| --- | --- | --- |
| `call()!` | Propagate an error from existing Go error returns | Current |
| `call() or err { ... }` | Handle an error locally | Current |
| `optional?` | Propagate absence from an Option value | Future |

Future `Option[T]` support needs an explicit representation of presence and
absence. Its representation, constructors, zero-value policy, and exact type
syntax remain undecided. A `none` keyword is not required or selected; absence
could be a variant such as `None`. Nil alone cannot represent absence for every
payload type while distinguishing absence from valid zero values.

Do not implement Option, Result, sum types, or new lambdas as prerequisites for
the current error-handling feature. Do not reinterpret existing pointers or Go
error-return tuples as Option or Result implicitly.

## Implemented semantics

### Types and success results

The operand of either construct must be a function or method call whose last
result has exactly the predeclared `error` type. Aliases of `error` are accepted.
A distinct named interface, a concrete error implementation, a type parameter,
or a non-final error result does not meet this requirement. Existing calls with
those signatures continue to work with ordinary explicit Go handling.

On success, the construct supplies all results preceding the final error:

```go
bytes := read()!       // read returns ([]byte, error).
name, size := stat()!  // stat returns (string, int, error).
flush()!              // flush returns only error.
```

Ordinary Go rules for zero, single, and multiple expression values then apply.
A construct may also be a statement, discarding its successful results. An
error-only construct cannot be used where a value is required. Neither construct
changes the called function's signature or representation.

### Propagation and enclosing returns

Postfix `!` is valid only inside a function whose final result has exactly type
`error`, including an alias. On a non-nil error, it immediately returns from
that function. Every preceding enclosing result is set to its own type's zero
value, and the final result is set to the original error. The operand's partial
results are discarded. The enclosing success-result types need not match those
of the operand; zero values include nil pointers, slices, maps, channels,
functions, and interfaces, as well as generic type parameters' zero values.

This rule also applies to named results that already hold values. They receive
the zero values and propagated error before deferred calls run. A defer may
observe or modify those named results according to ordinary Go rules. No
panic/recover machinery implements propagation; `recover()` in a defer does
not see a synthetic panic. Real panics retain their normal behavior.

The enclosing function is the nearest function literal or declaration. A call
inside a nested function never propagates through its caller's boundary. The
same source-level rule applies to range-over-function loops: propagation leaves
the enclosing function and stops iteration, while preserving iterator cleanup
and ordinary defer behavior.

Error comparison has the normal meaning of `err != nil`. An interface containing
a typed nil pointer is a non-nil error and propagates. No error is wrapped,
converted into a panic, inspected via reflection, or implicitly treated as an
Option/Result value.

### Local handlers and scope

`call() or problem { ... }` evaluates the call once and runs the block only when
its final error is non-nil. `problem` is a new variable of type `error`, scoped
to that block. It may shadow an outer name without changing the outer binding,
and it may be unused. Declarations being initialized by the expression are not
in scope inside the handler, following normal short-declaration scope rules.

A handler belongs to the enclosing source function. Its `return` statements
return from that function; its `defer` statements are ordinary function defers.
It can wrap errors explicitly, return different values, panic explicitly, or
contain its own terminating control flow. The function containing a local
handler does not need an error result.

If the call has any successful results, the handler must be a terminating
statement sequence under Go's usual termination rules, even when those results
are discarded by an expression statement. A conditional return without a
terminating alternative is insufficient. The compiler cannot resume an
assignment without values for all its successful-result positions. Error-only
handlers may fall through; in that case execution resumes after the construct.

The first implementation rejects labels and `goto` in a handler, as well as
`break`, `continue`, or `fallthrough` that would escape the handler. Loops and
switches entirely inside the handler may use their normal control transfers.
These are restrictions on the new construct; existing Go control flow is
unchanged. Replacement values and outward handler transfers are future design
work, not silently synthesized behavior.

### Expressions, evaluation, and invalid contexts

Postfix propagation is a primary-expression operation: `read()! + suffix`
propagates before addition. The call's successful value can participate in
ordinary expression contexts, including arguments, assignments, composite
literals, conditions, channel sends, and nested calls. Existing multi-value
rules still apply. An `or` block likewise supplies only the call's successful
results when the call succeeds.

The operand and its arguments are evaluated once, preserving Go's evaluation
order guarantees. An error skips evaluation of the remaining expression and
statement work when propagation or the handler transfers control. Short-circuit
`&&` and `||`, loop-condition reevaluation, lazy switch-case evaluation, and
select operand evaluation retain their ordinary behavior. Ordinary unevaluated
operands, such as a constant `unsafe.Sizeof` argument, remain unevaluated.

Calls in arguments to `defer` and `go` are evaluated at the ordinary call site:
for example, a failure in `defer consume(read()!)` returns before that defer is
registered. `defer read()!` and `go read()!` are invalid, because their outer
operand is no longer the function call required by those statements. Use an
explicit function body when error handling belongs to a deferred call or a new
goroutine.

Propagation outside a function, propagation in a function without the required
error result, non-call operands, missing handler bindings, and handlers that
would resume without required values are compile errors. Applying propagation
twice does not turn a returned value into another implicit error-bearing call.
The normal rules for constants still apply; no new constant call evaluation is
introduced. `or` remains a contextual identifier, and neither `error` nor other
existing identifiers gain new reserved-keyword status.

Automatic propagation deliberately drops partial results. It is suitable only
when that matches the desired explicit handling. Code that processes useful
data together with an error must retain its explicit tuple handling. The
operators do not redefine conventional tuples as exclusive success/failure
alternatives.

## Required validation

Follow the paired-example requirement in [AGENTS.md](../../AGENTS.md): each
implemented pain point must have automated executable examples in both legacy
Go and modern syntax, and both must work on the modified toolchain. Also run the
legacy examples on a compatible unmodified toolchain.

For this feature, cover propagation, local handling and wrapping, success and
failure paths, error-only and multi-value returns, handler scope, operand side
effects, named results, defers, and partial-result handling as supported by the
final specification. Test that legacy identifiers named `or`, same-line prefix
`!`, and `!=` still work. Test automatic semicolon insertion after postfix `!`
at newline/EOF, including trailing comments. Add rejection tests for the removed
multiline prefix-negation form, including comment variants, and for invalid new
syntax and contexts.

The automated executable pair is
[`test/errorhandling.go`](../../test/errorhandling.go), with fixtures in
[`test/errorhandling.dir`](../../test/errorhandling.dir). Each program asserts
expected results and observable traces; the harness then compares their output.
It checks the modern program both with normal optimization and with inlining
disabled, runs `go vet`, and verifies separate-package export/import and generic
instantiation. Independent invalid programs must produce source diagnostics,
never an internal compiler error.

From the repository root, after building the fork:

```sh
GO_ERROR_HANDLING_BASELINE=/path/to/unmodified/go \
  ./bin/go test cmd/internal/testdir -run='Test/errorhandling.go$' -count=1
```

`GO_ERROR_HANDLING_BASELINE` additionally runs both legacy programs with that
unmodified toolchain and compares their output with the fork. Use a baseline
that supports range-over-function (Go 1.23 or later). Without this environment
variable, the same test remains part of the ordinary Go regression suite and
checks both spellings on the fork. The baseline check must be supplied during
feature validation; omitting it is not evidence of upstream compatibility.

Documentation snippets are explanatory examples; the executable fixture pair,
compiler/public-tooling tests, and the full Go test suite provide validation.
See [VALIDATION.md](VALIDATION.md) for the recorded platform, baseline, commands,
and results.
