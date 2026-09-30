// errorcheck


package main

const A = complex(0()) // ERROR "cannot call .* not a function"
