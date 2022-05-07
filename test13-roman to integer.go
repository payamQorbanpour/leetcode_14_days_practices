package main

import (
	"fmt"
)

func main() {
	// s := "III"
	// s := "IV"
	// s := "XL"
	// s := "LVIII"
	s := "MCMXCIV"
	// s := "DCXXI"

	fmt.Println(romanToInt(s))
}

func romanToInt(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	output := 0
	for i, _ := range s {
		integer, _ := romanMap[s[i]]
		if i != 0 && (s[i] < s[i-1]) && (s[i] == 'V' || s[i] == 'X' || s[i] == 'L' || s[i] == 'C' || s[i] == 'D' || s[i] == 'M') {
			if s[i-1] == 'I' || s[i-1] == 'X' || s[i-1] == 'C' {
				integer = romanMap[s[i]] - (romanMap[s[i-1]] + romanMap[s[i-1]])
			}
		}
		output += integer
	}
	return output
}
