package main

import "fmt"

// calculator function to perform basic arithmetic operations
func calculator(property string, a, b int) int {
    var x int

    // Perform the appropriate calculation based on the property value
    switch property {
    case "add":
        x = a + b
    case "subtract":
        x = a - b
    case "multiply":
        x = a * b
    case "division":
        if b != 0 {
            x = a / b
        } else {
            fmt.Println("Error: Division by zero")
        }
	case "modulus":
        if b != 0 {
            x = a % b
        } else {
            fmt.Println("Error: Division by zero in modulus operation")
        }
    default:
        fmt.Println("Unsupported operation")
    }

    return x
}
