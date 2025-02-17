package main

import "fmt"

const (
    A = iota // A == 0
    B        // B == 1
    C        // C == 2
)

const (
    X = iota // X == 0 (iota resets to 0)
    Y        // Y == 1
    Z        // Z == 2
)

func main() {
    fmt.Println("A:", A)
    fmt.Println("B:", B)
    fmt.Println("C:", C)
    fmt.Println("X:", X)
    fmt.Println("Y:", Y)
    fmt.Println("Z:", Z)
}
