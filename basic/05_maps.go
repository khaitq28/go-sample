package main

import (
	"fmt"
)

func main() {
	// 1. Creating Maps
	fmt.Println("Creating Maps:")

	// Empty map
	var emptyMap map[string]int
	fmt.Printf("Empty map: %v\n", emptyMap)

	// Map with make
	ages := make(map[string]int)
	ages["Alice"] = 25
	ages["Bob"] = 30
	fmt.Printf("Ages map: %v\n", ages)

	// Map literal
	scores := map[string]int{
		"Alice":   95,
		"Bob":     87,
		"Charlie": 92,
	}
	fmt.Printf("Scores map: %v\n", scores)

	// 2. Map Operations
	fmt.Println("\nMap Operations:")

	// Adding elements
	scores["David"] = 88
	fmt.Printf("After adding David: %v\n", scores)

	// Updating elements
	scores["Alice"] = 98
	fmt.Printf("After updating Alice: %v\n", scores)

	// Deleting elements
	delete(scores, "Bob")
	fmt.Printf("After deleting Bob: %v\n", scores)

	// 3. Map Access
	fmt.Println("\nMap Access:")

	// Direct access
	fmt.Printf("Alice's score: %d\n", scores["Alice"])

	// Check if key exists
	if score, exists := scores["Charlie"]; exists {
		fmt.Printf("Charlie's score: %d\n", score)
	} else {
		fmt.Println("Charlie not found")
	}

	// 4. Map Iteration
	fmt.Println("\nMap Iteration:")

	// Iterate over keys and values
	for name, score := range scores {
		fmt.Printf("%s: %d\n", name, score)
	}

	// 5. Nested Maps
	fmt.Println("\nNested Maps:")

	// Create nested map
	students := map[string]map[string]int{
		"Alice": {
			"Math":    95,
			"Science": 88,
		},
		"Bob": {
			"Math":    87,
			"Science": 92,
		},
	}
	fmt.Printf("Nested map: %v\n", students)

	// Access nested map
	fmt.Printf("Alice's Math score: %d\n", students["Alice"]["Math"])

	// 6. Map Functions
	fmt.Println("\nMap Functions:")

	// Function to check if map is empty
	isEmpty := len(scores) == 0
	fmt.Printf("Is scores map empty? %v\n", isEmpty)

	// Function to get all keys
	keys := make([]string, 0, len(scores))
	for k := range scores {
		keys = append(keys, k)
	}
	fmt.Printf("Map keys: %v\n", keys)

	// 7. Map Safety
	fmt.Println("\nMap Safety:")

	// Safe access with default value
	score := scores["Unknown"]
	fmt.Printf("Accessing non-existent key: %d\n", score)

	// Safe access with existence check
	if score, exists := scores["Unknown"]; exists {
		fmt.Printf("Score: %d\n", score)
	} else {
		fmt.Println("Key does not exist")
	}

	// 8. Map as Set
	fmt.Println("\nMap as Set:")

	// Using map as a set
	unique := make(map[string]bool)
	words := []string{"apple", "banana", "apple", "orange", "banana"}

	for _, word := range words {
		unique[word] = true
	}

	fmt.Printf("Unique words: ")
	for word := range unique {
		fmt.Printf("%s ", word)
	}
	fmt.Println()
}
