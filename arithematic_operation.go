package main

import "fmt"

func arithematic_opertn() {
	var a int = 8
	var b int = 4

	add := func() {
		fmt.Println("Addition is:", a+b)
	}
	sub := func() {
		fmt.Println("Subtraction is:", a-b)
	}
	multiply := func() {
		fmt.Println("Multiplication is:", a*b)
	}
	div := func() {
		if b != 0 {
			fmt.Println("Division is:", a/b)
		} else {
			fmt.Println("Number is not divisible")
		}
	}

	add()
	sub()
	multiply()
	div()
}
