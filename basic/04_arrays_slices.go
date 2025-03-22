package main

/*
Arrays and Slices Tutorial
-------------------------
This tutorial covers arrays and slices in Go:
1. Arrays
   - Fixed-size arrays
   - Array initialization
   - Array operations
2. Slices
   - Slice creation
   - Slice operations (append, copy)
   - Slice tricks
3. Advanced Topics
   - 2D slices
   - Slice capacity and length
   - Slice memory management
*/

import (
	"fmt"
)

func main() {
	// 1. Arrays
	fmt.Println("Arrays:")

	// Fixed-size array
	var numbers [5]int
	fmt.Printf("Empty array: %v\n", numbers)

	// Array with values
	primes := [5]int{2, 3, 5, 7, 11}
	fmt.Printf("Primes array: %v\n", primes)

	// Array with size inference
	fruits := [...]string{"apple", "banana", "orange"}
	fmt.Printf("Fruits array: %v\n", fruits)

	// 2. Array Operations
	fmt.Println("\nArray Operations:")

	// Accessing elements
	fmt.Printf("First prime: %d\n", primes[0])
	fmt.Printf("Last prime: %d\n", primes[len(primes)-1])

	// Modifying elements
	primes[0] = 1
	fmt.Printf("Modified primes: %v\n", primes)

	// Array length
	fmt.Printf("Array length: %d\n", len(primes))

	// 3. Slices
	fmt.Println("\nSlices:")

	// Creating slices
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("Slice 1: %v\n", slice1)

	// Creating slice from array
	slice2 := primes[1:4]
	fmt.Printf("Slice from array: %v\n", slice2)

	// Empty slice
	var emptySlice []int
	fmt.Printf("Empty slice: %v\n", emptySlice)

	// 4. Slice Operations
	fmt.Println("\nSlice Operations:")

	// Append
	slice1 = append(slice1, 6, 7)
	fmt.Printf("After append: %v\n", slice1)

	// Append slice to slice
	slice3 := []int{8, 9, 10}
	slice1 = append(slice1, slice3...)
	fmt.Printf("After appending slice: %v\n", slice1)

	// Make slice with capacity
	slice4 := make([]int, 3, 5)
	fmt.Printf("Made slice: %v, Length: %d, Capacity: %d\n",
		slice4, len(slice4), cap(slice4))

	// 5. Slice Tricks
	fmt.Println("\nSlice Tricks:")

	// Copy slice
	slice5 := make([]int, len(slice1))
	copy(slice5, slice1)
	fmt.Printf("Copied slice: %v\n", slice5)

	// Delete element
	slice6 := []int{1, 2, 3, 4, 5}
	slice6 = append(slice6[:2], slice6[3:]...)
	fmt.Printf("After deleting element: %v\n", slice6)

	// 6. 2D Slices
	fmt.Println("\n2D Slices:")

	// Create 2D slice
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Printf("2D slice: %v\n", matrix)

	// Access 2D slice elements
	fmt.Printf("Element at [1][2]: %d\n", matrix[1][2])

	// 7. Slice Capacity and Length
	fmt.Println("\nSlice Capacity and Length:")
	slice7 := make([]int, 3, 6)
	fmt.Printf("Initial: Length=%d, Capacity=%d\n", len(slice7), cap(slice7))

	slice7 = append(slice7, 1, 2, 3)
	fmt.Printf("After append: Length=%d, Capacity=%d\n", len(slice7), cap(slice7))

	slice7 = append(slice7, 4)
	fmt.Printf("After exceeding capacity: Length=%d, Capacity=%d\n",
		len(slice7), cap(slice7))
}
