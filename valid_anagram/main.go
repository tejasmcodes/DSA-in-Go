package main

import "fmt"

func validAnagram(s, t string) bool {
	if len(s) != len(t){
		return false
	}
	// freq map for the string s
	freqS := make(map[rune]int)
	for _, char := range s {
		freqS[char]++
	}

	for _, char := range t {
		// for every rune in t, if it's in freqS decrement it
		freqS[char]--

		// if the char count is neg, it means that count of of a certain char of t is greater that that of s
		if freqS[char] < 0{
			return false
		}
	}

	return true
}

func main( ){
	s := "anagram"
	t := "nagaram"
	fmt.Println(validAnagram(s,t))
}