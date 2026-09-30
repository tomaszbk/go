// errorcheck

// Issue 4847: initialization cycle is not detected.

package p

type (
	E int
	S int
)

type matcher func(s *S) E

func matchList(s *S) E { return matcher(matchAnyFn)(s) }

var foo = matcher(matchList)

var matchAny = matcher(matchList) // ERROR "initialization cycle|depends upon itself"

func matchAnyFn(s *S) (err E) { return matchAny(s) }
