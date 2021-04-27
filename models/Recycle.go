package models

import (
	"awesomeProject3/utils/errmsg"
	"github.com/go-sql-driver/mysql"
)

//获取被删除的页面列表
func GetRecycleList(itemId int,uid uint)([]Page,int){
	var page []Page
	if checkItemCreator(uid,itemId) != errmsg.SUCCESE{
		return page,errmsg.ERROR_USER_NO_RIGHT
	}
	if itemId > 0 {
		db.Unscoped().Where("item_id =?",itemId).Where("is_del =1").Find(&page)
	}
return page, errmsg.SUCCESE
}

//页面恢复
func Recover (itemId int,pageId int,uid uint)int{
	//var page Page

if checkItemCreator(uid,itemId) != errmsg.SUCCESE{
	return errmsg.ERROR_USER_NO_RIGHT
}
if itemId > 0{
	err := db.Unscoped().Model(Page{}).Where("id =?",pageId).Update("deleted_at",mysql.NullTime{}).Update("is_del",0).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESE
}
return 0
}