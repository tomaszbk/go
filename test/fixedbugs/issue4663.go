// errorcheck


// Issue 4663.
// Make sure 'not used' message is placed correctly.

package main

func a(b int) int64 {
  b // ERROR "not used"
  return 0
}
