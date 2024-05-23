package post_service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/urcane/post-test-golang/exception"
	"github.com/urcane/post-test-golang/helper"
	"github.com/urcane/post-test-golang/model/domain"
	"github.com/urcane/post-test-golang/model/web"
	"github.com/urcane/post-test-golang/repository/post_repository"
	"github.com/urcane/post-test-golang/repository/post_status_repository"
)

type PostServiceImpl struct {
	PostRepository post_repository.PostRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewPostService(postRepository post_repository.PostRepository, db *sql.DB, validate *validator.Validate) PostService {
	return &PostServiceImpl{
		PostRepository: postRepository,
		DB:             db,
		Validate:       validate,
	}
}

func (service *PostServiceImpl) Create(ctx context.Context, request web.PostCreateRequest) web.PostResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	status, err := post_status_repository.NewPostStatusRepository().FindByName(ctx, tx, request.Status)
	helper.PanicIfError(err)

	post := domain.Post{
		Title:       request.Title,
		Content:     request.Content,
		Status:      &status,
		PublishDate: request.PublishDate,
	}

	post = service.PostRepository.Save(ctx, tx, post)

	return web.PostResponse{
		Id:          post.ID,
		Title:       post.Title,
		Content:     post.Content,
		Status:      post.Status.Status,
		PublishDate: post.PublishDate,
	}
}

func (service *PostServiceImpl) Update(ctx context.Context, request web.PostUpdateRequest) web.PostResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	status, err := post_status_repository.NewPostStatusRepository().FindByName(ctx, tx, request.Status)
	helper.PanicIfError(err)

	post := domain.Post{
		ID:          request.Id,
		Title:       request.Title,
		Content:     request.Content,
		Status:      &status,
		PublishDate: request.PublishDate,
	}

	post = service.PostRepository.Update(ctx, tx, post)

	return web.PostResponse{
		Id:          post.ID,
		Title:       post.Title,
		Content:     post.Content,
		Status:      post.Status.Status,
		PublishDate: post.PublishDate,
	}
}

func (service *PostServiceImpl) Delete(ctx context.Context, postId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	post, err := service.PostRepository.FindById(ctx, tx, postId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.PostRepository.Delete(ctx, tx, post)
}

func (service *PostServiceImpl) FindById(ctx context.Context, postId int) web.PostResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	post, err := service.PostRepository.FindById(ctx, tx, postId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return web.PostResponse{
		Id:          post.ID,
		Title:       post.Title,
		Content:     post.Content,
		Status:      post.Status.Status,
		PublishDate: post.PublishDate,
	}
}

func (service *PostServiceImpl) FindAll(ctx context.Context) []web.PostResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	posts := service.PostRepository.FindAll(ctx, tx)

	var postResponses []web.PostResponse
	for _, post := range posts {
		postResponses = append(postResponses,
			web.PostResponse{
				Id:          post.ID,
				Title:       post.Title,
				Content:     post.Content,
				Status:      post.Status.Status,
				PublishDate: post.PublishDate,
			},
		)
	}
	return postResponses
}

func (service *PostServiceImpl) CreateWithTag(ctx context.Context, request web.PostCreateRequest) web.PostWithTagResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	status, err := post_status_repository.NewPostStatusRepository().FindByName(ctx, tx, request.Status)
	helper.PanicIfError(err)

	post := domain.Post{
		Title:       request.Title,
		Content:     request.Content,
		Status:      &status,
		PublishDate: request.PublishDate,
	}

	post = service.PostRepository.SaveWithTags(ctx, tx, post)

	var tags []web.TagResponse
	for _, tag := range post.Tags {
		tags = append(tags, web.TagResponse{
			Id:    tag.ID,
			Label: tag.Label,
		})
	}

	return web.PostWithTagResponse{
		Id:          post.ID,
		Title:       post.Title,
		Content:     post.Content,
		Status:      post.Status.Status,
		PublishDate: post.PublishDate,
		Tags:        tags,
	}
}

func (service *PostServiceImpl) FindByIdWithTag(ctx context.Context, postId int) web.PostWithTagResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	post, err := service.PostRepository.FindByIDWithTags(ctx, tx, postId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	var tags []web.TagResponse
	for _, tag := range post.Tags {
		tags = append(tags, web.TagResponse{
			Id:    tag.ID,
			Label: tag.Label,
		})
	}

	return web.PostWithTagResponse{
		Id:          post.ID,
		Title:       post.Title,
		Content:     post.Content,
		Status:      post.Status.Status,
		PublishDate: post.PublishDate,
		Tags:        tags,
	}
}

func (service *PostServiceImpl) FindByTagLabel(ctx context.Context, tagLabel string) []web.PostWithTagResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	posts, err := service.PostRepository.FindByTagLabel(ctx, tx, tagLabel)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	var postResponses []web.PostWithTagResponse
	for _, post := range posts {
		var tags []web.TagResponse
		for _, tag := range post.Tags {
			tags = append(tags, web.TagResponse{
				Id:    tag.ID,
				Label: tag.Label,
			})
		}

		postResponses = append(postResponses, web.PostWithTagResponse{
			Id:          post.ID,
			Title:       post.Title,
			Content:     post.Content,
			Status:      post.Status.Status,
			PublishDate: post.PublishDate,
			Tags:        tags,
		})
	}
	return postResponses
}
