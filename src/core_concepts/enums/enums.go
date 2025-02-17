package enums

import "fmt"

// In Go, enums are implemented using constants and iota.
// iota is a special identifier that simplifies the definition of incrementing numbers.

type Day int

const (
	Sunday    Day = iota // iota starts at 0
	Monday               // iota increments to 1
	Tuesday              // iota increments to 2
	Wednesday            // iota increments to 3
	Thursday             // iota increments to 4
	Friday               // iota increments to 5
	Saturday             // iota increments to 6
)

// String method to convert the enum value to a string
func (d Day) String() string {
	// Define a slice of strings representing the days of the week
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	// Return the string corresponding to the enum value
	return days[d]
}

// Execute function to demonstrate the usage of enums
func Execute() {
	// Using the enum
	today := Wednesday
	fmt.Println("Today is:", today)

	// Iterating over enum values
	for d := Sunday; d <= Saturday; d++ {
		fmt.Println(d)
	}
}

// Note: Enums in Go are implemented using constants and iota.
// The iota identifier simplifies the definition of incrementing numbers.
// The String method is used to convert the enum value to a string for easy printing.
// The slice of strings represents the names of the days, and the enum value is used as an index to access the corresponding string.
