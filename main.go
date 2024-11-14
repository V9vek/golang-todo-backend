package main

import (
	"fmt"
	"net/http"
	"todo-backend/config"
	"todo-backend/controller"
	"todo-backend/repository"
	"todo-backend/router"
	"todo-backend/service"
	"todo-backend/utils"
)

func main() {
	fmt.Println("Start server")

	utils.LoadEnvFile()

	// database connection
	db := config.DatabaseConnection()

	// repository
	todoRepository := repository.NewTodoRepository(db)

	// service
	todoService := service.NewTodoServiceImpl(todoRepository)

	// controller
	todoController := controller.NewTodoController(todoService)

	// router handler
	router := router.NewRouter(todoController)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	utils.PanicIfError(err)
}
