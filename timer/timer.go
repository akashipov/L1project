package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

func producer(w *sync.WaitGroup, ctx context.Context, r chan int) {
	defer w.Done()
loop:
	for {
		select {
		case <-ctx.Done():
			close(r)
			break loop
		default:
			r <- rand.Intn(10)
		}
	}
	fmt.Println("Producer is finished")
}

func consumer(w *sync.WaitGroup, r chan int) {
	defer w.Done()
loop:
	for {
		select {
		default:
			x, ok := <-r
			if !ok {
				break loop
			}
			fmt.Println(x)
		}
	}
	fmt.Println("Consumer is finished")
}

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}
	ctx := context.Background()
	ctx, close := context.WithTimeout(ctx, time.Second*time.Duration(n))
	defer close()
	r := make(chan int)
	var w sync.WaitGroup
	t := time.Now()
	w.Add(2)
	go producer(&w, ctx, r)
	go consumer(&w, r)
	w.Wait()
	fmt.Println(time.Since(t))
}
