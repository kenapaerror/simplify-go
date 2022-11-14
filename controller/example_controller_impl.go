package controller

import (
	"net/http"
	"simplify-go/helper"
	"simplify-go/model/web"
	"simplify-go/service"

	"github.com/julienschmidt/httprouter"
)

type ExampleControllerImpl struct {
	ExampleService service.ExampleService
}

func NewExampleControllerImpl(exampleService service.ExampleService) ExampleController {
	return &ExampleControllerImpl{
		ExampleService: exampleService,
	}
}

func (controller *ExampleControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	exampleCreateRequest := web.ExampleCreateRequest{}
	helper.ReadFromRequestBody(request, &exampleCreateRequest)

	exampleResponse := controller.ExampleService.Create(request.Context(), exampleCreateRequest)
	response := web.WebResponse{
		Error:   false,
		Message: "OK",
		Data:    exampleResponse,
	}

	helper.WriteToResponseBody(writer, response)
}

func (controller *ExampleControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	exampleUpdateRequest := web.ExampleUpdateRequest{}
	helper.ReadFromRequestBody(request, &exampleUpdateRequest)

	exampleUpdateRequest.Id = params.ByName("exampleId")

	exampleResponse := controller.ExampleService.Update(request.Context(), exampleUpdateRequest)
	response := web.WebResponse{
		Error:   false,
		Message: "OK",
		Data:    exampleResponse,
	}

	helper.WriteToResponseBody(writer, response)
}

func (controller *ExampleControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	exampleId := params.ByName("exampleId")

	controller.ExampleService.Delete(request.Context(), exampleId)
	response := web.WebResponse{
		Error:   false,
		Message: "OK",
	}

	helper.WriteToResponseBody(writer, response)
}

func (controller *ExampleControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	exampleId := params.ByName("exampleId")

	exampleResponse := controller.ExampleService.FindById(request.Context(), exampleId)
	response := web.WebResponse{
		Error:   false,
		Message: "OK",
		Data:    exampleResponse,
	}

	helper.WriteToResponseBody(writer, response)
}

func (controller *ExampleControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	exampleResponse := controller.ExampleService.FindAll(request.Context())
	response := web.WebResponse{
		Error:   false,
		Message: "OK",
		Data:    exampleResponse,
	}

	helper.WriteToResponseBody(writer, response)
}
