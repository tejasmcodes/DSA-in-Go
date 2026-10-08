package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool{

	start := 0
	end := len(s)-1

	for start < end {
		left := rune(s[start])
		right := rune(s[end])
		
		if !unicode.IsLetter(left) && !unicode.IsDigit(left){
			start++
			continue
		}
		if !unicode.IsLetter(right) && !unicode.IsDigit(right){
			end--
			continue
		}
		if unicode.ToLower(left) != unicode.ToLower(right){
			return false
		}
		start++
		end--
	}
	return true
}


func main(){
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
}