package v1

import (
	"awesomeProject3/models"
	"awesomeProject3/utils"
	"awesomeProject3/utils/errmsg"
	"encoding/base64"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"strconv"
)

func AddPage(c *gin.Context){
	var data models.Page
	_ =c.ShouldBindJSON(&data)
	code = models.CreatePage(&data)
	c.JSON(http.StatusOK,gin.H{
		"status" : code ,
		"data" : data,
		"message" : errmsg.GetErrMsg(code),
	})

}

func EditPage(c *gin.Context){
	var data models.Page
	id ,_ := strconv.Atoi(c.Param("page_id"))
	c.ShouldBindJSON(&data)
	code = models.EditPage(id,&data)
	c.JSON(http.StatusOK,gin.H{
		"status" : code ,
		"message" : errmsg.GetErrMsg(code),
	})
}

func SavePage (c * gin.Context){
	var page models.Page
	var pagehistory models.PageHistory
	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)
	id,_ := strconv.Atoi(c.PostForm("page_id" ))
	item_id,_ := strconv.Atoi(c.PostForm("item_id"))
	s_number,_ := strconv.Atoi(c.PostForm("s_number"))
	page_title := c.PostForm("page_title")
	page_content := c.PostForm("page_content")
	cat_id,_  := strconv.Atoi(c.PostForm("cat_id"))
	Urlencode,_ := strconv.Atoi(c.PostForm("is_urlencode"))
	langId , _ := strconv.Atoi(c.PostForm("cid"))
	if Urlencode == 1{
		page_content, _ = UrlDecode(page_content)
	}


	page.ID = uint(id)
	page.ItemId = uint(item_id)
	page.SNumber = uint(s_number)
	page.PageTitle = page_title
	page.PageContent = page_content
	page.CatId = uint(cat_id)
	page.Cid = langId
	page.AuthorUid = v

	pagehistory.PageId = int(page.ID)
	pagehistory.ItemId = int(page.ItemId)
	pagehistory.CatId = int(page.CatId)
	pagehistory.PageTitle = page.PageTitle
	s, _ := utils.GZipData([]byte(page.PageContent))
	pagehistory.PageContent = base64.StdEncoding.EncodeToString(s)
	pagehistory.PageComment = page.PageComments
	pagehistory.SNumber = int(page.SNumber)
	pagehistory.AuthorUid = int(page.AuthorUid)


	data , code := page.SavePage()
	_, _ = pagehistory.SaveHistoryPage()

    c.JSON(http.StatusOK,gin.H{
    	"status" : code ,
    	"data" : data,
    	"message" : errmsg.GetErrMsg(code),
	})

}

func GetInfo(c *gin.Context){
	id ,_ := strconv.Atoi(c.PostForm("page_id"))
	data , code := models.GetPage(id)
	c.JSON(http.StatusOK,gin.H{
		"status" : code,
		"data" : data,
		"message" : errmsg.GetErrMsg(code),
	})
}

func DeletePage(c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)
	id ,_ := strconv.Atoi(c.PostForm("page_id"))
	code = models.DeletePage(id,v)
	c.JSON(http.StatusOK,gin.H{
		"status" : code,
		"message" : errmsg.GetErrMsg(code),
	})
}

func IsLock(c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)
	id ,_ := strconv.Atoi(c.PostForm("page_id"))
	code ,data := models.IsLock(id ,v)
	c.JSON(http.StatusOK,gin.H{
		"status" : code,
		"data" : data,
		"message" : errmsg.GetErrMsg(code),
	})
}

func SetLock(c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)
	id ,_ := strconv.Atoi(c.PostForm("page_id"))
	itemId ,_ := strconv.Atoi(c.PostForm("item_id"))
	lockto ,_ := strconv.Atoi(c.PostForm("lock_to"))
	code , data := models.SetLock(id,itemId,lockto,v)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}
func UrlDecode(str string) (string, error) {
	return url.QueryUnescape(str)
}

func History(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)
	id ,_ := strconv.Atoi(c.PostForm("page_id"))
	code , data := models.History(id , v)
	c.JSON(http.StatusOK,gin.H{
		"status":code ,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}
func Sort(c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)
	itemId ,_ := strconv.Atoi(c.PostForm("item_id"))
	pages := c.PostForm("pages")
	code = models.Sort(pages,itemId,v)
	c.JSON(http.StatusOK,gin.H{
		"status":code ,
		"message":errmsg.GetErrMsg(code),
	})



}