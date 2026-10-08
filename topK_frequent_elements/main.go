package main

import (
	"fmt"
	"slices"
	"cmp"
)

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int, len(nums))
	for _, num := range nums {
		freq[num]++
	}

	numsList := make([][2]int, 0, len(nums))

	for key, value := range freq {
		numsList = append(numsList,[2]int{key, value})
	}


	slices.SortFunc(numsList, func(a,b [2]int) int {
		return cmp.Compare(b[1], a[1])
	})

	result := make([]int, 0, k)
	for i := range k {
		result = append(result, numsList[i][0])
	}

	return result
}

func main() {
	nums := []int{1,2,1,2,1,2,3,1,3,2}
	k := 2
	fmt.Println(topKFrequent(nums,k))
}