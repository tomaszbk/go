//go:build ppc64le || ppc64

package reflect

func archFloat32FromReg(reg uint64) float32
func archFloat32ToReg(val float32) uint64
