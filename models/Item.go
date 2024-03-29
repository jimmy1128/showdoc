package models

import (
	"awesomeProject3/utils"
	"awesomeProject3/utils/errmsg"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"sort"
	"strconv"
	"time"
)

type Item struct {
	gorm.Model
	Title       string `gorm:"type:varchar(255);not null" json:"titlename"`
	Description string `gorm:"type:varchar(255);not null" json:"description"`
	Password    string `gorm:"type:varchar(255)" json:"password"`
	UserId      uint   `gorm:"type:int;not null" json:"userid"`
	Types       int    `gorm:"type:int;default:1" json:"types"`
	Creator     int    `gorm:"type:int" json:"creator"`
	Isarchived  int    `gorm:"type:int" json:"isarchived"`
	Lang        int    `gorm:"type:int" json:"lang"`
	Link        int    `gorm:"type:int" json:"link"`
	IsPrivate   int    `gorm:"type:int;default:0" json:"is_private"`
	MemberNum   int    `gorm:"type:int;default:0" json:"MemberNum"`
	Username    string `gorm:"type:varchar(255)" json:"username"`
	ItemDomain string `gorm:"type:varchar(255)" json:"itemdomain"`
	LangList string  `gorm:"type:text;not null" json:"lang_list"`
	IsComment int `gorm:"type:int;default:0" json:"is_comment"`
}
type ItemInfo struct {
	Id            int    `gorm:"primaryKey" json:"id"`
	Title         string `gorm:"type:varchar(255);not null" json:"itemtitle"`
	IsArchived    string `gorm:"type:varchar(255);not null" json:"isarchived"`
	DefaultPageId string `gorm:"type:varchar(10)" json:"defaultpageid"`
	DefaultCatId2 string `gorm:"type:varchar(10)" json:"defaultcatid2"`
	DefaultCatId3 string `gorm:"type:varchar(10)" json:"defaultcatid3"`
	DefaultCatId4 string `gorm:"type:varchar(10)" json:"defaultcatid4"`
	LangList string  `gorm:"type:text;not null" json:"lang_list"`
	ItemType      uint   `gorm:"type:int;not null" json:"itemtype"`
	IsLogin       bool   `gorm:"type:bool" json:"is_login"`
	ItemPermn     bool   `gorm:"type:bool" json:"itempermn"`
	ItemCreator   bool   `gorm:"type:bool" json:"itemcreator"`
	Lang          int    `gorm:"type:int" json:"lang"`
	Link          int    `gorm:"type:int" json:"link"`
	IsComment int `gorm:"type:int" json:"is_comment"`
	Menu          Menu   `json:"menu"`
}
type Menu struct {
	Page     []*Page    `json:"pages"`
	Catalogs []Catalogs `json:"catalogs"`
}
type ItemSort struct {
	Id           int    `gorm:"primaryKey;auto_increment" json:"id"`
	Uid          int    `gorm:"type:int;not null;default:0" json:"uid"`
	ItemSortData string `gorm:"type:text" json:"itemsortdata"`
	CreatedAt    string `gorm:"type:date" json:"created_at"`
}

func (data *Item) GetItemInfo(keyword string, langId int, uid uint , itemId uint,defaultpageid string) ItemInfo {
	var defaultid Page
	var defaultpage Page
	db.Model(Page{}).Where("id = ?",defaultpageid).Find(&defaultid)
		if langId != 0 {
			if langId != defaultid.Cid {
				db.Model(Page{}).Where("cid =?", langId).Where("group_id =?", defaultid.GroupId).Where("item_id = ?", itemId).Find(&defaultpage)
			}
		}

	var page Page
	var cat1 Catalogs
	var cat2 Catalogs
	var itemInfo ItemInfo
	itemInfo.Id = int(data.ID)
	itemInfo.ItemType = uint(data.Types)
	itemInfo.IsLogin = false
	itemInfo.ItemPermn = false
	itemInfo.ItemCreator = false
	if defaultpageid != "0" {
			b := strconv.Itoa(int(defaultpage.ID))
			if b == "0" && uid ==0{
				b = defaultpageid
			}
			if uid > 0 && b=="0"{
				b = "0"
			}
			//内链接无法跳转， 暂时优化
			itemInfo.DefaultPageId = defaultpageid
	}
    itemInfo.IsComment = data.IsComment
	itemInfo.LangList = data.LangList
	itemInfo.Title = data.Title
	itemInfo.Link = data.Link
	itemInfo.Lang = data.Lang
	itemInfo.Menu = itemInfo.GetMenu(keyword, langId, uid, itemId)
	if itemInfo.DefaultPageId != "" {
		db.Model(Page{}).Where("id =? ", itemInfo.DefaultPageId).Find(&page)
		if &page != nil {
			itemInfo.DefaultCatId4 = string(page.ID)
			db.Model(Catalogs{}).Where("id = ?", itemInfo.DefaultCatId4).Where("parent_cat_id > 0").Find(&cat1)
			if &cat1 != nil {
				itemInfo.DefaultCatId3 = string(cat1.ParentCatId)
			} else {
				itemInfo.DefaultCatId2 = itemInfo.DefaultCatId4
				itemInfo.DefaultCatId4 = "0"
			}
			db.Model(Catalogs{}).Where("id =?", itemInfo.DefaultCatId4).Where("parent_cat_id > 0").Find(&cat2)
			if &cat2 != nil {
				itemInfo.DefaultCatId2 = string(cat2.ParentCatId)
			} else {
				itemInfo.DefaultCatId2 = itemInfo.DefaultCatId3
				itemInfo.DefaultCatId3 = "0"
			}
		}
	}
	return itemInfo
}

func (data *ItemInfo) GetMenu(keyword string, langId int,uid uint , itemId uint) Menu {
	var menu Menu
	var menu2 Menu
	if keyword == "" {
		menu = getMenu(uint(data.Id), langId)
		menu2 = FilterMemberItem(uid ,itemId , menu)
		return menu2
	}
	menu.Page = GetPagesByItemId(uint(data.Id), keyword, langId)
	//menu.Catalogs, _ = GetCatalogsByItemId(uint(data.Id), keyword, langId)
	return menu
}

func GetMyItem(uid uint) ([]Item, int) {
	var item []Item
	var teamItemMember []TeamItemMember
	var itemMember []ItemMember
	var member_item_ids []uint
	var item_sort ItemSort
	var user User
	db.Model(User{}).Where("id = ?",uid).Find(&user)
	// 如果是管理员就显示全部项目
	if user.Role == 1 {
		db.Where("deleted_at is NULL").Find(&item)

	}else {

		err = db.Where("user_id =?", uid).Where("deleted_at is NULL").Find(&item).Error
	}
	for _, i2 := range item {
		member_item_ids = append(member_item_ids, i2.ID)
	}
	db.Model(ItemMember{}).Where("uid =?", uid).Find(&itemMember)
	for _, members := range itemMember {
		member_item_ids = append(member_item_ids, uint(members.Item_id))
	}

	db.Model(TeamItemMember{}).Where("member_uid=?", uid).Find(&teamItemMember)
	for _, member := range teamItemMember {
		member_item_ids = append(member_item_ids, member.Item_id)
	}

	db.Model(Item{}).Where("user_id=?", uid).Or("id IN (?)", member_item_ids).Scan(&item)
	if uid <= 0 {
		return nil, errmsg.ERROR_USER_NO_RIGHT
	}
	var result []Item
	var result3 []Item
	for _, i3 := range item {
		if user.Role == 1 {
			i3.Creator = 1
		} else if i3.UserId == uid {
				i3.Creator = 1
			} else {
				i3.Creator = 0
			}
			if i3.Password != "" {
				i3.IsPrivate = 1
			} else {
				i3.IsPrivate = 0
			}

		result = append(result, i3)

	}

	db.Model(ItemSort{}).Where("uid = ?", uid).Find(&item_sort)
	if item_sort.Id > 0{
		var result2 map[int]int
		err = json.Unmarshal([]byte(item_sort.ItemSortData), &result2)
		type kv struct {
			Key   int
			Value int
		}
		var ss []kv
		var arrss []kv
		for _, i2 := range result {
			ss = append(ss , kv{int(i2.ID),99})
		}
		for k, v := range result2 {
			ss = append(ss, kv{k, v})
		}
		amap := make(map[int]int)
		for _, s := range ss {
           amap[s.Key] = s.Value
		}
		for k, i2 := range amap {
			arrss = append(arrss , kv{k,i2})
		}

		sort.Slice(arrss, func(i, j int) bool {
			if arrss [i].Value == arrss[i].Value{
				return arrss[i].Key < arrss[j].Key
			}
			return arrss[i].Value < arrss[j].Value
		})
		if err != nil {
			return nil, errmsg.ERROR
		}
		for _,k := range arrss {
			for _, i2 := range result {
				if int(i2.ID)== k.Key {

					result3 = append(result3, i2)
				}
			}
		}
	}else {
		result3= result
	}

	return result3, errmsg.SUCCESE
}
func GetOneItem(uid uint, id int, sid string) (Item, int,int) {
	var item Item
	var user User
	db.Model(User{}).Where("id =?",uid).Find(&user)
	err = db.Where("id =?", id).First(&item).Error
	if err != nil {
		return item, errmsg.ERROR_USER_NO_RIGHT,0
	}
	if checkItemVisit(uid, id, sid) != errmsg.SUCCESE {
		return item, errmsg.REQUIRE_PASSWORD_PERMISSION,0
	}
	return item, errmsg.SUCCESE,user.Role
}
func GetItemDetail(id int, uid uint) (Item, int) {
	var item Item
	if checkItemCreator(uid, id) != errmsg.SUCCESE {
		return item, errmsg.SUCCESE
	}
	db.Model(Item{}).Where("id = ?", id).Find(&item)
	return item, errmsg.SUCCESE
}
func DeleteItem(id int, password string, v uint) int {
	var item Item
	var user User
	db.Where("id =?", v).Find(&user)
	_, code,_ := CheckLogin(user.Username, password)
	if code == errmsg.SUCCESE {
		err := db.Where("id =?", id).Delete(&item).Error
		if err != nil {
			return errmsg.ERROR
		}
	}else {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	return errmsg.SUCCESE
}
func CreateItem(data *Item, password string) int {
	var page Page
	if password != "" {
		data.Password = ScryptPw(password)
	}
	if password != "" {
		data.IsPrivate = 1
	}
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	if data.Types == 2 {
		page.AuthorUid = data.UserId
		page.PageTitle = data.Title
		page.ItemId = data.ID
		page.CatId = 0
		page.PageContent = "欢迎使用doc。点击右上方的编辑按钮进行编辑吧！"
		err := db.Create(&page).Error
		if err != nil {
			return errmsg.ERROR
		}
	} else if data.Types == 4 {
		page.AuthorUid = data.UserId
		page.PageTitle = data.Title
		page.ItemId = data.ID
		page.CatId = 0
		page.PageContent = ""
		err := db.Create(&page).Error
		if err != nil {
			return errmsg.ERROR
		}
	}
	return errmsg.SUCCESE
}
func EditItem(id int, title string, description string, password string,langlist string,isComment string) int {
    var comment int
	var maps = make(map[string]interface{})
	if isComment == "true" {
		comment =1
	}else {
		comment = 0
	}
	maps["is_comment"] = comment
	maps["title"] = title
	maps["description"] = description
	maps["lang_list"] = langlist

	if password != "" {
		maps["password"] = ScryptPw(password)
		maps["is_private"] = 1
	}
	if password == "" {
		maps["is_private"] = 0
		maps["password"] = ""
	}
	err = db.Model(Item{}).Where("id =?", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESE
}
func ExitTeam(id int, uid uint) int {
	var item Item
	var teamItemMember TeamItemMember
	var teamMember TeamMember
	err = db.Model(Item{}).Where("id =?", id).Where("user_id = ?", uid).Delete(&item).Error
	if err != nil {
		return errmsg.ERROR
	}
	err = db.Model(TeamItemMember{}).Joins("LEFT JOIN team on team.id = team_item_member.team_id").Where("item_id=?", id).Where("member_uid=?", uid).Find(&teamItemMember).Error
	if err == nil {
		db.Model(TeamItemMember{}).Where("member_uid = ?", uid).Where("team_id =?", teamItemMember.Team_id).Delete(&teamItemMember)
		db.Model(TeamMember{}).Where("member_uid =?", uid).Where("team_id =?", teamItemMember.Team_id).Delete(&teamMember)
	}
	return errmsg.SUCCESE
}
func GetKey(itemId int, uid uint) (Data, int) {
	var item Item
	var token Data
	db.Model(Item{}).Where("id =?", itemId).Find(&item)
	if checkItemCreator(uid, itemId) != errmsg.SUCCESE {
		return token, errmsg.ERROR_USER_NO_RIGHT

	}
	token = GetTokenByItemId(itemId)
	return token, errmsg.SUCCESE
}

func ResetKey(itemId int, uid uint) (Data, int) {
	var item Item
	var data Data
	db.Where("id =?", itemId).Find(&item)
	if checkItemCreator(uid, itemId) != errmsg.SUCCESE {
		return data, errmsg.ERROR_USER_NO_RIGHT
	}
	err := db.Where("item_id =?", itemId).Delete(&data).Error
	if err != nil {
		return data, errmsg.ERROR
	}
	data, _ = GetKey(itemId, uid)
	return data, errmsg.SUCCESE
}
func Archive(itemId int, password string, uid uint) int {
	var item Item
	var user User
	db.Model(Item{}).Where("id =?", itemId).Find(&item)
	if checkItemCreator(uid, itemId) != errmsg.SUCCESE {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	db.Model(User{}).Where("id=?", uid).Find(&user)
	_, code,_ := CheckLogin(user.Username, password)
	if code != errmsg.SUCCESE {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	err := db.Model(Item{}).Where("id=?", itemId).Update("isarchived", 1).Find(&item).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESE
}

//转让项目
func AttornItem(itemId int, username string, password string, uid uint) int {
	var item Item
	var user User
	var user2 User

	db.Model(User{}).Where("id =?", uid).Find(&user)
	_, code,_ := CheckLogin(user.Username, password)
	if code != errmsg.SUCCESE {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	err := db.Model(User{}).Where("username =?", username).Find(&user2)
	if err == nil {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	db.Model(Item{}).Where("id = ?", itemId).Update("user_id", user2.ID).Scan(&item)

	return errmsg.SUCCESE
}
func LinkItem(id int, itemId int, lang int) int {
	var item Item
	err := db.Model(Item{}).Where("id =?", id).Update("lang", lang).Update("link", itemId).Find(&item).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESE
}
func CheckItemPermn(uid uint, itemId int) int {
	var item Item
	var itemMember ItemMember
	var itemTeamMember TeamItemMember
	var user User
	db.Model(User{}).Where("id =?",uid).Find(&user)
	if user.Role == 1 {
		 return errmsg.SUCCESE
	}
	if uid <= 0 {
		return errmsg.ERROR
	}
	db.Model(Item{}).Where("id = ?", itemId).Find(&item)
	if item.UserId == uid {
		return errmsg.SUCCESE
	}
	db.Where("item_id =?", itemId).Where("uid = ?", uid).Where("member_group_id = 1").Find(&itemMember)
	if itemMember.ID > 0 {
		return errmsg.SUCCESE
	}
	db.Model(TeamItemMember{}).Where("item_id =?", itemId).Where("member_uid =?", uid).Where("member_group_id = 1").Find(&itemTeamMember)

	if itemTeamMember.ID  >0{
		return errmsg.SUCCESE
	}
	return errmsg.ERROR
}

func checkItemVisit(uid uint, itemId int, sid string) int {
	var itemMember ItemMember
	var teamitemMember TeamItemMember
	var item Item
	var user User
	db.Model(User{}).Where("id =?",uid).Find(&user)
	if user.Role == 1 {
		return errmsg.SUCCESE
	}
	if checkItemCreator(uid, itemId) == errmsg.SUCCESE {
		return errmsg.SUCCESE
	}
	db.Model(ItemMember{}).Where("item_id =? ", itemId).Where("uid =?", uid).Find(&itemMember)
	if itemMember.ID != 0{
		return errmsg.SUCCESE
	}
	if utils.Md5(string(itemMember.Item_id)) == sid {
		return errmsg.SUCCESE
	}
	db.Model(TeamItemMember{}).Where("item_id =?", itemId).Where("member_uid =?", uid).Find(&teamitemMember)
	if teamitemMember.ID != 0{
		return errmsg.SUCCESE
	}
	if utils.Md5(string(teamitemMember.Item_id)) == sid {
		return errmsg.SUCCESE
	}

	db.Model(Item{}).Where("id =? ", itemId).Find(&item)
	if utils.Md5(string(itemId)) == sid {
		return errmsg.SUCCESE
	}
	if item.Password != "" {
		return errmsg.REQUIRE_PASSWORD_PERMISSION
	} else {
		return errmsg.SUCCESE
	}
}

func Pwd(itemId int, password string) int {
	var item Item
	db.Model(Item{}).Where("id =?", itemId).Find(&item)
	if item.Password == ScryptPw(password) {
		return errmsg.SUCCESE
	}
	return errmsg.ERROR_PASSWORD_WRONG
}

func getMenu(itemId uint, langId int) Menu {
	var menu Menu
	menu = getContent(itemId, langId)
	return menu
}
func getContent(itemId uint, langId int) Menu {
	var page []*Page
	var catalogs []Catalogs
	var catalogs2 []Catalogs
	var result []Catalogs
	if langId <= 0 {
		db.Model(Catalogs{}).Where("item_id = ?", itemId).Order("s_number asc , id asc").Preload("Lang").Find(&catalogs)
		db.Model(Page{}).Where("item_id =?", itemId).Order("s_number asc , id asc").Find(&page)
	} else {
		db.Model(Catalogs{}).Where("item_id = ?", itemId).Order("s_number asc , id asc").Where("cid =?", langId).Preload("Lang").Find(&catalogs)
		db.Model(Page{}).Where("item_id =?", itemId).Where("cid =?",langId).Order("s_number asc , id asc").Find(&page)
	}
	if &catalogs != nil {
		for _, catalog := range catalogs {
			if catalog.Level == 2 {
				catalogs2 = append(catalogs2, catalog)
			}
		}
	}
	if &catalogs2 != nil {
		for _, c := range catalogs2 {
			catalogs2 = getCat(c, page, catalogs)
			for _, c2 := range catalogs2 {
				result = append(result, c2)
			}
		}
	}
	var menu Menu
	menu.Page = GetPagesByItemId(itemId, "", langId)
	menu.Catalogs = result
	return menu
}
func getCat(catalogData Catalogs, page []*Page, catalogs []Catalogs) []Catalogs {
	var sub_catalogs []Catalogs
	var catalogsSubData []Catalogs
	var mainCatalogs []Catalogs
	var tmpCatalogs []Catalogs
	var tmpCatalogs1 []Catalogs

	catalogData.Page = _getPageByCatId(catalogData.ID, page)
	sub_catalogs = _getCatByCatId(catalogData.ID, catalogs)
	if sub_catalogs != nil {
		for _, catalog := range sub_catalogs {
			tmpCatalogs = getCat(catalog, page, catalogs)
			for _, tmpCatalog := range tmpCatalogs {
				tmpCatalogs1 = append(tmpCatalogs1 , tmpCatalog)
			}
			catalogData.Catalogs =tmpCatalogs1
				catalogsSubData = append(mainCatalogs, catalogData)
		}
	}
	if catalogData.Catalogs == nil {
		catalogsSubData = append(catalogsSubData, catalogData)
	}

	return catalogsSubData
}

func _getPageByCatId(catId uint, page []*Page) []*Page {
	var page2 []*Page
	if &page != nil {
		for _, p := range page {
			if p.CatId == catId {
				page2 = append(page2, p)
			}
		}
	}
	return page2
}
func _getCatByCatId(catId uint, catalogs []Catalogs) []Catalogs {
	var catalogs2 []Catalogs
	if &catalogs != nil {
		for _, catalog := range catalogs {
			if catalog.ParentCatId == catId {
				catalogs2 = append(catalogs2, catalog)
			}
		}
	}
	return catalogs2
}

func SortByItem(uid int, data string) int {
	var itemsort ItemSort
	db.Model(ItemSort{}).Where("uid = ?", uid).Delete(&itemsort)

	itemsort.ItemSortData = data
	itemsort.Uid = uid
	t := time.Now()
	itemsort.CreatedAt = t.Format("2006-01-02 15:04:05")
	err = db.Create(&itemsort).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESE
}

func FilterMemberItem (uid uint ,itemId uint , menuData Menu)Menu {
	var itemMember ItemMember
	var teamItemMember TeamItemMember
	var menuData2 Menu
	var menuData3 Menu

	var catId uint
	if &menuData == nil {
       return menuData2
	}

db.Model(ItemMember{}).Where("uid =?",itemId).Where("item_id = ?",itemId).Find(&itemMember)

if itemMember.Cat_id > 0 {
	catId = uint(itemMember.Cat_id)
}
db.Model(TeamItemMember{}).Where("member_uid =?",uid).Where("item_id =?",itemId).Find(&teamItemMember)
if teamItemMember.Cat_id >0{
	catId = uint(teamItemMember.Cat_id)
}
if catId >0 {
	for _, catalog := range menuData.Catalogs {
		if catalog.ID == catId {
			menuData3.Catalogs = append(menuData3.Catalogs,catalog)
			menuData3.Page = GetPagesByItemId(itemId, "", 0)
		}
	}
	return menuData3
}
return  menuData
}

func CheckItemEdit (uid uint , itemId uint)bool {
	var itemMember ItemMember
	var teamItemMember TeamItemMember
	var item Item
	if uid == 0 {
		return false
	}
	db.Model(Item{}).Where("user_id =?",uid).Where("id = ?",itemId).Find(&item)
	if &item != nil {
		return true
	}
	db.Model(ItemMember{}).Where("item_id =?",itemId).Where("uid =?",uid).Where("member_group_id = 1").Find(&itemMember)
	if &itemMember != nil{
		return true
	}
	db.Model(TeamItemMember{}).Where("item_id =?",itemId).Where("member_uid =?",uid).Where("member_group_id =1").Find(&teamItemMember)
	if &teamItemMember != nil {
		return true
	}
	return false
}
