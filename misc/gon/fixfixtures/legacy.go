package main

import (
	"errors"
	"fmt"
)

var failure = errors.New("failure")
var events string

func read(fail bool) (int, error) {
	events += "read;"
	if fail {
		return 7, failure
	}
	return 21, nil
}

func propagate(fail bool) (int, error) {
	value, err := read(fail)
	if err != nil {
		return 0, err
	}
	return value * 2, nil
}

func wrapped(fail bool) (int, error) {
	value, err := read(fail)
	if err != nil {
		events += "wrap;"
		return 0, fmt.Errorf("wrapped: %w", err)
	}
	return value, nil
}

func flush(fail bool) error {
	events += "flush;"
	if fail {
		return failure
	}
	return nil
}

func errorOnly(fail bool) error {
	if err := flush(fail); err != nil {
		return err
	}
	return nil
}

func localHandler(fail bool) {
	if err := flush(fail); err != nil {
		events += err.Error() + ";"
	}
	events += "continue;"
}

func named(fail bool) (result int, problem error) {
	result = 99
	defer func() { events += fmt.Sprintf("defer:%d:%v;", result, problem) }()
	value, err := read(fail)
	if err != nil {
		return 0, err
	}
	return value, nil
}

// These results are useful even on failure, so this must stay explicit.
func partial(fail bool) (int, error) {
	value, err := read(fail)
	if err != nil {
		return value, err
	}
	return value, nil
}

// Keeping err alive after its guard prevents removing the binding.
func reused(fail bool) (int, error) {
	value, err := read(fail)
	if err != nil {
		return 0, err
	}
	events += fmt.Sprintf("later:%v;", err)
	return value, nil
}

func touch(label string, value int) int {
	events += label + ";"
	return value
}

func choose(flag bool) int {
	if flag {
		return touch("yes", 1)
	} else {
		return touch("no", 2)
	}
}

func assign(flag bool) int {
	var result int
	if flag {
		result = touch("yes", 1)
	} else {
		result = touch("no", 2)
	}
	return result
}

type record struct{ Value int }

func defaults() *record {
	events += "default;"
	return &record{Value: 8}
}

func initialize(p *record) *record {
	if p == nil {
		p = defaults()
	}
	return p
}

func coalesce(p *record) *record {
	if p != nil {
		return p
	} else {
		return defaults()
	}
}

func field(p *record) int {
	if p != nil {
		return p.Value
	} else {
		return touch("field-default", 9)
	}
}

func callback(f func(int) int) int {
	if f != nil {
		return f(touch("arg", 3))
	} else {
		return touch("call-default", 4)
	}
}

func interfaceDefault(p any) any {
	if p == nil {
		p = defaults()
	}
	return p
}

func boxed(p *record, fallback any) any {
	if p != nil {
		return p
	} else {
		return fallback
	}
}

func boxedConstant(p *record) any {
	if p != nil {
		return p
	} else {
		return 1
	}
}

func mixed(flag bool) any {
	if flag {
		return int8(1)
	} else {
		return 0
	}
}

func differentConstants(flag bool) any {
	if flag {
		return 1
	} else {
		return "other"
	}
}

func apply(f func(int) int, x int) int { return f(x) }

func closures() int {
	n := 5
	var f func(int) int = func(x int) int {
		defer func() { events += "lambda-defer;" }()
		return n + x
	}
	n = 7
	return f(1) + apply(func(x int) int { return x * 2 }, 3)
}

func expect(ok bool) {
	if !ok {
		panic("unexpected result or evaluation order: " + events)
	}
}

func main() {
	for _, fail := range []bool{false, true} {
		events = ""
		v, err := propagate(fail)
		expect(events == "read;" && ((fail && v == 0 && err == failure) || (!fail && v == 42 && err == nil)))
		events = ""
		v, err = wrapped(fail)
		expect((fail && v == 0 && errors.Is(err, failure) && err.Error() == "wrapped: failure" && events == "read;wrap;") || (!fail && v == 21 && err == nil && events == "read;"))
		events = ""
		err = errorOnly(fail)
		expect(events == "flush;" && ((fail && err == failure) || (!fail && err == nil)))
		events = ""
		localHandler(fail)
		expect((fail && events == "flush;failure;continue;") || (!fail && events == "flush;continue;"))
		events = ""
		v, err = named(fail)
		expect((fail && v == 0 && err == failure && events == "read;defer:0:failure;") || (!fail && v == 21 && err == nil && events == "read;defer:21:<nil>;"))
		events = ""
		v, err = partial(fail)
		expect(events == "read;" && ((fail && v == 7 && err == failure) || (!fail && v == 21 && err == nil)))
		events = ""
		v, err = reused(fail)
		expect((fail && v == 0 && err == failure && events == "read;") || (!fail && v == 21 && err == nil && events == "read;later:<nil>;"))
	}
	for _, selectValue := range []func(bool) int{choose, assign} {
		events = ""
		expect(selectValue(true) == 1 && events == "yes;")
		events = ""
		expect(selectValue(false) == 2 && events == "no;")
	}
	p := &record{Value: 0}
	for _, selectPointer := range []func(*record) *record{initialize, coalesce} {
		events = ""
		expect(selectPointer(p) == p && events == "")
		events = ""
		expect(selectPointer(nil).Value == 8 && events == "default;")
	}
	events = ""
	expect(field(p) == 0 && events == "")
	expect(field(nil) == 9 && events == "field-default;")
	events = ""
	expect(callback(nil) == 4 && events == "call-default;")
	events = ""
	expect(callback(func(x int) int { events += "call;"; return x * 2 }) == 6 && events == "arg;call;")
	var typedNil *record
	events = ""
	expect(interfaceDefault(typedNil) == typedNil && events == "")
	expect(interfaceDefault(nil).(*record).Value == 8 && events == "default;")
	expect(boxed(nil, nil) == nil)
	expect(boxed(nil, "missing") == "missing")
	expect(boxed(p, "missing") == p)
	expect(boxedConstant(nil).(int) == 1 && boxedConstant(p) == p)
	expect(mixed(true).(int8) == 1 && mixed(false).(int) == 0)
	expect(differentConstants(true).(int) == 1 && differentConstants(false).(string) == "other")
	events = ""
	expect(closures() == 14 && events == "lambda-defer;")
	fmt.Println("PASS: errors, partial results, named returns, lazy branches, nil interfaces, closures")
}
