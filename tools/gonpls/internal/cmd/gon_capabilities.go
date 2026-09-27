package cmd

import (
	"context"
	"fmt"
	"io"
	"strings"
)

type gonCapabilitiesResult struct {
	gonEnvelope
	Capabilities gonCapabilities `json:"capabilities"`
	Commands     []string        `json:"commands"`
}

type gonCapabilities struct {
	Queries       []string `json:"queries"`
	Rename        bool     `json:"rename"`
	RefactorApply bool     `json:"refactorApply"`
	Check         bool     `json:"check"`
	Explain       bool     `json:"explain"`
	// Persistent reports whether commands reuse state kept between
	// invocations. Every command analyzes in its own process, so it is false.
	Persistent bool   `json:"persistent"`
	Positions  string `json:"positions"`
}

func (r *gonRequest) capabilities(ctx context.Context) (gonResult, error) {
	result := &gonCapabilitiesResult{
		gonEnvelope: *r.envelope("capabilities"),
		Capabilities: gonCapabilities{
			Queries: []string{"def", "refs", "impls", "type", "symbols"},
			Rename:  true, RefactorApply: true, Check: true, Explain: true,
			Positions: "1-based line and UTF-8 byte column; 0-based byte offset",
		},
	}
	for _, cmd := range gonCommands() {
		result.Commands = append(result.Commands, cmd.path)
	}
	return result, nil
}

func (c *gonCapabilitiesResult) exitCode() int { return gonExitOK }

func (c *gonCapabilitiesResult) text(w io.Writer, r *gonRequest) {
	fmt.Fprintf(w, "gon tooling schema %d, %s\ntoolchain %s\n", c.SchemaVersion, c.Toolchain.Gonpls, c.Toolchain.Root)
	fmt.Fprintf(w, "queries: %s\nrename: %v, check: %v, explain: %v\n",
		strings.Join(c.Capabilities.Queries, ", "), c.Capabilities.Rename, c.Capabilities.Check, c.Capabilities.Explain)
}
