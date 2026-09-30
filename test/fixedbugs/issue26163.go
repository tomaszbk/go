// compile -N -d=softfloat


// Issue 26163: dead store generated in late opt messes
// up store chain calculation.

package p

var i int
var A = ([]*int{})[i]

var F func(float64, complex128) int
var C chan complex128
var B = F(1, 1+(<-C))
