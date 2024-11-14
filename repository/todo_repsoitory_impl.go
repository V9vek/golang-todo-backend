package repository

import (
	"context"
	"database/sql"
	"fmt"
	"todo-backend/model"
	"todo-backend/utils"
)

type TodoRepositoryImpl struct {
	Db *sql.DB
}

// implementing TodoRepository
func NewTodoRepository(Db *sql.DB) TodoRepository {
	return &TodoRepositoryImpl{Db: Db}
}

func (t *TodoRepositoryImpl) Add(ctx context.Context, todo model.Todo) {
	tx, err := t.Db.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	SQL_QUERY := "INSERT INTO todos(title, description) VALUES (?, ?)"
	_, errExec := tx.ExecContext(ctx, SQL_QUERY, todo.Title, todo.Description)
	utils.PanicIfError(errExec)
}

func (t *TodoRepositoryImpl) Delete(ctx context.Context, todoId int) {
	tx, err := t.Db.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	SQL_QUERY := "delete from todos where id = ?"
	_, errExec := tx.ExecContext(ctx, SQL_QUERY, todoId)
	utils.PanicIfError(errExec)
}

func (t *TodoRepositoryImpl) FindAll(ctx context.Context) []model.Todo {
	tx, err := t.Db.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	SQL_QUERY := "select id, title, description, status, created_at from todos"
	result, errExec := tx.QueryContext(ctx, SQL_QUERY)
	utils.PanicIfError(errExec)
	defer result.Close()

	var todos []model.Todo

	for result.Next() {
		todo := model.Todo{}
		err := result.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Status, &todo.CreatedAt)
		utils.PanicIfError(err)

		todos = append(todos, todo)
	}

	return todos
}

func (t *TodoRepositoryImpl) FindById(ctx context.Context, todoId int) (model.Todo, error) {
	tx, err := t.Db.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	SQL_QUERY := "select id, title, description, status, created_at from todos where id = ?"
	result, err := tx.QueryContext(ctx, SQL_QUERY, todoId)
	utils.PanicIfError(err)
	defer result.Close()

	todo := model.Todo{}

	if result.Next() {
		err := result.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Status, &todo.CreatedAt)
		utils.PanicIfError(err)
		return todo, nil
	} else {
		return todo, fmt.Errorf("todo id %d not found", todoId)
	}
}

func (t *TodoRepositoryImpl) Update(ctx context.Context, todo model.Todo) {
	tx, err := t.Db.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	SQL_QUERY := "update todos set title=?, description=?, status=? where id=?"
	_, errExec := tx.ExecContext(ctx, SQL_QUERY, todo.Title, todo.Description, todo.Status, todo.Id)
	utils.PanicIfError(errExec)
}
