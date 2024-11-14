package controller

import (
	"net/http"
	"strconv"
	"todo-backend/data/request"
	"todo-backend/data/response"
	"todo-backend/service"
	"todo-backend/utils"
)

type TodoController struct {
	TodoService service.TodoService
}

func NewTodoController(todoService service.TodoService) *TodoController {
	return &TodoController{TodoService: todoService}
}

func (controller *TodoController) Create(writer http.ResponseWriter, requests *http.Request) {
	todoCreateRequest := request.TodoCreateRequest{}
	utils.ReadRequestBody(requests, &todoCreateRequest)

	controller.TodoService.Create(requests.Context(), todoCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	utils.WriteResponseBody(writer, webResponse)
}

func (controller *TodoController) Update(writer http.ResponseWriter, requests *http.Request) {
	params := utils.GetParams(requests)
	todoUpdateRequest := request.TodoUpdateRequest{}
	utils.ReadRequestBody(requests, &todoUpdateRequest)

	todoId := params["todoId"]
	id, err := strconv.Atoi(todoId)
	utils.PanicIfError(err)
	todoUpdateRequest.Id = id

	controller.TodoService.Update(requests.Context(), todoUpdateRequest)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	utils.WriteResponseBody(writer, webResponse)
}

func (controller *TodoController) Delete(writer http.ResponseWriter, requests *http.Request) {
	params := utils.GetParams(requests)
	id, err := strconv.Atoi(params["todoId"])
	utils.PanicIfError(err)

	controller.TodoService.Delete(requests.Context(), id)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	utils.WriteResponseBody(writer, webResponse)
}

func (controller *TodoController) FindAll(writer http.ResponseWriter, requests *http.Request) {
	result := controller.TodoService.FindAll(requests.Context())

	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	utils.WriteResponseBody(writer, webResponse)
}

func (controller *TodoController) FindById(writer http.ResponseWriter, requests *http.Request) {
	params := utils.GetParams(requests)
	id, err := strconv.Atoi(params["todoId"])
	utils.PanicIfError(err)

	result := controller.TodoService.FindById(requests.Context(), id)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	utils.WriteResponseBody(writer, webResponse)
}
