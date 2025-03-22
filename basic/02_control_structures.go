package main

/*
Control Structures Tutorial
--------------------------
This tutorial demonstrates various control structures in Go:
1. If Statements
   - Basic if
   - if-else
   - if with initialization
2. For Loops
   - Basic for loop
   - While-style loop
   - Infinite loop with break
   - Range-based loop
3. Switch Statements
   - Basic switch
   - Switch with multiple cases
   - Switch without expression
*/

import (
	"fmt"
)

func main() {
	// 1. If Statements
	age := 18

	// Basic if
	if age >= 18 {
		fmt.Println("You are an adult")
	}

	// If with else
	if age >= 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}

	// If with initialization
	if score := 85; score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: C")
	}

	// 2. For Loops
	// Basic for loop
	fmt.Println("\nBasic for loop:")
	for i := 0; i < 5; i++ {
		fmt.Printf("Count: %d\n", i)
	}

	// While-style for loop
	fmt.Println("\nWhile-style loop:")
	count := 0
	for count < 3 {
		fmt.Printf("Count: %d\n", count)
		count++
	}

	// Infinite loop with break
	fmt.Println("\nInfinite loop with break:")
	number := 0
	for {
		if number >= 3 {
			break
		}
		fmt.Printf("Number: %d\n", number)
		number++
	}

	// 3. Range-based for loop
	fmt.Println("\nRange-based loop:")
	fruits := []string{"apple", "banana", "orange"}
	for index, fruit := range fruits {
		fmt.Printf("Index: %d, Fruit: %s\n", index, fruit)
	}

	// 4. Switch Statement
	fmt.Println("\nSwitch statement:")
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("It's Monday!")
	case "Tuesday":
		fmt.Println("It's Tuesday!")
	case "Wednesday":
		fmt.Println("It's Wednesday!")
	default:
		fmt.Println("It's another day!")
	}

	// Switch with multiple cases
	fmt.Println("\nSwitch with multiple cases:")
	month := 3
	switch month {
	case 12, 1, 2:
		fmt.Println("Winter")
	case 3, 4, 5:
		fmt.Println("Spring")
	case 6, 7, 8:
		fmt.Println("Summer")
	case 9, 10, 11:
		fmt.Println("Fall")
	default:
		fmt.Println("Invalid month!")
	}

	// Switch without expression
	fmt.Println("\nSwitch without expression:")
	score := 85
	switch {
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 80:
		fmt.Println("Grade: B")
	case score >= 70:
		fmt.Println("Grade: C")
	default:
		fmt.Println("Grade: F")
	}
}
