package migration

import (
	"database/sql"
	"fmt"
	"github.com/urcane/post-test-golang/helper"
)

func CreatePostTable(db *sql.DB) {
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS posts (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			content TEXT NOT NULL,
			status_id INT NOT NULL REFERENCES post_statuses(id),
			publish_date TIMESTAMP
		);

		CREATE INDEX idx_post_publish_date ON posts (publish_date);
		CREATE INDEX idx_post_title ON posts (title);
	`

	_, err := db.Exec(createTableQuery)
	helper.PanicIfError(err)

	fmt.Println("Post Migration completed successfully!")
}
