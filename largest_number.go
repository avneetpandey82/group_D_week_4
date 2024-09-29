package main

import "fmt"

func largestnumber() {
	var a int
	fmt.Printf("Please enter the number of elements: ")
	fmt.Scan(&a)

	numbers := make([]int, a)
	fmt.Println("Please provide the numbers:")

	for i := 0; i < a; i++ {
		fmt.Scan(&numbers[i])
	}

	largest := numbers[0]
	for _, num := range numbers {
		if num > largest {
			largest = num
		}
	}

	fmt.Printf("The largest number among the given number is: %d\n", largest)
}
