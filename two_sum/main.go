package main

import "fmt"

func twoSum(nums []int, target int) []int{
	combo := make(map[int]int)
	for index, num := range nums{
		if value, present := combo[num]; present{
			return []int{value, index}
		} else{
			combo[target-num] = index
		}
	}
	return []int{-1,-1}
}

func main(){
	nums := []int{2,11,15,7}
	target := 9
	result := twoSum(nums, target)
	fmt.Printf("Two sum combination: %v\n",result)
}