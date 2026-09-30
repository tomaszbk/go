//go:build !windows

package os

func rootCleanPath(s string, prefix, suffix []string) (string, error) {
	return s, nil
}
