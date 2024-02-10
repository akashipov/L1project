package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	i, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}
	var a int64 = 5
	var b int64 = 1
	b <<= i

	if os.Args[2] == "1" {
		a |= b
	} else {
		a &^= b
	}

	fmt.Println(a)
}
