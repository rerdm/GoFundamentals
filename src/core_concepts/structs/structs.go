package structs

import (
    "fmt"
    "reflect"
)

// Define a struct
type Person struct {
    Name string `json:"name"` // Field name with a JSON tag
    Age  int    `json:"age"`  // Field age with a JSON tag
}

// Function to demonstrate struct field names, types, and tags
func printStructInfo(p Person) {
    // Get the type of the struct
    t := reflect.TypeOf(p)

    // Iterate over the struct fields
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        fmt.Printf("Field Name: %s, Field Type: %s, Field Tag: %s\n", field.Name, field.Type, field.Tag)
    }
}

// Execute function to be called from main.go
func Execute() {
    // Create an instance of the struct
    p := Person{Name: "Alice", Age: 30}
    fmt.Println("Name:", p.Name)
    fmt.Println("Age:", p.Age)

    // Print struct field names, types, and tags
    printStructInfo(p)
}

// Note: Struct tags are used to attach metadata to struct fields.
// In this example, the `json` tags specify the key names to be used
// when serializing the struct to JSON. This is useful for controlling
// how the struct is encoded/decoded by libraries such as encoding/json.