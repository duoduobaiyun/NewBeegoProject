package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

//func (c *MainController) Get() {
//	c.Data["Website"] = "beego.me"
//	c.Data["Email"] = "astaxie@gmail.com"
//	c.TplName = "index.tpl"
//}

//func (c *MainController) Post()  {
//	//1.解析前端提交的json格式的数据
//	var info models.Info
//	databytes,err:=ioutil.ReadAll(c.Ctx.Request.Body)
//	if err != nil {
//		fmt.Println(err.Error())
//		c.Ctx.WriteString("数据解析错误")
//		return
//	}
//	err=json.Unmarshal(databytes,&info)
//	if err != nil {
//		c.Ctx.WriteString("数据解析失败")
//		return
//	}
//	fmt.Println("用户",info.User)
//	fmt.Println("密码",info.Password)
//	fmt.Println("学校",info.School)
//	fmt.Println("班级",info.Class)
//	fmt.Println("姓名",info.Name)
//	fmt.Println("性别",info.Sex)
//	fmt.Println("爱好",info.Hobby)
//	fmt.Println("喜欢的颜色",info.Color)
//	fmt.Println("数据解析成功")
//}