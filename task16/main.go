package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Intn(n int) int {
	if n == 0 {
		return 0
	}
	return rand.Intn(n)
}

func split(nums []int64, value int64) int {
	left := 0
	right := len(nums) - 1
	var idx int
	i := 0
	for i <= right {
		v := nums[i]
		if v <= value {
			nums[left] = v
			if v == value {
				idx = left
			}
			left += 1
			i += 1
		} else {
			nums[right], nums[i] = nums[i], nums[right]
			right -= 1
		}
	}
	nums[left-1], nums[idx] = nums[idx], nums[left-1]
	return left - 1
}

func sortLogic(nums []int64, value int64) []int64 {
	if (len(nums) == 1) || (len(nums) == 0) {
		return nums
	}
	// fmt.Println("Before nums:", nums)
	idx := split(nums, value)
	// fmt.Println("After nums:", nums)
	left := idx + 1
	n := Intn(idx)
	res := sortLogic(nums[:idx], nums[n])
	n = Intn(len(nums) - left)
	res = append(res, nums[idx])
	if left > len(nums)-1 {
		return res
	}
	res = append(res, sortLogic(nums[left:], nums[left+n])...)
	return res
}

func quicksort(nums []int64) []int64 {
	if (len(nums) == 1) || (len(nums) == 0) {
		return nums
	}
	i := Intn(len(nums))
	return sortLogic(nums, nums[i])
}

func main() {
	num1 := []int64{3, 2}
	fmt.Println(quicksort(num1))
	num1 = []int64{2, 3}
	fmt.Println(quicksort(num1))
	num1 = []int64{1, 2, 4, 5, 1, 8}
	fmt.Println(quicksort(num1))
	t := time.Now()
	num4 := []int64{8, 4, 2, 4, 6, 8, 3, 2, 678, 45, 1, 3}
	fmt.Println(quicksort(num4))
	fmt.Println(time.Since(t))
	num1 = []int64{5, 4, 3, 2, 1}
	fmt.Println(quicksort(num1))
	num1 = []int64{2, 2, 2, 2, 2, 2, 2, 2, 2, 2}
	fmt.Println(quicksort(num1))
	t = time.Now()
	n := len(num4) * 30
	num2 := make([]int64, n)
	for i := 0; i < 30; i++ {
		copy(num2[i*len(num4):(i+1)*len(num4)], num4)
	}
	fmt.Println(num2)
	fmt.Println(quicksort(num2))
	fmt.Println(time.Since(t))
}
