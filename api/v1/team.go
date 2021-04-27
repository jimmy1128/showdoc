package v1

import (
	"awesomeProject3/models"
	"awesomeProject3/utils/errmsg"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func SaveTeam (c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	v := user.(uint)
	if user == nil {
		return
	}else {
		id , _ := strconv.Atoi(c.PostForm("id"))
		teamName :=c.PostForm("team_name")
		var team models.Team
		team.ID = uint(id)
		team.Teamname = teamName
		team.Uid = uint(v)
		data , code := team.SaveTeam()
		c.JSON(http.StatusOK,gin.H{
			"status" : code ,
			"data" : data,
			"message" : errmsg.GetErrMsg(code),
		})
	}
}

func GetTeamList(c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)

    data, code := models.List(int(v))
	c.JSON(http.StatusOK,gin.H{
		"status" : code ,
		"data" : data,
		"message" : errmsg.GetErrMsg(code),
	})
}

func DeleteTeam(c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)
	id , _ := strconv.Atoi(c.PostForm("id"))
   code := models.DeleteTeam(int(v),id)
	c.JSON(http.StatusOK,gin.H{
		"status" : code ,
		"message" : errmsg.GetErrMsg(code),
	})
}

func AttornTeam (c *gin.Context){
	session :=sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)
	id , _ := strconv.Atoi(c.PostForm("team_id"))
	username := c.PostForm("username")
	password := c.PostForm("password")
	code := models.Attorn(id, int(v),username,password)
	c.JSON(http.StatusOK,gin.H{
		"status" : code ,
		"message": errmsg.GetErrMsg(code),
	})
}
