package db_mysql

import (
	"NewBeegoProject/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	fmt.Println("连接mysql数据库")
	config := beego.AppConfig                  //配置文件所在的路径
	dbDriver := config.String("db_driverName") //数据库的驱动,比如说Oracle,或者Mysql
	dbUser := config.String("db_user")         //用户名
	dbPassword := config.String("db_password") //密码
	dbIp := config.String("db_ip")             //本机的IpConfig,ip和端口
	dbName := config.String("db_name")       //数据库的库名

	connUrl := dbUser + ":" + dbPassword + "@tcp(" + dbIp + ")/" + dbName + "?charset=utf8" //连接数据库
	db, err := sql.Open(dbDriver, connUrl)                                                  //判断数据库是否连接失败或者成功
	if err != nil {
		fmt.Println("错错错是我的错", err.Error()) //打印错误
		panic("数据库连接错误,请检查配置")
	}
	//为全局变量赋值
	DB = db
}

//将用户信息保存到数据库表当中

func InsertInfo(info models.Info) (int64, error) {
	//panic("我错了")
	//1.将用户密码进行hash脱敏,使用md5计算hash值,并存储hash值
	hashMd5 := md5.New()
	//hashMd5.Write([]byte(info.User))
	hashMd5.Write([]byte(info.Password))
	//hashMd5.Write([]byte(info.School))
	//hashMd5.Write([]byte(info.Class))
	//hashMd5.Write(([]byte(info.Name)))
	//hashMd5.Write([]byte(info.Sex))
	//hashMd5.Write([]byte(info.Hobby))
	//hashMd5.Write([]byte(info.Color))
	bytes:=hashMd5.Sum(nil)
	//info.User=hex.EncodeToString(bytes)
	info.Password=hex.EncodeToString(bytes)
	//info.School=hex.EncodeToString(bytes)
	//info.Class=hex.EncodeToString(bytes)
	//info.Name=hex.EncodeToString(bytes)
	//info.Sex=hex.EncodeToString(bytes)
	//info.Hobby=hex.EncodeToString(bytes)
	//info.Color=hex.EncodeToString(bytes)
	//fmt.Println("用户名",info.User,"密码",info.Password)
	//fmt.Println("学校",info.School,"班级",info.Class)
	//fmt.Println("姓名",info.Name,"性别",info.Sex)
	//fmt.Println("爱好",info.Hobby,"喜欢的额颜色",info.Color)
	//panic("我错了")
	fmt.Println("用户",info.Nick,"密码",info.Password)
	//DB.SetMaxOpenConns(1)
    result,err:=DB.Exec("insert into info (nick,class,sex,users,password,color,hobby,school) values (?,?,?,?,?,?,?,?)",info.Nick,info.Class,info.Sex,info.User,info.Password,info.Color,info.Hobby,info.School)
	if err != nil {
		panic(err.Error())
		fmt.Println("错错错是我的错", err.Error())
		return -1, err
	}
	id,err:=result.RowsAffected()
	if err != nil {
		panic(err.Error())
		fmt.Println("错错错是我的错",err.Error())
		return -1,err
	}
	return id,err
}
