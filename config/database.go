package config

import (
	"database/sql"
	"fmt"
	"todo-backend/utils"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = ""
	dbName   = "todo_app"
)

func DatabaseConnection() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbName)
	driverName := "mysql"

	db, err := sql.Open(driverName, dsn)
	utils.PanicIfError(err)

	err = db.Ping()
	utils.PanicIfError(err)

	fmt.Println("Successfully connected to database")
	return db
}
