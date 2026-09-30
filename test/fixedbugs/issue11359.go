// errorcheck


// identifiers beginning with non-ASCII digits were incorrectly accepted.
// issue 11359.

package p
var ۶ = 0 // ERROR "identifier cannot begin with digit"
