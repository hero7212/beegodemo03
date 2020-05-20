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

	// 5、遍历结构体类型的切片

	c.Data["articleList"] = []Article{
		{
			Title:   "新闻111",
			Content: "新闻内容111",
		},
		{
			Title:   "新闻222",
			Content: "新闻内容222",
		},
		{
			Title:   "新闻333",
			Content: "新闻内容333",
		},
		{
			Title:   "新闻444",
			Content: "新闻内容444",
		},
	}

	// 6、结构体类型的另一种定义方法

	/*
		匿名结构，它就是一个类型
		struct {
		Title string
		}

	*/

	c.Data["cmsList"] = []struct {
		Title string
	}{

		{
			Title: "新闻11111111",
		},
		{
			Title: "新闻22222222",
		},
		{
			Title: "新闻33333333",
		},
		{
			Title: "新闻44444444",
		},
	}

	// 7、模板中的条件判断
	c.Data["isLogin"] = true
	c.Data["isHome"] = true
	c.Data["isAbout"] = true

	// 8、 大于、小于、等于

	c.Data["n1"] = 12
	c.Data["n2"] = 6

	c.TplName = "index.html"

}
