package main

import "fmt"

func vowels(ch string) {
	if ch == "a" || ch == "e" || ch == "i" || ch == "o" || ch == "u" {
		fmt.Printf("The character %s is Vowel. \n ", ch)
	} else {
		fmt.Println("The character %s is Consonent. \n" , ch)
	}
}