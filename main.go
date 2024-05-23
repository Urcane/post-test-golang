package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/urcane/post-test-golang/app"
	"github.com/urcane/post-test-golang/controller"
	"github.com/urcane/post-test-golang/database"
	"github.com/urcane/post-test-golang/helper"
	"github.com/urcane/post-test-golang/middleware"
	"github.com/urcane/post-test-golang/repository/post_repository"
	"github.com/urcane/post-test-golang/repository/tag_repository"
	"github.com/urcane/post-test-golang/repository/user_repository"
	"github.com/urcane/post-test-golang/service/post_service"
	"github.com/urcane/post-test-golang/service/tag_service"
	"github.com/urcane/post-test-golang/service/user_service"
	"net/http"
)

func main() {
	db := database.NewDB()
	validate := validator.New()

	userRepository := user_repository.NewUserRepository()
	userService := user_service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	postRepository := post_repository.NewPostRepository()
	postService := post_service.NewPostService(postRepository, db, validate)
	postController := controller.NewPostController(postService)

	tagRepository := tag_repository.NewTagRepository()
	tagService := tag_service.NewTagService(tagRepository, db, validate)
	tagController := controller.NewTagController(tagService)

	router := app.NewRouter(userController, postController, tagController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
