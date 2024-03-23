package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

func producer(w *sync.WaitGroup, done chan struct{}, r chan int) {
	defer w.Done()
loop:
	for {
		select {
		case <-done:
			close(r)
			break loop
		default:
			r <- rand.Intn(10)
		}
	}
}

func worker(w *sync.WaitGroup, r chan int, name string, f *os.File) {
	defer f.Close()
	defer w.Done()
loop:
	for {
		x, ok := <-r
		if !ok {
			f.Write([]byte(fmt.Sprintf("%s has finished\n", name)))
			break loop
		}
		f.Write([]byte(fmt.Sprintf("%s: %d\n", name, x)))
	}
}

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	r := make(chan int)
	var w sync.WaitGroup
	w.Add(n + 2)
	for i := 0; i < n; i++ {
		filename := fmt.Sprintf("file%d.txt", i)
		f, err := os.Create(filename)
		if err != nil {
			close(r)
			fmt.Println(err.Error())
		}
		go worker(&w, r, fmt.Sprintf("Worker %d", i), f)
	}
	done := make(chan struct{}, 1)
	go producer(&w, done, r)
	go func() {
		defer w.Done()
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
		sig := <-sigint
		fmt.Printf("\nSignal: %v\n", sig)
		done <- struct{}{}
	}()
	w.Wait()
}
