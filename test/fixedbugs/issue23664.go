// errorcheck


// Verify error messages for incorrect if/switch headers.

package p

func f() {
	if f() true { // ERROR "unexpected name true, expected {"
	}

	switch f() true { // ERROR "unexpected name true, expected {"
	}
}
