// db/db.go

package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "modernc.org/sqlite" // Import go-sqlite3 library
)

type User struct {
	ID           int
	Username     string
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
	stmt, err := DB.Prepare("INSERT INTO users (email, password_hash) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(email, passwordHash)
	if err != nil {
		return err
	}

	return err
}

// GetUser retrieves a user by username
func GetUser(email string) (*User, error) {
	// User struct to hold the data
	var user User

	// SQL query to select the user by username
	query := `SELECT id, username, email, password_hash, creation_time FROM users WHERE email = ?`

	// Query the database
	row := DB.QueryRow(query, email)
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreationTime)
	if err != nil {
		if err == sql.ErrNoRows {
			// No result (not an error in this context)
			return nil, nil
		}
		log.Printf("Error querying for user by username: %v", err)
		return nil, err
	}

	return &user, nil
}
