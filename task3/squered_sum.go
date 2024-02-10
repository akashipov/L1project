package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func composer(w *sync.WaitGroup, r chan int) {
	s := 0
	for {
		x, ok := <-r
		if !ok {
			fmt.Println(s)
			w.Done()
			break
		}
		s += x
	}
}

func worker(w *sync.WaitGroup, i chan int, r chan int) {
	s := 0
	for {
		x, ok := <-i
		if ok {
			s += x * x
		} else {
			r <- s
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
	r := make(chan int, n)
	for i := 0; i < n; i++ {
		h := make(chan int, 1)
		w.Add(1)
		go worker(&w, h, r)
		m[i] = h
	}
	go composer(&w, r)
	nums := [...]int{1, 2, 3, 5, 1, 2, 3, 5, 1, 2, 3, 5, 1, 2, 3, 5, 1, 2, 3, 5, 1, 2, 3, 5, 1, 2, 3, 5}
	for idx, v := range nums {
		b := idx % n
		m[b] <- v
	}
	for _, v := range m {
		close(v)
	}
	w.Wait()
	w.Add(1)
	close(r)
	w.Wait()
}
