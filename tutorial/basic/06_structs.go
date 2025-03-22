package main

/*
Structs and Methods Tutorial
---------------------------
This tutorial covers structs and methods in Go:
1. Basic Structs
   - Struct definition
   - Struct initialization
   - Struct fields
2. Advanced Struct Features
   - Embedded structs
   - Struct tags
   - Anonymous structs
3. Methods
   - Method declaration
   - Pointer receivers
   - Method chaining
4. Interfaces
   - Interface definition
   - Interface implementation
   - Interface satisfaction
*/

import (
	"fmt"
	"math"
)

// Basic struct
type Person struct {
	Name    string
	Age     int
	Address string
}

// Struct with embedded struct
type Employee struct {
	Person
	Salary   float64
	Position string
}

// Struct with tags
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"` // "-" means this field will be omitted in JSON
}

// Method on Person struct
func (p Person) SayHello() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

// Method with pointer receiver
func (p *Person) HaveBirthday() {
	p.Age++
	fmt.Printf("Happy Birthday! Now I am %d years old.\n", p.Age)
}

// Method on Employee struct
func (e Employee) GetInfo() {
	fmt.Printf("Name: %s\nAge: %d\nPosition: %s\nSalary: $%.2f\n",
		e.Name, e.Age, e.Position, e.Salary)
}

// Interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Struct implementing Shape interface
type Rectangle struct {
	Width  float64
	Height float64
}

// Methods for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Another struct implementing Shape interface
type Circle struct {
	Radius float64
}

// Methods for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	// 1. Creating and using basic struct
	fmt.Println("Basic Struct:")
	person := Person{
		Name:    "Alice",
		Age:     25,
		Address: "123 Main St",
	}
	fmt.Printf("Person: %+v\n", person)

	// 2. Using struct methods
	fmt.Println("\nStruct Methods:")
	person.SayHello()
	person.HaveBirthday()

	// 3. Embedded struct
	fmt.Println("\nEmbedded Struct:")
	employee := Employee{
		Person: Person{
			Name:    "Bob",
			Age:     30,
			Address: "456 Oak St",
		},
		Salary:   75000,
		Position: "Developer",
	}
	employee.GetInfo()

	// 4. Struct with tags
	fmt.Println("\nStruct with Tags:")
	user := User{
		ID:       1,
		Username: "john_doe",
		Email:    "john@example.com",
		Password: "secret123",
	}
	fmt.Printf("User: %+v\n", user)

	// 5. Interface implementation
	fmt.Println("\nInterface Implementation:")

	// Create shapes
	rect := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 2.5}

	// Use shapes through interface
	shapes := []Shape{rect, circle}

	for _, shape := range shapes {
		fmt.Printf("Area: %.2f\n", shape.Area())
		fmt.Printf("Perimeter: %.2f\n", shape.Perimeter())
		fmt.Println()
	}

	// 6. Struct comparison
	fmt.Println("\nStruct Comparison:")
	p1 := Person{Name: "Alice", Age: 25}
	p2 := Person{Name: "Alice", Age: 25}
	p3 := Person{Name: "Bob", Age: 30}

	fmt.Printf("p1 == p2: %v\n", p1 == p2)
	fmt.Printf("p1 == p3: %v\n", p1 == p3)

	// 7. Anonymous struct
	fmt.Println("\nAnonymous Struct:")
	car := struct {
		Make  string
		Model string
		Year  int
	}{
		Make:  "Toyota",
		Model: "Camry",
		Year:  2020,
	}
	fmt.Printf("Car: %+v\n", car)
}
