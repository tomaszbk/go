// errorcheck

// issue 5609: overflow when calculating array size

package pkg

const Large uint64 = 18446744073709551615

var foo [Large]uint64 // ERROR "array bound is too large|array bound overflows|invalid array length"
