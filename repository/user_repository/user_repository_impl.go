package user_repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/urcane/post-test-golang/helper"
	"github.com/urcane/post-test-golang/model/domain"
)

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	query := "INSERT INTO users (username, password, email, token) VALUES (?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, query, user.Username, user.Password, user.Email, user.Token)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.ID = int(id)
	return user
}

func (UserRepositoryImpl) FindByUsername(ctx context.Context, tx *sql.Tx, username string) (domain.User, error) {
	query := "SELECT id, username, password, email, token FROM users WHERE username = ?"
	rows, err := tx.QueryContext(ctx, query, username)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Token)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

func (UserRepositoryImpl) UpdateToken(ctx context.Context, tx *sql.Tx, userID int, token string) {
	query := "UPDATE users SET token = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, token, userID)
	helper.PanicIfError(err)
}
