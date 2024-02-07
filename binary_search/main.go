package main

import "fmt"

func BinarySearch(nums []int64, value int64, shift int) int {
	l := len(nums) / 2
	if nums[l] < value {
		if l+1 > len(nums)-1 {
			return -1
		}
		return BinarySearch(nums[l+1:], value, shift+l+1)
	} else if nums[l] > value {
		if l == 0 {
			return -1
		}
		return BinarySearch(nums[0:l], value, shift)
	} else {
		return shift + l
	}
}

func main() {
	nums := []int64{1, 2, 3, 4, 4, 4, 5, 6, 8, 15, 15}
	fmt.Println(BinarySearch(nums, -1, 0))
}
