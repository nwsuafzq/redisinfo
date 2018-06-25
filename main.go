package main

import (
	"github.com/asiainfoLDP/redisinfo/info"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"fmt"
	"flag"
	"os"
	"regexp"
)
const CORRECTIP = "((25[0-5]|2[0-4]\\d|((1\\d{2})|([1-9]?\\d)))\\.){3}(25[0-5]|2[0-4]\\d|((1\\d{2})|([1-9]?\\d)))"
var serverPort string = ":"

func init(){
	port := flag.String("server","10012","binding port")
	auth := flag.String("auth","","redis password")
	flag.Parse()
	info.Parm = flag.Args()
	if len(info.Parm) < 1{
		fmt.Printf("usage: %s [options] HOST[:PORT]...\n",os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}
	for _,addr := range info.Parm{
		tIp, _ := regexp.MatchString(CORRECTIP, addr)
		if !tIp{
			fmt.Println("input Ip is not correct IP ",addr)
			os.Exit(1)
		}
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

