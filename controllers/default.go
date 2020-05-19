package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}
type Article struct {
	Title   string
	Content string
}

func (c *MainController) Get() {

	// 1、模板中绑定基本类型
	c.Data["website"] = "beego 朱莉"
	c.Data["title"] = "你好beego"
	c.Data["num"] = 12
	c.Data["falg"] = true

	// 2、模板中绑定结构体数据
	acticle := Article{
		Title:   "我是golang教程",
		Content: "beego小米商城",
	}
	c.Data["article"] = acticle

	// 3、 range
	c.Data["sliceList"] = []string{"php", "java", "golang"}

	// 4、map
	userinfo := make(map[string]interface{})
	userinfo["username"] = "张三"
	userinfo["age"] = 20
	userinfo["sex"] = "男"
	c.Data["userinfo"] = userinfo

	// 4、遍历结构体类型的切片

	c.TplName = "index.html"

}
