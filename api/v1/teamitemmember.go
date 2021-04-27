package v1

import (
	"awesomeProject3/models"
	"awesomeProject3/utils/errmsg"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func SaveTeamItemMember(c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	v := user.(uint)
	id , _ := strconv.Atoi(c.PostForm("id"))
	catId , _ := strconv.Atoi(c.PostForm("cat_id"))
	memberGroupId , _ := strconv.Atoi(c.PostForm("memberGroupId"))
	data,code := models.SaveTeamItemMember(id , catId, v,memberGroupId)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}
func GetTeamItemMemberList(c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	v := user.(uint)
	teamId , _ := strconv.Atoi(c.PostForm("team_id"))
	itemId , _ := strconv.Atoi(c.PostForm("item_id"))
	data , code := models.GetTeamItemMemberList(v,itemId,teamId)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}