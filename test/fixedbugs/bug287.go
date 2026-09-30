// errorcheck


// Used to die dividing by zero; issue 879.

package main

var mult [3][...]byte = [3][5]byte{}	// ERROR "\.\.\."
