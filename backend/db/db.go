// db/db.go

package db

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "modernc.org/sqlite" // Import go-sqlite3 library
)

type User struct {
	ID           int
	Username     *string
	Email        string
	PasswordHash string
	CreationTime time.Time
}

var DB *sql.DB

// InitDB initializes the DB variable
func InitDB(dataSourceName string) error {
	var err error
	DB, err = sql.Open("sqlite", dataSourceName)
	if err != nil {
		return fmt.Errorf("error opening database: %w", err)
	}

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("error pinging database: %w", err)
	}

	return nil
}

// CloseDB closes the database connection
func CloseDB() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}

// AddUser adds a new user to the database
func AddUser(email string, passwordHash string) error {
	query := `INSERT INTO users (email, password_hash) VALUES (?, ?)`

	stmt, err := DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(email, passwordHash)
	if err != nil {
		return fmt.Errorf("error executing statement: %w", err)
	}

	return nil
}

// GetUser retrieves a user by username
func GetUser(identifier string) (*User, error) {
	// User struct to hold the data
	var user User

	// SQL query to select the user by username
	query := `SELECT id, username, email, password_hash, creation_time FROM users WHERE email = ? OR id = ?`

	// Execute the query
	err := DB.QueryRow(query, identifier, identifier).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreationTime)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// No result (not an error in this context)
			return nil, fmt.Errorf("user not found: %w", err)
		}
		return nil, fmt.Errorf("error querying for user: %w", err)
	}

	return &user, nil
}
