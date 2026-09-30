package time

func ForceAusFromTZIForTesting() {
	ResetLocalOnceForTest()
	localOnce.Do(func() { initLocalFromTZI(&aus) })
}

func ForceUSPacificFromTZIForTesting() {
	ResetLocalOnceForTest()
	localOnce.Do(func() { initLocalFromTZI(&usPacific) })
}

func ToEnglishName(stdname, dstname string) (string, error) {
	return toEnglishName(stdname, dstname)
}
