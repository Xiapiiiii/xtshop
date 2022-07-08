package main

import (
	"github.com/beego/beego/v2/server/web"
)

type MainController struct {
	web.Controller
}

func (m *MainController) Get() {
	m.Ctx.WriteString("hello world")
}
func main() {
	ctrl := &MainController{}
	web.Router("/", ctrl)
	web.Run()

}
