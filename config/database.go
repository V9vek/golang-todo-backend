package config

import (
	"database/sql"
	"fmt"
	"os"
	"todo-backend/utils"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host   = "localhost"
	port   = 3306
	user   = "root"
	dbName = "todo_app"
)

func DatabaseConnection() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, os.Getenv("DB_PASSWORD"), host, port, dbName)
	driverName := "mysql"

	db, err := sql.Open(driverName, dsn)
	utils.PanicIfError(err)

	err = db.Ping()
	utils.PanicIfError(err)

	fmt.Println("Successfully connected to database")
	return db
}
