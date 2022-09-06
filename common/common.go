package common

import (
	"ms-go-blog/config"
	"ms-go-blog/models"
	"sync"
)

var Template models.HtmlTemplate
func LoadTemplate() {
	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		//耗时
		var err error
		Template ,err = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/")
		if err != nil {
			panic(err)
		}
		//标识为已完成
		w.Done()
	}()
	//等待上方完成
	w.Wait()
}
