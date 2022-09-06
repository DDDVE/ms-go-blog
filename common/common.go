package common

import (
	"ms-go-blog/config"
	"ms-go-blog/models"
)

var Template models.HtmlTemplate
func LoadTemplate() {
	Template = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/")
}
