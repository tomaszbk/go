package main

import "fmt"

func propagate(seed int, e *effects) (int, error) {
	value, err := readAmount(seed, e)
	if err != nil {
		return 0, err
	}
	value, err = validateAmount(value, e)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func handle(seed int, e *effects) (int, error) {
	value, err := readAmount(seed, e)
	if err != nil {
		return 0, fmt.Errorf("read amount: %w", err)
	}
	return value + 7, nil
}

func conditional(active bool, value int, e *effects) int {
	if active {
		return trueValue(value, e)
	}
	return falseValue(value, e)
}

func capturedLambda(values []int, offset int) int {
	var transform func(int) int = func(value int) int { return value*2 + offset }
	offset += 2 // The closure observes the updated captured variable.
	return sumMapped(values, transform)
}

func safeField(rule *policy, e *effects) int {
	if rule == nil {
		return defaultNumber(e)
	}
	return rule.Factor
}

func safeCall(callback func(int) int, value int, e *effects) int {
	if callback == nil {
		return defaultNumber(e)
	}
	return callback(argument(value, e))
}

func coalesceAssign(slots []*policy, index int, e *effects) int {
	index = location(index, e)
	if slots[index] == nil {
		slots[index] = defaultPolicy(e)
	}
	return slots[index].Factor
}

func pipeline(item input, e *effects) (int, error) {
	value, err := parseAmount(item.Text, e)
	if err != nil {
		return 0, fmt.Errorf("parse amount: %w", err)
	}
	rule := item.Rule
	if rule == nil {
		rule = defaultPolicy(e)
	}
	if item.Active {
		value = trueValue(value, e)
	} else {
		value = falseValue(value, e)
	}
	var adjust func(int) int = func(value int) int { return value * rule.Factor }
	value = adjust(value)
	if item.Callback != nil {
		value = item.Callback(argument(value, e))
	}
	value, err = validateAmount(value, e)
	if err != nil {
		return 0, err
	}
	return value, nil
}
