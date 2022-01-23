package main

import "log"

func DNAStrand(dna string) string {
	var convertedDnaStrand string

	for _, c := range dna {
		convertedDnaStrand += convertLetter(c)
	}

	return convertedDnaStrand
}

func convertLetter(c rune) string {
	var convertedLetter string

	switch c {
	case 'A':
		convertedLetter = "T"
	case 'T':
		convertedLetter = "A"
	case 'C':
		convertedLetter = "G"
	case 'G':
		convertedLetter = "C"
	default:
		log.Fatalf("Invalid DNA character: %v", c)
	}

	return convertedLetter
}
