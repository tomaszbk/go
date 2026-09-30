package sysinfo_test

import (
	. "internal/sysinfo"
	"testing"
)

func TestCPUName(t *testing.T) {
	t.Logf("CPUName: %s", CPUName())
	t.Logf("osCPUInfoName: %s", XosCPUInfoName())
}
