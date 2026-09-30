package time

func ForceAndroidTzdataForTest() (undo func()) {
	allowGorootSource = false
	origLoadFromEmbeddedTZData := loadFromEmbeddedTZData
	loadFromEmbeddedTZData = nil

	return func() {
		allowGorootSource = true
		loadFromEmbeddedTZData = origLoadFromEmbeddedTZData
	}
}
