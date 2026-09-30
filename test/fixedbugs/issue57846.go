// compile


package main

func Float64D3(list [][][]float64, value float64) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Float64D2(listValue, value)
	}
	return valueCount
}

func Float64(list []float64, value float64) int {
	valueCount := 0
	for _, listValue := range list {
		if listValue == value {
			valueCount++
		}
	}
	return valueCount
}

func Float64D2(list [][]float64, value float64) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Float64(listValue, value)
	}
	return valueCount
}
