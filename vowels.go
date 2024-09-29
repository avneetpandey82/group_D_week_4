package main

import "fmt"

func vowels(ch string) {
	if ch == "a" || ch == "e" || ch == "i" || ch == "o" || ch == "u" || ch == "A" || ch == "E" || ch == "I" || ch == "O" || ch == "U" {
		fmt.Printf("The character %s is Vowel. \n ", ch)
	} else {
		fmt.Printf("The character %s is Consonent. \n", ch)
	}
}
