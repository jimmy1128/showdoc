package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode string
	HttpPort string
	JwtKey string

	Db string
	DbHost string
	DbPort string
	DbUser string
	DbPassWord string
	DbName string

	Domain string
	MaxAge int
)

func init (){
	file , err := ini.Load("conf/config.ini")
	if err !=nil {
		fmt.Println("配置文件读取错误")
		fmt.Println(file)
	}
	LoadServer(file)
	LoadData(file)
}

func LoadServer(file *ini.File){
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString("debug")
	JwtKey = file.Section("server").Key("JwtKey").MustString("89js82js72")
}
func LoadData (file *ini.File){
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost1")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser  = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord  = file.Section("database").Key("DbPassWord").MustString("Dreamfuture88")
	DbName  = file.Section("database").Key("DbName").MustString("doc")
}
func LoadCookie (file *ini.File){
	Domain = file.Section("cookie").Key("Domain").MustString("*")
	MaxAge = file.Section("cookie").Key("MaxAge").MustInt(3600)
}