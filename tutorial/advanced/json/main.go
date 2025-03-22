package main

/*
JSON Tutorial
------------
This tutorial covers JSON handling in Go:
1. Basic JSON Operations
   - Marshaling
   - Unmarshaling
   - Custom types
2. Advanced Features
   - Custom marshaling
   - JSON tags
   - Raw messages
3. Best Practices
   - Error handling
   - Validation
   - Performance
*/

import (
	"encoding/json"
	"fmt"
	"time"
)

// User represents a user in our system
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// CustomTime demonstrates custom JSON marshaling
type CustomTime struct {
	time.Time
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, ct.Format("2006-01-02 15:04:05"))), nil
}

func (ct *CustomTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	now, err := time.ParseInLocation(`"2006-01-02 15:04:05"`, string(data), time.Local)
	*ct = CustomTime{now}
	return err
}

// Product represents a product with custom marshaling
type Product struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Price       float64    `json:"price"`
	CreatedAt   time.Time  `json:"created_at"`
	CustomTime  CustomTime `json:"custom_time"`
	IsAvailable bool       `json:"is_available"`
}

// Order represents a complex nested structure
type Order struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Items     []Item    `json:"items"`
	Total     float64   `json:"total"`
	CreatedAt time.Time `json:"created_at"`
}

type Item struct {
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

// RawMessageExample demonstrates using json.RawMessage
type RawMessageExample struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

func main() {
	// Basic marshaling
	user := User{
		ID:        1,
		Name:      "John Doe",
		Email:     "john@example.com",
		CreatedAt: time.Now(),
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("Error marshaling user: %v\n", err)
		return
	}
	fmt.Printf("Marshaled user: %s\n", string(userJSON))

	// Basic unmarshaling
	var unmarshaledUser User
	if err := json.Unmarshal(userJSON, &unmarshaledUser); err != nil {
		fmt.Printf("Error unmarshaling user: %v\n", err)
		return
	}
	fmt.Printf("Unmarshaled user: %+v\n", unmarshaledUser)

	// Custom type marshaling
	product := Product{
		ID:          1,
		Name:        "Laptop",
		Price:       999.99,
		CreatedAt:   time.Now(),
		CustomTime:  CustomTime{time.Now()},
		IsAvailable: true,
	}

	productJSON, err := json.Marshal(product)
	if err != nil {
		fmt.Printf("Error marshaling product: %v\n", err)
		return
	}
	fmt.Printf("Marshaled product: %s\n", string(productJSON))

	// Complex nested structure
	order := Order{
		ID:     1,
		UserID: 1,
		Items: []Item{
			{ProductID: 1, Quantity: 2, Price: 999.99},
			{ProductID: 2, Quantity: 1, Price: 49.99},
		},
		Total:     2049.97,
		CreatedAt: time.Now(),
	}

	orderJSON, err := json.Marshal(order)
	if err != nil {
		fmt.Printf("Error marshaling order: %v\n", err)
		return
	}
	fmt.Printf("Marshaled order: %s\n", string(orderJSON))

	// Raw message example
	rawJSON := `{
		"type": "user",
		"payload": {
			"id": 1,
			"name": "John Doe",
			"email": "john@example.com"
		}
	}`

	var rawExample RawMessageExample
	if err := json.Unmarshal([]byte(rawJSON), &rawExample); err != nil {
		fmt.Printf("Error unmarshaling raw message: %v\n", err)
		return
	}

	fmt.Printf("Type: %s\n", rawExample.Type)
	fmt.Printf("Raw payload: %s\n", string(rawExample.Payload))

	// Pretty printing
	prettyJSON, err := json.MarshalIndent(order, "", "  ")
	if err != nil {
		fmt.Printf("Error pretty printing: %v\n", err)
		return
	}
	fmt.Printf("Pretty printed order:\n%s\n", string(prettyJSON))
}
