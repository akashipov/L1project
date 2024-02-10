package main

import (
	"fmt"
	"strings"
)

func unique(s string) bool {
	m := make(map[rune]rune)
	s = strings.ToLower(s)
	for _, v := range s {
		_, ok := m[v]
		if ok {
			return false
		} else {
			m[v] = v
		}
	}
	return true
}

func main() {
	s := "asdasdffcsx"
	fmt.Println(unique(s))
	s = "abc"
	fmt.Println(unique(s))
	s = "abcbf FAs"
	fmt.Println(unique(s))
	s = "abcf s"
	fmt.Println(unique(s))
	s = "abcf As"
	fmt.Println(unique(s))

}
