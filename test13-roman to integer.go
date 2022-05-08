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

	output := romanMap[s[0]]
	for i := 1; i < len(s); i++ {
		if romanMap[s[i-1]] < romanMap[s[i]] {
			output += romanMap[s[i]] - (2 * romanMap[s[i-1]])
		} else {
			output += romanMap[s[i]]
		}
	}
	return output
}
