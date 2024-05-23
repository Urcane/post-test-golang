package database

import (
	migration "github.com/urcane/post-test-golang/database/migration"
	"github.com/urcane/post-test-golang/database/seeder"
)

func Migrate() {
	db := NewDB()

	// Register the migration
	migration.CreateUserTable(db)
	migration.CreatePostStatusesTable(db)
	migration.CreatePostTable(db)
	migration.CreateTagsTable(db)
	migration.CreatePostTagsTable(db)

	//Register the seeder
	seeder.PostStatusSeeder(db)

	defer db.Close()
}
