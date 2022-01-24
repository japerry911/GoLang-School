package main

func Divisors(n int) int {
	d := 1

	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			d++
		}
	}

	return d
}

//func Divisors(n int) int {
//	var d int
//
//	for i := 1; i <= n; i++ {
//		if n%i == 0 {
//			d++
//		}
//	}
//
//	return d
//}
