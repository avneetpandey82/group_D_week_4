package main

import "fmt"

func license(age int) {

	if age >= 18 {
		fmt.Println("You are allowed to drive.")
	} else {
		fmt.Println("You are not allowed to drive.")
	}
}
