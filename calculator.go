package main

import "fmt"

func calculator(property string, a, b int) int {
    var x int

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
        }
	case "modulus":
        if b != 0 {
            x = a % b
        }
    default:
        fmt.Println("undefined")
    }

    return x
}
