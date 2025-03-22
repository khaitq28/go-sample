package calculator

/*
Testing Tutorial
---------------
This tutorial covers testing in Go:
1. Basic Testing
   - Unit tests
   - Table-driven tests
   - Test coverage
2. Advanced Testing
   - Benchmarks
   - Examples
   - Subtests
3. Testing Best Practices
   - Mocking
   - Test helpers
   - Test organization
*/

import (
	"errors"
	"fmt"
	"testing"
)

// TestAdd tests the Add method
func TestAdd(t *testing.T) {
	calc := &Calculator{}

	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -2, -3, -5},
		{"zero", 0, 0, 0},
		{"mixed numbers", 2, -3, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%v, %v) = %v; want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// TestSubtract tests the Subtract method
func TestSubtract(t *testing.T) {
	calc := &Calculator{}

	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 5, 3, 2},
		{"negative numbers", -2, -3, 1},
		{"zero", 0, 0, 0},
		{"mixed numbers", 2, -3, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Subtract(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Subtract(%v, %v) = %v; want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// TestGetLastResult tests the GetLastResult method
func TestGetLastResult(t *testing.T) {
	calc := &Calculator{}

	// Test after addition
	calc.Add(2, 3)
	if calc.GetLastResult() != 5 {
		t.Errorf("GetLastResult() = %v; want 5", calc.GetLastResult())
	}

	// Test after subtraction
	calc.Subtract(10, 4)
	if calc.GetLastResult() != 6 {
		t.Errorf("GetLastResult() = %v; want 6", calc.GetLastResult())
	}
}

// BenchmarkAdd benchmarks the Add method
func BenchmarkAdd(b *testing.B) {
	calc := &Calculator{}

	for i := 0; i < b.N; i++ {
		calc.Add(2, 3)
	}
}

// Example_calculator demonstrates using the Calculator
func Example_calculator() {
	calc := &Calculator{}

	// Add numbers
	result := calc.Add(2, 3)
	fmt.Printf("2 + 3 = %v\n", result)

	// Subtract numbers
	result = calc.Subtract(10, 4)
	fmt.Printf("10 - 4 = %v\n", result)

	// Get last result
	fmt.Printf("Last result: %v\n", calc.GetLastResult())

	// Output:
	// 2 + 3 = 5
	// 10 - 4 = 6
	// Last result: 6
}

// TestError demonstrates error testing
func TestError(t *testing.T) {
	err := errors.New("test error")
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if err.Error() != "test error" {
		t.Errorf("Expected error message 'test error', got '%s'", err.Error())
	}
}
