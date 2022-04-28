package v1

import (
	"awesomeProject3/models"
	"awesomeProject3/utils/errmsg"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteCatalogs(c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)
		itemid ,_ := strconv.Atoi(c.PostForm("item_id"))
		id ,_ := strconv.Atoi(c.Param("id"))
	code =models.DeleteCatalogs(itemid,id, int(v))

	c.JSON(http.StatusOK,gin.H{
		"status" : code,
		"message" : errmsg.GetErrMsg(code),
	})
}

func CatListGroup(c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)
	id, _ := strconv.Atoi(c.PostForm("item_id"))
	langId, _ := strconv.Atoi(c.PostForm("cid"))
	data, code := models.CatListName(uint(id),"",v,langId)
	c.JSON(http.StatusOK,gin.H{
		"status": code,
		"data" : data,
		"message" : errmsg.GetErrMsg(code),

	})
}

func EditCat(c *gin.Context){
	var data models.Catalogs
	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)
	id ,_:= strconv.Atoi(c.Param("item_id"))
	c.ShouldBindJSON(&data)
	code  =models.EditCatalogs(id, int(v))
	c.JSON(http.StatusOK,gin.H{
		"status" : code,
		"message" : errmsg.GetErrMsg(code),
	})

}

func AddCat(c *gin.Context){
var data models.Catalogs
_ =c.ShouldBindJSON(&data)
code =models.CreateCatalogs(&data)
c.JSON(http.StatusOK,gin.H{
	"status" : code,
	"data" : data,
	"message" : errmsg.GetErrMsg(code),
})
}

func SecondCatList(c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)
	id ,_:= strconv.Atoi(c.PostForm("item_id"))
	data, code := models.GetSecondCatalogsItemId(id,2, int(v))


	c.JSON(http.StatusOK,gin.H{
		"status" : code,
		"data" : data,
		"message" : errmsg.GetErrMsg(code),
	})

}

func ChildCatList(c *gin.Context){
	id ,_ := strconv.Atoi(c.PostForm("cat_id"))
	data,code := models.GetChildCatalogsByCatId(id)
	c.JSON(http.StatusOK,gin.H{
		"status" : code,
		"data" : data,
		"message" : errmsg.GetErrMsg(code),
	})

}

func GetDefaultCat(c *gin.Context){
	//var temp struct{
	//	DefaultCatId2 string `json:"default_cat_id2"`
	//	DefaultCatId3 string `json:"default_cat_id3"`
	//}
	session := sessions.Default(c)
	user := session.Get("id")
	 v := user.(uint)

	pageId ,_ := strconv.Atoi(c.PostForm("page_id"))
	itemId ,_ := strconv.Atoi(c.PostForm("item_id"))
	copyPageId ,_ := strconv.Atoi(c.PostForm("copy_page_id"))
	page_history_id ,_ := strconv.Atoi(c.PostForm("page_history_id"))

	data , code := models.Getdefaultcat(pageId,copyPageId,v,itemId,page_history_id)

	c.JSON(http.StatusOK,gin.H{
		"status" : 200 ,
		"data" : data,
		"message" : errmsg.GetErrMsg(code),
	})

}

func SaveCatalogs(c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	// v := user.(int)
	if user == nil {
		return
	}else {
		item_id ,_ := strconv.Atoi(c.PostForm("itemid"))
		cat_name  := c.PostForm("catname")
		s_number ,_ := strconv.Atoi(c.PostForm("snumber"))
		cat_id ,_ := strconv.Atoi(c.PostForm("catid"))
		parent_cat_id ,_ := strconv.Atoi(c.PostForm("parentcatid"))
		langid,_ := strconv.Atoi(c.PostForm("cid"))

		var catalog models.Catalogs
		catalog.Name = cat_name
		catalog.SNumber = uint(s_number)
		catalog.ID = uint(cat_id)
		catalog.ItemId = uint(item_id)
		catalog.Cid = langid

		var parent_id int
		if parent_cat_id == 0 {
			parent_id = 2
		}else {
			parent_id = 3
		}
		catalog.ParentCatId = uint(parent_cat_id)
		catalog.Level = uint(parent_id)
		data , code := catalog.Save()
		c.JSON(http.StatusOK,gin.H{
			"status" : code ,
			"data" : data,
			"message" : errmsg.GetErrMsg(code),
		})

		}
	}
	func GetPagesByCat(c *gin.Context){
		item_id ,_ := strconv.Atoi(c.PostForm("item_id"))
		cat_id ,_ := strconv.Atoi(c.PostForm("cat_id"))
		session := sessions.Default(c)
		user := session.Get("id")
		v , _ := user.(uint)
		fmt.Println(item_id,cat_id)
		data , code := models.GetPagesbycat(cat_id , item_id , v)
		c.JSON(http.StatusOK,gin.H{
			"status" : code ,
			"data" : data,
			"message" : errmsg.GetErrMsg(code),
		})
	}
	func BatUpdates(c *gin.Context){
		arrCat := c.PostForm("cats")
		itemId ,_ := strconv.Atoi(c.PostForm("item_id"))
		data, code := models.BatUpdate(arrCat , itemId)
		c.JSON(http.StatusOK,gin.H{
			"status":code ,
			"data":data,
			"message": errmsg.GetErrMsg(code),
		})
	}
