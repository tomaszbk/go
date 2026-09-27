package cmd

import (
	"context"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/tools/gopls/internal/settings"
)

type gonExplainResult struct {
	gonEnvelope
	Explanations []*gonExplanation `json:"explanations"`
}

// gonExplanation documents a diagnostic code or analyzer of the selected
// toolchain.
type gonExplanation struct {
	Code           string        `json:"code"`
	Kind           string        `json:"kind,omitempty"` // "type-error" or "analyzer"
	Number         *int          `json:"number,omitempty"`
	Summary        string        `json:"summary,omitempty"`
	Documentation  string        `json:"documentation,omitempty"`
	Cases          []gonCase     `json:"cases,omitempty"`
	DefaultEnabled *bool         `json:"defaultEnabled,omitempty"`
	URL            string        `json:"url,omitempty"`
	Source         string        `json:"source,omitempty"`
	References     []string      `json:"references,omitempty"`
	Error          *gonErrorInfo `json:"error,omitempty"`
}

// gonCase explains one message of a Gon diagnostic code.
type gonCase struct {
	Message string `json:"message"`
	Meaning string `json:"meaning"`
	Fix     string `json:"fix"`
}

// gonErrorHandlingCases are the diagnostics of postfix ! and "or name { }".
// They mirror the messages of types2 and go/types.
var gonErrorHandlingCases = []gonCase{
	{"error handling requires a function or method call",
		"The operand of ! or of an or handler must be a call, not a value, conversion or receive.",
		"Call the function directly, or handle the value with explicit Go code."},
	{"error handling is only permitted inside a function",
		"! and or handlers need an enclosing function, for example not in a package-level var.",
		"Move the call into a function, such as init or a helper returning error."},
	{"error handling requires a final result of type error",
		"The called function's last result must have exactly type error (an alias is accepted). " +
			"Named interfaces, concrete error types, type parameters and non-final errors do not qualify.",
		"Keep the explicit form: v, err := f(); if err != nil { ... }."},
	{"error propagation requires an enclosing function with a final result of type error",
		"Postfix ! returns from the nearest enclosing function literal or declaration, whose last result must have type error.",
		"Handle the error locally with 'or err { ... }'. Adding an error result changes the function's " +
			"contract; review its callers before choosing that design alternative."},
	{"error handler requires an explicit error binding",
		"An or handler must name the error it receives.",
		"Write 'call() or err { ... }'; the binding may be left unused."},
	{"labels are not permitted in error handlers",
		"Labeled statements inside an or handler are not supported yet.",
		"Move the labeled statement outside the handler or into a function literal."},
	{"branches to labels are not permitted in error handlers",
		"goto and labeled break/continue cannot leave or target an or handler.",
		"Restructure the control flow outside the handler; unlabeled loops inside the handler are allowed."},
	{"error handler for a value-producing call must terminate",
		"When the call has results besides its error, every path through the handler must end in " +
			"return, panic or another terminating statement, because execution cannot resume without those values.",
		"End the handler with a return or panic. Handlers of error-only calls may fall through."},
}

func (r *gonRequest) explain(ctx context.Context) (gonResult, error) {
	root := gonRoot(r.inv.env)
	codes, codesFile, err := gonTypeErrorCodes(root)
	if err != nil {
		return nil, gonErrorf(gonKindWorkspace, gonExitInfra, "reading the toolchain's error codes: %v", err)
	}
	result := &gonExplainResult{}
	for _, arg := range r.args {
		name := strings.TrimSuffix(strings.TrimPrefix(arg, "ErrorCode("), ")")
		ex := &gonExplanation{Code: arg}
		if n, err := strconv.Atoi(name); err == nil {
			for _, c := range codes {
				if c.number == n {
					name = c.name
				}
			}
		}
		if c := gonFindCode(codes, name); c != nil {
			n := c.number
			ex.Code, ex.Kind, ex.Number = c.name, "type-error", &n
			ex.Documentation = c.doc
			ex.Summary = gonFirstSentence(c.doc)
			ex.Source = fmt.Sprintf("%s:%d", codesFile, c.line)
			if c.name == "InvalidErrorHandling" {
				ex.Cases = gonErrorHandlingCases
				if doc := filepath.Join(root, "design", "error-handling", "README.md"); gonExists(doc) {
					ex.References = append(ex.References, doc)
				}
			}
		} else if a := gonFindAnalyzer(name); a != nil {
			enabled := a.Enabled(settings.DefaultOptions())
			ex.Code, ex.Kind, ex.DefaultEnabled = a.Analyzer().Name, "analyzer", &enabled
			ex.Documentation = strings.TrimSpace(a.Analyzer().Doc)
			ex.Summary = gonFirstSentence(ex.Documentation)
			ex.URL = a.Analyzer().URL
		} else {
			ex.Error = &gonErrorInfo{Kind: gonKindNotFound,
				Message: fmt.Sprintf("unknown code %q; use a code or analyzer name reported by 'gon check'", arg)}
		}
		result.Explanations = append(result.Explanations, ex)
	}
	result.gonEnvelope = *r.envelope("explain")
	return result, nil
}

func gonExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

type gonCodeDoc struct {
	name   string
	number int
	doc    string
	line   int
}

// gonTypeErrorCodes reads the type-checker error codes of the selected
// toolchain, so that explanations always match the compiler in use.
func gonTypeErrorCodes(root string) ([]gonCodeDoc, string, error) {
	path := filepath.Join(root, "src", "internal", "types", "errors", "codes.go")
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, path, err
	}
	var codes []gonCodeDoc
	for _, decl := range f.Decls {
		gen, ok := decl.(*ast.GenDecl)
		if !ok || gen.Tok != token.CONST {
			continue
		}
		iotaValue := true // whether the implicit repetition is iota-based
		for i, spec := range gen.Specs {
			vs := spec.(*ast.ValueSpec)
			value, known := i, iotaValue
			if len(vs.Values) == 1 {
				switch v := vs.Values[0].(type) {
				case *ast.Ident:
					known, iotaValue = v.Name == "iota", v.Name == "iota"
				case *ast.UnaryExpr:
					if lit, ok := v.X.(*ast.BasicLit); ok && v.Op == token.SUB {
						n, err := strconv.Atoi(lit.Value)
						value, known, iotaValue = -n, err == nil, false
					}
				default:
					known, iotaValue = false, false
				}
			}
			for _, id := range vs.Names {
				if id.Name == "_" || !known {
					continue
				}
				codes = append(codes, gonCodeDoc{
					name: id.Name, number: value,
					doc:  strings.TrimSpace(vs.Doc.Text()),
					line: fset.Position(id.Pos()).Line,
				})
			}
		}
	}
	return codes, path, nil
}

func gonFindCode(codes []gonCodeDoc, name string) *gonCodeDoc {
	for i := range codes {
		if codes[i].name == name {
			return &codes[i]
		}
	}
	return nil
}

func gonFindAnalyzer(name string) *settings.Analyzer {
	for _, a := range settings.AllAnalyzers {
		if a.Analyzer().Name == name {
			return a
		}
	}
	return nil
}

func gonFirstSentence(doc string) string {
	para, _, _ := strings.Cut(doc, "\n\n")
	para = strings.Join(strings.Fields(para), " ")
	if i := strings.Index(para, ". "); i >= 0 {
		return para[:i+1]
	}
	return para
}

func (e *gonExplainResult) exitCode() int {
	for _, ex := range e.Explanations {
		if ex.Error != nil {
			return gonExitFindings
		}
	}
	return gonExitOK
}

func (e *gonExplainResult) text(w io.Writer, r *gonRequest) {
	for i, ex := range e.Explanations {
		if i > 0 {
			fmt.Fprintln(w)
		}
		if ex.Error != nil {
			fmt.Fprintf(w, "%s: %s\n", ex.Code, ex.Error.Message)
			continue
		}
		fmt.Fprintf(w, "%s (%s", ex.Code, ex.Kind)
		if ex.Number != nil {
			fmt.Fprintf(w, " %d", *ex.Number)
		}
		fmt.Fprintf(w, ")\n\n%s\n", ex.Documentation)
		for _, c := range ex.Cases {
			fmt.Fprintf(w, "\n%q\n  %s\n  Fix: %s\n", c.Message, c.Meaning, c.Fix)
		}
		for _, ref := range ex.References {
			fmt.Fprintf(w, "\nSee %s\n", ref)
		}
		if ex.URL != "" {
			fmt.Fprintf(w, "\n%s\n", ex.URL)
		}
	}
}
