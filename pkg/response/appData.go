package response

import (
	"gin-blog/pkg/globals"
	"html/template"
)

type AppData struct {
	Code globals.AppCode `json:"code"`
	Msg  string          `json:"msg"`
	Data Data            `json:"data"`
	Name string          `json:"name"`
	Csrf template.HTML   `json:"csrf"`
}

type Data struct {
	Header  interface{} `json:"header"`
	Content interface{} `json:"content"`
}
