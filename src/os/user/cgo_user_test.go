//go:build cgo && !osusergo

package user

func init() {
	hasCgo = true
}
