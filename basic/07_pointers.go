package main

/*
Pointers Tutorial
----------------
This tutorial covers pointers in Go:
1. Basic Pointers
   - Pointer declaration
   - Address operator (&)
   - Dereference operator (*)
2. Advanced Pointer Features
   - Pointers to structs
   - Pointers to arrays/slices
   - Pointers to functions
3. Pointer Safety
   - Nil pointers
   - Pointer comparison
   - Safe pointer dereferencing
4. Special Cases
   - Double pointers
   - Pointer arithmetic (not supported)
   - new() function
*/

import (
	"fmt"
)

// Function to demonstrate pointer basics
func pointerBasics() {
	// 1. Basic pointer declaration and usage
	x := 42
	ptr := &x
	fmt.Printf("Value of x: %d\n", x)
	fmt.Printf("Address of x: %p\n", ptr)
	fmt.Printf("Value at ptr: %d\n", *ptr)

	// 2. Modifying value through pointer
	*ptr = 100
	fmt.Printf("New value of x: %d\n", x)

	// 3. Nil pointer
	var nilPtr *int
	fmt.Printf("Nil pointer: %v\n", nilPtr)

	// 4. New keyword
	newPtr := new(int)
	*newPtr = 200
	fmt.Printf("Value at newPtr: %d\n", *newPtr)
}

// Function to demonstrate pointer to struct
type Person struct {
	Name string
	Age  int
}

func structPointers() {
	// 1. Creating struct and pointer
	person := Person{Name: "Alice", Age: 25}
	personPtr := &person

	// 2. Accessing struct fields through pointer
	fmt.Printf("Name: %s, Age: %d\n", personPtr.Name, personPtr.Age)

	// 3. Modifying struct through pointer
	personPtr.Age = 26
	fmt.Printf("Updated age: %d\n", person.Age)

	// 4. Creating pointer to struct directly
	personPtr2 := &Person{Name: "Bob", Age: 30}
	fmt.Printf("Person 2: %+v\n", *personPtr2)
}

// Function to demonstrate pointer to array/slice
func arrayPointers() {
	// 1. Array pointer
	arr := [3]int{1, 2, 3}
	arrPtr := &arr
	fmt.Printf("Array: %v\n", *arrPtr)

	// 2. Modifying array through pointer
	(*arrPtr)[0] = 10
	fmt.Printf("Modified array: %v\n", arr)

	// 3. Slice (already contains pointer)
	slice := []int{1, 2, 3}
	slicePtr := &slice
	fmt.Printf("Slice: %v\n", *slicePtr)

	// 4. Modifying slice through pointer
	(*slicePtr)[0] = 10
	fmt.Printf("Modified slice: %v\n", slice)
}

// Function to demonstrate pointer to function
func functionPointers() {
	// 1. Function type
	type MathFunc func(int, int) int

	// 2. Assigning function to variable
	add := func(a, b int) int { return a + b }
	multiply := func(a, b int) int { return a * b }

	// 3. Using function pointer
	var mathFunc MathFunc
	mathFunc = add
	fmt.Printf("Add result: %d\n", mathFunc(5, 3))

	mathFunc = multiply
	fmt.Printf("Multiply result: %d\n", mathFunc(5, 3))
}

// Function to demonstrate pointer to interface
func interfacePointers() {
	// 1. Interface type
	var writer interface{}

	// 2. Pointer to interface
	writerPtr := &writer

	// 3. Assigning value to interface through pointer
	*writerPtr = "Hello"
	fmt.Printf("Interface value: %v\n", *writerPtr)
}

func main() {
	fmt.Println("Pointer Basics:")
	pointerBasics()

	fmt.Println("\nStruct Pointers:")
	structPointers()

	fmt.Println("\nArray/Slice Pointers:")
	arrayPointers()

	fmt.Println("\nFunction Pointers:")
	functionPointers()

	fmt.Println("\nInterface Pointers:")
	interfacePointers()

	// Additional pointer examples
	fmt.Println("\nAdditional Pointer Examples:")

	// 1. Double pointer
	x := 42
	ptr := &x
	doublePtr := &ptr
	fmt.Printf("Value: %d\n", **doublePtr)

	// 2. Pointer comparison
	y := 42
	ptr2 := &y
	fmt.Printf("Pointers equal: %v\n", ptr == ptr2)

	// 3. Pointer arithmetic (not supported in Go)
	// This would cause a compilation error:
	// ptr++

	// 4. Safe pointer dereferencing
	var safePtr *int
	if safePtr != nil {
		fmt.Printf("Value: %d\n", *safePtr)
	} else {
		fmt.Println("Pointer is nil")
	}
}
