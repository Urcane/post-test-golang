package migration

import (
	"database/sql"
	"fmt"
	"github.com/urcane/post-test-golang/helper"
)

func CreatePostStatusesTable(db *sql.DB) {
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS post_statuses (
			id SERIAL PRIMARY KEY,
			status TEXT NOT NULL CHECK (status IN ('draft', 'publish'))
		);
		CREATE INDEX idx_post_statuses_status ON post_statuses (status);
	`

	_, err := db.Exec(createTableQuery)
	helper.PanicIfError(err)

	fmt.Println("Post Status Migration completed successfully!")
}
