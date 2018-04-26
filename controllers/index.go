package controllers

import(
	"fmt"
	"github.com/astaxie/beego"
	"user/models"
)

type IndexController struct {
	beego.Controller
}

func (index *IndexController) Get()  {
	sess := index.StartSession()
	username := sess.Get("username")
	fmt.Println(username)
	if username == nil || username == "" {
		index.TplName = "index.html"
	} else {
		index.TplName = "success.html"
	}
}