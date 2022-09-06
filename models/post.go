package models

import (
	"html/template"
	"ms-go-blog/config"
	"time"
)

type Post struct {
	Pid 		int 		`json:"pid"`		//文章id
	Title 		string 		`json:"title"`		//文章标题
	Content 	string 		`json:"content"`	//文章的 html
	Markdown 	string 		`json:"markdown"`	//文章的 markdown
	CategoryId 	int 		`json:"categoryId"`	//分类 id
	UserId 		int 		`json:"userId"`		//用户 id
	ViewCount 	int 		`json:"viewCount"`	//被查看次数
	Type 		int 		`json:"type"`		//文章类型。0=普通 1=自定义
	Slug 		string 		`json:"slug"`		//自定义页面 path
	CreateAt 	time.Time 	`json:"createAt"`	//创建时间
	UpdateAt 	time.Time 	`json:"updateAt"`	//更新时间
}

type PostMore struct {
	Pid 			int 			`json:"pid"` //文章 id
	Title 			string 			`json:"title"` //文章标题
	Slug 			string 			`json:"slug"`//自定义页面 path
	Content 		template.HTML 	`json:"content"` //文章的 html
	CategoryId 		int 			`json:"categoryId"` //文章的 markdown
	CategoryName 	string 			`json:"categoryName"` //分类名
	UserId 			int 			`json:"userId"` //用户 id
	UserName 		string 			`json:"userName"` // 用户名
	ViewCount 		int 			`json:"viewCount"` //查看次数
	Type 			int 			`json:"type"` //文章类型。0=普通 1=自定义
	CreateAt 		string 			`json:"createAt"`
	UpdateAt 		string 			`json:"updateAt"`
}

type PostReq struct {
	Pid 		int 	`json:"pid"`
	Title 		string 	`json:"title"`
	Slug 		string 	`json:"slug"`
	Content 	string 	`json:"content"`
	Markdown 	string 	`json:"markdown"`
	CategoryId 	int 	`json:"categoryId"`
	UserId 		int 	`json:"userId"`
	Type 		int 	`json:"type"`
}

type SearchResp struct {
	Pid 	int 	`json:"pid" orm:"pid"`
	Title 	string 	`json:"title" orm:"title"`
}

type PostRes struct {
	config.Viewer
	config.SystemConfig
	Article PostMore
}
