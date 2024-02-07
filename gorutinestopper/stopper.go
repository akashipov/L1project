package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func f() {
	for {
		time.Sleep(time.Second)
		fmt.Println("I am working...")
	}
}

func main() {
	var goid int
	var w sync.WaitGroup
	w.Add(1)
	go func() {
		defer func() {
			fmt.Println("It was finished by goexit")
			w.Done()
		}()
		go f()
		// runtime.Goexit()
		fmt.Println("We are here")
	}()
	goid = runtime.NumGoroutine()
	fmt.Println(goid)
	w.Wait()
}
