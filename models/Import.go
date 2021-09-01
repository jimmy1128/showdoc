package models

import (
	"archive/zip"
	"awesomeProject3/utils/errmsg"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
)

type ImportItemInfo struct {
	Types      string   `json:"item_type"`
	Title         string `json:"item_name"`
	Description string `json:"item_description"`
	Password    string `json:"password"`
	UserId      int `json:"user_id"`
	Username string `json:"username"`
	ItemDomain string `json:"item_domain"`
	Pages ImportMenu `json:"pages"`
}

type ImportPageInfo struct {
	PageTitle string `json:"page_title"`
	Cat_id        string                  `json:"cat_id"`
	Page_content  string                `json:"page_content"`
	S_number      string                  `json:"s_number"`
	Page_comments string                `json:"page_comments"`
}

type ImportCatalogInfo struct{
	Id      string       `json:"cat_id"`
	Name        string     `json:"cat_name"`
	ItemId      string       `json:"item_id"`
	ParentCatId string       `json:"parent_cat_id"`
	Level       string       `json:"level"`
	SNumber     string       `json:"s_number"`
	Pages []*ImportPageInfo `json:"pages"`
	Catalogs []ImportCatalogInfo `json:"catalogs"`

}
type ImportMenu struct {
	Pages []*ImportPageInfo `json:"pages"`
	Catalogs []ImportCatalogInfo `json:"catalogs"`
}
type ImportCatalogs struct {
	gorm.Model
	Lang        Lang       `gorm:"foreignKey:cid"`
	Name        string     `gorm:"type:varchar(255);not null" json:"catname"`
	ItemId      uint       `gorm:"type:int;not null" json:"itemid"`
	SNumber     uint       `gorm:"type:int;not null" json:"snumber"`
	ParentCatId uint       `gorm:"type:int;not null" json:"parentcatid"`
	Level       uint       `gorm:"type:int;not null" json:"level"`
	Cid         int        `gorm:"type:int;not null" json:"cid"`
	Pages        []Page    `json:"pages"`
	Catalogs    Catalogs `json:"catalogs"`
}

func ImportItem (jsonData []byte,uid uint , itemName string , itemDescription string , itemPassowrd string , itemDomain string)int{
	var item Item

	var user User
	var importInfo ImportItemInfo
	db.Model(User{}).Where("id = ?",uid).Find(&user)
	err = json.Unmarshal(jsonData,&importInfo )

	item.Title = importInfo.Title
	item.Description = importInfo.Description
	item.Types,_ = strconv.Atoi(importInfo.Types)
	item.Password=importInfo.Password
	item.UserId = user.ID
	item.ItemDomain = itemDomain
	err = db.Model(Item{}).Create(&item).Error

	if err != nil {
		return errmsg.ERROR
	}
	if &importInfo.Pages != nil {
		for _, info := range importInfo.Pages.Pages {
			var page Page
			page.PageContent = info.Page_content
			page.PageComments = info.Page_comments
			page.PageTitle = info.PageTitle
			i, _ := strconv.ParseUint(info.S_number, 10, 64)
			page.SNumber = uint(i)
			page.AuthorUid = user.ID
			page.ItemId = item.ID
			page.CatId = 0
			err = db.Model(Page{}).Create(&page).Error
			if err != nil {
				return errmsg.ERROR
			}
		}
	}
	if importInfo.Pages.Catalogs != nil {
		_insertCat(item.ID, importInfo.Pages.Catalogs ,user, 0 , 2)
	}
return errmsg.SUCCESE

}

func _insertCat (itemId uint , catalogs []ImportCatalogInfo , user User , parentCatId uint , level uint)uint{

	if &catalogs == nil {
		return errmsg.ERROR
	}
	for _, catalog := range catalogs {
		var catalogsData ImportCatalogs
		catalogsData.Name= catalog.Name
		catalogsData.Level = level
		i, _ := strconv.ParseUint(catalog.SNumber, 10, 64)
		catalogsData.SNumber = uint(i)
		catalogsData.ItemId = itemId
		catalogsData.ParentCatId = parentCatId

		db.Table("catalogs").Create(&catalogsData)
		for _, info := range catalog.Pages {
			var pages Page
			pages.AuthorUid = user.ID
			pages.PageTitle = info.PageTitle
			pages.PageComments = info.Page_comments
			pages.PageContent = info.Page_content
			i, _ := strconv.ParseUint(info.S_number, 10, 64)
			pages.SNumber = uint(i)
			pages.ItemId = itemId
			pages.CatId = catalogsData.ID
			db.Model(Page{}).Create(&pages)
		}
		if catalog.Catalogs != nil{
			_insertCat(itemId , catalog.Catalogs , user,catalogsData.ID,level+1)
		}
	}
	return errmsg.SUCCESE
}

func Auto(filename string,uid uint)int{
	var result map[string]interface{}
	var filenameWithSuffix string
	var fileSuffix string
	filenameWithSuffix = path.Base(filename)
	fileSuffix = path.Ext(filenameWithSuffix)
	if fileSuffix == ".zip" {
		err,filedir := DeCompress(filename,"./upload/")

		if err != nil {
			return errmsg.ERROR
		}
		data, err := ioutil.ReadFile(filedir+"/prefix_info.json")
		if err != nil {
			fmt.Println(err)
			return errmsg.ERROR
		}
		if data != nil {
			err =json.Unmarshal(data , &result)
			if err !=nil {
				return errmsg.ERROR
			}
			if result != nil {
					markdown(data,uid)
			}

		}
	}
	if fileSuffix == ".json"{
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			return errmsg.ERROR
		}
		if data != nil {
			err =json.Unmarshal(data , &result)
		}
		for s, _ := range result {
			if s == "swagger" || s == "openapi" {
				return errmsg.FUNCTION_UNDER_DEVELOP
			}
			if s == "id"{
				return errmsg.FUNCTION_UNDER_DEVELOP
			}
			if s =="info"{
				return errmsg.FUNCTION_UNDER_DEVELOP
			}
		}
		return errmsg.ERROR
	}
return errmsg.SUCCESE
}
func DeCompress(zipFile, dest string) (error,string){
	var filedir string
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err,""
	}
	defer reader.Close()
	for _, file := range reader.File {
		rc, err := file.Open()
		if err != nil {
			return err,""
		}
		defer rc.Close()
		filename := dest + file.Name
		filedir = getDir(filename)
		err = os.MkdirAll(getDir(filename), 0755)
		if err != nil {
			return err,""
		}
		w, err := os.Create(filename)
		if err != nil {
			return err,""
		}
		defer w.Close()
		_, err = io.Copy(w, rc)
		if err != nil {
			return err,""
		}
		w.Close()
		rc.Close()
	}

	return nil,filedir
}
func getDir(path string) string {
	return subString(path, 0, strings.LastIndex(path, "/"))
}

func subString(str string, start, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < start || end > length {
		panic("end is wrong")
	}

	return string(rs[start:end])
}

func markdown (info []byte, uid uint)int {
	if &info != nil {
		ImportItem(info,uid,"","","","")
		return errmsg.SUCCESE
	}

return errmsg.ERROR
}
