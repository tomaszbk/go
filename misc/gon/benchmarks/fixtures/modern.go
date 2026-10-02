package main

import "fmt"

func propagate(seed int, e *effects) (int, error) {
	value := readAmount(seed, e)!
	value = validateAmount(value, e)!
	return value, nil
}

func handle(seed int, e *effects) (int, error) {
	value := readAmount(seed, e) or err {
		return 0, fmt.Errorf("read amount: %w", err)
	}
	return value + 7, nil
}

func conditional(active bool, value int, e *effects) int {
	return if active { trueValue(value, e) } else { falseValue(value, e) }
}

func capturedLambda(values []int, offset int) int {
	var transform func(int) int = (value) => value*2 + offset
	offset += 2 // The closure observes the updated captured variable.
	return sumMapped(values, transform)
}

func safeField(rule *policy, e *effects) int {
	return rule?.Factor ?? defaultNumber(e)
}

func safeCall(callback func(int) int, value int, e *effects) int {
	return callback?(argument(value, e)) ?? defaultNumber(e)
}

func coalesceAssign(slots []*policy, index int, e *effects) int {
	slots[location(index, e)] ??= defaultPolicy(e)
	return slots[index].Factor
}

func pipeline(item input, e *effects) (int, error) {
	value := parseAmount(item.Text, e) or err {
		return 0, fmt.Errorf("parse amount: %w", err)
	}
	rule := item.Rule ?? defaultPolicy(e)
	value = if item.Active { trueValue(value, e) } else { falseValue(value, e) }
	var adjust func(int) int = (value) => value * rule.Factor
	value = adjust(value)
	value = item.Callback?(argument(value, e)) ?? value
	value = validateAmount(value, e)!
	return value, nil
}
