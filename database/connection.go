package database

import (
	"database/sql"
	"fmt"
	"github.com/urcane/post-test-golang/app"
	"github.com/urcane/post-test-golang/helper"
	"time"

	_ "github.com/lib/pq"
)

func NewDB() *sql.DB {
	config := *app.NewConfig()

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DBName,
	)

	db, err := sql.Open("postgres", psqlInfo)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
