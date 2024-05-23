package post_service

import (
	"context"
	"github.com/urcane/post-test-golang/model/web"
)

type PostService interface {
	Create(ctx context.Context, request web.PostCreateRequest) web.PostResponse
	Update(ctx context.Context, request web.PostUpdateRequest) web.PostResponse
	Delete(ctx context.Context, postId int)
	FindById(ctx context.Context, postId int) web.PostResponse
	FindAll(ctx context.Context) []web.PostResponse
	CreateWithTag(ctx context.Context, request web.PostCreateRequest) web.PostWithTagResponse
	FindByIdWithTag(ctx context.Context, postId int) web.PostWithTagResponse
	FindByTagLabel(ctx context.Context, tagLabel string) []web.PostWithTagResponse
}
