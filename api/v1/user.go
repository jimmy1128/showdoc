package v1

import (
	"awesomeProject3/models"
	"awesomeProject3/utils/errmsg"
	"awesomeProject3/utils/validator"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var code int
// 添加用户
func AddUser(c *gin.Context){
	var data models.User
	var msg string

	_ = c.ShouldBindJSON(&data)
	msg,code =validator.Validate(&data)
	if code != errmsg.SUCCESE{
		c.JSON(http.StatusOK,gin.H{
			"status" : code,
			"message" : msg,
		})
		return
	}
	if models.RegisterSetting() == "1"{
		code = errmsg.ERROR_APPLY_NEW_ACCOUNT
	}
	data.Name = data.Username
	code = models.CheckUser(data.Username)
	if code == errmsg.SUCCESE{
		models.CreateUser(&data)
	}
	if code == errmsg.ERROR_USERNAME_USED{
		code = errmsg.ERROR_USERNAME_USED
	}
	c.JSON(http.StatusOK ,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}
//查询单个用户
func GetUserInfo(c*gin.Context){

	session := sessions.Default(c)
	user := session.Get("id")
	redirect :=c.PostForm("redirect_login")
	v , _ := user.(uint)

	data,code := models.GetUser(v,redirect)
	//username =data.Username
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data" : data,
		"message":errmsg.GetErrMsg(code),
	})
}
func GetUsers(c *gin.Context) {

	username :=c.PostForm("username")
	data,code:=models.GetUsers(username)


	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}
func EditUser(c *gin.Context){
	var data models.User
	id ,_:= strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code = models.CheckUpUser(id,data.Email)
	if code == errmsg.SUCCESE{
		models.EditUser(id,&data)
	}
	if code == errmsg.ERROR_USERNAME_USED{
		c.Abort()
	}
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}
func ChangeUserPassword(c *gin.Context){
	var data models.User
	password := c.PostForm("password")
	//password := c.GetString("password")
	newPassword := c.PostForm("new_password")

	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)
	data,_ = models.GetUser(v,"")
	if models.ScryptPw(password) == data.Password {
		data.Password = models.ScryptPw(newPassword)
		code = models.ChangePassword(int(v),&data)
	}else{
		code = errmsg.ERROR_PASSWORD_WRONG
	}
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data.Username,
		"message":errmsg.GetErrMsg(code),
	})
}
func DeleteUser(c *gin.Context){
	id ,_:= strconv.Atoi(c.Param("id"))
	code = models.DeleteUser(id)

	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}
func AddGuest(c *gin.Context){
	var data models.Guest
	var msg string
	_ = c.ShouldBindJSON(&data)
	msg,code =validator.Validate(&data)
	if code != errmsg.SUCCESE{
		c.JSON(http.StatusOK,gin.H{
			"status" : code,
			"message" : msg,
		})
		return
	}
	code = models.CheckGuest(data.GName)
	if code == errmsg.SUCCESE{
		models.CreateGuest(&data)
	}
	if code == errmsg.ERROR_USERNAME_USED{
		code = errmsg.ERROR_USERNAME_USED
	}
	c.JSON(http.StatusOK ,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}







