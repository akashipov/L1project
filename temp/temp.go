package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []float64{1, 2, 3, 15, 18, -7, -9, 0, -15, -4, -4}
	m := make(map[float64][]float64)
	for _, v := range nums {
		k := math.Mod(v, 10)
		k = v - k
		if v < 0 {
			k -= 10
		} else {
			k += 10
		}
		value, ok := m[k]
		if !ok {
			m[k] = []float64{v}
			continue
		}
		value = append(value, v)
		m[k] = value
	}
	fmt.Printf("%+v\n", m)
}
