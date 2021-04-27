package models

import (
	"awesomeProject3/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Catalogs struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);not null" json:"catname"`
	ItemId uint `gorm:"type:int;not null" json:"itemid"`
	SNumber uint `gorm:"type:int;not null" json:"snumber"`
	ParentCatId uint `gorm:"type:int;not null" json:"parentcatid"`
	Level uint `gorm:"type:int;not null" json:"level"`
	Page []*Page `json:"pages"`
	Catalogs []Catalogs `json:"catalogs"`

}

func GetOneCatalogs(id int)(Catalogs,int){
	var catalogs Catalogs
	err = db.Where("id =?",id).First(&catalogs).Error
	if err != nil{
		return catalogs,errmsg.ERROR
	}
	return catalogs,errmsg.SUCCESE
}

func CreateCatalogs(data *Catalogs)int{

	err := db.Create(&data).Error
	if err !=nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESE
}

func EditCatalogs (id int,v int) int{
	var catalogs Catalogs
	var maps = make(map[string]interface{})
	maps["item_id"]=catalogs.ItemId
	maps["name"] = catalogs.Name
	maps["s_number"] = catalogs.SNumber
	maps["parent_cat_id"] = catalogs.ParentCatId
	maps["level"] = catalogs.Level
	err = db.Model(&catalogs).Where("id =?",id).Update(maps).Error
	if v <= 1 {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESE
}

func DeleteCatalogs(itemid int , id int,v int)int{
	var catalogs Catalogs
	err = db.Where("item_id=?",itemid).Where("id =?",id).Delete(&catalogs).Error
	if v <= 1 {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESE
}

func GetCatalogsByItemId(id uint,keyword string)([]Catalogs,int){
	var catalogs []Catalogs
	var catalogsData [] Catalogs
	err = db.Where("item_id=?",id).Where("level =?",2).Find(&catalogs).Error
	if err != nil {
		return catalogs,errmsg.ERROR
	}
	for _,value := range catalogs{
		//value .Catalogs,_ =
		value.Page = GetPagesByCatId(value.ID,keyword,value.ItemId)
		catalogsData = append(catalogsData, value)
	}
	return catalogsData,errmsg.SUCCESE
}

func GetChildCatalogsByCatId(id int)([]*Catalogs,int){
	var catalogs []*Catalogs
	err = db.Where("parent_cat_id =?",id).Find(&catalogs).Error
	if err != nil {
		return catalogs,errmsg.ERROR
	}
	for _,value := range catalogs{
		value.Page = GetPagesByCatId(value.ID,"",value.ItemId)
	}
	return catalogs,errmsg.SUCCESE
}

func GetSecondCatalogsItemId(id int,level int,v int) ([]*Catalogs,int){
	var catalogs []*Catalogs
	err = db.Where("item_id=?",id).Where("level=?",level).Find(&catalogs).Error
	if v == 0 {
		return nil ,errmsg.ERROR_USER_NO_RIGHT
	}
	if err != nil{
		return catalogs,errmsg.ERROR
	}
	for _,value := range catalogs{
		value.Page = GetPagesByCatId(value.ID,"",value.ItemId)
	}
	return  catalogs,errmsg.SUCCESE
}

func (data *Catalogs)Save ()(*Catalogs,int){
	if data.ID == 0 {
		err := db.Create(&data).Error
		if err !=nil{
			return nil,errmsg.ERROR
		}
		return data,errmsg.SUCCESE
	}else {
		err := db.Model(Catalogs{}).Where("id =?",data.ID).Update(data).Error
		if err != nil {
			return nil, errmsg.ERROR
		}
	}
	return data,errmsg.SUCCESE
}

func Getdefaultcat(pageId int,copyPageId int,uid uint) (int,int){
	var tmp Template
	var copyPage Page
	var page Page
	var last_page Page
	var default_cat_id int
	var tmpId int



	if pageId > 0 {
		if tmpId != 0{
			db.Model(Template{}).Where("id =?",tmpId).Find(&tmp)
		}else {
			db.Model(Page{}).Where("id =?",pageId).Find(&tmp)
	}
		default_cat_id = int(tmp.CatId)
	}else if copyPageId != 0{
		//如果是复制接口
		db.Model(Page{}).Where("id =?",copyPageId).Find(&copyPage)
		page.ItemId = copyPage.ItemId
		default_cat_id = int(copyPage.CatId)
	}else {
		//查找用户上一次设置的目录
		db.Model(Page{}).Where("author_uid =?", uid).Where("item_id =?", page.ItemId).Order("created_at desc").Limit(1).Find(&last_page)
		default_cat_id = int(last_page.CatId)
	}
	return default_cat_id,errmsg.SUCCESE

}

func GetPagesbycat(catId int,itemId int,uid uint)([]Page,int){
	var page []Page
	if CheckItemPermn(uid , itemId) != errmsg.SUCCESE{
		return page,errmsg.ERROR_USER_NO_RIGHT
	}
	db.Model(Page{}).Where("cat_id = ?",catId).Where("item_id =?",itemId).Where("is_del is NULL").Select("id ,page_title ,s_number").Order("s_number asc , id asc").Find(&page)
return page ,errmsg.SUCCESE

}

func fileMemberCat(uid uint, catData []*Catalogs)[]*Catalogs{
	var itemMember ItemMember
	var teamItemMember TeamItemMember
	var itemId int
	var catId uint
	if catData != nil{
		return catData
	}
	for _, datum := range catData {
		itemId = int(datum.ItemId)
	}

	// 首先看是否被添加为项目成员
	db.Model(ItemMember{}).Where("uid =?").Where("item_id = ?",itemId).Find(&itemMember)
	if itemMember.Cat_id > 0{
		catId = uint(itemMember.Cat_id)
	}
	//再看是否添加为团队-项目成员
	db.Model(TeamItemMember{}).Where("member_uid = ?",uid).Where("item_id =?",itemId).Find(&teamItemMember)
	if teamItemMember.Cat_id > 0 {
		catId = uint(teamItemMember.Cat_id)
	}

	if catId > 0 {
		for i, datum := range catData {
			if datum.ID != catId{
				index := 1
				array := catData[:i]
				catData = append(array[:index])
			}

		}
		return catData
	}
	return catData
}
