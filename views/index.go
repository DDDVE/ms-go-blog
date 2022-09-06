package views

import (
	"errors"
	"log"
	"ms-go-blog/common"
	"ms-go-blog/service"
	"net/http"
)



func (h *HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
	//页面上涉及到的所有数据必须有定义
	//数据库查询
	hr , err := service.GetAllIndexInfo()
	if err != nil {
		log.Println("首页获取数据出错: ", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员"))
	}

	index.WriteData(w, hr)
}
