// errorcheck


// Offending character % must not be interpreted as
// start of format verb when emitting error message.

package% // ERROR "unexpected %|package name must be an identifier|after package clause|expected declaration"
