package models

import (
	"awesomeProject3/utils/errmsg"
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

type ItemMember struct {
	gorm.Model
	Username string `gorm:"type:varchar(255)" json:"username"`
	Uid uint `gorm:"type:int;not null" json:"uid"`
	Item_id int `gorm:"type:int;not null" json:"item_id"`
	Member_group_id int `gorm:"type:int;not null" json:"member_group_id"`
	Member_group string `gorm:"type:varchar(255)" json:"member_group"`
	Cat_id int `gorm:"type:int;not null" json:"cat_id"`
	Cat_name string`gorm:"type:varchar(255)" json:"cat_name"`
}

func SaveMember(itemId int ,catId int ,uid uint ,username string , membergroupid int)(ItemMember, int){
	var usernameArray [] string
	var itemMember ItemMember
	//var itemMember2 ItemMember
	var userdata []User
	var itemMemberdata []ItemMember
	if checkItemCreator(uid , itemId) != errmsg.SUCCESE{
		return itemMember, errmsg.ERROR_USER_NO_RIGHT
	}
     usernameArray = strings.Split(username,",")

	for _, s := range usernameArray {
		var user User
		db.Where("username in (?)",s).Find(&user)
		userdata = append(userdata , user)
	}

	//db.Where("uid = ?" ,uid).Where("item_id =?",itemId).Find(&itemMember)

	for _, userdatum := range userdata {
		itemMember.Username = userdatum.Username
		itemMember.Uid = userdatum.ID
		itemMember.Item_id = itemId
		itemMember.Member_group_id = membergroupid
		itemMember.Cat_id = catId
		itemMemberdata= append(itemMemberdata , itemMember)
	}
	for _, memberdatum := range itemMemberdata {

		fmt.Println(memberdatum,"1")
		err := db.Model(ItemMember{}).Create(&memberdatum).Error
		if err != nil {
			return itemMember,errmsg.ERROR
		}
	}

	return itemMember,errmsg.SUCCESE

}

func GetList(itemId int,uid uint)([]ItemMember,int) {
	var itemMember []ItemMember
	var catalogs Catalogs
	var result []ItemMember
	if checkItemCreator(uid,itemId) != errmsg.SUCCESE{
		return result,errmsg.ERROR_USER_NO_RIGHT
	}
	 if itemId > 0 {
	 	db.Model(ItemMember{}).Where("item_id =?",itemId).Joins("LEFT JOIN user on user.id = item_member.uid").Select("item_member.*,user.name as name").Order("created_at asc").Scan(&itemMember)
	 }
	for _, member := range itemMember {
		if member.Member_group_id == 1 {
			member.Member_group = "编制/目录"
		}else{
			member.Member_group = "只读/目录"
		}
		member.Cat_name = "所有目录"
		if member.Cat_id > 0 {
			db.Where("id =?",member.Cat_id).Find(&catalogs)
			member.Cat_name = catalogs.Name
		}
		if member.Member_group_id == 1 {
			member.Member_group = "编制/目录:"+ member.Cat_name
		}else {
			member.Member_group = "只读/目录" + member.Cat_name
		}
	result = append(result,member)
	}
	return result ,errmsg.SUCCESE
}

func DeleteMember (itemId int,uid uint,itemMemberId int)int {
	var itemMember ItemMember

	if checkItemCreator(uid, itemId) != errmsg.SUCCESE {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	db.Model(ItemMember{}).Where("item_id=?",itemId).Where("id =?", itemMemberId).Find(&itemMember)
	err := db.Model(ItemMember{}).Where("item_id=?",itemId).Where("id =?", itemMemberId).Delete(&itemMember).Error
	if err != nil{
		return errmsg.ERROR
	}
return errmsg.SUCCESE
}
