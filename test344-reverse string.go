package main

import "fmt"

func reverseString(s []byte)  {
	j := len(s) - 1
	for i := 0; i < len(s) / 2; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
}

func main() {
	data := "hello"
	s := []byte(data)

	reverseString(s)
}