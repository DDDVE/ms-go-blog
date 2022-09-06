package views

import (
	"ms-go-blog/common"
	"ms-go-blog/config"
	"ms-go-blog/models"
	"net/http"
)



func (h *HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
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
	index := common.Template.Index
	index.WriteData(w, hr)
}
