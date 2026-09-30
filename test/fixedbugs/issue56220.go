// errorcheck


package p

func f() int {
	return int(1 - .0000001) // ERROR "cannot convert 1 - \.0000001 \(untyped float constant 0\.9999999\) to type int"
}

func g() int64 {
	return int64((float64(0.03) - float64(0.02)) * 1_000_000) // ERROR "cannot convert \(float64\(0\.03\) - float64\(0\.02\)\) \* 1_000_000 \(constant 9999\.999999999998 of type float64\) to type int64"
}
