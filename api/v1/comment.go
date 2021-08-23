package v1

import (
	"awesomeProject3/models"
	"awesomeProject3/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddComment(c *gin.Context){
	var data models.Comment
	pageId,_ := strconv.Atoi(c.PostForm("page_id"))
	guestId,_ := strconv.Atoi(c.PostForm("guest_id"))
	username := c.PostForm("username")
	rate,_ := strconv.Atoi(c.PostForm("rate"))
	content := c.PostForm("content")

	data.Username = username
	data.PageId = uint(pageId)
	data.Rate = rate
	data.Content = content
	data.GuestId = uint(guestId)

	code = models.AddComment(&data)
	c.JSON (http.StatusOK,gin.H{
		"status" : code,
		"data" : data,
		"message":errmsg.GetErrMsg(code),
	})
}

func GetComment(c *gin.Context){
	id,_ := strconv.Atoi(c.PostForm("page_id"))
	data , code := models.GetComment(id)
	c.JSON(http.StatusOK,gin.H{
		"status" : code,
		"data" : data,
		"message":errmsg.GetErrMsg(code),
	})
}

func DeleteComment(c *gin.Context){
	id,_ := strconv.Atoi(c.PostForm("comment_id"))
	code = models.DeleteComment(uint(id))
	c.JSON(http.StatusOK,gin.H{
		"status" : code,
		"message" : errmsg.GetErrMsg(code),
	})
}

func GetCommentCount(c *gin.Context){
	id,_ :=strconv.Atoi(c.Param("id"))
	total := models.GetCommentCount(id)
	c.JSON(http.StatusOK,gin.H{
		"total" : total,
	})
}

func GetCommentListFront(c *gin.Context){
	id,_ := strconv.Atoi(c.PostForm("page_id"))
	pageSize,_ := strconv.Atoi(c.PostForm("count"))
	pageNum,_ := strconv.Atoi(c.PostForm("page"))

	switch{
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0 :
		pageSize = 10
	}
	if pageNum == 0{
		pageNum = 1
	}
	data ,total,code := models.GetCommentListFront(id,pageSize,pageNum)
	c.JSON(http.StatusOK,gin.H{
		"status" : code,
		"data" : data,
		"total" : total,
		"message" : errmsg.GetErrMsg(code),

	})
}