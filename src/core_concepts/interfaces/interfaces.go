package interfaces

import "fmt"

// An interface is a type that specifies a set of method signatures.
// A type implements an interface by implementing its methods.
// Interfaces are used to define the behavior that different types must have.

type Shape interface {
    // Any type that has an Area method with this signature
    // (no parameters, returns a float64) satisfies the Shape interface.
    Area() float64
}

// A struct is a collection of fields.
// Here, we define a struct named Circle with a single field Radius.
type Circle struct {
    Radius float64
}

// Implement the interface by defining the method Area for the Circle type.
// This means that Circle now satisfies the Shape interface.
func (c Circle) Area() float64 {
    // Calculate the area of the circle using the formula πr²
    return 3.14 * c.Radius * c.Radius 
}

func Execute() {
    // Declare a variable of type Shape and assign it a Circle.
    // Since Circle implements the Shape interface, this is valid.
    var s Shape = Circle{Radius: 5}
    // Call the Area method on the Shape interface, which will use
    // the implementation provided by the Circle type.
    fmt.Println("Area of the circle:", s.Area())
}
