package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func worker(w *sync.WaitGroup, i chan int) {
	for {
		x, ok := <-i
		if ok {
			fmt.Println(x * x)
		} else {
			w.Done()
			break
		}
	}
}

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	m := make(map[int]chan int, n)
	var w sync.WaitGroup
	for i := 0; i < n; i++ {
		h := make(chan int, 1)
		w.Add(1)
		go worker(&w, h)
		m[i] = h
	}
	nums := [...]int{1, 2, 3, 4}
	for idx, v := range nums {
		b := idx % n
		m[b] <- v
	}
	time.Sleep(time.Second * 10)
	for _, v := range m {
		close(v)
	}
	w.Wait()
}
