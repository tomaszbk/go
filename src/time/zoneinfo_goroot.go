//go:build !ios && !android

package time

func gorootZoneSource(goroot string) (string, bool) {
	if goroot == "" {
		return "", false
	}
	return goroot + "/lib/time/zoneinfo.zip", true
}
