// errorcheck


// Issue 1411.

package main

const ui uint = 0
const i int = ui // ERROR "type"
