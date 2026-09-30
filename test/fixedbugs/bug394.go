// errorcheck


// Issue 2598
package foo

return nil // ERROR "non-declaration statement outside function body|expected declaration"
