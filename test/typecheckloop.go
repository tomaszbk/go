// errorcheck


// Verify that constant definition loops are caught during
// typechecking and that the errors print correctly.

package main

const A = 1 + B // ERROR "constant definition loop\n.*A uses B\n.*B uses C\n.*C uses A|initialization cycle"
const B = C - 1 // ERROR "constant definition loop\n.*B uses C\n.*C uses B|initialization cycle"
const C = A + B + 1
