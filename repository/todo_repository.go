package repository

import (
	"context"
	"todo-backend/model"
)

type TodoRepository interface {
	Add(ctx context.Context, todo model.Todo)
	Update(ctx context.Context, todo model.Todo)
	Delete(ctx context.Context, todoId int)
	FindById(ctx context.Context, todoId int) (model.Todo, error)
	FindAll(ctx context.Context) []model.Todo
}
