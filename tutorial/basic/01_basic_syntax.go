package main

/*
Basic Syntax Tutorial
--------------------
This tutorial covers the fundamental syntax elements in Go:
1. Variables and Data Types
   - Variable declaration (var, :=)
   - Basic data types (int, float64, string, bool, rune)
   - Type conversion
   - Zero values
2. Constants
   - Constant declaration
   - Multiple constant declaration
3. Print Statements
   - fmt.Println
   - fmt.Printf with format verbs
*/

import (
	"fmt"
)

// This is a comment
/*
   This is a multi-line comment
*/

func main() {
	// 1. Variables and Data Types
	var name string = "John"
	age := 25 // Short variable declaration
	var (
		height float64 = 1.75
		weight float64 = 70.5
	)

	// 2. Constants
	const PI = 3.14159
	const (
		Monday    = 1
		Tuesday   = 2
		Wednesday = 3
	)

	// 3. Basic Print Statements
	fmt.Println("Hello, Go!")
	fmt.Printf("Name: %s, Age: %d\n", name, age)
	fmt.Printf("Height: %.2fm, Weight: %.1fkg\n", height, weight)

	// 4. Basic Data Types
	var (
		integer   int     = 42
		float     float64 = 3.14
		boolean   bool    = true
		character rune    = 'A'
		text      string  = "Hello, World!"
	)

	// 5. Type Conversion
	var intToFloat float64 = float64(integer)
	var floatToInt int = int(float)

	// 6. Zero Values
	var (
		zeroInt    int
		zeroFloat  float64
		zeroString string
		zeroBool   bool
	)

	// Print all examples
	fmt.Println("\nBasic Data Types:")
	fmt.Printf("Integer: %d\n", integer)
	fmt.Printf("Float: %.2f\n", float)
	fmt.Printf("Boolean: %v\n", boolean)
	fmt.Printf("Character: %c\n", character)
	fmt.Printf("Text: %s\n", text)

	fmt.Println("\nType Conversion:")
	fmt.Printf("Int to Float: %.2f\n", intToFloat)
	fmt.Printf("Float to Int: %d\n", floatToInt)

	fmt.Println("\nZero Values:")
	fmt.Printf("Zero Int: %d\n", zeroInt)
	fmt.Printf("Zero Float: %.2f\n", zeroFloat)
	fmt.Printf("Zero String: %q\n", zeroString)
	fmt.Printf("Zero Bool: %v\n", zeroBool)
}
