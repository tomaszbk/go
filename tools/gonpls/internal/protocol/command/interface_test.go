package command_test

import (
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"golang.org/x/tools/gopls/internal/protocol/command"
	"golang.org/x/tools/gopls/internal/protocol/command/gen"
	"golang.org/x/tools/internal/testenv"
)

// TestGenerated ensures that we haven't forgotten to update command_gen.go.
func TestGenerated(t *testing.T) {
	testenv.NeedsGoPackages(t)
	testenv.NeedsLocalXTools(t)

	onDisk, err := os.ReadFile("command_gen.go")
	if err != nil {
		t.Fatal(err)
	}

	generated, err := gen.Generate()
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(string(generated), string(onDisk)); diff != "" {
		t.Errorf("command_gen.go is stale -- regenerate (-generated +on disk)\n%s", diff)
	}
}

// TestNamespace ensures that every command is in the gonpls.* namespace, so
// that the Gon and Go (gopls.*) extensions can be active in one VS Code window
// without registering the same command IDs.
func TestNamespace(t *testing.T) {
	if len(command.Commands) == 0 {
		t.Fatal("no commands")
	}
	for _, c := range command.Commands {
		if !strings.HasPrefix(string(c), "gonpls.") {
			t.Errorf("command %q is outside the gonpls.* namespace", c)
		}
	}
}
