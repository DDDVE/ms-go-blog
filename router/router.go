package router

import (
	"ms-go-blog/api"
	"ms-go-blog/views"
	"net/http"
)

func Router() {
	//1.页面views 2.api 数据（json）3.静态资源
	http.HandleFunc("/", views.HTML.Index)
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	//将请求路径为 /resource 映射到 public/resource/
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
