package main

import "math"

func maxF(a float64, b float64) float64 {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	return math.Max(a, b)
}
func minF(a float64, b float64) float64 {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	return math.Min(a, b)
}

func maxI(a int64, b int64) int64 {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	if a > b {
		return a
	}
	return b
}
func minI(a int64, b int64) int64 {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	if a < b {
		return a
	}
	return b
}
