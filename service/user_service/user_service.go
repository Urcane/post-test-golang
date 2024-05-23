package user_service

import (
	"context"
	webpost "github.com/urcane/post-test-golang/model/web"
)

type TagService interface {
	Create(ctx context.Context, request webpost.TagCreateRequest) webpost.TagResponse
	Update(ctx context.Context, request webpost.TagUpdateRequest) webpost.TagResponse
	Delete(ctx context.Context, tagId int)
	FindById(ctx context.Context, tagId int) webpost.TagResponse
	FindAll(ctx context.Context) []webpost.TagResponse
	CreateWithPost(ctx context.Context, request webpost.TagCreateRequest) webpost.TagWithPostResponse
	FindByIdWithPost(ctx context.Context, tagId int) webpost.TagWithPostResponse
}
