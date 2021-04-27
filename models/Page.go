package models

import (
	"awesomeProject3/utils/errmsg"
	"github.com/jinzhu/gorm"
	"gorm.io/plugin/soft_delete"
	"time"
)

type Page struct {
	gorm.Model
	ItemId uint `gorm:"type:int;not null" json:"itemid"`
	AuthorUid uint `gorm:"type:int;not null" json:"authoruid"`
	CatId uint `gorm:"type:int;not null" json:"catid"`
	Name string `gorm:"type:varchar(255);not null" json:"catname"`
	PageTitle string `gorm:"type:varchar(255);not null" json:"pagetitle"`
	PageContent string `gorm:"type:longtext" json:"pagecontent"`
	SNumber uint `gorm:"type:int;not null" json:"snumber"`
	IsDel soft_delete.DeletedAt `gorm:"softDelete:flag;default:0"`
	DelName string `gorm:"type:varchar(255)" json:"del_name"`
}
type PageLock struct {
	gorm.Model
	PageId       int `gorm:"type:int;default:0" json:"pageId"`
	LockUid      int `gorm:"type:int;default:0" json:"lockUid"`
	LockUsername string `gorm:"type:varchar(255);default:''" json:"lockUsername"`
	LockTo       int `gorm:"type:int;default:0" json:"lockTo"`

}
type CheckLock struct {
	Lock         int `gorm:"type:int" json:"lock"`
	LockUid      int `gorm:"type:int" json:"lockUid"`
	LockUsername string `gorm:"type:varchar(255)" json:"lockUsername"`
	IsCurUser    int `gorm:"type:int" json:"isCurUser"`
}


func GetPagesByCatId(id uint ,keyword string ,itemId uint)[]*Page{
	var pages []*Page
	if keyword == "" {
		err := db.Where("item_id =?",itemId).Where("cat_id =?",id).Find(&pages).Error
		if err != nil{
			return pages
		}
		return pages
	}
	err := db.Where("item_id =?",itemId).Where("cat_id =?",id).Where("page_content LIKE ? OR page_title LIKE ?","%"+keyword+"%","%"+keyword+"%").Find(&pages).Error
	if err != nil{
		return pages
	}
	return pages
}

func GetPagesByItemId(id uint , keyword string)[]*Page{
	var pages []*Page
	if keyword == "" {
		err := db.Where("item_id =?",id).Where("cat_id =?",0).Find(&pages).Error
		if err != nil{
			return pages
		}
		return pages
	}
	err := db.Where("item_id =?",id).Where("cat_id =?",0).Where("page_content LIKE ? OR page_title LIKE ?","%"+keyword+"%","%"+keyword+"%").Find(&pages).Error
	if err != nil{
		return pages
	}
	return pages
}

func GetPage(id int)(Page ,int){
	var page Page
	err := db.Where("id =?",id).First(&page).Error
	if err !=nil{
		return page ,errmsg.ERROR
	}
	return page ,errmsg.SUCCESE
}

func DeletePage (id int,v uint)int{
	var page Page
	var user User
db.Where("id=?",v).Find(&user)
	db.Unscoped().Model(page).Where("id =?",id).Update("is_del",1).Update("del_name",user.Username)
	err := db.Where("id =?",id).Delete(&page).Error

	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESE
}

func CreatePage(page *Page)int{
	err := db.Create(&page).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESE
}

func EditPage(id int , page *Page)int{
	var maps = make(map[string]interface{})
	maps["item_id"]= page.ID
	maps["author_uid"] = page.AuthorUid
	maps["cat_id"] = page.CatId
	maps["page_title"] = page.PageTitle
	maps["page_content"] = page.PageContent
	maps["s_number"] = page.SNumber
	err = db.Model(&Catalogs{}).Where("id =?",id).Update(maps).Error
	if err !=nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESE
}

func (data *Page) SavePage()(*Page,int){
	if data.ID == 0 {
		err := db.Create(&data).Error
		if err != nil {
			return data, errmsg.ERROR
		}
		return data, errmsg.SUCCESE
	} else {
		err = db.Model(Page{}).Where("id =?",data.ID).Update(&data).Error
		if err !=nil{
			return data,errmsg.ERROR
		}
		return data,errmsg.SUCCESE
	}
}

func IsLock(pageId int,uid uint)(int,CheckLock) {

var pagelock []PageLock
	var checkLock CheckLock
	var curUser int
	var start int
	t := time.Now()
	start = int(t.Unix())
lock := 0
db.Model(PageLock{}).Where("page_id =?",pageId).Where("lock_to > ?",start).Find(&pagelock)

if &pagelock != nil {
	lock =1
}
	for _, pageLock1 := range pagelock {
		if pageLock1.LockUid == int(uid) {
			curUser = 1
		}else {
			curUser = 0
		}
		checkLock.Lock = lock
		checkLock.LockUid= pageLock1.LockUid
		checkLock.LockUsername = pageLock1.LockUsername
		checkLock.IsCurUser = curUser
	}
return errmsg.SUCCESE,checkLock
}

func SetLock (pageId int , itemId int ,lockTo int, uid uint)(int,PageLock){
	var pageLock PageLock
	var user User

	db.Model(User{}).Where("id = ?",uid).Select("username").Scan(&user)
	if CheckItemPermn(uid,itemId) != errmsg.SUCCESE {
		return errmsg.ERROR,pageLock
	}
	t := time.Now()
	lockTo = int(t.Unix()) // seconds since 1970
	db.Model(PageLock{}).Where("page_id = ?",pageId ).Delete(&pageLock)
	pageLock.PageId = pageId
	pageLock.LockUid = int(uid)
	pageLock.LockUsername = user.Username
	pageLock.LockTo = lockTo + 5*60
	err := db.Model(PageLock{}).Create(&pageLock).Error
	if err != nil {
		return errmsg.ERROR,pageLock
	}
	db.Model(PageLock{}).Where("lock_to < ?",lockTo).Delete(&pageLock)
	return errmsg.SUCCESE,pageLock
}