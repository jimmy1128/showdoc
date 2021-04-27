package v1

import (
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// gin session key
const KEY = "Your Secret Key"

// 使用 Cookie 保存 session
func EnableCookieSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte(KEY))
	return sessions.Sessions("SESSIONID", store)
}
// 使用 内存 保存 session
func EnableMemorySession() gin.HandlerFunc {
	store := memstore.NewStore([]byte(KEY))
	return sessions.Sessions("SESSIONID", store)
}

//func HomeGet(c *gin.Context) {
//	//获取session，判断用户是否登录
//	islogin := GetSession(c)
//	c.HTML(http.StatusOK, "home.html", gin.H{"IsLogin": islogin})
//}

func ExitGet(c *gin.Context) {

	//清除该用户登录状态的数据
	session := sessions.Default(c)
	session.Delete("id")
	session.Save()
	//session.Clear()

}
