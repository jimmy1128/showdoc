package v1

import (
	"fmt"
	"github.com/gin-contrib/sessions"
"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	// login success
	if password != "" {
		session := sessions.Default(c)
		session.Set("username", username)
		session.Save()

		c.String(200, "登录成功")

	} else {
		c.String(401, "密码错误")
	}
}

func CurrentUserHandler(c *gin.Context) {
	session := sessions.Default(c)
	var user string
	sessionValue := session.Get("username")
	user = sessionValue.(string)
	fmt.Println(user)

	c.String(200, user)
}
