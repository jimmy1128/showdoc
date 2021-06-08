package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cookie (c *gin.Context){
	cookie, err := c.Cookie("gin_cookie")
	if err !=nil{
		cookie ="lang"
		//c.SetCookie("lang","CN",3600,"/","localhost",false,true)
		c.SetCookie("itemid","1",600,"/","localhost",false,true)
	}
	fmt.Printf("Cookie value: %s \n", cookie)
}
func NoResponse(c *gin.Context) {
	//返回404状态码
	c.JSON(http.StatusNotFound, gin.H{
		"status": 404,
		"error":  "404, page not exists!",
	})
}

