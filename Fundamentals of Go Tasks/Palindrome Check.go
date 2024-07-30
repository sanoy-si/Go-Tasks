package main

import (
	"strings"
)

func isPalindrome(input string) bool {
	input = strings.ToLower(input)
	words := []string{}
	currentWord := []string{}

	for _, c := range input {
		if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9'){
			currentWord = append(currentWord, string(c))
		} else {
			if len(currentWord) != 0{
				words = append(words, strings.Join(currentWord,""))
				currentWord = []string{}
			}
		}
	}
	if len(currentWord) != 0 {
		words = append(words, strings.Join(currentWord,""))
	}

	s := strings.Join(words, "")
	left, right := 0, len(s) - 1

	for left < right{
		if s[left] != s[right]{
			return false
		}else{
			left += 1
			right -= 1
		}
	}

	return true

}
