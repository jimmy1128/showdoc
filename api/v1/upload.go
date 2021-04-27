package v1

import (
	"awesomeProject3/utils"
	"awesomeProject3/utils/errmsg"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func UpLoad(c *gin.Context) {
	var code int
	var ext string
	var filename string //文件名称
	var url string //文件 url

	// single file

	file, err := c.FormFile("editormd-image-file")
	if err != nil {
		code = errmsg.ERROR
	}
	ext = file.Filename[strings.LastIndex(file.Filename,"."):]
	filename = utils.UniqueId() +ext


	path := "./upload/" + filename
	url = c.Request.URL.Host + "upload/" + filename
	log.Println(file.Filename,url)


	err = c.SaveUploadedFile(file, path)
	if err != nil {
		code = errmsg.ERROR
	}else {
		code = errmsg.SUCCESE
	}
	c.JSON(http.StatusOK, gin.H{
		"success":1,
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"url":     url,
	})

}
