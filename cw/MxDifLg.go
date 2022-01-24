package main

import "math"

func MxDifLg(a1 []string, a2 []string) int {
	if len(a1) <= 0 || len(a2) <= 0 {
		return -1
	}

	minA1 := math.MaxInt
	var maxA1 int
	for _, s := range a1 {
		if len(s) < minA1 {
			minA1 = len(s)
		}

		if len(s) > maxA1 {
			maxA1 = len(s)
		}
	}

	minA2 := math.MaxInt
	var maxA2 int
	for _, s := range a2 {
		if len(s) < minA2 {
			minA2 = len(s)
		}

		if len(s) > maxA2 {
			maxA2 = len(s)
		}
	}

	return int(math.Max(
		math.Abs(float64(minA1)-float64(maxA2)),
		math.Abs(float64(maxA1)-float64(minA2)),
	),
	)
}
