package main

import (
	"fmt"
	"log"
)

func foo(a int) (x int) {
	defer func() {
		fmt.Println("foo exit")
	}()
	x = a
	return x
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()
	foo(5)
}
