// Package label provides common declarations used by both the [runtime] and [runtime/pprof] packages.
// The [Set] type is used for goroutine labels, and is duplicated as
// [runtime/pprof.LabelSet]. The type is duplicated due to go.dev/issue/65437
// preventing the use of a type-alias in an existing public interface.
package label

// Label is a key/value pair of strings.
type Label struct {
	Key   string
	Value string
}

// Set is a set of labels.
type Set struct {
	List []Label
}

// NewSet constructs a LabelSet that wraps the provided labels.
func NewSet(list []Label) Set {
	return Set{List: list}
}
