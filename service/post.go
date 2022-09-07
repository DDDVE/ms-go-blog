package service

import (
	"html/template"
	"log"
	"ms-go-blog/config"
	"ms-go-blog/dao"
	"ms-go-blog/models"
)

func GetPostDetail(pid int) (*models.PostRes, error) {
	post, err := dao.GetPostById(pid)

	if err != nil {
		return nil, err
	}
	categoryName := dao.GetCategoryNameById(post.CategoryId)
	userName := dao.GetUserNameById(post.UserId)
	postMore := models.PostMore{
		Pid: post.Pid,
		Title: post.Title,
		Slug: post.Slug,
		Content: template.HTML(post.Content),
		CategoryId: post.CategoryId,
		CategoryName: categoryName,
		UserName: userName,
		ViewCount: post.ViewCount,
		Type: post.Type,
		CreateAt: models.DateDay(post.CreateAt),
		UpdateAt: models.DateDay(post.UpdateAt),
	}
	var postRes =  &models.PostRes{
		config.Cfg.Viewer,
		config.Cfg.System,
		postMore,
	}
	return postRes, nil
}

func Writing() (wr models.WritingRes) {
	wr.Title = config.Cfg.Viewer.Title
	wr.CdnURL = config.Cfg.System.CdnURL
	category, err := dao.GetAllCategory()
	if err != nil {
		log.Println(err)
		return
	}
	wr.Categorys = category
	return
}

func SavePost(post *models.Post) {
	dao.SavePost(post)
}
