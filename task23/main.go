package main

import "fmt"

func delete(sl []int, idx int) []int {
	if (idx > len(sl)-1) || (idx < 0) {
		panic("Wrong idx")
	}
	return append(sl[:idx], sl[idx+1:]...)
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	a = delete(a, 3)
	fmt.Println(a, len(a), cap(a))
	a = delete(a, 4)
}
