// errorcheck

// Issue 11362: prints empty canonical import path

package main

import _ "unicode//utf8" // GC_ERROR "non-canonical import path .unicode//utf8. \(should be .unicode/utf8.\)"

func main() {
}
