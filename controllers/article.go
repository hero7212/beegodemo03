package controllers

import (
	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

func (c *MainController) Get() {

	c.TplName = "article.html"
}
