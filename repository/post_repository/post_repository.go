package post_repository

import (
	"context"
	"database/sql"
	"github.com/urcane/post-test-golang/model/domain"
)

type PostRepository interface {
	Save(ctx context.Context, tx *sql.Tx, post domain.Post) domain.Post
	Update(ctx context.Context, tx *sql.Tx, post domain.Post) domain.Post
	Delete(ctx context.Context, tx *sql.Tx, post domain.Post)
	FindById(ctx context.Context, tx *sql.Tx, postId int) (domain.Post, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Post
	SaveWithTags(ctx context.Context, tx *sql.Tx, post domain.Post) domain.Post
	FindByIDWithTags(ctx context.Context, tx *sql.Tx, postID int) (domain.Post, error)
	FindByTagLabel(ctx context.Context, tx *sql.Tx, tagLabel string) ([]domain.Post, error)
}
