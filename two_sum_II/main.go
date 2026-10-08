package main

import "fmt"

func twoSum(nums []int, target int) []int {
	start := 0
	end := len(nums) - 1

	for start < end {
		if nums[start]+nums[end] > target {
			end--
		} else if nums[start]+nums[end] < target {
			start++
		} else{
			return []int{start+1, end+1}
		}
	}
	return []int{}
}

func main(){
	nums := []int{-1,0}
	target := -1
	fmt.Println(twoSum(nums, target))
}