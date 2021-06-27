package v1

import "C"
import (
	"awesomeProject3/models"
	"awesomeProject3/utils"
	"awesomeProject3/utils/errmsg"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteItem(c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)
	id ,_:= strconv.Atoi(c.PostForm("item_id"))
	password := c.PostForm("password")
	code =models.DeleteItem(id,password,v)
	c.JSON(http.StatusOK,gin.H{
		"status" :code,
		"message" : errmsg.GetErrMsg(code),
	})
}

func ItemsInfo(c *gin.Context){
	var access int
	id ,_ := strconv.Atoi(c.PostForm("id"))
	keyword := c.PostForm("keyword")
	langId,_ := strconv.Atoi(c.PostForm("lang_id"))
	session := sessions.Default(c)
	user := session.Get("id")
	defaultpageid :=c.PostForm("defaultpageid")
	v , _ := user.(uint)
	i , _ :=c.Cookie("visit_item")
	item, code := models.GetOneItem(v,id,i)
	itemInfo := item.GetItemInfo(keyword,langId,v,uint(id), defaultpageid)
	if user == nil {
		itemInfo.IsLogin = false

	} else {
		itemInfo.IsLogin = true
		if itemInfo.IsLogin == true {
			access = 1
		}
	}
	if item.UserId == user{
		itemInfo.ItemCreator = true
		itemInfo.ItemPermn = true
	}

	if item.Isarchived == 1 {
		itemInfo.ItemCreator = false
		itemInfo.ItemPermn = false
	}else if models.CheckItemPermn(v,id) == errmsg.SUCCESE{
		itemInfo.ItemPermn = true
	}

	c.JSON(http.StatusOK,gin.H{
		"status" : code,
		"access" : access,
		"data" : itemInfo,
		"message" : errmsg.GetErrMsg(code),
	})
}

func MyList(c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)
	data ,code := models.GetMyItem(v)
	c.JSON(http.StatusOK,gin.H{
		"status" : code ,
		"data" : data,
		"message" : errmsg.GetErrMsg(code),
	})
}

func ItemDetail(c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)
	id ,_ := strconv.Atoi(c.PostForm("item_id"))
	data ,code := models.GetItemDetail(id,v)
	c.JSON(http.StatusOK,gin.H{
		"status": code ,
		"data" : data ,
		"message" : errmsg.GetErrMsg(code),
	})
}

func UpdateItem(c *gin.Context) {

	id ,_ := strconv.Atoi(c.PostForm("item_id"))
	title := c.PostForm("title")
	description  := c.PostForm("description")
	password := c.PostForm("password")
	langlist := c.PostForm("langlist")


	code = models.EditItem(id,title,description,password,langlist)

	c.JSON(http.StatusOK,gin.H{
		"status" : code ,
		"message" : errmsg.GetErrMsg(code),
	})
}

func AddItem (c *gin.Context){
	var data models.Item
	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)

	item_type , _ := strconv.Atoi(c.PostForm("types"))
	item_name  := c.PostForm("titlename")
	item_description := c.PostForm("description")
	password := c.PostForm("password")
	langlist := c.PostForm("langlist")

	data.Title = item_name
	data.Types = item_type
	data.Description = item_description
	data.LangList = langlist
	data.UserId = v

	code =models.CreateItem(&data,password )
	c.JSON(http.StatusOK,gin.H{
		"status" : code ,
		"data" : data,
		"message" : errmsg.GetErrMsg(code),
	})
}

func ExitItem (c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)
	id ,_ := strconv.Atoi(c.PostForm("item_id"))
	code = models.ExitTeam(id , v)

	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}

func GetKey (c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)
	id ,_ := strconv.Atoi(c.PostForm("item_id"))
	data,code := models.GetKey(id,v)
	c.JSON(http.StatusOK,gin.H{
		"status" : code ,
		"data" : data,
		"message" : errmsg.GetErrMsg(code),
	})
}

func GetReset(c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)
	id ,_ := strconv.Atoi(c.PostForm("item_id"))
	data,code := models.ResetKey(id,v)
	c.JSON(http.StatusOK,gin.H{
		"status" : code ,
		"data" : data,
		"message" : errmsg.GetErrMsg(code),
	})

}

func AttornItem(c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)
	id ,_ := strconv.Atoi(c.PostForm("item_id"))
	username := c.PostForm("username")
	password:= c.PostForm("password")

	code := models.AttornItem(id,username,password,v)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}

func ArchiveItem (c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)
	id ,_ := strconv.Atoi(c.PostForm("item_id"))
	password:= c.PostForm("password")
	code := models.Archive(id,password,v)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}

func LinkItem (c *gin.Context){
	itemId ,_ := strconv.Atoi(c.PostForm("item_id"))
	id ,_ := strconv.Atoi(c.PostForm("id"))
	lang ,_ := strconv.Atoi(c.PostForm("lang"))
	code := models.LinkItem(itemId,id,lang)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})

}

func Pwd (c *gin.Context){
	itemId ,_ := strconv.Atoi(c.PostForm("item_id"))
	//pageId ,_ := strconv.Atoi(c.PostForm("page_id"))
	password:= c.PostForm("password")
	code := models.Pwd(itemId,password)
	if code == errmsg.SUCCESE{
		c.SetCookie("visit_item", utils.Md5(string(itemId)),utils.MaxAge,"/",utils.Domain,false,true)
	}
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}
func SortByItem(c *gin.Context){
	session := sessions.Default(c)
	user := session.Get("id")
	v , _ := user.(uint)
	data:= c.PostForm("data")
	code := models.SortByItem(int(v), data)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}