package control_structures

import "fmt"

func Execute() {
    // If-else statement
    num := 10
    if num%2 == 0 {
        fmt.Println(num, "is even")
    } else {
        fmt.Println(num, "is odd")
    }

    // For loop
    for i := 0; i < 5; i++ {
        fmt.Println("i:", i)
    }

    // Switch statement
    day := "Monday"
    switch day {
    case "Monday":
        fmt.Println("Start of the work week")
    case "Friday":
        fmt.Println("End of the work week")
    default:
        fmt.Println("Midweek")
    }
}
