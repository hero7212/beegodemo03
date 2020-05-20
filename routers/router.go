package routers

import (
	"beegodemo03/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/article", &controllers.ArticleController{})
}
