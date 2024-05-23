package controller

import (
	"github.com/julienschmidt/httprouter"
	"github.com/urcane/post-test-golang/helper"
	"github.com/urcane/post-test-golang/model/web"
	"github.com/urcane/post-test-golang/service/post_service"
	"net/http"
	"strconv"
)

type PostController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type PostControllerImpl struct {
	PostService post_service.PostService
}

func NewPostController(postService post_service.PostService) PostController {
	return &PostControllerImpl{
		PostService: postService,
	}
}

func (controller *PostControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	postCreateRequest := web.PostCreateRequest{}
	helper.ReadFromRequestBody(request, &postCreateRequest)

	postResponse := controller.PostService.CreateWithTag(request.Context(), postCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   postResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PostControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	postUpdateRequest := web.PostUpdateRequest{}
	helper.ReadFromRequestBody(request, &postUpdateRequest)

	postId := params.ByName("postId")
	id, err := strconv.Atoi(postId)
	helper.PanicIfError(err)

	postUpdateRequest.Id = id

	postResponse := controller.PostService.Update(request.Context(), postUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   postResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PostControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	postId := params.ByName("postId")
	id, err := strconv.Atoi(postId)
	helper.PanicIfError(err)

	controller.PostService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PostControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	postId := params.ByName("postId")
	id, err := strconv.Atoi(postId)
	helper.PanicIfError(err)

	postResponse := controller.PostService.FindByIdWithTag(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   postResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PostControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Check for the presence of the "tag" query parameter
	tagLabel := request.URL.Query().Get("tag")

	var webResponse web.WebResponse
	if tagLabel != "" {
		// If tagLabel is provided, call the service method to find posts by tag label
		postResponses := controller.PostService.FindByTagLabel(request.Context(), tagLabel)
		webResponse = web.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   postResponses,
		}
	} else {
		// If tagLabel is not provided, call the service method to find all posts
		postResponses := controller.PostService.FindAll(request.Context())
		webResponse = web.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   postResponses,
		}
	}

	helper.WriteToResponseBody(writer, webResponse)
}
