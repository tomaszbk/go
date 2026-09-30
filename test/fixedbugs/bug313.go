// errorcheckdir


// Issue 1284

package bug313

/*
6g bug313.dir/[ab].go

Before:
bug313.dir/b.go:7: internal compiler error: fault

Now:
bug313.dir/a.go:10: undefined: fmt.DoesNotExist
*/
