package routers

import (
	"github.com/Peterliang233/Function/middlerware"
	v1 "github.com/Peterliang233/Function/router/v1"
	"github.com/Peterliang233/Function/router/v1/api/user/controller"
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine { //定义路由组
	router := gin.Default()
	//实现一个登录接口
	v1Group := router.Group("/v1")
	v1Group.POST("/login", controller.Login) //通过post发送登录请求
	api := v1Group.Group("/user")
	api.Use(middlerware.JWTAuthMiddleware()) //调用中间件
	{
		api.GET("/employee", v1.FindEmployee)
		api.POST("/employee", v1.AddEmployee)
		//新建一个对信息进行更新的路由组
		update := api.Group("/update")
		{
			update.PUT("/employer", v1.UpdateEmployee)
			update.PUT("/password", v1.UpdatePassword)
			update.PUT("/role", v1.UpdateRole)
			update.PUT("/department", v1.UpdateDepartment)
		}
		api.DELETE("/employee", v1.DeleteEmployee)
	}
	return router
}
