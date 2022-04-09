package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	output := ""
    words := strings.Fields(s)
	for i, _ := range(words) {
		l := len(words[i]) - 1
		
		wordInByte := []byte(words[i])
		for j := 0; j < len(words[i]) / 2; j++ {
			wordInByte[j], wordInByte[l] =wordInByte[l], wordInByte[j]
			l--
		}
		output += string(wordInByte)

		if i != len(words) - 1 {
			output += " "
		}
	}
	return output
}

func main() {
	s := "Let's take LeetCode contest"

	fmt.Println(reverseWords(s))
}