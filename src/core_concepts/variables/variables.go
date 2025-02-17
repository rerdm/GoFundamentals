package variables

import "fmt"

func Execute() {
    // Declare a variable with a specific type and initialize it
    var a int = 10
    fmt.Println("a:", a)

    // Declare a variable without specifying the type (type inference)
    var b = 20
    fmt.Println("b:", b)

    // Short variable declaration (only inside functions)
    c := 30
    fmt.Println("c:", c)

    // Declare multiple variables at once
    var d, e, f int = 1, 2, 3
    fmt.Println("d:", d, "e:", e, "f:", f)

    // Declare multiple variables with type inference
    var g, h, i = 4, "hello", 5.5
    fmt.Println("g:", g, "h:", h, "i:", i)

    // Short variable declaration for multiple variables
    j, k, l := 6, "world", 7.7
    fmt.Println("j:", j, "k:", k, "l:", l)

    // Declare a variable without initializing it (zero value)
    var m int
    fmt.Println("m (zero value):", m)

    // Declare a constant
    const n = 100
    fmt.Println("n (constant):", n)
}

// Note: 
// 1. `var a int = 10`: Declares a variable `a` of type `int` and initializes it with the value 10.
// 2. `var b = 20`: Declares a variable `b` with type inference, meaning Go will infer the type based on the value (int in this case).
// 3. `c := 30`: Short variable declaration, only allowed inside functions. It declares and initializes `c` with the value 30.
// 4. `var d, e, f int = 1, 2, 3`: Declares multiple variables `d`, `e`, and `f` of type `int` and initializes them with values 1, 2, and 3, respectively.
// 5. `var g, h, i = 4, "hello", 5.5`: Declares multiple variables with type inference. `g` is an `int`, `h` is a `string`, and `i` is a `float64`.
// 6. `j, k, l := 6, "world", 7.7`: Short variable declaration for multiple variables. `j` is an `int`, `k` is a `string`, and `l` is a `float64`.
// 7. `var m int`: Declares a variable `m` of type `int` without initializing it. It gets the zero value for its type (0 for int).
// 8. `const n = 100`: Declares a constant `n` with the value 100. Constants cannot be changed once declared.
