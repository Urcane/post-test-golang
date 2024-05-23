package post_repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/urcane/post-test-golang/helper"
	"github.com/urcane/post-test-golang/model/domain"
)

type PostRepositoryImpl struct{}

func NewPostRepository() PostRepository {
	return &PostRepositoryImpl{}
}

func (PostRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, post domain.Post) domain.Post {
	query := "INSERT INTO posts (title,content,status_id,publish_date) VALUES (?, ?, ?, ?, ?)"

	result, err := tx.ExecContext(ctx, query, post.Title, post.Content, post.Status.ID, post.PublishDate)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	post.ID = int(id)
	return post
}

func (PostRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, post domain.Post) domain.Post {
	query := "UPDATE posts SET title = ? ,content = ?,status_id = ?,publish_date = ?) WHERE id = ?"

	_, err := tx.ExecContext(ctx, query, post.Title, post.Content, post.Status.ID, post.PublishDate, post.ID)
	helper.PanicIfError(err)

	return post
}

func (PostRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, post domain.Post) {
	query := "DELETE FROM posts WHERE id = ?"

	_, err := tx.ExecContext(ctx, query, post.ID)
	helper.PanicIfError(err)
}

func (PostRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, postId int) (domain.Post, error) {
	query := "SELECT id, title, content, status_id, publish_date FROM posts WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, postId)
	helper.PanicIfError(err)
	defer rows.Close()

	post := domain.Post{}
	if rows.Next() {
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Status.ID, &post.PublishDate)
		helper.PanicIfError(err)
		return post, nil
	} else {
		return post, errors.New("post is not found")
	}
}

func (PostRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Post {
	query := "SELECT id, name FROM posts"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)
	defer rows.Close()

	var posts []domain.Post
	for rows.Next() {
		post := domain.Post{}
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Status.ID, &post.PublishDate)
		helper.PanicIfError(err)
		posts = append(posts, post)
	}
	return posts
}

func (PostRepositoryImpl) SaveWithTags(ctx context.Context, tx *sql.Tx, post domain.Post) domain.Post {
	post = PostRepositoryImpl{}.Save(ctx, tx, post)

	for _, tag := range post.Tags {
		_, err := tx.ExecContext(ctx, "INSERT INTO post_tags (post_id, tag_id) VALUES (?, ?)", post.ID, tag.ID)
		helper.PanicIfError(err)
	}

	return post
}

func (PostRepositoryImpl) FindByIDWithTags(ctx context.Context, tx *sql.Tx, postID int) (domain.Post, error) {
	post, err := PostRepositoryImpl{}.FindById(ctx, tx, postID)
	if err != nil {
		return domain.Post{}, err
	}

	rows, err := tx.QueryContext(ctx, "SELECT t.id, t.label FROM tags t JOIN post_tags pt ON t.id = pt.tag_id WHERE pt.post_id = ?", postID)
	helper.PanicIfError(err)
	defer rows.Close()

	for rows.Next() {
		tag := domain.Tag{}
		err := rows.Scan(&tag.ID, &tag.Label)
		helper.PanicIfError(err)
		post.Tags = append(post.Tags, &tag)
	}

	return post, nil
}

func (PostRepositoryImpl) FindByTagLabel(ctx context.Context, tx *sql.Tx, tagLabel string) ([]domain.Post, error) {
	query := `
        SELECT p.id, p.title, p.content, p.status_id, p.publish_date,
               t.id, t.label
        FROM posts p
        JOIN post_tags pt ON p.id = pt.post_id
        JOIN tags t ON pt.tag_id = t.id
        WHERE t.label = ?`

	rows, err := tx.QueryContext(ctx, query, tagLabel)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := make(map[int]*domain.Post)
	for rows.Next() {
		var post domain.Post
		var tag domain.Tag
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Status.ID, &post.PublishDate, &tag.ID, &tag.Label)
		if err != nil {
			return nil, err
		}

		// Check if Post is Already Existing, if yes append tags to that post
		if existingPost, ok := posts[post.ID]; ok {
			existingPost.Tags = append(existingPost.Tags, &tag)
		} else {
			post.Tags = []*domain.Tag{&tag}
			posts[post.ID] = &post
		}
	}

	// Convert map to slice
	var result []domain.Post
	for _, post := range posts {
		result = append(result, *post)
	}

	return result, nil
}
