package main

import "fmt"

func arithematic_opertn() {
	var a int = 8
	var b int = 3
	var c float32

	println()
	println("Basic Arthematic Operations of 2 integres-")

	add := func() {
		fmt.Println("Addition of", a, "+", b, "is", a+b)
	}
	sub := func() {
		fmt.Println("Subtraction of", a, "-", b, "is", a-b)
	}
	multiply := func() {
		fmt.Println("Multiplication of", a, "*", b, "is", a*b)
	}
	div := func() {
		if b != 0 {
			c = float32(a) / float32(b)
			fmt.Printf("Division of %d / %d is %.2f\n", a, b, c)
		} else {
			fmt.Println("Number is not divisible")
		}
	}
	compare := func() {
		if a > b {
			fmt.Println(a, "is Greater than", b)
		} else if a < b {
			fmt.Println(a, "is Smaller than", b)
		} else {
			fmt.Println(a, "is Equal to", b)
		}
	}
	even := func() {
		if a%2 == 0 {
			fmt.Println(a, "is Even Number")
		} else {
			fmt.Println(a, "is Odd Number")
		}
	}
	odd := func() {
		if b%2 == 0 {
			fmt.Println(b, "is Even Number")
		} else {
			fmt.Println(b, "is Odd Number")
		}
	}

	println()

	add()
	sub()
	multiply()
	div()
	compare()
	even()
	odd()

}
