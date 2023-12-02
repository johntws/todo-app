package route

import (
	"github.com/gin-gonic/gin"
	service "todo-app/service"
)

func SetupRouter(todoService service.TodoService) *gin.Engine {
	Router := gin.Default()

	Router.POST("/login", service.Login)

	protected := Router.Group("/protected")
	//protected.Use(middleware.JWT())
	{
		v1 := protected.Group("/v1")
		{
			v1.GET("todo", todoService.GetTodos)
			v1.POST("todo", todoService.CreateTodo)
			v1.GET("todo/:id", todoService.GetTodo)
			v1.PUT("todo/:id", todoService.UpdateTodo)
			v1.DELETE("todo/:id", todoService.DeleteTodo)
		}
	}

	return Router
}
