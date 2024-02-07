package main

import (
	"strconv"
	"strings"
)

type A struct {
	Values []int
}

func (a *A) GetJson() string {
	b := strings.Builder{}
	b.WriteString("[")
	for idx, v := range a.Values {
		b.WriteString(strconv.Itoa(v))
		if idx != len(a.Values)-1 {
			b.WriteString(", ")
		}
	}
	b.WriteString("]")
	return b.String()
}

func main() {

}
