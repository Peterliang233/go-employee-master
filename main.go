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
}

func main() {
	database.ReadFile() //read the history information
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
	err:=server.ListenAndServe()
	if err != nil{
		fmt.Println(err)
	}
}