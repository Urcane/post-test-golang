package user_repository

import (
	"context"
	"database/sql"
	"github.com/urcane/post-test-golang/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	FindByUsername(ctx context.Context, tx *sql.Tx, username string) (domain.User, error)
	UpdateToken(ctx context.Context, tx *sql.Tx, userID int, token string)
}
