package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type IndexData struct {
	Title string	`json:"title"`
	Desc string		`json:"desc"`

}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//res := "status:200, message:\"good\""
	var indexData IndexData
	indexData.Title = "dddve码神之路"
	indexData.Desc = "现在是入门阶段"
	jsonStr, err := json.Marshal(indexData)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonStr)
}

func main() {
	//程序入口
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", index)
	//开始启动并监听
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}


