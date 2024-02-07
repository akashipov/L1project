package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func Revert(s string) (string, error) {
	a := strings.Builder{}
	n := make([]rune, utf8.RuneCountInString(s))
	for _, v := range s {
		n = append(n, v)
	}
	for i := len(n) - 1; i >= 0; i-- {
		a.WriteRune(n[i])
	}
	return a.String(), nil
}

func main() {
	a := "adsfds🦍"
	fmt.Println(a, utf8.RuneCountInString(a))
	b, _ := Revert(a)
	fmt.Println(b, utf8.RuneCountInString(a))
}
