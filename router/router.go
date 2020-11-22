package routers

import (
	middlewares "github.com/Peterliang233/Function/middleware"
	v1 "github.com/Peterliang233/Function/router/v1"
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	api.Use(middlewares.Adapter)
	{
		api.GET("/employee", v1.FindEmployee)
		api.POST("/employee", v1.AddEmployee)
		api.PUT("/employee", v1.UpdateEmployee)
		api.DELETE("/employee", v1.DeleteEmployee)
	}
	return router
}