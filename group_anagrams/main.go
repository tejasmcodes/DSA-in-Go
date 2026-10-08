package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
    groupsMap := make(map[[26]int][]string)
    for _, word := range strs {
      freq := [26]int{}
      for i:= range word {
        freq[word[i] - 'a']++
      }
      groupsMap[freq] = append(groupsMap[freq], word)
    }

	result := [][]string{}
	for _, value := range groupsMap {
		result = append(result,value)
	}
  return result
}


func main() {
  strs := []string{"eat","tea","tan","ate","nat","bat"}
  fmt.Printf("%+v",groupAnagrams(strs))
}