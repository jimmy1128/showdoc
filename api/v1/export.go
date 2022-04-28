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
	defer os.RemoveAll("showdoc.zip")

}

func ExportWord(c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	v ,_ := user.(uint)
	itemId,_ := strconv.Atoi(c.Query("item_id"))
	CatId,_ := strconv.Atoi(c.Query("cat_id"))
	PageId,_ := strconv.Atoi(c.Query("page_id"))
	code = models.ExportWord(uint(itemId), uint(CatId), uint(PageId),v)

	c.Header("Content-type","application/octet-steam")
	c.Header("Content-Disposition","filename="+"1"+".doc")
	c.Header("Content-Description","File Transfer")
	c.Header("Content-Type","application/octet-stream")
	c.Header("Content-Transfer-Encoding","binary")
	c.Header("Expires","0")
	c.Header("Cache-Control","must-revalidate,post-check=0,pre-check=0")
	c.Header("Pragma","public")

}

func ExportPdf (c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)
	itemId,_ := strconv.Atoi(c.Param("item_id"))

	code = models.ExportPdf(uint(itemId), v)

	c.Header("Cache-Control","max-age=0")
	c.Header("Content-Description","File Transfer")
	c.Header("Content-Type","application/zip")
	c.Header("Content-Transfer-Encoding","binary")
	c.Header("Content-disposition","attachment; filename=showdoc.zip")
	stats, _ :=os.Stat("showdoc.zip")
	y := strconv.Itoa(int(stats.Size()))

	c.Header("Content-Length", y)
	c.File("showdoc.zip")
	defer os.RemoveAll("showdoc.zip")
}

