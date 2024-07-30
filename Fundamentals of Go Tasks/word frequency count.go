package main

import (
	"strings"
)

func wordCounter(input string) map[string]int {
	input = strings.ToLower(input)
	words := []string{}
	currentWord := []string{}

	for _, c := range input {
		if (c >= 'a' && c <= 'z') || c == ' ' || (c >= '0' && c <= '9'){
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

	wordFreq := make(map[string]int)
	for _, word := range words {
		wordFreq[word]++
	}

	return wordFreq
}
