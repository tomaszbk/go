// errorcheck


// Issue 2500

package foo

// Check that we only get root cause message, no further complaints about r undefined
func (r *indexWriter) foo() {}  // ERROR "undefined.*indexWriter"
