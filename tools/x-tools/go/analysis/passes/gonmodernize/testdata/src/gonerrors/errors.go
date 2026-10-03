package gonerrors

import "fmt"

func read() (int, error) { return 7, nil }
func many() (int, string, error) { return 7, "ok", nil }
func flush() error { return nil }
func consume(...any) {}

func propagation() (int, error) {
	// want +1 "replace error check with Gon ! propagation"
	x, err := read()
	if err != nil {
		return 0, err
	}
	return x, nil
}

func multiple() (int, string, error) {
	// want +1 "replace error check with Gon ! propagation"
	x, s, err := many()
	if nil != err {
		return 0, "", err
	}
	return x, s, nil
}

func errorOnly() error {
	// want +1 "replace error check with Gon ! propagation"
	err := flush()
	if err != nil {
		return err
	}
	// want +1 "replace error check with Gon ! propagation"
	if problem := flush(); problem != nil {
		return problem
	}
	return nil
}

func wrapped() (int, error) {
	// want +1 "replace error check with a Gon or handler"
	x, err := read()
	if err != nil {
		return 0, fmt.Errorf("read: %w", err)
	}
	return x, nil
}

func sideEffects() (int, error) {
	// want +1 "replace error check with a Gon or handler"
	x, err := read()
	if err != nil {
		consume(err)
		return 0, err
	}
	return x, nil
}

func partialResult() (int, error) {
	x, err := read()
	if err != nil {
		return x, err
	}
	return x, nil
}

func laterError() (int, error) {
	x, err := read()
	if err != nil {
		return 0, err
	}
	return x, err
}

func reusedError() (int, error) {
	var err error
	x, err := read()
	if err != nil {
		return 0, err
	}
	return x, nil
}

func reusedValue() (x int, result error) {
	defer func() { consume(x) }()
	x, err := read()
	if err != nil {
		return 0, err
	}
	return x, nil
}

func namedResults() (x int, result error) {
	x = 123
	defer func() { consume(x, result) }()
	// want +1 "replace error check with Gon ! propagation"
	v, err := read()
	if err != nil {
		return 0, err
	}
	return v, nil
}

func bareReturn() (x int, result error) {
	// want +1 "replace error check with a Gon or handler"
	v, err := read()
	if err != nil {
		result = err
		return
	}
	return v, nil
}

func commented() (int, error) {
	// want +1 "replace error check with a Gon or handler"
	x, err := read()
	if err != nil {
		// Keep this context.
		return 0, err
	}
	return x, nil
}

func commentBetween() (int, error) {
	x, err := read()
	// Keep this context too.
	if err != nil {
		return 0, err
	}
	return x, nil
}

func elseBranch() (int, error) {
	x, err := read()
	if err != nil {
		return 0, err
	} else {
		consume(x)
	}
	return x, nil
}

func recoverError() (int, error) {
	x, err := read()
	if err != nil {
		consume(err)
	}
	return x, nil
}

func blank() error {
	// want +1 "replace error check with Gon ! propagation"
	_, err := read()
	if err != nil {
		return err
	}
	return nil
}

func interfaceResult() (any, error) {
	// want +1 "replace error check with a Gon or handler"
	x, err := read()
	if err != nil {
		return 0, err // boxing zero must not become a nil interface
	}
	return x, nil
}

func interfaceNil() (any, error) {
	// want +1 "replace error check with Gon ! propagation"
	x, err := read()
	if err != nil {
		return nil, err
	}
	return x, nil
}

func emptySlice() ([]int, error) {
	// want +1 "replace error check with a Gon or handler"
	x, err := read()
	if err != nil {
		return []int{}, err
	}
	return []int{x}, nil
}

func zeroStruct() (struct{ X int }, error) {
	// want +1 "replace error check with Gon ! propagation"
	x, err := read()
	if err != nil {
		return struct{ X int }{}, err
	}
	return struct{ X int }{x}, nil
}

func differentErrorResult() (int, any) {
	// want +1 "replace error check with a Gon or handler"
	x, err := read()
	if err != nil {
		return 0, err
	}
	return x, nil
}

func panicHandler() {
	// want +1 "replace error check with a Gon or handler"
	x, err := read()
	if err != nil {
		panic(err)
	}
	consume(x)
}

func unusedHandlerError() {
	// want +1 "replace error check with a Gon or handler"
	x, err := read()
	if err != nil {
		panic("failed")
	}
	consume(x)
}

func shadowedPanic(panic func(error)) {
	x, err := read()
	if err != nil {
		panic(err)
	}
	consume(x)
}

func labels() (int, error) {
	x, err := read()
	if err != nil {
		goto failed
	}
	return x, nil
failed:
	return 0, err
}

func breaks() error {
	for {
		err := flush()
		if err != nil {
			if true { break }
			return err
		}
		break
	}
	return nil
}

type ownError interface { Error() string }
func own() (int, ownError) { return 0, nil }

func namedErrorInterface() (int, error) {
	x, err := own()
	if err != nil {
		return 0, err
	}
	return x, nil
}

type aliasError = error
func aliased() (int, aliasError) { return 1, nil }

func errorAlias() (int, aliasError) {
	// want +1 "replace error check with Gon ! propagation"
	x, err := aliased()
	if err != nil {
		return 0, err
	}
	return x, nil
}

func closureBoundary() (int, error) {
	f := func() (int, error) {
		// want +1 "replace error check with Gon ! propagation"
		x, err := read()
		if err != nil {
			return 0, err
		}
		return x, nil
	}
	consume(f)
	return 0, nil
}

func capturesPartialResult() (int, error) {
	x, err := read()
	if err != nil {
		consume(func() int { return x })
		return 0, err
	}
	return x, nil
}

func capturesError() (int, error) {
	// want +1 "replace error check with a Gon or handler"
	x, err := read()
	if err != nil {
		consume(func() error { return err })
		return 0, err
	}
	return x, nil
}

func shadowedValue(x int) (int, error) {
	{
		// want +1 "replace error check with Gon ! propagation"
		x, err := func(i int) (int, error) { return i, nil }(x)
		if err != nil {
			return 0, err
		}
		return x, nil
	}
}

func shadowedNamedError() (value int, err error) {
	value = 42
	defer func() { consume(value, err) }()
	{
		// want +1 "replace error check with Gon ! propagation"
		v, err := read()
		if err != nil {
			return 0, err
		}
		return v, nil
	}
}

func shadowedNil(nil error) (int, error) {
	x, err := read()
	if err != nil {
		return 0, err
	}
	return x, nil
}

func errorConversion() error {
	err := error(nil)
	if err != nil {
		return err
	}
	return nil
}

func lambdaBoundary() {
	var f func() (int, error) = () => {
		// want +1 "replace error check with Gon ! propagation"
		x, err := read()
		if err != nil {
			return 0, err
		}
		return x, nil
	}
	consume(f)
}

func overlapping() (int, error) {
	// want +1 "replace error check with a Gon or handler"
	x, err := read()
	if err != nil {
		// This nested opportunity is deferred to the next fix invocation.
		if inner := flush(); inner != nil {
			return 0, inner
		}
		return 0, err
	}
	return x, nil
}

func errorOnlyFallthrough() {
	// want +1 "replace error check with a Gon or handler"
	if err := flush(); err != nil {
		consume(err)
	}
	consume("continued")
	// want +1 "replace error check with a Gon or handler"
	err := flush()
	if err != nil {
		defer func() { consume(err) }()
	}
	consume("continued again")
}

func errorOnlyEmptyHandler() {
	// want +1 "replace error check with a Gon or handler"
	if err := flush(); err != nil {
	}
}

func errorOnlyFallthroughLaterUse() {
	err := flush()
	if err != nil {
		consume(err)
	}
	consume(err)
}

func errorOnlyFallthroughReusedBinding() {
	var err error
	err = flush()
	if err != nil {
		consume(err)
	}
}

func errorOnlyFallthroughBreak() {
	for {
		if err := flush(); err != nil {
			break
		}
	}
}
