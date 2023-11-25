package route

import (
	"github.com/gin-gonic/gin"
	"todo-app/middleware"
	"todo-app/service"
)

func SetupRouter() *gin.Engine {
	Router := gin.Default()

	//r.Use(middleware.GlobalErrorHandler())

	//v1 := Router.Group("/v1")
	//{
	//	v1.GET("todo", service.GetTodos)
	//	v1.POST("todo", service.CreateTodo)
	//	v1.GET("todo/:id", service.GetTodo)
	//	v1.PUT("todo/:id", service.UpdateTodo)
	//	v1.DELETE("todo/:id", service.DeleteTodo)
	//}

	Router.POST("/login", service.Login)

	protected := Router.Group("/protected")
	protected.Use(middleware.JWT())
	{
		v1 := protected.Group("/v1")
		{
			v1.GET("todo", service.GetTodos)
			v1.POST("todo", service.CreateTodo)
			v1.GET("todo/:id", service.GetTodo)
			v1.PUT("todo/:id", service.UpdateTodo)
			v1.DELETE("todo/:id", service.DeleteTodo)
		}
	}

	return Router
}
