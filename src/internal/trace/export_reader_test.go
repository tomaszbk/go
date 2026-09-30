package trace

import "internal/trace/version"

// GoVersion is the version set in the trace header.
func (r *Reader) GoVersion() version.Version {
	return r.version
}
