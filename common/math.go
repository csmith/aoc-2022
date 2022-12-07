package common

import "math"

// Abs returns the absolute value of x.
func Abs(x int64) int64 {
	y := x >> 63
	return (x ^ y) - y
}

// Max returns the highest of the two given ints.
func Max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

// Min returns the lowest of the two given ints.
func Min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

// GCD finds the greatest common divisor (GCD) via Euclidean algorithm
// Source: https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd
func GCD(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM finds the Least Common Multiple (LCM) via GCD
// Source: https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd
func LCM(a, b int64, integers ...int64) int64 {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

// Range finds the minimum and maximum of the given ints
func Range(ints []int) (min, max int) {
	min = math.MaxInt64
	max = math.MinInt64
	for _, i := range ints {
		min = Min(min, i)
		max = Max(max, i)
	}
	return
}
