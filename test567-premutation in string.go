package main

import (
	"fmt"
	"reflect"
)

func main() {
	s1 := "adc"
	s2 := "dcda"

	fmt.Println(checkInclusion(s1, s2))
}

func checkInclusion(s1 string, s2 string) bool {
	m1 := make(map[byte]int)

	for i := 0; i < len(s1); i++ {
		m1[s1[i]] += 1
	}

	for l, r := 0, len(s1); r <= len(s2); r++ {
		window := s2[l:r]
		m2 := make(map[byte]int)

		for j, _ := range window {
			m2[window[j]] += 1
		}

		if reflect.DeepEqual(m1, m2) {
			return true
		} else {
			l++
		}
	}
	return false
}
