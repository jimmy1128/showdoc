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


	path := utils.Upload + filename
	url = c.Request.URL.Host + "docupload/" + filename
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
func UpLoadImg(c *gin.Context) {
	var code int
	var url string //文件 url
	// single file
	file, err := c.FormFile("file")
	if err != nil {
		code = errmsg.ERROR
	}
	log.Println(file.Filename)

	err = c.SaveUploadedFile(file, utils.Upload+file.Filename)
	url = c.Request.URL.Host + "docupload/" + file.Filename
	if err != nil {
		code = errmsg.ERROR
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"url":     url,
	})
}
