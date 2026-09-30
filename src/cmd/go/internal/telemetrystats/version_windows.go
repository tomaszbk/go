//go:build !cmd_go_bootstrap && windows

package telemetrystats

import (
	"fmt"
	"internal/syscall/windows"

	"cmd/internal/telemetry/counter"
)

func incrementVersionCounters() {
	major, minor, build := windows.Version()
	counter.Inc(fmt.Sprintf("go/platform/host/windows/major-version:%d", major))
	counter.Inc(fmt.Sprintf("go/platform/host/windows/version:%d-%d", major, minor))
	counter.Inc(fmt.Sprintf("go/platform/host/windows/build:%d", build))
}
