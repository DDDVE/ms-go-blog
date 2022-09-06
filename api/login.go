package api

import (
	"ms-go-blog/common"
	"ms-go-blog/service"
	"net/http"
)

func (api *Api) Login(w http.ResponseWriter, r *http.Request) {
	//接收用户名和密码，返回对应的 json 数据
	params := common.GetRequestJsonParam(r)
	userName := params["username"].(string)
	passwd := params["passwd"].(string)
	loginRes, err := service.Login(userName, passwd)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, loginRes)
}
