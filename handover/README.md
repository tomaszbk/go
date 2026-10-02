# Handover: Gon features #8, #6, #7 (paused 2026-09-30)

Written for a new agent session that resumes this work. Read
[AGENTS.md](../AGENTS.md) first; its rules are binding. This folder is a
working note, not part of the toolchain: delete it once the work is merged.

The current conditional implementation, tooling infrastructure and upstream
integration are recorded together. Use `git status` for subsequent local work.
The user's own untracked files (`a.go`, `gon_logo.png`, `AGENTS.md`, `.agents/`,
and the pre-existing `design/` docs) must be left alone unless a task names them.

## Go upstream integration (2026-10-02)

The 96 Go master commits after the original fork point have been integrated,
through `67c1d421161d3d1ae9f5fd005e84c29fd0d9f896`. x/tools now uses
`v0.50.1-0.20260929192349-98444708d405`, shared by vet and gonpls.
See [the integration and validation record](upstream-2026-10-02.md) and
`misc/gon/UPSTREAM.json` before another upstream update. The integration is a Gon source commit; upstream
merge ancestry is not recorded. Existing conditional-feature gaps remain.

## Infrastructure update (2026-10-02)

This update supersedes the dependency/patch workflow below. User authorized the
integration checklist/runner, maintained modules, and shared AST traversal plus
analyzer tests. Sources now live in `tools/x-tools` and `tools/staticcheck`;
gonpls and vet share the same x/tools baseline. The three dependency patches and
`sources.json` have been retired after importing their changes. Run
`python3 misc/gon/vendor.py` to regenerate cmd vendor and `--check` to check it;
never use `--save` or edit the old `pkg/gon-tools` copies. Vendored test files have
moved to the maintained module; run x/tools tests there.

`ast.Children` shares its child inventory with Walk. The real gonpls analyzer
registry is exercised by `TestGonAnalyzers`, including optional analyzers and all
current Gon nodes. Staticcheck effects/copy/equality helpers are conservative;
the inliner explicitly refuses Gon control-flow callee bodies. Conditional
package-level typerefs are supported. Explicit semantics, inspector cursor
metadata and mutable AST slots still need feature-specific integration.

Use `GON_BASELINE_GO=/opt/homebrew/bin/go python3 misc/gon/validate.py tooling`.
The checklist is `misc/gon/INTEGRATION.md`; `features.json` records remaining
conditional work. No lambda/null-safety implementation has started. This does
not close #8: cgo, cover, editor completion/tokens/queries/extraction, and the
unconfirmed specification clarifications remain. See the appended validation
record in `misc/gon/VALIDATION.md` for current results.

## User decisions (genuine; do not reopen)

- **#8 conditional expression** `if c { a } else { b }`: spec accepted,
  [design/conditional-expression/README.md](../design/conditional-expression/README.md).
- **#6 lambdas**: all five decisions accepted on 2026-09-30, recorded in
  [design/lambda/README.md](../design/lambda/README.md) ("Status: accepted; not
  implemented"):
  - `(a, b) => expr` / `(a, b) => { ... }`, token `=>`;
  - parentheses always, also with one parameter (`(x) => …`);
  - parameter types are never written; the context supplies them
    (`func(x int) int => …` is a possible future extension, not now);
  - with a no-result target, an expression body is allowed only if the
    expression itself yields no values;
  - generic result inference from expression bodies only.
  [design/lambda/OPTIONS.md](../design/lambda/OPTIONS.md) is the decision aid.
- **#7 null safety** (`?.`, `??`): **parked** by the user. Its proposal,
  [design/null-safety/README.md](../design/null-safety/README.md), is not
  accepted. Do not implement it.
- **Testing policy**: focused tests only, meaning the modified packages, `-run`
  filters, and the single testdir test of a feature. Never run `all.bash`,
  `run.bash`, `go tool dist test`, `go test std`, `go test cmd/...`, or all of
  `test/`. This applies to subagents too.
- The user wants Opus subagents to split the work. Proposals of design choices
  must show Go code per option, with pros/cons that assume we control gofmt and
  gonpls.

## Environment

- Fork toolchain: `./bin/go` (GOROOT = repo root), currently installed with all
  changes. After compiler edits: `./bin/go install cmd/compile` (also `cmd/vet`,
  `cmd/gofmt`, `cmd/cover`, `cmd/cgo` as relevant). If broken:
  `cd src && GOROOT_BOOTSTRAP=/opt/homebrew/opt/go/libexec ./make.bash`.
- Unmodified baseline Go: `/opt/homebrew/bin/go` (go1.27.1).
- `pkg/gon-tools` is disposable build output of `python3 misc/gon/build.py`,
  rebuilt from the module cache plus `misc/gon/patches/{tools,staticcheck}.patch`.
  The patches are up to date; `build.py` has **not** been re-run since they
  changed. Vendored x/tools changes are persisted with
  `python3 misc/gon/vendor.py --save` into `misc/gon/patches/cmd-vendor.patch`
  (keep all existing hunks).
- Precedent for every integration point: the ErrorExpr feature (commits
  `256d7a160a`, `311e16f8b6` cover, `2e64ee4ef0` + `0d90e86926` cgo, and
  `misc/gon` history). Grep for `ErrorExpr` and `CondExpr` to find the hooks.

## #8 conditional expression: state

### Done and verified

| Area | Files | Verified with |
| --- | --- | --- |
| Compiler parser | `syntax/{nodes,parser,positions,printer,walk}.go`, `condexpr_test.go`; `midway/deepcopy.go`. Node: `CondExpr{Cond, Then, Else Expr; Rbrace Pos}` | `./bin/go test cmd/compile/internal/syntax` |
| Type checking | `types2/condexpr.go` (+ `expr.go`, `call.go`, `builtins.go`, `conversions.go`); `go/types` mirror, partly generated (`generate_test.go` filemap has `condexpr.go`); shared testdata `src/internal/types/testdata/check/condexpr.go` | types2 and go/types `-short`, `TestGenerate`, `TestStdlib`, `TestStdFixed`, `TestStdKen` |
| Public syntax | `go/ast` (`CondExpr`: If, Cond, Lbrace, Then, Rbrace, ElsePos, ElseLbrace, Else, ElseRbrace), `go/parser`, `go/printer` (incl. parentheses for programmatic ASTs at statement start), gofmt goldens, `api/fork.txt` | `go test go/ast go/parser go/printer go/format cmd/gofmt`; old-vs-new gofmt over all 38,182 repo files: 0 diffs |
| Lowering | `noder/condexpr.go`, `codes.go` (`exprCond` after `exprError`), writer/reader hooks. Lowers to `ir.InlinedCallExpr` with `if cond { tmp = a } else { tmp = b }`; no closures | below |
| Executable pair | `test/conditional.go`, `test/conditional.dir/*`: legacy == modern (also `-l`, `-N -l`), vet, gofmt, export/cross-package inlining, generics, 36 invalid programs | `GON_BASELINE_GO=/opt/homebrew/bin/go ./bin/go test cmd/internal/testdir -run='Test/conditional.go$' -count=1` |
| vet / vendored x/tools | inspector, edge, astutil, cfg (CondThen/CondElse/CondDone blocks), satisfy, lostcancel, copylock, shift; 2 pre-existing ErrorExpr vet false positives fixed; `src/cmd/vet/condexpr_test.go` | `cd src/cmd && ../../bin/go test golang.org/x/tools/go/cfg golang.org/x/tools/go/ast/inspector golang.org/x/tools/refactor/satisfy golang.org/x/tools/go/ast/astutil`; `GO_CONDITIONAL_EXPRESSION_BASELINE=/opt/homebrew/bin/go ./bin/go test -run 'TestCondExpr$' cmd/vet` |
| Pinned x/tools + staticcheck (in patches) | same AST helpers; `go/ssa/gon.go` condExpr (blocks + phi); staticcheck `go/ir` condExpr, plus a fix of a pre-existing upstream bug (switch tag evaluated after capturing the entry block); fixtures `misc/gon/analysisfixtures/conditional_{legacy,modern}.go` | in `pkg/gon-tools/tools`: `go test ./go/ssa ./go/ast/inspector ./go/ast/astutil ./go/cfg ./refactor/satisfy -run 'TestGon\|CondExpr\|ErrorHandling'`; in `pkg/gon-tools/staticcheck`: `go test -mod=mod ./go/ir -run TestGon` |

### Implementation clarifications (spec-silent rules already implemented)

These must be added to the #8 spec doc at closure, and reported to the user.
The first four were chosen by the coordinator, and the user has not
objected yet.

1. Conversions distribute: `T(if c {a} else {b})` ≡ `if c {T(a)} else {T(b)}`.
2. append element args, the delete key and the panic arg get a target type.
3. Generic calls: an argument with exactly one untyped-nil branch, for a
   parameter whose inferred type is a different interface, is rejected
   (strict, relaxable).
4. Interface targets: untyped non-nil branches first take their no-target type
   (so `Printf("%.2f", if c { f } else { 0 })` prints `0.00`), then each branch
   converts to the target.
5. Other rules:
   - parentheses keep the target;
   - two untyped nils are an error;
   - untyped non-constant results get their type later through updateExprType;
   - a statement-position conditional expression is "not used";
   - the compiler prefixes these parse errors with `syntax error:`;
   - nesting is structural (an operand of a larger expression inside a branch
     is allowed);
   - no new error codes.

To rebuild the full list, compare `types2/condexpr.go` comments and the
testdata.

### Remaining for #8

1. **cmd/cgo**: `walk` panics on `CondExpr`. Add a `condexpr.go` /
   `condexpr_bootstrap.go` split like `errorexpr.go` (cgo builds at bootstrap
   with upstream go/ast), review `gcc.go` call rewriting for CondExpr args, and
   add tests like `cmd/cgo/internal/testerrorhandling`.
2. **cmd/cover**: probably no `cover.go` change is needed. A CondExpr is part
   of its statement; `!` and handlers in branches still split blocks. Add paired
   `testdata/condflow` profile tests modelled on errorflow, and re-run
   `TestErrorFlowCoverage` and `TestLegacyInstrumentationUnchanged`.
3. **`go fix` inliner** (`x/tools/internal/refactor/inline`, both the vendored
   and pinned copies): `calleefx.go visitExpr` panics on CondExpr **and the
   pre-existing ErrorExpr**, and `falcon.go` skips them. Make it conservative,
   add tests, and re-save both patches.
4. **Pinned x/tools `needsParens`**: a CondExpr at statement start must be
   parenthesized by inline/refactorings. `NoEffects` should accept CondExpr.
5. **Staticcheck panics** in analyzers that gonpls enables by default:
   `code.MayHaveSideEffects` and astutil `CopyExpr`/`Equal` hit
   `default: panic` on CondExpr **and ErrorExpr**. These are reached by S1001,
   S1009, S1017, S1036, QF100x, SA4014 and SA5002.
6. **gonpls** (`tools/gonpls`):
   - semtok case for the `if`/`else` keywords (otherwise test.py fails on
     "failed to implement");
   - typerefs `visitExpr` panics on a package-level `var x = if …`;
   - a `gon query type` construct value (then update `misc/gon/CLI.md`);
   - completion inside branches;
   - refuse "extract variable" from inside a branch.
7. Extend `misc/gon/test.py` / `test_analysis.py`, run `build.py`, then those
   scripts.
8. Update `.agents/skills/gon/SKILL.md`: move the conditional expression to
   implemented; keep lambdas/null safety/Option/Result/sum types/pattern
   matching as not implemented. Repo-local only, never global.
9. **Unify the baseline env var**:
   - `test/conditional.go` reads `GON_BASELINE_GO`, falling back to
     `GO_ERROR_HANDLING_BASELINE`;
   - the cmd/vet test uses `GO_CONDITIONAL_EXPRESSION_BASELINE`.
   Pick one scheme and document it.
10. Minor: for a missing `}` before `else`, the compiler says "branch must be a
    single expression" while go/parser says "expected '}', found 'else'".
11. Closure:
    - review the whole diff;
    - set the #8 spec status to implemented and add the clarifications;
    - write `design/conditional-expression/VALIDATION.md` with the exact
      commands and results, as error-handling did;
    - update `misc/gon/README.md` and `VALIDATION.md`;
    - run a final focused validation.

## #6 lambdas: next

Implementation has not started. The detailed plan was written in a temporary
scratchpad that has since been lost. Its fixed interfaces were these;
re-derive the details from the accepted README's "Implementation notes":

- **Tokens**: `syntax._FatArrow` (after `_Star`) and `go/token.FATARROW`
  (after `TILDE`, so existing values don't change; `IsOperator` is true, with
  the lowest precedence).
- **Nodes**: `syntax.LambdaExpr` and `go/ast.LambdaExpr`. The syntax node has a
  `Lowered *FuncLit` filled by the noder after type checking.
- **Parsing**: both parsers decide at `(` in operand position, with no
  backtracking (`()`, `(x,`, `(x) =>`).
- **types.Info**: records the unnamed `*Signature`, Defs for each parameter
  (including `_`) and the function scope. No implicit return node is recorded;
  consumers use the signature's result count.
- **Typing**: reuse the target kinds from #8 (`assign`, `conv`, `infer`,
  `condOnly` in `types2/expr.go`). Conditional-expression branches pass the
  target to lambdas. Parentheses keep the target for CondExpr but not for
  lambdas. In generic calls a lambda waits until its parameter types are
  inferred, and its body is checked once.
- **Lowering**: rewrite to an ordinary FuncLit after type checking and
  **before** `prepareErrorPropagation` and `rangefunc.Rewrite`. Reuse
  `exprFuncLit`, so export data doesn't change. Package-level initializers
  must be lowered too.
- **Waves**:
  1. lambda-checker (syntax + types2 + go/types incl. inference + midway
     deepcopy), in parallel with lambda-ast (go/token, scanner, ast, parser,
     printer, gofmt, api; land go/token + go/ast first).
  2. lambda-lowering (noder + `test/lambda.go` pair with the go1.27.1
     baseline and invalid programs), in parallel with lambda-cmd (vendored
     x/tools + vet first, then cover, cgo) and lambda-gonpls (patched x/tools,
     SSA, staticcheck IR, gonpls, skill).
- **Risks**:
  - `prepareErrorPropagation` and rangefunc must treat a lambda as a function
    boundary with its own signature;
  - the label check in `types2/errorexpr.go` must stop at lambdas;
  - the hot `(` path in the parser;
  - inference ordering;
  - cover has no type information, so block bodies are instrumented like
    function literals and expression bodies count with their statement.

Suggested before #6 (proposed to the user, not yet decided): make AST
consumers generic so that a new node doesn't make them panic. Add a child
enumeration for any node, replace `default: panic` in walkers, and add a test
running every analyzer over a corpus file containing every Gon construct. That
would shrink the tooling long tail for #6 and later features. Ask the user
first.

## How the previous session worked

- One coordinator assigned waves with exclusive file ownership per agent, and
  agents shared a status board.
- Agents never commit or rewrite git state, and stay out of each other's files.
- Two agents reinstalling `cmd/compile` concurrently must keep their edits
  compile-consistent.
- Recreate a shared board (for example in this folder) if you work the same
  way.
