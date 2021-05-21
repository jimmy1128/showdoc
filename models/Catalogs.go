package models

import (
	"awesomeProject3/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Catalogs struct {
	gorm.Model
	Lang        Lang       `gorm:"foreignKey:cid"`
	Name        string     `gorm:"type:varchar(255);not null" json:"catname"`
	ItemId      uint       `gorm:"type:int;not null" json:"itemid"`
	SNumber     uint       `gorm:"type:int;not null" json:"snumber"`
	ParentCatId uint       `gorm:"type:int;not null" json:"parentcatid"`
	Level       uint       `gorm:"type:int;not null" json:"level"`
	Cid         int        `gorm:"type:int;not null" json:"cid"`
	Page        []*Page    `json:"pages"`
	Catalogs    []Catalogs `json:"catalogs"`
}
type CatalogsTitle struct {
	ID          uint             `gorm:"type:int;not null" json:"cat_id"`
	Name        string           `gorm:"type:varchar(255);not null" json:"cat_name"`
	ItemId      uint             `gorm:"type:int;not null" json:"item_id"`
	Level       uint             `gorm:"type:int;not null" json:"level"`
	ParentCatId uint             `gorm:"type:int;not null" json:"parent_cat_id"`
	SNumber     uint             `gorm:"type:int;not null" json:"s_number"`
	Cid         int              `gorm:"type:int;not null" json:"cid"`
	Sub         []*CatalogsTitle `json:"sub"`
}

func GetOneCatalogs(id int) (Catalogs, int) {
	var catalogs Catalogs
	err = db.Where("id =?", id).First(&catalogs).Error
	if err != nil {
		return catalogs, errmsg.ERROR
	}
	return catalogs, errmsg.SUCCESE
}

func CreateCatalogs(data *Catalogs) int {
	if data.Name == "" {
		return errmsg.ERROR_CATE_NOT_EXIST
	}

	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESE
}

func EditCatalogs(id int, v int) int {
	var catalogs Catalogs
	if catalogs.Name == "" {
		return errmsg.ERROR_CATE_NOT_EXIST
	}
	var maps = make(map[string]interface{})
	maps["item_id"] = catalogs.ItemId
	maps["name"] = catalogs.Name
	maps["s_number"] = catalogs.SNumber
	maps["parent_cat_id"] = catalogs.ParentCatId
	maps["level"] = catalogs.Level
	err = db.Model(&catalogs).Where("id =?", id).Update(maps).Error
	if v <= 1 {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESE
}

func DeleteCatalogs(itemid int, id int, v int) int {
	var catalogs Catalogs
	err := db.Where("item_id=?", itemid).Where("id =?", id).Delete(&catalogs).Error
	if v <= 1 {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESE
}

func GetCatalogsByItemId(id uint, keyword string, langId int) ([]Catalogs, int) {
	var catalogs []Catalogs
	var catalogsData []Catalogs
	if langId <= 0 {
		err = db.Where("item_id=?", id).Where("level =?", 2).Preload("Lang").Find(&catalogs).Error
		if err != nil {
			return catalogs, errmsg.ERROR
		}
		for _, value := range catalogs {
			//value .Catalogs,_ =
			value.Page = GetPagesByCatId(value.ID, keyword, value.ItemId)
			catalogsData = append(catalogsData, value)
		}
		return catalogsData, errmsg.SUCCESE
	}
	err = db.Where("item_id=?", id).Where("level =?", 2).Preload("Lang").Where("cid =?", langId).Find(&catalogs).Error
	if err != nil {
		return catalogs, errmsg.ERROR
	}
	for _, value := range catalogs {
		//value .Catalogs,_ =
		value.Page = GetPagesByCatId(value.ID, keyword, value.ItemId)
		catalogsData = append(catalogsData, value)
	}
	return catalogsData, errmsg.SUCCESE
}

func GetChildCatalogsByCatId(id int) ([]*Catalogs, int) {
	var catalogs []*Catalogs
	err = db.Where("parent_cat_id =?", id).Find(&catalogs).Error
	if err != nil {
		return catalogs, errmsg.ERROR
	}
	for _, value := range catalogs {
		value.Page = GetPagesByCatId(value.ID, "", value.ItemId)
	}
	return catalogs, errmsg.SUCCESE
}

func GetSecondCatalogsItemId(id int, level int, v int) ([]*Catalogs, int) {
	var catalogs []*Catalogs
	err = db.Where("item_id=?", id).Where("level=?", level).Find(&catalogs).Error
	if v == 0 {
		return nil, errmsg.ERROR_USER_NO_RIGHT
	}
	if err != nil {
		return catalogs, errmsg.ERROR
	}
	for _, value := range catalogs {
		value.Page = GetPagesByCatId(value.ID, "", value.ItemId)
	}
	return catalogs, errmsg.SUCCESE
}

func (data *Catalogs) Save() (*Catalogs, int) {
	if data.Cid == 0 {
		return data, errmsg.ERROR_LANG_EMPTY
	}
	if data.ID == 0 {
		err := db.Create(&data).Error
		if err != nil {
			return nil, errmsg.ERROR
		}
		return data, errmsg.SUCCESE
	} else {
		err := db.Model(Catalogs{}).Where("id =?", data.ID).Update(data).Error
		if err != nil {
			return nil, errmsg.ERROR
		}
	}
	return data, errmsg.SUCCESE
}

func Getdefaultcat(pageId int, copyPageId int, uid uint, itemId int, pageHistoryId int) (int, int) {
	var copyPage Page
	var page Page
	var lastPage Page
	var defaultCatId int
	var pageHistory PageHistory
	if pageId > 0 {
		if pageHistoryId != 0 {
			db.Model(PageHistory{}).Where("id =? ", pageHistoryId).Find(&pageHistory)
		} else {
			db.Model(Page{}).Where("id =?", pageId).Find(&page)
		}
		defaultCatId = int(page.CatId)
	} else if copyPageId == 0 {
		//如果是复制接口
		db.Model(Page{}).Where("id =?", copyPageId).Find(&copyPage)
		page.ItemId = copyPage.ItemId
		defaultCatId = int(copyPage.CatId)
	} else {
		//查找用户上一次设置的目录
		db.Model(Page{}).Where("author_uid =?", uid).Where("item_id =?", itemId).Order("created_at desc").Limit(1).Find(&lastPage)
		defaultCatId = int(lastPage.CatId)
	}
	return defaultCatId, errmsg.SUCCESE
}

func GetPagesbycat(catId int, itemId int, uid uint) ([]Page, int) {
	var page []Page
	if CheckItemPermn(uid, itemId) != errmsg.SUCCESE {
		return page, errmsg.ERROR_USER_NO_RIGHT
	}
	db.Model(Page{}).Where("cat_id = ?", catId).Where("item_id =?", itemId).Where("is_del is NULL").Select("id ,page_title ,s_number").Order("s_number asc , id asc").Find(&page)
	return page, errmsg.SUCCESE
}

func filterMemberCat(uid uint, catData []*CatalogsTitle) []CatalogsTitle {
	var itemMember ItemMember
	var teamItemMember TeamItemMember
	var catData2 []CatalogsTitle
	var itemId []uint
	var catId uint
	if catData == nil {
		return catData2
	}
	for _, datum := range catData {
		itemId = append(itemId, datum.ItemId)
	}
	// 首先看是否被添加为项目成员
	db.Model(ItemMember{}).Where("uid =?", uid).Where("item_id in (?)", itemId).Find(&itemMember)
	if itemMember.Cat_id > 0 {
		catId = uint(itemMember.Cat_id)
	}
	//再看是否添加为团队-项目成员
	db.Model(TeamItemMember{}).Where("member_uid = ?", uid).Where("item_id in (?)", itemId).Find(&teamItemMember)
	if teamItemMember.Cat_id > 0 {
		catId = uint(teamItemMember.Cat_id)
	}
	if catId > 0 {
		for _, datum := range catData {
			if datum.ID != catId{
			}
		}
		return catData2
	}
	return catData2
}

func CatListName(itemId uint, keyword string, uid uint) ([]*CatalogsTitle, int) {
	var catalogs []*CatalogsTitle
	var catalogsData []CatalogsTitle
	if itemId > 0 {
		catalogs = getList(itemId, true)
		catalogsData = filterMemberCat(uid, catalogs)
		if catalogsData == nil {
			return catalogs, errmsg.SUCCESE
		}
	}

	return catalogs, errmsg.SUCCESE
}

func getList(itemId uint, isGroup bool) []*CatalogsTitle {
	var ret []*CatalogsTitle
	var ret2 []*CatalogsTitle

	if itemId > 0 {
		db.Model(Catalogs{}).Where("item_id =?", itemId).Order("s_number,id asc").Scan(&ret)
	}
	if ret != nil {
		if isGroup == true {
			for _, catalogs := range ret {
				ret2 = _getChlid(catalogs.ID, ret)
				catalogs.Sub  =  ret2
			}
			for _, catalogs := range ret {
				if catalogs.ParentCatId == 0 {

					ret2 = append( ret2 , catalogs)
				}
			}
		}
		return ret2
	} else {
		return ret
	}

}
func _getChlid(catId uint, data []*CatalogsTitle) []*CatalogsTitle {
	var itemData2 []*CatalogsTitle

	amap := make(map[uint]*CatalogsTitle)
	if data != nil && catId > 0 {
		for _, datum := range data {
			if datum.ParentCatId == catId {
				amap[datum.ID] = datum
			}
		}
	}
	for _,v:= range amap{
		itemData2 = append(itemData2, v)
	}
	return itemData2
}
