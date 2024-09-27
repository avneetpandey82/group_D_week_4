package main

import "fmt"

func tablePrinting(num int) {
	fmt.Printf("Multiplication Table of %d:\n", num)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", num, i, num*i)
	}
}
