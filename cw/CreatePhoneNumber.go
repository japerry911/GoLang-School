package main

import "strconv"

func CreatePhoneNumber(numbers [10]uint) string {
	pn := "("

	for i, n := range numbers {
		pn += strconv.Itoa(int(n))

		if i == 2 {
			pn += ") "
		} else if i == 5 {
			pn += "-"
		}
	}

	return pn
}
