package seeder

import (
	"database/sql"
	"fmt"
	"github.com/urcane/post-test-golang/helper"
)

func PostStatusSeeder(db *sql.DB) {
	createTableQuery := `
		INSERT INTO post_statuses (status) VALUES ('draft'), ('publish')
	`

	_, err := db.Exec(createTableQuery)
	helper.PanicIfError(err)

	fmt.Println("Post Status Seeder completed successfully!")
}
