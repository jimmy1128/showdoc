package models

import (
	"awesomeProject3/utils/errmsg"
	"encoding/json"
)

type ItemAvatar struct {
	AvatarId     uint   `gorm:"primary_key;autoIncrement:true" json:"avatarId"`
	AvatarName   string `gorm:"type:varchar(255)" json:"avatarName"`
	AvatarUrl    string `gorm:"type:varchar(255)" json:"avatarUrl"`
	AvatarItemId uint   `gorm:"type:int" json:"avatarItemId"`
}

type ItemHeader struct {
	HeaderId   uint   `gorm:"primary_key;autoIncrement:true" json:"headerId"`
	HeaderName string `gorm:"type:varchar(255)" json:"headerName"`
	HeaderUrl  string `gorm:"type:varchar(255)" json:"headerUrl"`
	ItemId     uint   `gorm:"type:int" json:"itemId"`
	Cid        int    `gorm:"type:int" json:"cid"`
}

func SaveAvatar(itemId int, name string, url string, uid uint) int {
	var avatar ItemAvatar
	db.Model(ItemAvatar{}).Where("avatar_item_id = ?", itemId).Find(&avatar)
	avatar.AvatarItemId = uint(itemId)
	avatar.AvatarName = name
	avatar.AvatarUrl = url
	if avatar.AvatarId == 0 {
		err = db.Model(ItemAvatar{}).Create(&avatar).Error
		if err != nil {
			return errmsg.ERROR
		}
		return errmsg.SUCCESE
	}
	if avatar.AvatarId != 0 {
		err = db.Model(ItemAvatar{}).Update(&avatar).Error
		return errmsg.SUCCESE
	}

	return errmsg.SUCCESE
}
func GetAvatar(itemId int) (ItemAvatar, int) {
	var avatar ItemAvatar
	err = db.Model(ItemAvatar{}).Where("avatar_item_id = ?", itemId).Find(&avatar).Error

	return avatar, errmsg.SUCCESE
}

func SaveHeader(itemId int, v uint, data string) int {
	var itemHeader ItemHeader
	var tmpheader []ItemHeader
	var maps []interface{}
	var store []uint
	var tmpStore []uint
	if v < 0 {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	db.Model(ItemHeader{}).Where("item_id =?", itemId).Find(&tmpheader)
	for _, header := range tmpheader {
		store = append(store, header.HeaderId)
	}

	err = json.Unmarshal([]byte(data), &maps)
	for _, i := range maps {
		restlt := i.(map[string]interface{})
		if restlt["headerId"] != nil {
			itemHeader.HeaderId = uint(restlt["headerId"].(float64))
		}

		itemHeader.Cid = int(restlt["cid"].(float64))
		itemHeader.HeaderName = restlt["headerName"].(string)
		itemHeader.HeaderUrl = restlt["headerUrl"].(string)
		itemHeader.ItemId = uint(itemId)
		tmpStore = append(tmpStore, itemHeader.HeaderId)

		if itemHeader.HeaderId == 0 {
			err = db.Model(ItemHeader{}).Create(&itemHeader).Error
		} else {
			err = db.Model(ItemHeader{}).Where("header_id = ?", itemHeader.HeaderId).Update(&itemHeader).Error
		}
	}
	_, removed := Arrcmp(store, tmpStore)
	for _, u := range removed {
		db.Model(ItemHeader{}).Where("header_id =?", u).Delete(&itemHeader)
	}
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESE
}

func GetHeader(itemId int, langId int) ([]ItemHeader, int) {
	var itemHeader []ItemHeader
	db.Model(ItemHeader{}).Where("item_id = ?", itemId).Where("cid =?", langId).Find(&itemHeader)
	return itemHeader, errmsg.SUCCESE
}

func DeleteHeader(uid uint, headerId int) int {
	var itemHeader ItemHeader
	if uid < 0 {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	err = db.Model(ItemHeader{}).Where("header_id = ?", headerId).Delete(&itemHeader).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESE
}
func Arrcmp(src []uint, dest []uint) ([]uint, []uint) {
	msrc := make(map[uint]byte) //按源数组建索引
	mall := make(map[uint]byte) //源+目所有元素建索引

	var set []uint //交集

	//1.源数组建立map
	for _, v := range src {
		msrc[v] = 0
		mall[v] = 0
	}
	//2.目数组中，存不进去，即重复元素，所有存不进去的集合就是并集
	for _, v := range dest {
		l := len(mall)
		mall[v] = 1
		if l != len(mall) { //长度变化，即可以存
			l = len(mall)
		} else { //存不了，进并集
			set = append(set, v)
		}
	}
	//3.遍历交集，在并集中找，找到就从并集中删，删完后就是补集（即并-交=所有变化的元素）
	for _, v := range set {
		delete(mall, v)
	}
	//4.此时，mall是补集，所有元素去源中找，找到就是删除的，找不到的必定能在目数组中找到，即新加的
	var added, deleted []uint
	for v, _ := range mall {
		_, exist := msrc[v]
		if exist {
			deleted = append(deleted, v)
		} else {
			added = append(added, v)
		}
	}

	return added, deleted
}
