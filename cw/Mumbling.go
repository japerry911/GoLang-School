package main

import (
	"strings"
)

func Accum(s string) string {
	rs := make([]string, len(s))

	for i := 0; i < len(s); i++ {
		rs[i] = strings.ToUpper(string(s[i])) + strings.Repeat(strings.ToLower(string(s[i])), i)
	}

	return strings.Join(rs, "-")
}

//func Accum(s string) string {
//	var rs string
//
//	for i, c := range s {
//		if i > 0 {
//			rs += "-" + strings.ToUpper(string(c)) + strings.Repeat(strings.ToLower(string(c)), i)
//		} else {
//			rs += strings.ToUpper(string(c))
//		}
//	}
//
//	return rs
//}

//func Accum(s string) string {
//	var rs string
//
//	for i, c := range s {
//		for i2 := 0; i2 < i+1; i2++ {
//			if i2 == 0 && i != 0 {
//				rs += "-"
//			}
//
//			if i2 == 0 {
//				rs += strings.ToUpper(string(c))
//			} else {
//				rs += strings.ToLower(string(c))
//			}
//		}
//	}
//
//	return rs
//}
