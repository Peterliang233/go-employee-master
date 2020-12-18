package main

import (
	"fmt"
	"github.com/Peterliang233/Function/model"
	routers "github.com/Peterliang233/Function/router"
	"github.com/Peterliang233/Function/settings"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() { //initial
	settings.LoadSettings() //加载配置
	model.ConnectMysql()    //连接数据库
}

func main() {
	gin.SetMode(settings.ServerSetting.RunMode)
	r := routers.InitRouters()
	readTimeout := settings.ServerSetting.ReadTimeout
	writeTimeout := settings.ServerSetting.WriteTimeout
	maxHeaderBytes := 1 << 20
	server := &http.Server{
		Addr:           ":9090",
		Handler:        r,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	err := server.ListenAndServe() //监听端口
	if err != nil {
		fmt.Println(err)
	}
	model.CloseDatabase() //关闭数据库
}
