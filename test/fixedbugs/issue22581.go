// errorcheck


package p

func f() {
	if i := g()); i == j { // ERROR "unexpected \)"
	}

	if i == g()] { // ERROR "unexpected \]"
	}

	switch i := g()); i { // ERROR "unexpected \)"
	}

	switch g()] { // ERROR "unexpected \]"
	}

	for i := g()); i < y; { // ERROR "unexpected \)"
	}

	for g()] { // ERROR "unexpected \]"
	}
}
