package v1

import (
	"awesomeProject3/models"
	"awesomeProject3/utils/errmsg"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//var code int
//添加新的template
func AddTemplate(c *gin.Context){
	var data models.Template
	_ = c.ShouldBindJSON(&data)
	code = models.CreateTemplate(&data)
	//if code != errmsg.SUCCESE{
		c.JSON(http.StatusOK,gin.H{
			"status" : code ,
			"data" : data,
			"message" : errmsg.GetErrMsg(code),
		})
	//}
}

//修改template
func EditTemplate(c *gin.Context){
	var data models.Template
	id ,_:= strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code = models.EditTemplate(id,&data)
	c.JSON(http.StatusOK,gin.H{
		"status" : code ,
		"message" : errmsg.GetErrMsg(code),
	})

}
// 获取template
func GetList(c *gin.Context){

	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)

data , code := models.GetTemplateByUid(v)
c.JSON(http.StatusOK,gin.H{
	"status" : code ,
	"data" : data,
	"message" : errmsg.GetErrMsg(code),
})
}
// 删除template
func DeleteTemplate (c *gin.Context){
	id ,_ := strconv.Atoi(c.Param("id"))
	code = models.DeleteTemplate(id)
	c.JSON(http.StatusOK,gin.H{
		"status" : code,
		"message" : errmsg.GetErrMsg(code),
	})
}

func SaveTemp(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("id")
	v := user.(int)
	if user == nil {
		return
	}else {
		title := c.PostForm("template_title")
		content := c.PostForm("template_content")

		var template models.Template
		template.Title = title
		template.Content = content
		template.UserId = uint(v)

		data , code := template.Save()
		c.JSON(http.StatusOK,gin.H{
			"status" : code ,
			"data" : data,
			"message" : errmsg.GetErrMsg(code),
		})
	}

}






