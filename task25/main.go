package main

import (
	"fmt"
	"time"
)

func sleep(t time.Duration) {
	now := time.Now()
	for {
		if time.Since(now) > t {
			fmt.Println("Since V:", time.Since(now))
			break
		}
	}
}

func sleepAfterVersion(t time.Duration) {
	now := time.Now()
	n := time.After(t)
	<-n
	fmt.Println("After V:", time.Since(now))
}

func main() {
	t := time.Second * 5
	sleep(t)
	sleepAfterVersion(t)
}
