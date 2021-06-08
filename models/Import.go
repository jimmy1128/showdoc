package models

import (
	"archive/zip"
	"awesomeProject3/utils/errmsg"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
)

type ImportItemInfo struct {
	ItemType      string   `json:"item_type"`
	Title         string `json:"item_name"`
	Description string `json:"item_description"`
	Password    string `json:"paswsword"`
	UserId      int `json:"user_id"`
	Username string `json:"username"`
	ItemDomain string `json:"item_domain"`
	Pages ImportPageInfo `json:"pages"`
	Catalog []ImportCatalogInfo `json:"catalogs"`
}
type ImportPageInfo struct {
	PageTitle    string                `gorm:"type:varchar(255)" json:"pagetitle"`
	CatId        uint                  `gorm:"type:int" json:"cat_id"`
	PageContent  string                `gorm:"type:longtext" json:"page_content"`
	SNumber      int                  `gorm:"type:int" json:"s_number"`
	PageComments string                `gorm:"type:text" json:"page_comments"`
}
type ImportCatalogInfo struct{
	CatId      uint       `gorm:"type:int" json:"cat_id"`
	Name        string     `gorm:"type:varchar(255);not null" json:"cat_name"`
	ItemId      uint       `gorm:"type:int" json:"item_id"`
	ParentCatId uint       `gorm:"type:int" json:"parent_cat_id"`
	Level       uint       `gorm:"type:int" json:"level"`
	SNumber     uint       `gorm:"type:int" json:"s_number"`
	Pages []*ImportPageInfo `json:"pages"`
	Catalogs []ImportCatalogInfo `json:"catalogs"`

}

func ImportItem (jsonData []byte,uid uint , itemName string , itemDescription string , itemPassowrd string , itemDomain string)int{
	var item Item
	//var page Page
	var user User
	var result map[string]interface{}
	//result := make(map[string]interface{})
	//var importInfo ImportItemInfo
	db.Model(User{}).Where("id = ?",uid).Find(&user)
	err = json.Unmarshal(jsonData,&result )

	item.Title = result["item_name"].(string)
	item.Description = result["item_description"].(string)
	item.Types,_ = strconv.Atoi(result["item_type"].(string))
	item.Password=result["password"].(string)
	item.UserId = 1
	item.ItemDomain = itemDomain
	db.Model(Item{}).Create(&item)

	if err != nil {
		fmt.Println(err)
		return errmsg.ERROR
	}
	//db.Model(Item{}).Create(result)
	for s, i := range result {
		fmt.Println(s , i)
	}
	if result["pages"] != nil {
		//page.PageContent = result
		//page.PageContent = result["page_content"].(string)
		//page.PageTitle = result["page_title"].(string)
		//i, _ := strconv.ParseUint(result["s_number"].(string), 10, 64)
		//page.SNumber = uint(i)
		//db.Model(Page{}).Create(&page)

	}
		//page.AuthorUid = uid
		//page.PageTitle = result["page_title"]
		//page.PageContent = result["page_content"]
		//page.PageComments = result["page_comments"]
	    //i, _ := strconv.ParseUint(result["s_number"], 10, 64)
		//page.SNumber = uint(i)
		//page.ItemId = item.ID
		//page.CatId = 0



	//for _, info := range importInfo {
	//	//if &info.Catalog != nil {
	//	//	_insertCat(item.ID, info.Catalog ,user , 0 , 2)
	//	//}
	//
	//}
return errmsg.SUCCESE

}

func _insertCat (itemId uint , catalogs []ImportCatalogInfo , user User , parentCatId uint , level uint)uint{
	var catalogsData ImportCatalogInfo
	var pages Page
	if &catalogs == nil {
		return 0
	}
	for _, catalog := range catalogs {
		catalogsData.Name= catalog.Name
		catalogsData.Level = level
		catalogsData.SNumber = catalog.SNumber
		catalogsData.ItemId = itemId
		catalogsData.ParentCatId = parentCatId
		db.Model(Catalogs{}).Create(&catalogsData)
		for _, info := range catalogsData.Pages {
			pages.AuthorUid = user.ID
			pages.PageTitle = info.PageTitle
			pages.PageComments = info.PageTitle
			pages.SNumber = 1
			pages.ItemId = itemId
			pages.CatId = catalogsData.CatId
		}
		db.Model(Page{}).Create(&pages)
	}
	if &catalogsData.Catalogs != nil{
		_insertCat(itemId , catalogsData.Catalogs , user,catalogsData.CatId,level)
	}
	return catalogsData.CatId
}

func Auto(filename string,uid uint)int{
	var result map[string]interface{}
	//var result2 map[string]string
	var filenameWithSuffix string
	var fileSuffix string
	filenameWithSuffix = path.Base(filename)
	fileSuffix = path.Ext(filenameWithSuffix)
	if fileSuffix == ".zip" {
		err := DeCompress(filename,"./upload/")
		if err != nil {
			return errmsg.ERROR
		}
		data, err := ioutil.ReadFile("./upload/prefix_info.json")
		if err != nil {
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
return errmsg.SUCCESE
}
func DeCompress(zipFile, dest string) error {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer reader.Close()
	for _, file := range reader.File {
		rc, err := file.Open()
		if err != nil {
			return err
		}
		defer rc.Close()
		filename := dest + file.Name
		err = os.MkdirAll(getDir(filename), 0755)
		if err != nil {
			return err
		}
		w, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer w.Close()
		_, err = io.Copy(w, rc)
		if err != nil {
			return err
		}
		w.Close()
		rc.Close()
	}
	return nil
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
	if &info == nil {

	}
	if &info != nil {
		ImportItem(info,uid,"","","","")
		return errmsg.SUCCESE
	}

return errmsg.ERROR
}