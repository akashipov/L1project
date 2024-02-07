package main

import (
	"fmt"
	"reflect"
)

func replace(a interface{}) {
	fmt.Println(reflect.TypeOf(a))
}

func main() {
	a := 5
	replace(a)
	b := float64(10)
	replace(b)
	c := "safadfad"
	replace(c)
}
