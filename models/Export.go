package models

import (
	"archive/zip"
	"awesomeProject3/utils/errmsg"
	"encoding/json"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)
type ExportItemInfo struct {
	Item_type      string   `json:"item_type"`
	Item_name         string `json:"item_name"`
	Item_description string `json:"item_description"`
	Password    string `json:"password"`
	UserId      int `json:"user_id"`
	Username string `json:"username"`
	ItemDomain string `json:"item_domain"`
	Pages Menu `json:"pages"`
}

func ExportMarkdown (itemId uint , uid uint)int {
	var item Item
	var iteminfo ImportItemInfo
	var exportitem ImportCatalogInfo
   if CheckItemEdit(uid, itemId) == false {
   	return errmsg.ERROR
   }
   db.Model(Item{}).Where("id =?",itemId).Find(&item)
	exportJson,_ := export(item.ID,true)
	json.Unmarshal(exportJson,&iteminfo)
	dname, err := ioutil.TempDir("./", "showdoc_")
	if err != nil {
		return  errmsg.ERROR
	}
	defer os.RemoveAll(dname)
	fname := filepath.Join(dname , "prefix_info.json")
	fname1 := filepath.Join(dname , "prefix_readme.md")
	err = ioutil.WriteFile(fname,exportJson,0666)
	err = ioutil.WriteFile(fname1, []byte("由于页面标题可能含有特殊字符导致异常，所以markdown文件的命名均为英文（md5串），以下是页面标题和文件的对应关系：\n"),0666)
		_markdownTofile(iteminfo.Pages,dname,exportitem)
	Zip("./"+dname, "showdoc.zip")


return errmsg.SUCCESE
}
func export (itemId uint , uncompress bool)([]byte,int){
	var item ImportItemInfo
	var itemMember ItemMember
	db.Model(Item{}).Select("types,title,description,password").Where("id =?",itemId).Scan(&item)
	item.Pages = getExportContent(itemId ,uncompress)
    db.Model(ItemMember{}).Select("member_group_id,uid,username").Where("item_id = ?",itemId).Scan(&itemMember)
	jsonItem, err := json.Marshal(item)
	if err != nil {
		 return nil,errmsg.ERROR
	}
return jsonItem , errmsg.SUCCESE

}
func _markdownTofile (catalogData ImportMenu , tempDir string,catalogs ImportCatalogInfo)ImportCatalogInfo {
	var catalogData2 ImportMenu
	var mainCatalogs ImportCatalogInfo
	if catalogData.Pages !=nil {
		for _, page := range catalogData.Pages {
			t := time.Now()
			filename := "prefix_"+strconv.FormatInt(t.Unix(), 10)+ strconv.Itoa(rand.Intn(1000)) +".md"
			fname1 := filepath.Join(tempDir , filename)
            ioutil.WriteFile(fname1, []byte(page.Page_content),0666)
			fl, _ := os.OpenFile(tempDir+"/prefix_readme.md", os.O_APPEND|os.O_WRONLY, 0666)
			fl.Write([]byte("\n"+page.PageTitle+"-prefix_"+filename))
		}
	}
	if catalogData.Catalogs !=nil {
		for i, catalog := range catalogData.Catalogs {
			catalogData.Catalogs[i] = _markdownTofile(catalogData2,tempDir,catalog)
		}
	}
	if catalogs.Catalogs !=nil {
		for i, catalog := range catalogs.Catalogs {
			catalogs.Catalogs[i] = _markdownTofile(catalogData2,tempDir,catalog)
		}
	}
	if catalogs.Pages !=nil {
		for _, page := range catalogs.Pages {
			t := time.Now()
			filename := "prefix_"+strconv.FormatInt(t.Unix(), 10)+ strconv.Itoa(rand.Intn(1000)) +".md"
			fname1 := filepath.Join(tempDir , filename)
			ioutil.WriteFile(fname1, []byte(page.Page_content),0666)
			fl, _ := os.OpenFile(tempDir+"/prefix_readme.md", os.O_APPEND|os.O_WRONLY, 0666)
			fl.Write([]byte("\n"+page.PageTitle+"-prefix_"+filename))
		}
	}

return mainCatalogs
}
func getExportContent(itemId uint ,uncompress bool)ImportMenu{
	var catalogsData []ImportCatalogInfo
	var pagesData []*ImportPageInfo
	var catalogsData2 []ImportCatalogInfo
	var result []ImportCatalogInfo
	db.Model(Page{}).Where("item_id =?", itemId).Order("s_number asc , id asc").Scan(&pagesData)
	db.Model(Catalogs{}).Where("item_id = ?", itemId).Order("s_number asc , id asc").Scan(&catalogsData)
	if &catalogsData != nil {
		for _, catalog := range catalogsData {
			if catalog.Level == "2" {
				catalogsData2 = append(catalogsData2, catalog)
			}
		}
	}
	if &catalogsData2 != nil {
		for _, c := range catalogsData2 {
			catalogsData2 = getExportCat(c, pagesData, catalogsData)
			for _, c2 := range catalogsData2 {
				result = append(result, c2)
			}
		}
	}
	var menu ImportMenu
	menu.Pages = _getExportPagesByItemId(itemId)
	menu.Catalogs = result
	return menu
}
func getExportCat(catalogData ImportCatalogInfo, page []*ImportPageInfo, catalogs []ImportCatalogInfo) []ImportCatalogInfo {
	var sub_catalogs []ImportCatalogInfo
	var catalogsSubData []ImportCatalogInfo
	var mainCatalogs []ImportCatalogInfo
	catalogData.Pages = _getExportPageByCatId(catalogData.Id, page)
	sub_catalogs = _getExportCatByCatId(catalogData.Id, catalogs)
	if sub_catalogs != nil {
		for _, catalog := range sub_catalogs {
			catalogData.Catalogs = getExportCat(catalog, page, catalogs)
			catalogsSubData = append(mainCatalogs, catalogData)
		}
	}
	if catalogData.Catalogs == nil {
		catalogsSubData = append(catalogsSubData, catalogData)
	}
	return catalogsSubData
}
func _getExportPageByCatId(catId string, page []*ImportPageInfo) []*ImportPageInfo {
	var page2 []*ImportPageInfo
	if &page != nil {
		for _, p := range page {
			if p.Cat_id == catId {
				page2 = append(page2, p)
			}
		}
	}
	return page2
}

func _getExportCatByCatId(catId string, catalogs []ImportCatalogInfo) []ImportCatalogInfo {
	var catalogs2 []ImportCatalogInfo

	if &catalogs != nil {
		for _, catalog := range catalogs {
			if catalog.ParentCatId == catId {

				catalogs2 = append(catalogs2, catalog)
			}
		}
	}
	return catalogs2
}

func _getExportPagesByItemId(id uint) []*ImportPageInfo {
	var pages []*ImportPageInfo

		err := db.Model(Page{}).Where("item_id =?", id).Where("cat_id =?", 0).Scan(&pages).Error
		if err != nil {
			return pages
		}
		return pages
}
func Zip(src_dir string, zip_file_name string) {

	// 预防：旧文件无法覆盖
	os.RemoveAll(zip_file_name)

	// 创建：zip文件
	zipfile, _ := os.Create(zip_file_name)
	defer zipfile.Close()

	// 打开：zip文件
	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	// 遍历路径信息
	filepath.Walk(src_dir, func(path string, info os.FileInfo, _ error) error {

		// 如果是源路径，提前进行下一个遍历
		if path == src_dir {
			return nil
		}

		// 获取：文件头信息
		header, _ := zip.FileInfoHeader(info)
		header.Name = strings.TrimPrefix(path, src_dir+`\`)

		// 判断：文件是不是文件夹
		if info.IsDir() {
			header.Name += `/`
		} else {
			// 设置：zip的文件压缩算法
			header.Method = zip.Deflate
		}

		// 创建：压缩包头部信息
		writer, _ := archive.CreateHeader(header)
		if !info.IsDir() {
			file, _ := os.Open(path)
			defer file.Close()
			io.Copy(writer, file)
		}
		return nil
	})
}