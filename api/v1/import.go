package v1

import (
	"awesomeProject3/models"
	"awesomeProject3/utils/errmsg"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ImportItem(c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)
	file, _ := c.FormFile("file")
	path  := "./upload/" +file.Filename
	 c.SaveUploadedFile(file, path)
	code = models.Auto("./upload/" +file.Filename,v)

	c.JSON(http.StatusOK,gin.H{
		"status":code ,
		"message":errmsg.GetErrMsg(code),
	})
}
