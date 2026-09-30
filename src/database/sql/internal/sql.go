// Package internal contains internal symbols shared between
// database/sql and database/sql/driver.
package internal

// ScanContext is database/sql/driver.ScanContext.
// We define it here so driver.ScanContext can be opaque to users but
// visible to database/sql.
type ScanContext struct {
	v any
}

func NewScanContext(v any) ScanContext {
	return ScanContext{v}
}

func ScanContextValue(c ScanContext) any {
	return c.v
}
