package tag_service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/urcane/post-test-golang/exception"
	"github.com/urcane/post-test-golang/helper"
	"github.com/urcane/post-test-golang/model/domain"
	webpost "github.com/urcane/post-test-golang/model/web"
	"github.com/urcane/post-test-golang/repository/tag_repository"
)

type TagServiceImpl struct {
	TagRepository tag_repository.TagRepository
	DB            *sql.DB
	Validate      *validator.Validate
}

func NewTagService(tagRepository tag_repository.TagRepository, db *sql.DB, validate *validator.Validate) TagService {
	return &TagServiceImpl{
		TagRepository: tagRepository,
		DB:            db,
		Validate:      validate,
	}
}

func (service *TagServiceImpl) Create(ctx context.Context, request webpost.TagCreateRequest) webpost.TagResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	tag := domain.Tag{
		Label: request.Label,
	}

	tag = service.TagRepository.Save(ctx, tx, tag)

	return webpost.TagResponse{
		Id:    tag.ID,
		Label: tag.Label,
	}
}

func (service *TagServiceImpl) Update(ctx context.Context, request webpost.TagUpdateRequest) webpost.TagResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	tag := domain.Tag{
		ID:    request.Id,
		Label: request.Label,
	}

	tag = service.TagRepository.Update(ctx, tx, tag)

	return webpost.TagResponse{
		Id:    tag.ID,
		Label: tag.Label,
	}
}

func (service *TagServiceImpl) Delete(ctx context.Context, postId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	post, err := service.TagRepository.FindById(ctx, tx, postId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.TagRepository.Delete(ctx, tx, post)
}

func (service *TagServiceImpl) FindById(ctx context.Context, postId int) webpost.TagResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	tag, err := service.TagRepository.FindById(ctx, tx, postId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return webpost.TagResponse{
		Id:    tag.ID,
		Label: tag.Label,
	}
}

func (service *TagServiceImpl) FindAll(ctx context.Context) []webpost.TagResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	tags := service.TagRepository.FindAll(ctx, tx)

	var tagResponses []webpost.TagResponse
	for _, tag := range tags {
		tagResponses = append(tagResponses,
			webpost.TagResponse{
				Id:    tag.ID,
				Label: tag.Label,
			},
		)
	}
	return tagResponses
}

func (service *TagServiceImpl) CreateWithPost(ctx context.Context, request webpost.TagCreateRequest) webpost.TagWithPostResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	tag := domain.Tag{
		Label: request.Label,
	}

	tag = service.TagRepository.SaveWithPosts(ctx, tx, tag)

	var posts []webpost.PostResponse
	for _, post := range tag.Posts {
		posts = append(posts, webpost.PostResponse{
			Id:          post.ID,
			Title:       post.Title,
			Content:     post.Content,
			Status:      post.Status.Status,
			PublishDate: post.PublishDate,
		})
	}

	return webpost.TagWithPostResponse{
		Id:    tag.ID,
		Label: tag.Label,
		Posts: posts,
	}
}

func (service *TagServiceImpl) FindByIdWithPost(ctx context.Context, postId int) webpost.TagWithPostResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	tag, err := service.TagRepository.FindByIDWithPosts(ctx, tx, postId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	var posts []webpost.PostResponse
	for _, post := range tag.Posts {
		posts = append(posts, webpost.PostResponse{
			Id:          post.ID,
			Title:       post.Title,
			Content:     post.Content,
			Status:      post.Status.Status,
			PublishDate: post.PublishDate,
		})
	}

	return webpost.TagWithPostResponse{
		Id:    tag.ID,
		Label: tag.Label,
		Posts: posts,
	}
}
