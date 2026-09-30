// Export guts for testing.

package runtime

const SiginfoMaxSize = _si_max_size
const SigeventMaxSize = _sigev_max_size

var NewOSProc0 = newosproc0
var Mincore = mincore
var ParseRelease = parseRelease

type Siginfo siginfo
type Sigevent sigevent
