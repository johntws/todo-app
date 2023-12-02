package main

import (
	config "todo-app/config"
	"todo-app/repository"
	"todo-app/service"
)

func main() {
	db := config.InitializeMysql()
	repo := repository.NewTodoRepo(db)
	todoService := service.NewTodoService(repo)
	r := config.SetupRouter(todoService)
	r.Run()
}
