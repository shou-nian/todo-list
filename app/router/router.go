package router

import (
	"github.com/gin-gonic/gin"
	"github.com/riny/demo-go-gin/app/controller"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func New() *gin.Engine {
	router := gin.Default()
	t := controller.New()

	router.POST("/todo", t.AddTodoList)
	router.GET("/todo", t.QueryTodoList)
	router.PUT("/todo", t.UpdateTodoStatus)
	router.DELETE("/todo", t.DeleteTodoList)

	// swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
