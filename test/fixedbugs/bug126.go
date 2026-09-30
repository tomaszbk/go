// errorcheck


package main

// same const identifier declared twice should not be accepted
const none = 0  // GCCGO_ERROR "previous"
const none = 1;  // ERROR "redeclared|redef"
