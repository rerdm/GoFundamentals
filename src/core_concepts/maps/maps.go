package maps

import "fmt"

func Execute() {
    // Creating a map
    capitals := map[string]string{
        "France": "Paris",
        "Italy":  "Rome",
        "Japan":  "Tokyo",
    }

    // Adding an element
    capitals["Germany"] = "Berlin"

    // Accessing an element
    fmt.Println("The capital of Japan is", capitals["Japan"])

    // Iterating over a map
    for country, capital := range capitals {
        fmt.Println("The capital of", country, "is", capital)
    }

    // Deleting an element
    delete(capitals, "Italy")
    fmt.Println("Updated map:", capitals)
}
