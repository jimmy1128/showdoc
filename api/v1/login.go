package v1

import (
	"awesomeProject3/middleware"
	"awesomeProject3/models"
	"awesomeProject3/utils/errmsg"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var formData models.User
	var token string
	var code int
	c.ShouldBindJSON(&formData)
	password := formData.Password
	name := formData.Username
	formData, code ,ldapStatus := models.CheckLogin(formData.Username, formData.Password)
	if code != errmsg.SUCCESE && ldapStatus=="1"{
		formData, code = models.CheckLdapLogin(name, password)
	}

	if code == errmsg.SUCCESE {
		token, code = middleware.SetToken(formData.Username)
		session := sessions.Default(c)
		session.Set("id", formData.ID) // In real world usage you'd set this to the users ID
		session.Save()
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
		"data":    formData.Username,
		"id":      formData.ID,
	})
}

func LoginFront(c *gin.Context) {
	var formData models.Guest
	_ = c.ShouldBindJSON(&formData)
	var token string
	var code int

	formData, code = models.CheckLoginFront(formData.GUsername, formData.GPassword)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    formData.GName,
		"id":      formData.ID,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}

type UpToken struct {
	Token string `json:"token"`
}

//验证
func CheckToken(c *gin.Context) {
	var Token UpToken
	_ = c.ShouldBindJSON(&Token)
	_, code = middleware.CheckToken(Token.Token)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
