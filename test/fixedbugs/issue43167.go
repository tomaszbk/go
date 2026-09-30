// errorcheck


package p

import . "bytes"

var _ Buffer // use package bytes

var Index byte // ERROR "Index redeclared.*\n\tLINE-4: previous declaration during import .bytes.|already declared|redefinition"
