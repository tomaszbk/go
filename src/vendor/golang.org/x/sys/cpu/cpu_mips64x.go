//go:build mips64 || mips64le

package cpu

const cacheLineSize = 32

func initOptions() {
	options = []option{
		{Name: "msa", Feature: &MIPS64X.HasMSA},
	}
}
