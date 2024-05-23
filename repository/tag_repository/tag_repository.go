package tag_repository

import (
	"context"
	"database/sql"
	"github.com/urcane/post-test-golang/model/domain"
)

type TagRepository interface {
	Save(ctx context.Context, tx *sql.Tx, tag domain.Tag) domain.Tag
	Update(ctx context.Context, tx *sql.Tx, tag domain.Tag) domain.Tag
	Delete(ctx context.Context, tx *sql.Tx, tag domain.Tag)
	FindById(ctx context.Context, tx *sql.Tx, tagId int) (domain.Tag, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Tag
	SaveWithPosts(ctx context.Context, tx *sql.Tx, tag domain.Tag) domain.Tag
	FindByIDWithPosts(ctx context.Context, tx *sql.Tx, tagID int) (domain.Tag, error)
}
