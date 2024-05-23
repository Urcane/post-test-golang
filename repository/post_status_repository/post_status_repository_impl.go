package post_status_repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/urcane/post-test-golang/helper"
	"github.com/urcane/post-test-golang/model/domain"
)

type PostStatusRepositoryImpl struct{}

func NewPostStatusRepository() PostStatusRepository {
	return &PostStatusRepositoryImpl{}
}

func (PostStatusRepositoryImpl) FindByName(ctx context.Context, tx *sql.Tx, statusName string) (domain.Status, error) {
	SQL := "SELECT id, title, content, status_id, publish_date FROM post_statuses WHERE status = ?"
	rows, err := tx.QueryContext(ctx, SQL, statusName)
	helper.PanicIfError(err)
	defer rows.Close()

	status := domain.Status{}
	if rows.Next() {
		err := rows.Scan(&status.ID, &status.Status)
		helper.PanicIfError(err)
		return status, nil
	} else {
		return status, errors.New("status is not found")
	}
}
