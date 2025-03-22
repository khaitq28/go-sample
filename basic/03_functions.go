package main

/*
Functions Tutorial
----------------
This tutorial covers various aspects of functions in Go:
1. Basic Functions
   - Function declaration
   - Parameters and return values
2. Advanced Function Features
   - Multiple return values
   - Named return values
   - Variadic functions
3. Function Types
   - Functions as parameters
   - Anonymous functions
   - Closures
*/

import (
	"fmt"
)

// Basic function
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// Function with return value
func add(a, b int) int {
	return a + b
}

// Function with multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// Function with named return values
func rectangle(width, height float64) (area float64, perimeter float64) {
	area = width * height
	perimeter = 2 * (width + height)
	return // naked return
}

// Variadic function (accepts variable number of arguments)
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Function as a parameter
func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// Anonymous function
func main() {
	// 1. Basic function call
	fmt.Println("Basic function:")
	greet("Alice")

	// 2. Function with return value
	fmt.Println("\nFunction with return value:")
	result := add(5, 3)
	fmt.Printf("5 + 3 = %d\n", result)

	// 3. Function with multiple return values
	fmt.Println("\nFunction with multiple return values:")
	quotient, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", quotient)
	}

	// 4. Function with named return values
	fmt.Println("\nFunction with named return values:")
	area, perimeter := rectangle(5, 3)
	fmt.Printf("Rectangle: Area = %.2f, Perimeter = %.2f\n", area, perimeter)

	// 5. Variadic function
	fmt.Println("\nVariadic function:")
	total := sum(1, 2, 3, 4, 5)
	fmt.Printf("Sum of numbers: %d\n", total)

	// 6. Function as a parameter
	fmt.Println("\nFunction as a parameter:")
	multiply := func(a, b int) int {
		return a * b
	}
	result = applyOperation(4, 2, multiply)
	fmt.Printf("4 * 2 = %d\n", result)

	// 7. Anonymous function
	fmt.Println("\nAnonymous function:")
	func() {
		fmt.Println("This is an anonymous function!")
	}()

	// 8. Closure
	fmt.Println("\nClosure:")
	counter := func() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}()
	fmt.Printf("Counter: %d\n", counter())
	fmt.Printf("Counter: %d\n", counter())
	fmt.Printf("Counter: %d\n", counter())
}
