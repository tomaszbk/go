// errorcheck


package main

type T U	// bogus "invalid recursive type T" from 6g
type U int

const x T = 123

type V V	// ERROR "invalid recursive type"


