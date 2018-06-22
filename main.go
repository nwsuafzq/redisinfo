package main

import (
	"github.com/asiainfoLDP/redisinfo/info"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"os"
	"fmt"
)

func main() {
	info.Parm = os.Args[1:]
	if len(info.Parm) < 1{
		fmt.Println("Please input parm")
		return
	}
	router := handle()
	s := &http.Server{
		Addr:    ":10012",
		Handler: router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 0,
	}
	//监听端口
	s.ListenAndServe()
}

func handle() (router *gin.Engine) {
	gin.SetMode(gin.DebugMode)
	//获取路由实例
	router = gin.Default()
	router.GET("/metrics", info.GetAllInfo)
	router.GET("/metrics/:info", info.GetAppointInfo)
	return
}
