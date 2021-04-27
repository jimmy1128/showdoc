package v1

import (
	"awesomeProject3/models"
	"awesomeProject3/utils/errmsg"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetRecycleList (c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)
	itemId , _ := strconv.Atoi(c.PostForm("item_id"))
	data , code := models.GetRecycleList(itemId , v)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}

func Recover (c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)
	itemId , _ := strconv.Atoi(c.PostForm("item_id"))
	pageId , _ := strconv.Atoi(c.PostForm("page_id"))
	code := models.Recover(itemId ,pageId, v)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}
