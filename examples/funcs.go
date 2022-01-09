package main

import "fmt"

// func add(x int, y int) int {
// 	return x + y
// }

// Same type, only do last
func add (x, y int) int {
	return x + y
}

// return 2 returns
func swap(x, y string) (string, string) {
	return y, x
}

// named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// variable declaration
var c, python, java bool

// variable with initializers, doesn't need type
var i, j int = 1, 2

func main() {
	fmt.Println(add(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))

	var i1 int
	fmt.Println(i1, c, python, java)

	var c2, python2, java2 = true, false, "no!"
	fmt.Println(i, j, c2, python2, java2)

	// short variable declaration, can only be used inside func
	k := 3
	fmt.Println(k)
}
