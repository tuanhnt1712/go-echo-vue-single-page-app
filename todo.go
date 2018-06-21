package main

import (
	"database/sql"

	"go-echo-vue-single-page-app/handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// db := database.initDB("storage.db")
	// database.migrate(db)
	db := initDB("storage.db")
	migrate(db)

	e := echo.New()

	e.File("/", "public/index.html")
	e.GET("/tasks", handlers.GetTasks(db))
	e.POST("/tasks", handlers.PutTask(db))
	e.PUT("/tasks/:id", handlers.DoneTask(db))
	e.DELETE("/tasks/:id", handlers.DeleteTask(db))
	e.GET("/tasks_done", handlers.GetTasksDone(db))

	e.Run(standard.New(":8000"))
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	// Here we check for any db errors then exit
	if err != nil {
		panic(err)
	}

	// If we don't get any errors but somehow still don't get a db connection
	// we exit as well
	if db == nil {
		panic("db nil")
	}
	return db
}

func migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS tasks(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL,
		is_done BOOL DEFAULT false
	);
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL,
		email VARCHAR NOT NULL
	);
	`

	_, err := db.Exec(sql)
	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}
}
