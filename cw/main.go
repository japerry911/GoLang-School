package main

import "fmt"

func main() {
	dnaResult := DNAStrand("ATTGC")
	fmt.Printf("DNA Passed? %v\n", dnaResult == "TAACG")
}
