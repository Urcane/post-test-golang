package migration

import (
	"database/sql"
	"fmt"
	"github.com/urcane/post-test-golang/helper"
)

func CreatePostTagsTable(db *sql.DB) {
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS post_tags (
			post_id INT REFERENCES posts(id) ON DELETE CASCADE,
			tag_id INT REFERENCES tags(id) ON DELETE CASCADE,
			PRIMARY KEY (post_id, tag_id)
		);
		CREATE INDEX idx_post_tags_post_id ON post_tags (post_id);
		CREATE INDEX idx_post_tags_tag_id ON post_tags (tag_id);
	`

	_, err := db.Exec(createTableQuery)
	helper.PanicIfError(err)

	fmt.Println("Post Tag Migration completed successfully!")
}
