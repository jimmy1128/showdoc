package v1

import (
	"awesomeProject3/models"
	"awesomeProject3/utils/errmsg"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func SaveMember (c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)
	itemId , _ := strconv.Atoi(c.PostForm("item_id"))
	catId ,_ := strconv.Atoi(c.PostForm("cat_id"))
	membergroupid ,_ := strconv.Atoi(c.PostForm("member_group_id"))
	username := c.PostForm("username")
	data, code := models.SaveMember(itemId,catId,v,username,membergroupid)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}

func GetMemberList(c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)
	itemId , _ := strconv.Atoi(c.PostForm("item_id"))
	data ,code := models.GetList(itemId,v)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}

func DeleteMember(c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)
	itemId , _ := strconv.Atoi(c.PostForm("item_id"))
	itemMemberId , _ := strconv.Atoi(c.PostForm("item_member_id"))
	code = models.DeleteMember(itemId,v,itemMemberId)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}
