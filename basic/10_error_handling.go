package main

/*
Error Handling Tutorial
---------------------
This tutorial covers error handling in Go:
1. Basic Error Handling
   - Error interface
   - Error creation
   - Error checking
2. Advanced Error Features
   - Custom error types
   - Error wrapping
   - Error chains
3. Error Handling Patterns
   - Panic and recover
   - Timeout handling
   - Multiple error handling
4. Best Practices
   - Error context
   - Error cleanup
   - Error logging
*/

import (
	"errors"
	"fmt"
	"time"
)

// Custom error type
type ValidationError struct {
	Field string
	Issue string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Validation error in %s: %s", e.Field, e.Issue)
}

// Function that returns an error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Function that returns multiple errors
func validateUser(name string, age int) error {
	if name == "" {
		return &ValidationError{
			Field: "name",
			Issue: "name cannot be empty",
		}
	}
	if age < 0 {
		return &ValidationError{
			Field: "age",
			Issue: "age cannot be negative",
		}
	}
	return nil
}

// Function that wraps errors
func processUser(name string, age int) error {
	if err := validateUser(name, age); err != nil {
		return fmt.Errorf("failed to process user: %w", err)
	}
	return nil
}

// Function that demonstrates panic and recover
func riskyOperation() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()

	panic("something went wrong!")
}

// Function that demonstrates error handling with timeouts
func operationWithTimeout() error {
	done := make(chan error)
	go func() {
		// Simulate some work
		time.Sleep(2 * time.Second)
		done <- errors.New("operation timed out")
	}()

	select {
	case err := <-done:
		return err
	case <-time.After(1 * time.Second):
		return errors.New("operation timed out")
	}
}

func main() {
	// 1. Basic error handling
	fmt.Println("Basic Error Handling:")
	result, err := divide(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: %v\n", result)
	}

	// 2. Custom error type
	fmt.Println("\nCustom Error Type:")
	err = validateUser("", -1)
	if err != nil {
		if validationErr, ok := err.(*ValidationError); ok {
			fmt.Printf("Validation Error: %v\n", validationErr)
		}
	}

	// 3. Error wrapping
	fmt.Println("\nError Wrapping:")
	err = processUser("", -1)
	if err != nil {
		fmt.Printf("Wrapped Error: %v\n", err)
	}

	// 4. Panic and recover
	fmt.Println("\nPanic and Recover:")
	riskyOperation()

	// 5. Error handling with timeouts
	fmt.Println("\nError Handling with Timeouts:")
	err = operationWithTimeout()
	if err != nil {
		fmt.Printf("Timeout Error: %v\n", err)
	}

	// 6. Multiple error handling
	fmt.Println("\nMultiple Error Handling:")
	errs := make([]error, 0)

	if err := validateUser("", -1); err != nil {
		errs = append(errs, err)
	}
	if err := validateUser("John", -1); err != nil {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		fmt.Println("Multiple errors occurred:")
		for _, err := range errs {
			fmt.Printf("- %v\n", err)
		}
	}

	// 7. Error handling with cleanup
	fmt.Println("\nError Handling with Cleanup:")
	file := "example.txt"
	defer func() {
		fmt.Printf("Cleaning up file: %s\n", file)
	}()

	// Simulate file operation error
	err = errors.New("file operation failed")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// 8. Error handling with context
	fmt.Println("\nError Handling with Context:")
	type Context struct {
		Operation string
		Details   string
	}

	ctx := Context{
		Operation: "database query",
		Details:   "fetching user data",
	}

	err = errors.New("connection failed")
	if err != nil {
		fmt.Printf("Error in %s (%s): %v\n", ctx.Operation, ctx.Details, err)
	}
}
