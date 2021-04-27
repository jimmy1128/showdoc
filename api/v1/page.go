package v1

import (
	"awesomeProject3/models"
	"awesomeProject3/utils/errmsg"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
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
	id,_ := strconv.Atoi(c.PostForm("page_id" ))
	item_id,_ := strconv.Atoi(c.PostForm("item_id"))
	s_number,_ := strconv.Atoi(c.PostForm("s_number"))
	page_title := c.PostForm("page_title")
	page_content := c.PostForm("page_content")
	cat_id,_  := strconv.Atoi(c.PostForm("cat_id"))


	page.ID = uint(id)
	page.ItemId = uint(item_id)
	page.SNumber = uint(s_number)
	page.PageTitle = page_title
	page.PageContent = page_content
	page.CatId = uint(cat_id)


	data , code := page.SavePage()
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