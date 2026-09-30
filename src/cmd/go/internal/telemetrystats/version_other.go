//go:build !cmd_go_bootstrap && !unix && !windows

package telemetrystats

import "cmd/internal/telemetry/counter"

func incrementVersionCounters() {
	counter.Inc("go/platform:version-not-supported")
}
