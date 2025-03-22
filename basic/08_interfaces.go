package main

/*
Interfaces Tutorial
------------------
This tutorial covers interfaces in Go:
1. Basic Interfaces
   - Interface definition
   - Interface implementation
   - Interface satisfaction
2. Advanced Interface Features
   - Interface composition
   - Empty interfaces
   - Type assertions
3. Interface Patterns
   - Interface as function parameters
   - Interface type switches
   - Interface best practices
4. Common Interfaces
   - Reader/Writer interfaces
   - Stringer interface
   - Error interface
*/

import (
	"fmt"
	"math"
)

// Basic interface
type Speaker interface {
	Speak() string
}

// Interface with multiple methods
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Interface composition
type ReadWriter interface {
	Reader
	Writer
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

// Struct implementing Speaker interface
type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

// Struct implementing Speaker interface
type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return "Meow!"
}

// Struct implementing Shape interface
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Struct implementing Shape interface
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Interface type assertion
func describe(i interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", i, i)
}

// Interface with type switch
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	// 1. Basic interface implementation
	fmt.Println("Basic Interface:")
	dog := Dog{Name: "Rex"}
	cat := Cat{Name: "Whiskers"}

	var speaker Speaker
	speaker = dog
	fmt.Printf("%s says: %s\n", dog.Name, speaker.Speak())

	speaker = cat
	fmt.Printf("%s says: %s\n", cat.Name, speaker.Speak())

	// 2. Interface with multiple methods
	fmt.Println("\nShape Interface:")
	rect := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 2.5}

	shapes := []Shape{rect, circle}
	for _, shape := range shapes {
		fmt.Printf("Area: %.2f, Perimeter: %.2f\n",
			shape.Area(), shape.Perimeter())
	}

	// 3. Empty interface
	fmt.Println("\nEmpty Interface:")
	var i interface{}
	i = 42
	describe(i)

	i = "hello"
	describe(i)

	// 4. Type assertion
	fmt.Println("\nType Assertion:")
	var s interface{} = "hello"
	str, ok := s.(string)
	if ok {
		fmt.Printf("String value: %s\n", str)
	}

	// 5. Type switch
	fmt.Println("\nType Switch:")
	do(21)
	do("hello")
	do(true)

	// 6. Interface satisfaction
	fmt.Println("\nInterface Satisfaction:")
	var _ Speaker = (*Dog)(nil) // Compile-time check
	var _ Speaker = (*Cat)(nil) // Compile-time check

	// 7. Interface as function parameter
	fmt.Println("\nInterface as Parameter:")
	printArea(rect)
	printArea(circle)

	// 8. Interface composition
	fmt.Println("\nInterface Composition:")
	var rw ReadWriter // This would need a concrete implementation
	_ = rw            // Use the variable to avoid unused variable error
}

// Function accepting interface
func printArea(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
}
