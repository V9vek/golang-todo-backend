package service

import (
	"context"
	"todo-backend/data/request"
	"todo-backend/data/response"
)

type TodoService interface {
	Create(ctx context.Context, request request.TodoCreateRequest)
	Update(ctx context.Context, request request.TodoUpdateRequest)
	Delete(ctx context.Context, todoId int)
	FindById(ctx context.Context, todoId int) response.TodoResponse
	FindAll(ctx context.Context) []response.TodoResponse
}
