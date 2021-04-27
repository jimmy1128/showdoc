package v1

import (
	"awesomeProject3/models"
	"awesomeProject3/utils/errmsg"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func SaveTeamMember(c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	v := user.(uint)
	teamId,_:= strconv.Atoi(c.PostForm("teamid"))
	username := c.PostForm("member_username")
	data,code := models.SaveTeamMember(teamId, v,username)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}
func GetTeamMemberList(c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	v := user.(uint)
	teamId,_:= strconv.Atoi(c.PostForm("teamid"))
	data , code := models.GetTeamMemberList(teamId, uint(v))
	c.JSON(http.StatusOK,gin.H{
	"status":code,
	"data":data,
	"message":errmsg.GetErrMsg(code),
	})
}

func DeleteTeamMember (c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	v := user.(uint)
	teamMemberId,_:= strconv.Atoi(c.PostForm("id"))
	code := models.DeleteTeamMember(teamMemberId,v)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}