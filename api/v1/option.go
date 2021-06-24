package v1

import (
	"awesomeProject3/models"
	"awesomeProject3/utils/errmsg"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func SaveConfig (c *gin.Context){
	var id int
	session := sessions.Default(c)
	user := session.Get("id")
	conf := c.PostForm("register_open")
	v , _ := user.(uint)
	if conf == "true"{
		id = 0
	}else if conf == "false" {
		id = 1
	}
	code = models.SaveConfig(v , id)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}

func LoadConfig (c *gin.Context){
	data,code := models.LoadConfig()
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}

func SaveLangConfig(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("id")
	conf,_ := strconv.Atoi(c.PostForm("lang"))
	v , _ := user.(uint)

	code = models.SaveLangConfig(v , conf)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}
func LoadLangConfig(c *gin.Context) {
	var access int
	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)
	if v > 0 {
		access = 1
	}
	data,code := models.LoadLangConfig()
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"access":access,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}
