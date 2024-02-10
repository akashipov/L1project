package main

import "fmt"

func main() {
	nums1 := map[int]int{2: 2, 4: 4, 7: 7}
	nums2 := map[int]int{1: 1, 3: 3, 4: 4, 7: 7}
	m := make(map[int]int)
	for k := range nums1 {
		_, ok := nums2[k]
		if ok {
			m[k] = k
		}
	}
	fmt.Printf("%+v\n", m)
}
