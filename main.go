package main

import (
	"context"
	"github.com/Peterliang233/Function/model"
	routers "github.com/Peterliang233/Function/router"
	"github.com/Peterliang233/Function/settings"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() { //初始化
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
	//err := server.ListenAndServe() //监听端口
	//if err != nil {
	//	fmt.Println(err)
	//}
	//优雅地关闭和重启
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%s\n", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown;", err)
	}
	log.Println("Server exiting")
	model.CloseDatabase() //关闭数据库
}
