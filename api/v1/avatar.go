package v1

import (
	"awesomeProject3/models"
	"awesomeProject3/utils/errmsg"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func SaveAvatar(c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)
	itemId ,_ := strconv.Atoi(c.PostForm("item_id"))
	avatarName  := c.PostForm("avatar_name")
	uri  := c.PostForm("uri")
	code = models.SaveAvatar(itemId,avatarName,uri,v)
	c.JSON(http.StatusOK,gin.H{
		"status" : code,
		"message":errmsg.GetErrMsg(code),
	})
}

func GetAvatar (c *gin.Context){

	itemId ,_ := strconv.Atoi(c.PostForm("item_id"))
	data , code := models.GetAvatar( itemId,)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}

func SaveHeader (c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)
	itemId ,_ := strconv.Atoi(c.PostForm("item_id"))
	data := c.PostForm("data")
	code = models.SaveHeader(itemId,v,data)
	c.JSON(http.StatusOK,gin.H{
		"status":code ,
		"message" : errmsg.GetErrMsg(code),

	})
}
func GetHeader(c *gin.Context){
	itemId ,_ := strconv.Atoi(c.PostForm("item_id"))
	langId ,_ := strconv.Atoi(c.PostForm("lang_id"))
	data , code := models.GetHeader(itemId,langId)
	c.JSON(http.StatusOK,gin.H{
		"status":code ,
		"data":data,
		"message": errmsg.GetErrMsg(code),
	})
}

func DeleteHeader (c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)
	headerId ,_ := strconv.Atoi(c.PostForm("header_id"))
	code = models.DeleteHeader(v , headerId)
	c.JSON(http.StatusOK , gin.H{
		"status": code ,
		"message": errmsg.GetErrMsg(code),
	})

}
