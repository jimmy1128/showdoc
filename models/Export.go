package models

import (
	"archive/zip"
	"awesomeProject3/utils/errmsg"
	"container/list"
	"encoding/json"
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"io"
	"os/exec"
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
	//fname1 := filepath.Join(dname , "prefix_readme.md")
	err = ioutil.WriteFile(fname,exportJson,0666)
	//err = ioutil.WriteFile(fname1, []byte("由于页面标题可能含有特殊字符导致异常，所以markdown文件的命名均为英文（md5串），以下是页面标题和文件的对应关系：\n"),0666)
		_markdownTofile(iteminfo.Pages,dname,exportitem)
	Zip(dname, "showdoc.zip")

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
			//dirarr = append(dirarr , fname1)
			//new_page_content := strings.Replace(page.Page_content, "&quot;", " ", -1)

			ioutil.WriteFile(fname1, []byte("## "+page.PageTitle + page.Page_content),0666)
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
	var mainCatalogs1 []ImportCatalogInfo
	var mainCatalogs2 []ImportCatalogInfo
	catalogData.Pages = _getExportPageByCatId(catalogData.Id, page)
	sub_catalogs = _getExportCatByCatId(catalogData.Id, catalogs)
	//fmt.Println(sub_catalogs)
	if sub_catalogs != nil {
		for _, catalog := range sub_catalogs {
			//fmt.Println("2",catalog)
			//catalogData.Catalogs = sub_catalogs
			//catalog.Catalogs = getExportCat(catalog, page, catalogs)
			mainCatalogs1 = getExportCat(catalog, page, catalogs)
			//fmt.Println("3",catalogData.Catalogs )
			//catalogs2 = append(catalogs2, catalog)
			for _, info := range mainCatalogs1 {
				mainCatalogs2 = append(mainCatalogs2, info)
			}

			//fmt.Println("4",catalogsSubData)
		}
		catalogData.Catalogs = mainCatalogs2
		catalogsSubData = append(mainCatalogs, catalogData)
	}
	if catalogData.Catalogs == nil {
		catalogsSubData = append(catalogsSubData, catalogData)
		//fmt.Println("5",catalogsSubData)
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

func ExportWord (itemId uint , catId uint , pageId uint , uid uint)int {
	var item Item
	var page Page
	var menu Menu
	var pages []*Page
	var catalogs []Catalogs
	if CheckItemEdit(uid , itemId) != true {
		return errmsg.ERROR_USER_NO_RIGHT
	}

	db.Model(Item{}).Where("id =? ",itemId).Find(&item)
	menu = getContent(itemId,0)
	if pageId >0 {
		db.Model(Page{}).Where("id =?",pageId).Find(&page)

	}else if catId != 0  {
		for _, catalog := range menu.Catalogs {
			if catId == catalog.ID{
				pages= catalog.Page
				catalogs = catalog.Catalogs
			}else {
				if catalog.Catalogs != nil {
					for _, c := range catalog.Catalogs {
						if catId == c.ID{
							pages = c.Page
							catalogs = c.Catalogs
						}
					}

				}

			}
		}
	}else {
		pages = menu.Page
		catalogs = menu.Catalogs
	}
	var data = ""
	//var data1 interface{}
	//var data1 map[int]interface{}
	var parent = 1
	var count int
	if pages != nil {
		listHaiCoder := list.New()

			for _, p := range pages {
				if pageCount(pages) > 0 {
					parentStr := strconv.Itoa(parent)
					listHaiCoder.PushFront( "<h1>" + parentStr + "," + p.PageTitle + "</h1>")
				} else {
					listHaiCoder.PushFront( "<h1>" + p.PageTitle + "</h1>")
				}
				listHaiCoder.PushFront( `<div style=\"margin-left:20px;\">`)
				//tmp_content = runApiToMd(p.PageContent)
				unsafe := blackfriday.MarkdownCommon([]byte(p.PageContent))
				html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
				listHaiCoder.PushFront(string(html))
				//listHaiCoder.PushFront("</div>")
				parent++

			}
		//dname, err := ioutil.TempDir("./", "tmpfile")
		for i := listHaiCoder.Back(); i != nil; i = i.Prev() {

			bytevalue := []byte(fmt.Sprintf("%v",i.Value))
			err = ioutil.WriteFile("./dat2.doc",bytevalue,0666)
			if err != nil {
				return errmsg.ERROR
			}

		}

	}
	if catalogs != nil {
		for _, catalog := range catalogs {
			parentStr := strconv.Itoa(parent)
			data = "<h1>" + parentStr +"," + catalog.Name + "</h1>"
			data = `<div style=\"margin-left:0px;\">`
			child := 1
			if catalog.Page != nil {
				childStr := strconv.Itoa(child)
				for _, p := range catalog.Page {
					data = "<h2>" + parentStr+"." +childStr +"," + p.PageTitle+"</h2>"
					data = `<div style=\"margin-left:0px;\">`
					unsafe := blackfriday.MarkdownCommon([]byte(p.PageContent))
					html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
					data= string(html)
					data = "</div>"
					child ++
				}
			}
			if catalog.Catalogs != nil {
				parent2 := 1
				for _, c := range catalog.Catalogs {
					parent2Str := strconv.Itoa(parent2)
					data = "<h2>" + parentStr+"."+ parent2Str+ c.Name +"</h2>"
					data = `<div style="margin-left:0px;">`
					child2 := 1
					if c.Page != nil {
						for _, p := range c.Page {
							child2Str := strconv.Itoa(child2)
							data = "<h3>"+parentStr+"." + parent2Str +"." +child2Str+","+p.PageTitle+"</h3>"
							data = `<div style="margin-left:0px;">`
							unsafe := blackfriday.MarkdownCommon([]byte(p.PageContent))
							html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
							data= string(html)
							data = "</div>"
							child2 ++
						}
					}
					if c.Catalogs !=nil{
						parent3 := 1
						parent3Str := strconv.Itoa(parent3)
						for _, c2 := range c.Catalogs {
							data = "<h2>" + parent2Str +"." + parent2Str + "." + parent3Str+"、"+c2.Name+"</h2>"
							data = `<div style="margin-left:0px;">`
							child3 := 1
							if c2.Page !=nil {
									child3Str := strconv.Itoa(child3)
									for _, p := range c2.Page {
										data = "<h3>" + parentStr + "." + parent2Str + "."+parent3Str+"."+child3Str+","+p.PageContent+"</h3>"
										data = `<div style="margin-left:0px;">`
										unsafe := blackfriday.MarkdownCommon([]byte(p.PageContent))
										html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
										data= string(html)
										data = "</div>"
										child3 ++
									}
								}
								data= "</div>"
								parent3++
						}
					}
					data = "</div>"
					parent2++
				}
			}
			data = "</div>"
			parent++
		}
	}
	fmt.Println(data, parent,count)

return 200
}
func pageCount(pages []*Page) int{
var count int
	for i := range pages {
		count = i
	}
	return count

}
func runApiToMd ( content string){
	if content != "" {


	}
}
var (
	ebookConvert = "mdout"
)
func ExportPdf (itemId uint , uid uint)int {
	var item Item
	var iteminfo ImportItemInfo
	var exportitem ImportCatalogInfo
	if CheckItemEdit(uid, itemId) == false {
		return errmsg.ERROR
	}
	db.Model(Item{}).Where("id =?",itemId).Find(&item)

	exportJson,_ := export(item.ID,true)
	json.Unmarshal(exportJson,&iteminfo)
	dname1, err := ioutil.TempDir("./", "showdoc_")
	if err != nil {
		return  errmsg.ERROR
	}
	defer os.RemoveAll(dname1)

	//fname := filepath.Join(dname1 , "prefix_info.json")
	//fname1 := filepath.Join(dname , "prefix_readme.md")
	//err = ioutil.WriteFile(fname,exportJson,0666)
	//err = ioutil.WriteFile(fname1, []byte("由于页面标题可能含有特殊字符导致异常，所以markdown文件的命名均为英文（md5串），以下是页面标题和文件的对应关系：\n"),0666)
	_markdownToPdf(iteminfo.Pages,dname1,exportitem)
	Zip(dname1, "showdoc.zip")

	return errmsg.SUCCESE
}
func _markdownToPdf (catalogData ImportMenu , tempDir string,catalogs ImportCatalogInfo)ImportCatalogInfo {
	var catalogData2 ImportMenu
	var mainCatalogs ImportCatalogInfo
	var dirarr []string
	var dirarrs []string
	if catalogData.Pages !=nil {
		for _, page := range catalogData.Pages {
			//t := time.Now()
			filename := "prefix_"+page.PageTitle +".md"

			fname1 := filepath.Join(tempDir , filename)
			fmt.Println(fname1)
			dirarrs = append(dirarrs , fname1)
			new_page_content := strings.Replace(page.Page_content, "&quot;", " ", -1)
			ioutil.WriteFile(fname1, []byte("## "+page.PageTitle + new_page_content),0666)
			fl, _ := os.OpenFile(tempDir+"/prefix_readme.md", os.O_APPEND|os.O_WRONLY, 0666)
			fl.Write([]byte("\n"+page.PageTitle+"-prefix_"+filename))

		}
	}
	for _, s := range dirarrs {
		cmd := exec.Command(ebookConvert, s)
		//if this.Debug {
		//	fmt.Println(cmd.Args)
		//}
		defer os.Remove(s)
		cmd.Run()
	}
	if catalogData.Catalogs !=nil {
		for i, catalog := range catalogData.Catalogs {
			catalogData.Catalogs[i] = _markdownToPdf(catalogData2,tempDir,catalog)
		}
	}
	if catalogs.Catalogs !=nil {
		for i, catalog := range catalogs.Catalogs {
			catalogs.Catalogs[i] = _markdownToPdf(catalogData2,tempDir,catalog)
		}
	}
	if catalogs.Pages !=nil {
		for _, page := range catalogs.Pages {
			//t := time.Now()
			filename := "prefix_"+page.PageTitle +".md"
			fname1 := filepath.Join(tempDir , filename)
			dirarr = append(dirarr , fname1)
			new_page_content := strings.Replace(page.Page_content, "&quot;", " ", -1)
			ioutil.WriteFile(fname1, []byte("## "+page.PageTitle+"\n" + new_page_content),0666)
			fl, _ := os.OpenFile(tempDir+"/prefix_readme.md", os.O_APPEND|os.O_WRONLY, 0666)
			fl.Write([]byte("\n"+page.PageTitle+"-prefix_"+filename))
		}
	}

	for _, s := range dirarr {
		//fmt.Println(s)
		cmd := exec.Command(ebookConvert, s)
		//if this.Debug {
		//	fmt.Println(cmd.Args)
		//}
		cmd.Run()
		defer os.Remove(s)
	}
	return mainCatalogs
}
