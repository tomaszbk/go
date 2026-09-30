package telemetry

import "golang.org/x/telemetry/internal/telemetry"

// Dir returns the telemetry directory.
func Dir() string {
	return telemetry.Default.Dir()
}
