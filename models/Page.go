package models

import (
	"awesomeProject3/utils"
	"awesomeProject3/utils/errmsg"
	"encoding/base64"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"gorm.io/plugin/soft_delete"
	"html"
	"time"
)

type Page struct {
	gorm.Model
	Lang         Lang                  `gorm:"foreignKey:cid"`
	ItemId       uint                  `gorm:"type:int;not null" json:"itemid"`
	AuthorUid    uint                  `gorm:"type:int;not null" json:"authoruid"`
	CatId        uint                  `gorm:"type:int;not null" json:"catid"`
	PageTitle    string                `gorm:"type:varchar(255);not null" json:"pagetitle"`
	PageContent  string                `gorm:"type:longtext" json:"pagecontent"`
	SNumber      uint                  `gorm:"type:int;not null" json:"snumber"`
	IsDel        soft_delete.DeletedAt `gorm:"softDelete:flag;default:0"`
	DelName      string                `gorm:"type:varchar(255)" json:"del_name"`
	PageComments string                `gorm:"type:text" json:"pagecomments"`
	Cid          int                   `gorm:"type:int;not null" json:"cid"`
	GroupId      uint                  `gorm:"type:int;not null" json:"groupId"`
}
type PageLock struct {
	gorm.Model
	PageId       int    `gorm:"type:int;default:0" json:"pageId"`
	LockUid      int    `gorm:"type:int;default:0" json:"lockUid"`
	LockUsername string `gorm:"type:varchar(255);default:''" json:"lockUsername"`
	LockTo       int    `gorm:"type:int;default:0" json:"lockTo"`
}
type CheckLock struct {
	Lock         int    `gorm:"type:int" json:"lock"`
	LockUid      int    `gorm:"type:int" json:"lockUid"`
	LockUsername string `gorm:"type:varchar(255)" json:"lockUsername"`
	IsCurUser    int    `gorm:"type:int" json:"isCurUser"`
}
type PageHistory struct {
	gorm.Model
	PageId         int    `gorm:"type:int;not null" json:"page_id"`
	ItemId         int    `gorm:"type:int;not null" json:"item_id"`
	CatId          int    `gorm:"type:int;not null" json:"cat_id"`
	PageTitle      string `gorm:"type:varchar(255);not null" json:"pagetitle"`
	PageContent    string `gorm:"type:longtext" json:"pagecontent"`
	SNumber        int    `gorm:"type:int;not null" json:"snumber"`
	AuthorUid      int    `gorm:"type:int;not null" json:"authoruid"`
	AuthorUsername string `gorm:"type:varchar(255)" json:"authorusername"`
	PageComment    string `gorm:"type:text" json:"pagecomment"`
}
type Pages struct {
	PageId  int
	SNumber int
}
type DiffPage struct {
	 Page Page
     PageHistory PageHistory
}



func GetPagesByCatId(id uint, keyword string, itemId uint) []*Page {
	var pages []*Page
	if keyword == "" {
		err := db.Where("item_id =?", itemId).Where("cat_id =?", id).Order("s_number asc , id asc").Find(&pages).Error
		if err != nil {
			return pages
		}
		return pages
	}
	err := db.Where("item_id =?", itemId).Where("cat_id =?", id).Where("page_content LIKE ? OR page_title LIKE ?", "%"+keyword+"%", "%"+keyword+"%").Order("s_number asc , id asc").Find(&pages).Error
	if err != nil {
		return pages
	}
	return pages
}

func GetPagesByLangId (itemId int , langId int)([]*Page,int){
	var pages []*Page
	err = db.Model(Page{}).Where("item_id =?",itemId).Where("cid =?",langId).Order("group_id").Find(&pages).Error
	if err != nil {
		return nil,errmsg.ERROR
	}
return pages,errmsg.SUCCESE
}

func GetPagesByItemId(id uint, keyword string, langId int) []*Page {
	var pages []*Page
	if langId <= 0 {
		db.Where("item_id =?", id).Preload("Lang").Where("cat_id =?", 0).Where("page_content LIKE ? OR page_title LIKE ?", "%"+keyword+"%", "%"+keyword+"%").Find(&pages)
		return pages
	}
	if keyword == "" {
		err := db.Where("item_id =?", id).Preload("Lang").Where("cat_id =?", 0).Where("cid = ?", langId).Find(&pages).Error
		if err != nil {
			return pages
		}
		return pages
	}
	err := db.Where("item_id =?", id).Preload("Lang").Where("cid = ?", langId).Where("page_content LIKE ? OR page_title LIKE ?", "%"+keyword+"%", "%"+keyword+"%").Find(&pages).Error
	if err != nil {
		return pages
	}
	return pages
}

func GetPage(id int) (Page, int) {
	var page Page
	err := db.Where("id =?", id).Preload("Lang").First(&page).Error
	if err != nil {
		return page, errmsg.ERROR
	}
	return page, errmsg.SUCCESE
}

func DeletePage(id int, v uint) int {
	var page Page
	var user User
	db.Where("id=?", v).Find(&user)
	db.Unscoped().Model(page).Where("id =?", id).Update("is_del", 1).Update("del_name", user.Username)
	err := db.Where("id =?", id).Delete(&page).Error

	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESE
}

func CreatePage(page *Page) int {
	err := db.Create(&page).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESE
}

func EditPage(id int, page *Page) int {
	var maps = make(map[string]interface{})
	maps["item_id"] = page.ID
	maps["author_uid"] = page.AuthorUid
	maps["cat_id"] = page.CatId
	maps["page_title"] = page.PageTitle
	maps["page_content"] = page.PageContent
	maps["s_number"] = page.SNumber
	err = db.Model(&Catalogs{}).Where("id =?", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESE
}

func (data *Page) SavePage() (*Page, int) {
	var option Options
	db.Model(Options{}).Where("option_name = ?","lang").Find(&option)
	if data.Cid == 0 && option.OptionId != 0   {
		return data, errmsg.ERROR_LANG_EMPTY
	}

	if data.ID == 0 {
		err := db.Create(&data).Error
		if err != nil {
			return data, errmsg.ERROR
		}
	} else {
		err = db.Model(Page{}).Where("id =?", data.ID).Update(&data).Error
		if err != nil {
			return data, errmsg.ERROR
		}
	}
	return data , errmsg.SUCCESE
}
func (data *PageHistory) SaveHistoryPage() ([]PageHistory, int) {
	var count int
	var pagehistory []PageHistory
	var user User
	db.Model(User{}).Where("id = ?", data.AuthorUid).Find(&user)
	data.AuthorUsername = user.Username
	err := db.Model(PageHistory{}).Create(&data).Error
	if err != nil {
		return pagehistory, errmsg.ERROR
	}
	db.Model(PageHistory{}).Where("page_id =?", data.PageId).Count(&count)
	if count > 20 {
		db.Model(PageHistory{}).Where("page_id =?", data.PageId).Limit(1).Last(&pagehistory).Delete(pagehistory)
	}
	return pagehistory, errmsg.SUCCESE
}

func IsLock(pageId int, uid uint) (int, CheckLock) {

	var pagelock []PageLock
	var checkLock CheckLock
	var curUser int
	var start int
	t := time.Now()
	start = int(t.Unix())
	lock := 0
	db.Model(PageLock{}).Where("page_id =?", pageId).Where("lock_to > ?", start).Find(&pagelock)

	if &pagelock != nil {
		lock = 1
	}
	for _, pageLock1 := range pagelock {
		if pageLock1.LockUid == int(uid) {
			curUser = 1
		} else {
			curUser = 0
		}
		checkLock.Lock = lock
		checkLock.LockUid = pageLock1.LockUid
		checkLock.LockUsername = pageLock1.LockUsername
		checkLock.IsCurUser = curUser
	}
	return errmsg.SUCCESE, checkLock
}

func SetLock(pageId int, itemId int, lockTo int, uid uint) (int, PageLock) {
	var pageLock PageLock
	var user User

	db.Model(User{}).Where("id = ?", uid).Select("username").Scan(&user)
	if CheckItemPermn(uid, itemId) != errmsg.SUCCESE {
		return errmsg.ERROR, pageLock
	}
	t := time.Now()
	lockTo = int(t.Unix()) // seconds since 1970
	db.Model(PageLock{}).Unscoped().Where("page_id = ?", pageId).Delete(&pageLock)
	pageLock.PageId = pageId
	pageLock.LockUid = int(uid)
	pageLock.LockUsername = user.Username
	pageLock.LockTo = lockTo + 5*60
	err := db.Model(PageLock{}).Create(&pageLock).Error
	if err != nil {
		return errmsg.ERROR, pageLock
	}
	db.Model(PageLock{}).Unscoped().Where("lock_to < ?", lockTo).Delete(&pageLock)
	return errmsg.SUCCESE, pageLock
}

func History(pageId int, uid uint) (int, []PageHistory) {
	var page Page
	var pagehistory []PageHistory
	db.Model(Page{}).Where("id = ?", pageId).Find(&page)
	if checkItemVisit(uid, int(page.ItemId), "") != errmsg.SUCCESE {
		return errmsg.ERROR, pagehistory
	}
	db.Model(PageHistory{}).Where("page_id =?", pageId).Order("created_at desc").Limit(20).Scan(&pagehistory)

	if &pagehistory != nil {
		for _, history := range pagehistory {
			PageContent, _ := utils.GUnzipData([]byte(history.PageContent))
			if PageContent != nil {
				history.PageContent = html.UnescapeString(string(PageContent))
			}
		}
		return errmsg.SUCCESE, pagehistory
	}
	return errmsg.SUCCESE, pagehistory
}

func UpdateHistoryComments(uid uint, pageId int, pageComments string, pageHistoryId int) int {
	var page Page
	db.Model(Page{}).Where("id =?", pageId).Find(&page)
	if CheckItemPermn(uid, int(page.ItemId)) != errmsg.SUCCESE {
		return errmsg.ERROR
	}
	err := db.Model(PageHistory{}).Where("id = ?", pageHistoryId).Update("pagecomment", pageComments).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESE
}
func Diff(pageId int, pagehistoryId int, uid uint) (bool, DiffPage) {
	var pageData DiffPage
	var page Page
	var pagehistory PageHistory
	if pageId == 0 {
		return false, pageData
	}
	db.Model(Page{}).Where("id = ?", pageId).Find(&page)
	if &page == nil {
		time.Sleep(1)
		return false, pageData
	}
	if checkItemVisit(uid, int(page.ItemId), "") != errmsg.SUCCESE {
		return true, pageData
	}
	db.Model(PageHistory{}).Where("id = ?", pagehistoryId).Find(&pagehistory)
	pageDecode,_ := base64.StdEncoding.DecodeString(pagehistory.PageContent)
	pageContent, _ := utils.GUnzipData(pageDecode)
	pagehistory.PageContent = string(pageContent)
	pageData.Page= page
	pageData.PageHistory = pagehistory

	return true, pageData
}

func Sort(pages string, itemId int, uid uint) int {
	var htmlpages []byte
	if CheckItemPermn(uid, itemId) != errmsg.SUCCESE {
		return errmsg.ERROR
	}
	var result map[string]int
	htmlpages = []byte(pages)
	err = json.Unmarshal(htmlpages, &result)
	if err != nil {
		return errmsg.ERROR
	}

	for s, p := range result {
		db.Model(Page{}).Where("id = ?", s).Where("item_id = ?", itemId).Update("s_number", p)
	}
	return errmsg.SUCCESE
}

func SortbyPage (pages string , itemId int)int {
	var htmlpages []byte
	var result map[string]int
	htmlpages = []byte(pages)
	err = json.Unmarshal(htmlpages, &result)
	if err != nil {
		return errmsg.ERROR
	}

	for s, p := range result {
		db.Model(Page{}).Where("id = ?", s).Where("item_id = ?", itemId).Update("groupId", p)
	}
	return errmsg.SUCCESE
}


