package migration

import (
	"database/sql"
	"fmt"
	"github.com/urcane/post-test-golang/helper"
)

func CreateTagsTable(db *sql.DB) {
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS tags (
			id SERIAL PRIMARY KEY,
			label TEXT NOT NULL
		);
	`

	_, err := db.Exec(createTableQuery)
	helper.PanicIfError(err)

	fmt.Println("Tag Migration completed successfully!")
}
