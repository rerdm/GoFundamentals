package loops

import "fmt"

func Execute() {
    // For loop
    for i := 0; i < 5; i++ {
        fmt.Println("i:", i)
    }

    // While-like loop
    j := 0
    for j < 5 {
        fmt.Println("j:", j)
        j++
    }

    // Infinite loop with break
    k := 0
    for {
        if k >= 5 {
            break
        }
        fmt.Println("k:", k)
        k++
    }
}
