package controller

import (
	"github.com/julienschmidt/httprouter"
	"github.com/urcane/post-test-golang/helper"
	"github.com/urcane/post-test-golang/model/web"
	"github.com/urcane/post-test-golang/service/tag_service"
	"net/http"
	"strconv"
)

type TagController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type TagControllerImpl struct {
	TagService tag_service.TagService
}

func NewTagController(tagService tag_service.TagService) TagController {
	return &TagControllerImpl{
		TagService: tagService,
	}
}

func (controller *TagControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	tagCreateRequest := web.TagCreateRequest{}
	helper.ReadFromRequestBody(request, &tagCreateRequest)

	tagResponse := controller.TagService.CreateWithPost(request.Context(), tagCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   tagResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TagControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	tagUpdateRequest := web.TagUpdateRequest{}
	helper.ReadFromRequestBody(request, &tagUpdateRequest)

	tagId := params.ByName("tagId")
	id, err := strconv.Atoi(tagId)
	helper.PanicIfError(err)

	tagUpdateRequest.Id = id

	tagResponse := controller.TagService.Update(request.Context(), tagUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   tagResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TagControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	tagId := params.ByName("tagId")
	id, err := strconv.Atoi(tagId)
	helper.PanicIfError(err)

	controller.TagService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TagControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	tagId := params.ByName("tagId")
	id, err := strconv.Atoi(tagId)
	helper.PanicIfError(err)

	tagResponse := controller.TagService.FindByIdWithPost(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   tagResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TagControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	tagResponses := controller.TagService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   tagResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
