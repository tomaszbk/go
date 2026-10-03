package lambdaimports

import (
	"lambdahelper"
	"time"
)

// These are time's only references. Neither signature can be removed, even
// though the other signature would keep the import used for an individual fix.
func callbacks() {
	lambdahelper.Consume(func(t time.Duration) {})
	lambdahelper.Consume(func(t time.Duration) {})
}
