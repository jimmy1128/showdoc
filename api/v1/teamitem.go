package v1

import (
	"awesomeProject3/models"
	"awesomeProject3/utils/errmsg"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func SaveTeamItem(c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	v := user.(uint)
	teamId , _ := strconv.Atoi(c.PostForm("team_id"))
	itemId := c.PostForm("item_id")
	data,code := models.SaveTeamItem(v,itemId,teamId)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})

}

func GetTeamItemList(c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	v := user.(uint)
	itemId , _ := strconv.Atoi(c.PostForm("item_id"))
	data , code := models.GetTeamItemList(v,itemId)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}

func GetListByTeam (c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	v := user.(uint)
	teamId , _ := strconv.Atoi(c.PostForm("team_id"))
	fmt.Println(teamId)
	data , code := models.GetListByTeam(teamId,v)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}

func DeleteTeamItem(c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	v := user.(uint)
	id , _ := strconv.Atoi(c.PostForm("id"))
	code := models.DeleteTeamItem(id , v)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}
