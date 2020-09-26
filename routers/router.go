package routers

import (
	"NewBeegoProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
    beego.Router("/login",&controllers.GoodController{})
}
