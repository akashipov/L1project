package main

import (
	"fmt"
	"time"
)

func fun2(b chan int) {
	for {
		x, ok := <-b
		if !ok {
			break
		}
		fmt.Println(x)
	}
}

func fun1(a chan int, b chan int) {
	for {
		x, ok := <-a
		if !ok {
			break
		}
		b <- x * x
	}
}

func main() {
	a := make(chan int)
	b := make(chan int)
	go fun1(a, b)
	go fun2(b)
	nums := [...]int{123, 123, 1, 2, 3, 5, 6, 7, 78}
	for _, v := range nums {
		a <- v
	}
	time.Sleep(time.Second)
}
