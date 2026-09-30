package time

// in wasip1 zoneinfo is managed by the runtime.
var platformZoneSources = []string{}

func initLocal() {
	localLoc.name = "Local"
}
