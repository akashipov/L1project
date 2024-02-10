package main

import "fmt"

func main() {
	a := []string{"cat", "cat", "dog", "cat", "tree"}
	m := make(map[string]string)
	for _, v := range a {
		_, ok := m[v]
		if !ok {
			m[v] = v
		}
	}
	fmt.Println(m)
}
