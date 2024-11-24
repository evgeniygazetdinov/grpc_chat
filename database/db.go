package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

const (
	host     = "localhost"
	port     = 45432
	user     = "postgres"
	password = "example"
	dbname   = "postgres"
)

func InitDB() error {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("error connecting to the database: %v", err)
	}
	if err != nil {
		return fmt.Errorf("error creating table: %v", err)
	}

	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			gender VARCHAR(1) NOT NULL,
			name VARCHAR(255) NOT NULL,
			location TEXT NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL,
			app_version VARCHAR(50) NOT NULL,
			bio TEXT,
			token TEXT,
			user_headers TEXT,
			id_user INTEGER,
			interview BOOLEAN DEFAULT false,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return fmt.Errorf("error creating users table: %v", err)
	}

	log.Println("Database connection established successfully")
	return nil
}

// StoreUser stores a user in the database
func StoreUser(gender, name, location, email, appVersion, bio, token, userHeaders string, idUser int, interview bool) error {
	query := `
		INSERT INTO users (gender, name, location, email, app_version, bio, token, user_headers, id_user, interview)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		ON CONFLICT (email) DO UPDATE SET
			gender = EXCLUDED.gender,
			name = EXCLUDED.name,
			location = EXCLUDED.location,
			app_version = EXCLUDED.app_version,
			bio = EXCLUDED.bio,
			token = EXCLUDED.token,
			user_headers = EXCLUDED.user_headers,
			id_user = EXCLUDED.id_user,
			interview = EXCLUDED.interview
		RETURNING id`

	var id int
	err := DB.QueryRow(query,
		gender, name, location, email, appVersion,
		bio, token, userHeaders, idUser, interview,
	).Scan(&id)

	if err != nil {
		return fmt.Errorf("error storing user: %v", err)
	}

	return nil
}
