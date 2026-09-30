//go:build ppc64 || ppc64le

package cpu

const cacheLineSize = 128

func initOptions() {
	options = []option{
		{Name: "darn", Feature: &PPC64.HasDARN},
		{Name: "scv", Feature: &PPC64.HasSCV},
	}
}
