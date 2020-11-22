package routers

import (
	middlewares "github.com/Peterliang233/Function/middleware"
	v12 "github.com/Peterliang233/Function/router/v1"
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.Adapter)
	api := router.Group("/api")
	{
		api.GET("/employee", v12.FindEmployee)
		api.POST("/employee", v12.AddEmployee)
		api.PUT("/employee", v12.UpdateEmployee)
		api.DELETE("/employee", v12.DeleteEmployee)
	}
	return router
}
