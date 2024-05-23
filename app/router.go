package app

import (
	"github.com/julienschmidt/httprouter"
	"github.com/urcane/post-test-golang/controller"
	"github.com/urcane/post-test-golang/exception"
)

func NewRouter(userController controller.UserController, postController controller.PostController, tagController controller.TagController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/register", userController.Register)
	router.POST("/api/login", userController.Login)

	router.GET("/api/posts", postController.FindAll)
	router.GET("/api/posts/:postId", postController.FindById)
	router.POST("/api/posts", postController.Create)
	router.PUT("/api/posts/:postId", postController.Update)
	router.DELETE("/api/posts/:postId", postController.Delete)

	router.GET("/api/tags", tagController.FindAll)
	router.GET("/api/tags/:tagId", tagController.FindById)
	router.POST("/api/tags", tagController.Create)
	router.PUT("/api/tags/:tagId", tagController.Update)
	router.DELETE("/api/tags/:tagId", tagController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
