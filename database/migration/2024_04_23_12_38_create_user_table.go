package migration

import (
	"database/sql"
	"fmt"
	"github.com/urcane/post-test-golang/helper"
)

func CreateUserTable(db *sql.DB) {
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(50) NOT NULL UNIQUE,
			password_hash TEXT NOT NULL,
			email VARCHAR(100) NOT NULL UNIQUE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
		CREATE INDEX idx_users_username ON users (username);
		CREATE INDEX idx_users_email ON users (email);
	`

	_, err := db.Exec(createTableQuery)
	helper.PanicIfError(err)

	fmt.Println("User Migration completed successfully!")
}
