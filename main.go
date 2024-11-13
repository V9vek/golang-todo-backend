package main

import (
	"fmt"
	"net/http"
	"todo-backend/utils"
)

func main() {
	fmt.Println("Start server")

	router := NewRouter()

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	utils.PanicIfError(err)
}
