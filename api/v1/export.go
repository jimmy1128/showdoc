package v1

import (
	"awesomeProject3/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

func ExportItem(c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)
	itemId,_ := strconv.Atoi(c.Param("item_id"))
   code = models.ExportMarkdown(uint(itemId), v)
   c.Header("Cache-Control","max-age=0")
   c.Header("Content-Description","File Transfer")
   c.Header("Content-Type","application/zip")
   c.Header("Content-Transfer-Encoding","binary")
	c.Header("Content-disposition","attachment; filename=showdoc.zip")
	stats, _ :=os.Stat("showdoc.zip")
	y := strconv.Itoa(int(stats.Size()))

	c.Header("Content-Length", y)
	c.File("showdoc.zip")

}

