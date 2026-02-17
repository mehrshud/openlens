package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"src/models"
)

// Database represents the database connection
type Database struct {
	Conn *sql.DB
}

// NewDatabase creates a new database connection
func NewDatabase(config models.Config) (*Database, error) {
	db, err := sql.Open("postgres", config.DBURL)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Connected to database")

	return &Database{Conn: db}, nil
}

// CreateUser creates a new user in the database
func (db *Database) CreateUser(user models.User) error {
	_, err := db.Conn.Exec("INSERT INTO users (id, name, email) VALUES ($1, $2, $3)", user.ID, user.Name, user.Email)
	if err != nil {
		return err
	}

	log.Printf("User created: %+v\n", user)

	return nil
}

// GetDashboard retrieves a dashboard from the database
func (db *Database) GetDashboard(id int) (*models.Dashboard, error) {
	var dashboard models.Dashboard

	row := db.Conn.QueryRow("SELECT * FROM dashboards WHERE id = $1", id)

	err := row.Scan(&dashboard.ID, &dashboard.Name)
	if err != nil {
		return nil, err
	}

	log.Printf("Dashboard retrieved: %+v\n", dashboard)

	return &dashboard, nil
}

// Close closes the database connection
func (db *Database) Close() {
	db.Conn.Close()
	log.Println("Database connection closed")
}
