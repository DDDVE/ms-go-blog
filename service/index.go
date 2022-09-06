package service

import (
	"ms-go-blog/config"
	"ms-go-blog/dao"
	"ms-go-blog/models"
)

func GetAllIndexInfo() (*models.HomeResponse, error) {
	var categorys, err = dao.GetAllCategory()
	if err != nil {
		return nil, err
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
	return hr, nil
}
