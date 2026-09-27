package cmd

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	goformat "go/format"
	"go/token"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"golang.org/x/tools/gopls/internal/protocol"
	"golang.org/x/tools/internal/diff"
)

// gonPlan is a previewable, revision-checked set of source edits.
type gonPlan struct {
	gonEnvelope
	Status  string         `json:"status"` // "planned" or "applied"
	Target  *gonTarget     `json:"target,omitempty"`
	NewName string         `json:"newName,omitempty"`
	Files   []*gonPlanFile `json:"files"`
	Summary gonPlanSummary `json:"summary"`
	old     map[string][]byte
}

type gonPlanSummary struct {
	Files int `json:"files"`
	Edits int `json:"edits"`
}

// gonPlanFile holds the edits for one file. Edits are sorted, do not overlap,
// and refer to the content whose SHA-256 is recorded.
type gonPlanFile struct {
	Path       string    `json:"path"`
	SHA256     string    `json:"sha256"`
	NewSHA256  string    `json:"newSha256"`
	Formatting string    `json:"formatting"` // "unchanged", "gofmt", or "skipped" (file was not gofmt-clean)
	Edits      []gonEdit `json:"edits"`
}

func gonHash(b []byte) string {
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:])
}

func (p *gonPlan) exitCode() int { return gonExitOK }

func (r *gonRequest) rename(ctx context.Context) (gonResult, error) {
	spec, newName := r.args[0], r.args[1]
	if !token.IsIdentifier(newName) {
		return nil, gonErrorf(gonKindUsage, gonExitUsage, "invalid new name %q: want a Go identifier", newName)
	}
	if _, err := r.ensureEngine(ctx); err != nil {
		return nil, err
	}
	t, err := r.resolveTarget(ctx, spec, nil)
	if err != nil {
		return nil, err
	}
	edit, err := r.engine.cli.server.Rename(ctx, &protocol.RenameParams{
		TextDocumentPositionParams: protocol.LocationTextDocumentPositionParams(t.loc),
		NewName:                    newName,
	})
	if err != nil {
		return nil, gonErrorf(gonKindRejected, gonExitFindings, "cannot rename %s: %v", spec, err)
	}
	plan := &gonPlan{Status: "planned", Target: t, NewName: newName, Files: []*gonPlanFile{}, old: make(map[string][]byte)}
	edits := make(map[protocol.DocumentURI][]protocol.TextEdit)
	if edit != nil {
		for _, change := range edit.DocumentChanges {
			if change.TextDocumentEdit == nil {
				return nil, gonErrorf(gonKindRejected, gonExitFindings,
					"renaming %s would create, move or delete files; this is not supported by gon refactor", spec)
			}
			uri := change.TextDocumentEdit.TextDocument.URI
			edits[uri] = append(edits[uri], protocol.AsTextEdits(change.TextDocumentEdit.Edits)...)
		}
		for uri, list := range edit.Changes {
			edits[uri] = append(edits[uri], list...)
		}
	}
	for uri, list := range edits {
		d, err := r.doc(ctx, uri)
		if err != nil {
			return nil, err
		}
		updated, byteEdits, err := protocol.ApplyEdits(d.mapper, list)
		if err != nil {
			return nil, gonErrorf(gonKindInternal, gonExitInfra, "%s: %v", uri.Path(), err)
		}
		pf := &gonPlanFile{Path: uri.Path(), SHA256: d.hash, Formatting: "unchanged"}
		// Format only files that were already gofmt-clean, so that applying
		// the plan never introduces unrelated formatting changes. When gofmt
		// realigns code, whole changed lines replace the precise edits.
		if clean, err := goformat.Source(d.content); err != nil || !bytes.Equal(clean, d.content) {
			pf.Formatting = "skipped"
		} else if formatted, err := goformat.Source(updated); err != nil {
			pf.Formatting = "skipped"
		} else if !bytes.Equal(formatted, updated) {
			updated, pf.Formatting = formatted, "gofmt"
			byteEdits = diff.Lines(string(d.content), string(updated))
		}
		pf.NewSHA256 = gonHash(updated)
		diff.SortEdits(byteEdits)
		for _, e := range byteEdits {
			pf.Edits = append(pf.Edits, gonEdit{Location: d.offsets(e.Start, e.End), NewText: e.New})
		}
		if len(pf.Edits) == 0 {
			continue
		}
		plan.old[pf.Path] = d.content
		plan.Files = append(plan.Files, pf)
		plan.Summary.Edits += len(pf.Edits)
	}
	sort.Slice(plan.Files, func(i, j int) bool { return plan.Files[i].Path < plan.Files[j].Path })
	plan.Summary.Files = len(plan.Files)
	plan.gonEnvelope = *r.envelope("refactor.rename")
	if !r.bool("dry-run") {
		if err := gonApplyPlan(plan); err != nil {
			return nil, err
		}
		plan.Status = "applied"
	}
	return plan, nil
}

// applyPlanFile implements "gon refactor apply".
func (r *gonRequest) applyPlanFile(ctx context.Context) (gonResult, error) {
	var data []byte
	var err error
	if r.args[0] == "-" {
		data, err = io.ReadAll(os.Stdin)
	} else {
		data, err = os.ReadFile(r.abs(r.args[0]))
	}
	if err != nil {
		return nil, gonErrorf(gonKindUsage, gonExitUsage, "reading plan: %v", err)
	}
	var plan gonPlan
	if err := json.Unmarshal(data, &plan); err != nil {
		return nil, gonErrorf(gonKindUsage, gonExitUsage, "invalid plan: %v", err)
	}
	if plan.SchemaVersion != gonSchemaVersion || !strings.HasPrefix(plan.Operation, "refactor.") || plan.Status != "planned" {
		return nil, gonErrorf(gonKindUsage, gonExitUsage,
			"not a schema version %d refactoring plan in state \"planned\"", gonSchemaVersion)
	}
	if err := gonApplyPlan(&plan); err != nil {
		return nil, err
	}
	plan.Status = "applied"
	plan.gonEnvelope = *r.envelope("refactor.apply")
	return &plan, nil
}

// gonApplyPlan writes a plan after verifying that every file still has the
// analyzed content. No file is modified if any check fails, and a failure
// while replacing files restores those already replaced.
func gonApplyPlan(plan *gonPlan) error {
	type change struct {
		path     string
		old, new []byte
		mode     os.FileMode
		temp     string
	}
	var changes []*change
	var stale []string
	for _, f := range plan.Files {
		content, err := os.ReadFile(f.Path)
		if err != nil || gonHash(content) != f.SHA256 {
			stale = append(stale, f.Path)
			continue
		}
		updated, err := gonApplyEdits(content, f.Edits)
		if err != nil || gonHash(updated) != f.NewSHA256 {
			return gonErrorf(gonKindUsage, gonExitUsage, "%s: plan edits do not produce the planned content", f.Path)
		}
		st, err := os.Stat(f.Path)
		if err != nil {
			return gonErrorf(gonKindInternal, gonExitInfra, "%v", err)
		}
		changes = append(changes, &change{path: f.Path, old: content, new: updated, mode: st.Mode().Perm()})
	}
	if len(stale) > 0 {
		return gonErrorf(gonKindStale, gonExitFindings,
			"files changed since the plan was computed; no file was modified: %s", strings.Join(stale, ", "))
	}

	cleanup := func() {
		for _, c := range changes {
			if c.temp != "" {
				os.Remove(c.temp)
			}
		}
	}
	for _, c := range changes {
		tmp, err := os.CreateTemp(filepath.Dir(c.path), "."+filepath.Base(c.path)+".gon-*")
		if err == nil {
			c.temp = tmp.Name()
			_, err = tmp.Write(c.new)
			if cerr := tmp.Close(); err == nil {
				err = cerr
			}
		}
		if err == nil {
			err = os.Chmod(c.temp, c.mode)
		}
		if err != nil {
			cleanup()
			return gonErrorf(gonKindInternal, gonExitInfra, "writing %s: %v", c.path, err)
		}
	}
	// Recheck immediately before replacing, to narrow the window for
	// concurrent modifications.
	for _, c := range changes {
		if content, err := os.ReadFile(c.path); err != nil || !bytes.Equal(content, c.old) {
			cleanup()
			return gonErrorf(gonKindStale, gonExitFindings, "%s changed while applying; no file was modified", c.path)
		}
	}
	for i, c := range changes {
		if err := os.Rename(c.temp, c.path); err != nil {
			for _, done := range changes[:i] {
				os.WriteFile(done.path, done.old, done.mode) // best-effort restoration
			}
			cleanup()
			return gonErrorf(gonKindInternal, gonExitInfra, "replacing %s: %v; earlier files were restored", c.path, err)
		}
		c.temp = ""
	}
	return nil
}

// gonApplyEdits applies sorted, non-overlapping byte-offset edits.
func gonApplyEdits(content []byte, edits []gonEdit) ([]byte, error) {
	var out bytes.Buffer
	last := 0
	for _, e := range edits {
		start, end := e.Location.Offset, e.Location.EndOffset
		if start < last || end < start || end > len(content) {
			return nil, fmt.Errorf("invalid edit range %d-%d", start, end)
		}
		out.Write(content[last:start])
		out.WriteString(e.NewText)
		last = end
	}
	out.Write(content[last:])
	return out.Bytes(), nil
}

func (p *gonPlan) text(w io.Writer, r *gonRequest) {
	if p.Status == "planned" && p.old != nil {
		for _, f := range p.Files {
			var edits []diff.Edit
			for _, e := range f.Edits {
				edits = append(edits, diff.Edit{Start: e.Location.Offset, End: e.Location.EndOffset, New: e.NewText})
			}
			name := filepath.ToSlash(r.rel(f.Path))
			if unified, err := diff.ToUnified("a/"+name, "b/"+name, string(p.old[f.Path]), edits, diff.DefaultContextLines); err == nil {
				fmt.Fprint(w, unified)
			}
		}
	}
	verb := "would change"
	if p.Status == "applied" {
		verb = "changed"
	}
	fmt.Fprintf(w, "%s %d edit(s) in %d file(s)", verb, p.Summary.Edits, p.Summary.Files)
	if p.Target != nil && p.NewName != "" {
		fmt.Fprintf(w, " renaming %s to %s", p.Target.Spec, p.NewName)
	}
	fmt.Fprintln(w)
	if p.Status == "applied" {
		for _, f := range p.Files {
			note := ""
			switch f.Formatting {
			case "gofmt":
				note = " (realigned by gofmt)"
			case "skipped":
				note = " (not gofmt-clean; left unformatted)"
			}
			fmt.Fprintf(w, "\t%s%s\n", r.rel(f.Path), note)
		}
	}
}
