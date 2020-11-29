package main

import (
	"fmt"
	"github.com/Peterliang233/Function/database"
	routers "github.com/Peterliang233/Function/router"
	"github.com/Peterliang233/Function/settings"
	"net/http"
)

func init(){  //initial
	settings.LoadSettings()
	database.ConnectMysql() //连接数据库
}

func main() {
	r := routers.InitRouters()
	readTimeout := settings.ServerSetting.ReadTimeout
	writeTimeout := settings.ServerSetting.WriteTimeout
	maxHeaderBytes := 1 << 20
	server := &http.Server{
		Addr : ":9090",
		Handler: r,
		ReadTimeout: readTimeout,
		WriteTimeout: writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	err:=server.ListenAndServe()  //监听端口
	if err != nil{
		fmt.Println(err)
	}
	database.Db.Close()  //关闭数据库
}