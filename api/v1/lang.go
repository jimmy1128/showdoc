package v1

import (
	"awesomeProject3/models"
	"awesomeProject3/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddLang(c *gin.Context){
	var data models.Lang
	id ,_ := strconv.Atoi(c.PostForm("id"))
	name  := c.PostForm("name")
	icon  := c.PostForm("icon")
	data.Name = name
	data.Icon = icon
	if id != 0 {
		code = models.EditLang(id,&data)
	}else {
		code = models.CheckLang(data.Name)
		if code == errmsg.SUCCESE{
			models.CreateLang(&data)
		}
	}

	if code == errmsg.ERROR_LANG_USED{
		code = errmsg.ERROR_LANG_USED
	}
	c.JSON(http.StatusOK ,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}
//查询单个用户
//TODO 查询分类下的所有文章
func GetLangInfo(c *gin.Context){
	id ,_ := strconv.Atoi(c.PostForm("id"))
	data,code := models.GetLangInfo(id)

	c.JSON(http.StatusOK ,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})

}
//用户列表
func GetLang(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.PostForm("count"))
	pageNum, _ := strconv.Atoi(c.PostForm("page"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageSize = -1
	}
	data,total :=models.GetLang(pageSize,pageNum)
	if data != nil{
		code = 200
	}

	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"total":total,
		"message":errmsg.GetErrMsg(code),
	})
}
//编辑分类
func EditLang(c *gin.Context){
	var data models.Lang
	id ,_:= strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
		code = models.EditLang(id,&data)
	if code == errmsg.ERROR_CATENAME_USED{
		c.Abort()
	}
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}
//删除分类
func DeleteLang(c *gin.Context){
	id ,_:= strconv.Atoi(c.PostForm("id"))
	code = models.DeleteLang(id)

	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}