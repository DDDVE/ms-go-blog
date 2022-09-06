package main

import (
	"log"
	"ms-go-blog/common"
	"ms-go-blog/router"
	"net/http"
)


func init() {
	//页面模板加载
	common.LoadTemplate()
}


func main() {
	//程序入口
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	router.Router()
	//开始启动并监听
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}


