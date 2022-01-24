package main

import "fmt"

func main() {
	dnaResult := DNAStrand("ATTGC")
	fmt.Printf("DNA Passed? %v\n", dnaResult == "TAACG")

	mumblingResult := Accum("abCd")
	fmt.Printf("Mumbling Passed? %v\n", mumblingResult == "A-Bb-Ccc-Dddd")

	positiveSumResult := PositiveSum([]int{-1, 2, 3, 4, -5})
	fmt.Printf("PositiveSum Passing? %v\n", positiveSumResult == 9)

	a1 := []string{
		"hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"}
	a2 := []string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"}
	maxLengthDifferenceResult := MxDifLg(a1, a2)
	fmt.Printf("MaxLengthDifference Passing? %v\n", maxLengthDifferenceResult == 13)
	fmt.Println(maxLengthDifferenceResult)
}
