package main

import "fmt"

func main() {
    var a int = 10
    var p *int = &a

    fmt.Println("Value of a:", a)
    fmt.Println("Address of a:", p)
    fmt.Println("Value at address p:", *p)
}
