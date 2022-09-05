package main

import (
	//"fmt"
	"html/template"
	"log"
	"ms-go-blog/config"
	"ms-go-blog/models"
	"net/http"
	"time"
)

type IndexData struct {
	Title string	`json:"title"`
	Desc string		`json:"desc"`

}


func IsODD(num int) bool {
	return num % 2 == 0
}

func GetNextName(strs []string, index int) string {
	return strs[index + 1]
}

func Date(format string) string {
	return time.Now().Format(format)
}

func index(w http.ResponseWriter, r *http.Request) {
	t := template.New("index.html")
	//1. 拿到当前的路径
	path := config.Cfg.System.CurrentDir
	//访问博客模板的时候，因为有多个模板嵌套，解析文件的时候，需要将涉及到的所有模板进行解析
	home := path + "/template/home.html"
	header := path + "/template/layout/header.html"
	footer := path + "/template/layout/footer.html"
	pagination := path + "/template/layout/pagination.html"
	personal := path + "/template/layout/personal.html"
	post := path + "/template/layout/post-list.html"
	t.Funcs(template.FuncMap{"isODD":IsODD, "getNextName":GetNextName, "date":Date})
	t, err := t.ParseFiles(path + "/template/index.html", home, header, pagination, personal, post, footer)
	if err != nil {
		log.Panicln("解析模板出错：", err)
	}
	//页面上涉及到的所有数据必须有定义
	var categorys = []models.Category{
		{
			Cid: 1,
			Name: "go",
		},
	}
	var posts = []models.PostMore{
		{
			Pid: 1,
			Title: "go博客",
			Content: "内容",
			UserName: "码神",
			ViewCount: 123,
			CreateAt: "2022-02-20",
			CategoryId:1,
			CategoryName: "go",
			Type:0,
		},
	}
	var hr = &models.HomeResponse{
		config.Cfg.Viewer,
		categorys,
		posts,
		1,
		1,
		[]int{1},
		true,
	}
	t.Execute(w, hr)
}

func main() {
	//程序入口
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", index)
	//将请求路径为 /resource 映射到 public/resource/
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
	//开始启动并监听
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}


