package api

import (
	"errors"
	"ms-go-blog/common"
	"ms-go-blog/models"
	"ms-go-blog/service"
	"ms-go-blog/utils"
	"net/http"
	"strconv"
	"time"
)

func (api Api) SaveAndUpdatePost(w http.ResponseWriter, r *http.Request) {
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

	}


}


