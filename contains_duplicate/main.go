package main

import "fmt"

func containsDuplicate(nums []int) bool {
	// create a exist map which holds a bool
	exist := make(map[int]bool, len(nums))

	for _, num := range nums {
		// check if the num exists
		if present := exist[num]; present {
			return true
		}
		// if it doesn't exist, it means it is appearing for the first time
		exist[num] = true
	}
	return false
}

func main(){
	nums := []int{1,2,3,1}
	nums1 := []int{1,2,3}
	fmt.Println(containsDuplicate(nums))
	fmt.Println(containsDuplicate(nums1))
}