package main

import (
	"fmt"

	"github.com/akashipov/L1project/pointer/pointer"
)

func main() {
	a := pointer.NewPointer(0.0, 0.0)
	b := pointer.NewPointer(4.0, 3.0)
	fmt.Println(b.Distance(a))
	fmt.Println(a.Distance(b))
}
