package api

import (
	"errors"
	"ms-go-blog/common"
	"ms-go-blog/dao"
	"ms-go-blog/models"
	"ms-go-blog/service"
	"ms-go-blog/utils"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (api *Api) GetPost(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	pIdStr := strings.TrimPrefix(path, "/api/v1/post/")
	pid, err := strconv.Atoi(pIdStr)
	if err != nil {
		common.Error(w, errors.New("不识别此请求路径"))
		return
	}
	post, err := dao.GetPostById(pid)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, post)
}

func (api *Api) SaveAndUpdatePost(w http.ResponseWriter, r *http.Request) {
	//获取用户 id，判断用户是否登录
	token := r.Header.Get("Authorization")
	_, claim, err := utils.ParseToken(token)
	if err != nil {
		common.Error(w, errors.New("登陆已过期"))
		return
	}
	uid := claim.Uid
	//post = save
	method := r.Method
	switch method {
	case http.MethodPost:
		params := common.GetRequestJsonParam(r)
		cId := params["categoryId"].(string)
		categoryId, _ := strconv.Atoi(cId)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)
		pType := int(postType)
		post := &models.Post{
			-1,
			title,
			slug,
			content,
			markdown,
			categoryId,
			uid,
			-1,
			pType,
			time.Now(),
			time.Now(),

		}
		service.SavePost(post)
		common.Success(w, post)
	case http.MethodPut:
		//update
		params := common.GetRequestJsonParam(r)
		cId := params["categoryId"].(string)
		categoryId, _ := strconv.Atoi(cId)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)
		pType := int(postType)
		pidFloat := params["pid"].(float64)
		pid := int(pidFloat)
		post := &models.Post{
			pid,
			title,
			slug,
			content,
			markdown,
			categoryId,
			uid,
			-1,
			pType,
			time.Now(),
			time.Now(),

		}
		service.UpdatePost(post)
		common.Success(w, post)
	}
}

func (*Api) SearchPost(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	condition := r.Form.Get("val")
	searchRes := service.SearchPost(condition)
	common.Success(w, searchRes)
}


