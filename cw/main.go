package main

import "fmt"

func main() {
	dnaResult := DNAStrand("ATTGC")
	fmt.Printf("DNA Passed? %v\n", dnaResult == "TAACG")

	mumblingResult := Accum("abCd")
	fmt.Printf("Mumbling Passed? %v\n", mumblingResult == "A-Bb-Ccc-Dddd")

	positiveSumResult := PositiveSum([]int{-1, 2, 3, 4, -5})
	fmt.Printf("PositiveSum Passing? %v\n", positiveSumResult == 9)
}
