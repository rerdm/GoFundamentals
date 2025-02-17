package constants

import "fmt"

const (
    Pi       = 3.14
    Language = "Go"
)

func Execute() {
    fmt.Println("Value of Pi:", Pi)
    fmt.Println("Programming Language:", Language)
}
