// errorcheck


package p

var _ = true == '\\' // ERROR "invalid operation: (cannot compare true)|(true) == '\\\\' \(mismatched types untyped bool and untyped rune\)"
var _ = true == '\'' // ERROR "invalid operation: (cannot compare true)|(true) == '\\'' \(mismatched types untyped bool and untyped rune\)"
var _ = true == '\n' // ERROR "invalid operation: (cannot compare true)|(true) == '\\n' \(mismatched types untyped bool and untyped rune\)"
