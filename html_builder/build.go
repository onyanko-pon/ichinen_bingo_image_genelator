package html_builder

import (
	"strings"
	"text/template"
)

type HTMLData struct {
	ImageURL string
	TodoList []string
}

func BuildHTML(data HTMLData) string {
	tmpl, _ := template.New("index.html").ParseFiles("./template/index.html")
	writer := new(strings.Builder)
	tmpl.Execute(writer, data)
	return writer.String()
}
