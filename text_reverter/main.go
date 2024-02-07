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

func TextRevert(s string) (string, error) {
	listS := strings.Split(s, " ")
	b := strings.Builder{}
	for _, v := range listS {
		v, err := Revert(v)
		if err != nil {
			return "", err
		}
		b.WriteString(v)
		b.WriteString(" ")
	}
	return b.String(), nil
}

func main() {
	a := "adsfds🦍 artem art🦍sevil"

	fmt.Println(a, utf8.RuneCountInString(a))
	b, _ := TextRevert(a)
	fmt.Println(b, utf8.RuneCountInString(a))
}
