package envs

import (
    "fmt"
    "os"
)

// Execute function to display specific environment variables
func Execute() {
    // Access environment variables
    envVar1 := os.Getenv("ENV_VAR1")
    envVar2 := os.Getenv("ENV_VAR2")

    // Display environment variables
    fmt.Println("ENV_VAR1:", envVar1)
    fmt.Println("ENV_VAR2:", envVar2)
}

// PrintAllEnv function to display all environment variables
func PrintAllEnv() {
    for _, env := range os.Environ() {
        fmt.Println(env)
    }
}
