package controllers

import (
	"NewBeegoProject/db_mysql"
	"NewBeegoProject/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type GoodController struct {
	beego.Controller
}

//该方法用于处理post类请求
func (c *GoodController) Post() {

	//1、获取请求数据
	userName := c.Ctx.Input.Query("user")
	password := c.Ctx.Input.Query("psd")
	//2、使用固定数据进行数据校验
	//admin 123456
	if userName != "admin" && password != "123546" {
		//代表错误处理
		c.Ctx.ResponseWriter.Write([]byte("对不起,数据校验错误"))
		return

		fmt.Println(c == nil)
		fmt.Println(c.Ctx == nil)
		fmt.Println(c.Ctx.Request == nil)
		//接收前端传递的数据
		bodyBytes, err := ioutil.ReadAll(c.Ctx.Request.Body)
		if err != nil {
			fmt.Println(err.Error(), "我错了")
			c.Ctx.WriteString("数据库接收错误,请认真检查")
			return
		}
		var info models.Info //调用info结构体
		err = json.Unmarshal(bodyBytes, &info)
		if err != nil {
			fmt.Println(err.Error(), "我错了")
			c.Ctx.WriteString("数据解析错误")
			return
		}

		//3.将解析到用户的数据,保存到数据库
		id, err := db_mysql.InsertInfo(info)
		if err != nil {
			fmt.Println(err.Error(), "我错了")
			c.Ctx.WriteString("数据解析错误")
			return
		}
		fmt.Println(id)

		result := models.ResponseResult{
			Code:    0,
			Message: "保存成功",
			Date:    nil,
		}
		c.Data["json"] = &result
		c.ServeJSON()
		c.Data["Website"] = "beego.me"
		c.Data["Email"] = "astaxie@gmail.com"
		c.TplName = "views/good.html"
	}
}
