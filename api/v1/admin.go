package v1

import (
	"awesomeProject3/models"
	"awesomeProject3/utils/errmsg"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AdminAddUser (c *gin.Context){
	username:= c.PostForm("username")
	password:= c.PostForm("password")
	uid , _ := strconv.Atoi(c.PostForm("uid"))
	name:= c.PostForm("name")

	code := models.AdminAddUser(uint(uid),username,password,name)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}

func AdminGetList (c *gin.Context) {
	username:= c.PostForm("username")
	pageSize , _ := strconv.Atoi(c.PostForm("count"))
	pageNum , _ := strconv.Atoi(c.PostForm("page"))
	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)
	switch {
	case pageSize >=100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageSize = -1
	}
	data,total := models.AdminGetList(v,pageSize,pageNum,username)
	code = errmsg.SUCCESE
c.JSON(http.StatusOK,gin.H{
	"status":code,
	"data":data,
	"total":total,
	"message":errmsg.GetErrMsg(code),
})
}
func AdminDeleteUser (c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("uid"))
	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)
	code := models.AdminDeleteUser(v,id)
	c.JSON(http.StatusOK,gin.H{
		"status" : code,
		"message": errmsg.GetErrMsg(code),
	})
}
func AdminEditPassword (c *gin.Context){
	id, _ := strconv.Atoi(c.PostForm("uid"))
	password:= c.PostForm("password")
	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)
	code := models.AdminEditPassword(v,password,id)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}

func AdminItemGetList(c *gin.Context){
	username:= c.PostForm("username")
	itemName:= c.PostForm("item_name")
	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)
	pageSize , _ := strconv.Atoi(c.PostForm("count"))
	pageNum , _ := strconv.Atoi(c.PostForm("page"))
	data,total ,code := models.AdminItemGetList(v,pageSize,pageNum,itemName,username)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"total":total,
		"message":errmsg.GetErrMsg(code),
	})
}
func AdminDeleteItem (c *gin.Context){
	itemId , _ := strconv.Atoi(c.PostForm("item_id"))
	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)
	code := models.AdminDeleteItem(v, itemId)
	c.JSON(http.StatusOK,gin.H{
		"status" : code,
		"message": errmsg.GetErrMsg(code),
	})
}

func AdminItemAttorn(c *gin.Context){
	username:= c.PostForm("username")
	itemId , _ := strconv.Atoi(c.PostForm("item_id"))
	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)
	code := models.AdminItemAttorn(v,username,itemId)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}