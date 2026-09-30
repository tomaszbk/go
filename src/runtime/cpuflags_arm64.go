package runtime

import (
	"internal/cpu"
)

var arm64UseAlignedLoads bool

func init() {
	if cpu.ARM64.IsNeoverse {
		arm64UseAlignedLoads = true
	}
}
