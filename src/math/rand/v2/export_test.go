package rand

func GetNormalDistributionParameters() (float64, [128]uint32, [128]float32, [128]float32) {
	return rn, kn, wn, fn
}

func GetExponentialDistributionParameters() (float64, [256]uint32, [256]float32, [256]float32) {
	return re, ke, we, fe
}
