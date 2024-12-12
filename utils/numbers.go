package utils

import "strconv"

func Atoi(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		panic(s + " is not a number")
	}
	return result
}

func PowInts(x, n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	y := PowInts(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return x * y * y
}
