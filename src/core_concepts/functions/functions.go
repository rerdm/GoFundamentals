package main

import "fmt"

// Function that adds two integers
func add(a int, b int) int {
    return a + b
}

// Variadic function that sums multiple integers
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

// Function with named return values
func divide(a, b float64) (result float64, err error) {
    if b == 0 {
        err = fmt.Errorf("division by zero")
        return
    }
    result = a / b
    return
}

// Function that modifies a value through a pointer
func increment(ptr *int) {
    *ptr++
}

func main() {
    result := add(3, 4)
    fmt.Println("Sum:", result)

    total := sum(1, 2, 3, 4, 5)
    fmt.Println("Total:", total)

    quotient, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Quotient:", quotient)
    }

    value := 10
    increment(&value)
    fmt.Println("Incremented value:", value)
}
