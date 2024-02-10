// Также я мог здесь показать пример с Mutext, RWMutex, Cond и так далее
// много вариаций с ctx и использовать существующие конструкции на каналы
// но я не вижу в этом смысла... я показал основные категории, которые я выделяю,
// на мой взгляд
package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var IsRun atomic.Bool

func AtomicGlobal() {
	for {
		fmt.Println("I am working AtomicGlobal ...")
		time.Sleep(time.Second)
		if !IsRun.Load() {
			fmt.Println("Gorutine AtomicGlobal is ended")
			break
		}
	}
}

func G() {
	IsRun.Store(true)
	go AtomicGlobal()
	time.Sleep(time.Second * 3)
	IsRun.Store(false)
	time.Sleep(time.Second * 3)
}

func CtxGorutine(ctx context.Context) {
loop:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Gorutine CtxGorutine is finished")
			break loop
		default:
			fmt.Println("I am working CtxGorutine ...")
			time.Sleep(time.Second)
		}
	}
}

func CtxGo() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	go CtxGorutine(ctx)
	time.Sleep(time.Second * 5)
}

func ChanGorutine(done chan struct{}) {
loop:
	for {
		select {
		case <-done:
			fmt.Println("Gorutine ChanGorutine is finished")
			break loop
		default:
			fmt.Println("I am working ChanGorutine ...")
			time.Sleep(time.Second)
		}
	}
}

func ChanGo() {
	done := make(chan struct{})
	go ChanGorutine(done)
	time.Sleep(time.Second * 3)
	done <- struct{}{}
	time.Sleep(time.Second * 2)
}

func f() {
	for {
		time.Sleep(time.Second)
		fmt.Println("I am working F...")
	}
}

func main() {
	var w sync.WaitGroup
	w.Add(1)
	go func() {
		defer func() {
			fmt.Println("It was finished by goexit")
			w.Done()
		}()
		go f()
		runtime.Goexit()
		fmt.Println("We are here")
	}()
	w.Wait()
	// G()
	// CtxGo()
	ChanGo()
}
