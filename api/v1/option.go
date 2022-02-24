package v1

import (
	"awesomeProject3/models"
	"awesomeProject3/utils/errmsg"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func SaveConfig(c *gin.Context) {
	var id int
	session := sessions.Default(c)
	user := session.Get("id")
	conf := c.PostForm("register_open")
	ldapOpen := c.PostForm("ldap_open")
	ldapForm := c.PostForm("ldap_form")
	if ldapOpen == "true" {
		id = 0
	} else if ldapOpen == "false" {
		id = 1
	}
	code1 := models.SaveLDapConfig(id, ldapForm)
	v, _ := user.(uint)
	if conf == "true" {
		id = 0
	} else if conf == "false" {
		id = 1
	}
	code2 := models.SaveConfig(v, id)
	if code1 != 200 {
		code = code1
	} else if code2 != 200 {
		code = code2
	} else {
		code = 200
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func LoadConfig(c *gin.Context) {
	data, code := models.LoadConfig()
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

func SaveLangConfig(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("id")
	conf, _ := strconv.Atoi(c.PostForm("lang"))
	v, _ := user.(uint)

	code = models.SaveLangConfig(v, conf)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
func SaveIconConfig(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("id")
	conf, _ := strconv.Atoi(c.PostForm("open_country_icon"))
	v, _ := user.(uint)

	code = models.SaveIconConfig(v, conf)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
func SaveCountryConfig(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("id")
	conf, _ := strconv.Atoi(c.PostForm("open_country_code"))
	v, _ := user.(uint)

	code = models.SaveCountryConfig(v, conf)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
func LoadLangConfig(c *gin.Context) {
	var access int
	session := sessions.Default(c)
	user := session.Get("id")
	v, _ := user.(uint)
	itemId, _ := strconv.Atoi(c.PostForm("id"))
	if v > 0 {
		access = 1
	}
	data, code := models.LoadLangConfig(itemId)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"access":  access,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

