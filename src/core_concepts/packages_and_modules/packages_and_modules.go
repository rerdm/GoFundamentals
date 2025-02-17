package packages_and_modules

import (
    "fmt"
    "math"
    "github.com/google/uuid" // Importing an external module from GitHub
)

// Execute function to be called from other packages
func Execute() {
    // Using the math package
    fmt.Println("Square root of 16 is", math.Sqrt(16))

    // Using the external module
    id := uuid.New()
    fmt.Println("Generated UUID:", id)

    // Example of using the UUID in a function
    processUUID(id)
}

// Function that takes a UUID and prints it
func processUUID(id uuid.UUID) {
    fmt.Println("Processing UUID:", id)
}

// Note: Ensure your go.mod file includes the external module dependency:
// module GoFundamentals

// run go mod tidy to update the file

// go 1.16

// require (
//     github.com/google/uuid v1.3.0 // Replace with the actual version
// )
