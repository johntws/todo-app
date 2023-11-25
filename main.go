package main

import config "todo-app/config"

func main() {
	config.InitializeMysql()
	r := config.SetupRouter()
	r.Run()

}
