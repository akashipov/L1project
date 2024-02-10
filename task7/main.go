package main

import (
	"fmt"
	"sync"
	"time"
)

// my own custom map with locker
type MapWithLock struct {
	M map[string]string
	L *sync.Mutex
}

func NewMapWithLock() *MapWithLock {
	l := sync.Mutex{}
	return &MapWithLock{
		M: map[string]string{},
		L: &l,
	}
}

func gorutine1(m *MapWithLock) {
	m.L.Lock()
	fmt.Printf("%p\n", m)
	m.L.Unlock()
	k := "key"
	for {
		m.L.Lock()
		m.M[k] = "g1"
		m.L.Unlock()
	}

}

func gorutine2(m *MapWithLock) {
	m.L.Lock()
	fmt.Printf("%p\n", m)
	m.L.Unlock()
	k := "key"
	for {
		m.L.Lock()
		m.M[k] = "g2"
		m.L.Unlock()
	}
}

// other solution via sync.Map
func gorutineMap1(m *sync.Map) {
	k := "key"
	for {
		m.Store(k, "g1")
	}

}

func gorutineMap2(m *sync.Map) {
	k := "key"
	for {
		m.Store(k, "g2")
	}
}

func main() {
	m := sync.Map{}
	go gorutineMap1(&m)
	time.Sleep(time.Second)
	go gorutineMap2(&m)
	time.Sleep(time.Second * 1)
	fmt.Println(m.Load("key"))
}
