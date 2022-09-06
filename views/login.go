package views

import (
	"ms-go-blog/common"
	"ms-go-blog/config"
	"net/http"
)

func (h *HTMLApi) Login(w http.ResponseWriter, r *http.Request) {
	login := common.Template.Login
	login.WriteData(w, config.Cfg.Viewer)
}
