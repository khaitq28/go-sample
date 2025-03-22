package main

/*
Database Tutorial
----------------
This tutorial covers database operations in Go:
1. Basic Database Operations
   - Connection setup
   - CRUD operations
   - Transactions
2. Advanced Features
   - Prepared statements
   - Connection pooling
   - Migration tools
3. Best Practices
   - Error handling
   - Connection management
   - Query optimization
*/

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// User represents a user in the database
type User struct {
	ID        int64
	Name      string
	Email     string
	CreatedAt time.Time
}

// Database handles database operations
type Database struct {
	db *sql.DB
}

// NewDatabase creates a new database connection
func NewDatabase(dbPath string) (*Database, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	return &Database{db: db}, nil
}

// InitSchema creates the necessary tables
func (d *Database) InitSchema() error {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			email TEXT UNIQUE NOT NULL,
			created_at DATETIME NOT NULL
		)
	`

	_, err := d.db.Exec(query)
	return err
}

// CreateUser adds a new user to the database
func (d *Database) CreateUser(user *User) error {
	query := `
		INSERT INTO users (name, email, created_at)
		VALUES (?, ?, ?)
	`

	result, err := d.db.Exec(query, user.Name, user.Email, user.CreatedAt)
	if err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error getting last insert id: %v", err)
	}

	user.ID = id
	return nil
}

// GetUser retrieves a user by ID
func (d *Database) GetUser(id int64) (*User, error) {
	query := `
		SELECT id, name, email, created_at
		FROM users
		WHERE id = ?
	`

	user := &User{}
	err := d.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("error getting user: %v", err)
	}

	return user, nil
}

// UpdateUser updates an existing user
func (d *Database) UpdateUser(user *User) error {
	query := `
		UPDATE users
		SET name = ?, email = ?
		WHERE id = ?
	`

	result, err := d.db.Exec(query, user.Name, user.Email, user.ID)
	if err != nil {
		return fmt.Errorf("error updating user: %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %v", err)
	}

	if rows == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

// DeleteUser removes a user from the database
func (d *Database) DeleteUser(id int64) error {
	query := `DELETE FROM users WHERE id = ?`

	result, err := d.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting user: %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %v", err)
	}

	if rows == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

// ListUsers retrieves all users from the database
func (d *Database) ListUsers() ([]User, error) {
	query := `
		SELECT id, name, email, created_at
		FROM users
		ORDER BY created_at DESC
	`

	rows, err := d.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying users: %v", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning user: %v", err)
		}
		users = append(users, user)
	}

	return users, nil
}

// TransactionExample demonstrates using transactions
func (d *Database) TransactionExample() error {
	tx, err := d.db.Begin()
	if err != nil {
		return fmt.Errorf("error beginning transaction: %v", err)
	}
	defer tx.Rollback()

	// Insert first user
	_, err = tx.Exec(`
		INSERT INTO users (name, email, created_at)
		VALUES (?, ?, ?)
	`, "Alice", "alice@example.com", time.Now())
	if err != nil {
		return fmt.Errorf("error inserting first user: %v", err)
	}

	// Insert second user
	_, err = tx.Exec(`
		INSERT INTO users (name, email, created_at)
		VALUES (?, ?, ?)
	`, "Bob", "bob@example.com", time.Now())
	if err != nil {
		return fmt.Errorf("error inserting second user: %v", err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}

	return nil
}

func main() {
	// Create database connection
	db, err := NewDatabase("users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.db.Close()

	// Initialize schema
	if err := db.InitSchema(); err != nil {
		log.Fatal(err)
	}

	// Create a user
	user := &User{
		Name:      "John Doe",
		Email:     "john@example.com",
		CreatedAt: time.Now(),
	}

	if err := db.CreateUser(user); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created user: %+v\n", user)

	// Get the user
	retrievedUser, err := db.GetUser(user.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Retrieved user: %+v\n", retrievedUser)

	// Update the user
	user.Name = "John Smith"
	if err := db.UpdateUser(user); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated user: %+v\n", user)

	// List all users
	users, err := db.ListUsers()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All users:")
	for _, u := range users {
		fmt.Printf("- %+v\n", u)
	}

	// Demonstrate transaction
	if err := db.TransactionExample(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Transaction completed successfully")
}
