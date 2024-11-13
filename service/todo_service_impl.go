package service

import (
	"context"
	"todo-backend/data/request"
	"todo-backend/data/response"
	"todo-backend/model"
	"todo-backend/repository"
	"todo-backend/utils"
)

type TodoServiceImpl struct {
	TodoRepository repository.TodoRepository
}

// implementing TodoService
func NewTodoServiceImpl(todoRepository repository.TodoRepository) TodoService {
	return &TodoServiceImpl{TodoRepository: todoRepository}
}

func (t *TodoServiceImpl) Create(ctx context.Context, request request.TodoCreateRequest) {
	todo := model.Todo{
		Title:       request.Title,
		Description: request.Description,
	}
	t.TodoRepository.Add(ctx, todo)
}

func (t *TodoServiceImpl) Delete(ctx context.Context, todoId int) {
	todo, err := t.TodoRepository.FindById(ctx, todoId)
	utils.PanicIfError(err)
	t.TodoRepository.Delete(ctx, todo.Id)
}

func (t *TodoServiceImpl) FindAll(ctx context.Context) []response.TodoResponse {
	todos := t.TodoRepository.FindAll(ctx)

	var todosResponse []response.TodoResponse

	for _, value := range todos {
		todo := response.TodoResponse{
			Id:          value.Id,
			Title:       value.Title,
			Description: value.Description,
			Status:      value.Status,
			CreatedAt:   value.CreatedAt,
		}
		todosResponse = append(todosResponse, todo)
	}

	return todosResponse
}

func (t *TodoServiceImpl) FindById(ctx context.Context, todoId int) response.TodoResponse {
	todo, err := t.TodoRepository.FindById(ctx, todoId)
	utils.PanicIfError(err)
	return response.TodoResponse(todo)
}

func (t *TodoServiceImpl) Update(ctx context.Context, request request.TodoUpdateRequest) {
	todo, err := t.TodoRepository.FindById(ctx, request.Id)
	utils.PanicIfError(err)

	todo.Title = request.Title
	todo.Description = request.Description
	todo.Status = request.Status

	t.TodoRepository.Update(ctx, todo)
}
