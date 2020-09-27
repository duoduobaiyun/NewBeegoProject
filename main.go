package main

import (
	_ "NewBeegoProject/routers"
	"github.com/astaxie/beego"
)

//func init() {
//	beego.BConfig.CopyRequestBody = true
//}
//

func main() {
	beego.Run()

}

