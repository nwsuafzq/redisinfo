package main

import (
	"github.com/asiainfoLDP/redisinfo/info"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"fmt"
	"flag"
)

var serverPort string = ":"

func init(){
	port := flag.String("server","10012","input your port")
	auth := flag.String("auth","","input your auth")
	flag.Parse()
	info.Parm = flag.Args()
	if len(info.Parm) < 1{
		panic("Please input Ip")
	}
	info.Secret = *auth
	serverPort += *port
	fmt.Println("serverPort is ",serverPort)
}

func main(){
	router := handle()
	s := &http.Server{
		Addr:    serverPort,
		Handler: router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 0,
	}
	//监听端口
	err := s.ListenAndServe()
	fmt.Println("Sever error ",err)
}

func handle() (router *gin.Engine) {
	gin.SetMode(gin.DebugMode)
	//获取路由实例
	router = gin.Default()
	router.GET("/metrics", info.GetAllInfo)
	router.GET("/metrics/:info", info.GetAppointInfo)
	return
}