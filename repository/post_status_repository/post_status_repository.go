package post_status_repository

import (
	"context"
	"database/sql"
	"github.com/urcane/post-test-golang/model/domain"
)

type PostStatusRepository interface {
	FindByName(ctx context.Context, tx *sql.Tx, statusName string) (domain.Status, error)
}
