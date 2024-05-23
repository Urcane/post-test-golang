package tag_repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/urcane/post-test-golang/helper"
	"github.com/urcane/post-test-golang/model/domain"
)

type TagRepositoryImpl struct{}

func NewTagRepository() TagRepository {
	return &TagRepositoryImpl{}
}

func (TagRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, tag domain.Tag) domain.Tag {
	query := "INSERT INTO tags (label) VALUES (?)"

	result, err := tx.ExecContext(ctx, query, tag.Label)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	tag.ID = int(id)
	return tag
}

func (TagRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, tag domain.Tag) domain.Tag {
	query := "UPDATE tags SET label = ? WHERE id = ?"

	_, err := tx.ExecContext(ctx, query, tag.Label, tag.ID)
	helper.PanicIfError(err)

	return tag
}

func (TagRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, tag domain.Tag) {
	query := "DELETE FROM tags WHERE id = ?"

	_, err := tx.ExecContext(ctx, query, tag.ID)
	helper.PanicIfError(err)
}

func (TagRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, tagId int) (domain.Tag, error) {
	query := "SELECT id, label FROM tags WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, tagId)
	helper.PanicIfError(err)
	defer rows.Close()

	tag := domain.Tag{}
	if rows.Next() {
		err := rows.Scan(&tag.ID, &tag.Label)
		helper.PanicIfError(err)
		return tag, nil
	} else {
		return tag, errors.New("tag is not found")
	}
}

func (TagRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Tag {
	query := "SELECT id, label FROM tags"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)
	defer rows.Close()

	var tags []domain.Tag
	for rows.Next() {
		tag := domain.Tag{}
		err := rows.Scan(&tag.ID, &tag.Label)
		helper.PanicIfError(err)
		tags = append(tags, tag)
	}
	return tags
}

func (TagRepositoryImpl) SaveWithPosts(ctx context.Context, tx *sql.Tx, tag domain.Tag) domain.Tag {
	// Save the Tag
	tag = TagRepositoryImpl{}.Save(ctx, tx, tag)

	// Link Tag with Posts
	for _, post := range tag.Posts {
		_, err := tx.ExecContext(ctx, "INSERT INTO post_tags (post_id, tag_id) VALUES (?, ?)", post.ID, tag.ID)
		helper.PanicIfError(err)
	}

	return tag
}

func (TagRepositoryImpl) FindByIDWithPosts(ctx context.Context, tx *sql.Tx, tagID int) (domain.Tag, error) {
	// Find the Tag
	tag, err := TagRepositoryImpl{}.FindById(ctx, tx, tagID)
	if err != nil {
		return domain.Tag{}, err
	}

	// Find Posts associated with the Tag
	rows, err := tx.QueryContext(ctx, "SELECT p.id, p.title, p.content, p.status_id, p.publish_date FROM posts p JOIN post_tags pt ON p.id = pt.post_id WHERE pt.tag_id = ?", tagID)
	helper.PanicIfError(err)
	defer rows.Close()

	for rows.Next() {
		var post domain.Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Status.ID, &post.PublishDate)
		helper.PanicIfError(err)
		tag.Posts = append(tag.Posts, &post)
	}

	return tag, nil
}
