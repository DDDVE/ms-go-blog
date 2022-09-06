package models

import (
	"html/template"
	"io"
	"log"
	"time"
)

type TemplateBlog struct {
	*template.Template
}

type HtmlTemplate struct {
	Index TemplateBlog
	Category TemplateBlog
	Custom TemplateBlog
	Detail TemplateBlog
	Login TemplateBlog
	Pigeonhole TemplateBlog
	Writing TemplateBlog
}

func (t *TemplateBlog) WriteData(w io.Writer, data interface{}) {
	err := t.Execute(w, data)
	if err != nil {
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
		}
	}
}

func (t *TemplateBlog) WriteError(w io.Writer, err error) {
	if err != nil {
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
		}
	}
}

func InitTemplate(templateDir string) (HtmlTemplate, error) {
	tp, err := readTemplate(
		[]string{"index", "category", "custom", "detail", "login", "pigeonhole", "writing"},
		templateDir,
		)
	var htmlTemplate HtmlTemplate
	if err != nil {
		return htmlTemplate, err
	}
	htmlTemplate.Index = tp[0]
	htmlTemplate.Category = tp[1]
	htmlTemplate.Custom = tp[2]
	htmlTemplate.Detail = tp[3]
	htmlTemplate.Login = tp[4]
	htmlTemplate.Pigeonhole = tp[5]
	htmlTemplate.Writing = tp[6]
	return htmlTemplate, nil
}

func IsODD(num int) bool {
	return num % 2 == 0
}

func GetNextName(strs []string, index int) string {
	return strs[index + 1]
}

func Date(format string) string {
	return time.Now().Format(format)
}

func DateDay(date time.Time) string {
	return date.Format("2006-01-02 15:05:04")
}

func readTemplate(templates []string, templateDir string) ([]TemplateBlog, error) {
	var tbs []TemplateBlog
	for _, view := range templates {
		viewName := view + ".html"
		t := template.New(viewName)
		//访问博客模板的时候，因为有多个模板嵌套，解析文件的时候，需要将涉及到的所有模板进行解析
		home := templateDir + "home.html"
		header := templateDir + "layout/header.html"
		footer := templateDir + "layout/footer.html"
		pagination := templateDir + "layout/pagination.html"
		personal := templateDir + "layout/personal.html"
		post := templateDir + "layout/post-list.html"
		t.Funcs(template.FuncMap{"isODD":IsODD, "getNextName":GetNextName, "date":Date, "dateDay":DateDay})
		t, err := t.ParseFiles(templateDir + viewName, home, header, pagination, personal, post, footer)
		if err != nil {
			log.Panicln("解析模板出错：", err)
			return nil, err
		}
		var tb TemplateBlog
		tb.Template = t
		tbs = append(tbs, tb)
	}
	return tbs, nil
}
