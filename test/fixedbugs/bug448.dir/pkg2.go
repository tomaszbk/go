// Issue 3843: inlining bug due to wrong receive operator precedence.

package pkg2

import "./pkg1"

func F() {
	pkg1.Do()
}

