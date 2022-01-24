package main

func PositiveSum(numbers []int) int {
	var t int

	for _, n := range numbers {
		if n > 0 {
			t += n
		}
	}

	return t
}
