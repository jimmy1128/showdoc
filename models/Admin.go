package models

import (
	"awesomeProject3/utils/errmsg"
	"fmt"
	"github.com/jinzhu/gorm"
)

type ItemListInfo struct {
	ID int `gorm:"type:int" json:"id"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	Password string `gorm:"type:varchar(255)" json:"password"`
	UserId uint `gorm:"type:int" json:"userid"`
	Username string `gorm:"type:varchar(255)" json:"username"`
	Member_num int `gorm:"type:int" json:"member_num"`
	Title string `gorm:"type:varchar(255)" json:"titlename"`
	CreatedAt string `gorm:"type:date" json:"created_at"`
	DeletedAt string `gorm:"type:date" json:"deleted_at"`
}

type TotalMember struct {
	ItemId int `gorm:"type:int" json:"item_id"`
	MemberNum int `gorm:"type:int" json:"member_num"`
}

func AdminGetList(uid uint , pageSize int , pageNum int,username string)([]User,int) {
	var total int
	var user []User
if checkAdmin(uid) != errmsg.SUCCESE{
	return nil,0
}
if username != "" {
	err := db.Model(User{}).Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Where("username LIKE ? ",username + "%").Find(&user).Error
	db.Model(User{}).Where("username LIKE ? ",username + "%").Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil  , 0
	}
	return user,total
}
	err := db.Limit(pageSize).Offset((pageNum -1) * pageSize).Where("1=1").Find(&user).Error
	db.Model(&user).Where("1=1").Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound{
		 return nil , 0
	}
return user,total
}
func checkAdmin (uid uint)int{
	var user User
	if uid != 0 {
		db.Model(User{}).Where("id = ?",uid).Find(&user)
		if user.Role == 1 {
			return errmsg.SUCCESE
		}
	}else {
		return errmsg.ERROR
	}
	return errmsg.SUCCESE
}

func AdminDeleteUser (uid uint,id int)int{
	 var item Item
	 var user User
	if checkAdmin(uid) != errmsg.SUCCESE{
		return errmsg.ERROR
	}
	db.Model(Item{}).Where("user_id = ?",id).Find(&item)
	 if item.ID < 0 {
	 	return errmsg.ERROR_EXITING_ITEM
	 }
	 err := db.Model(User{}).Where("id = ?",id).Delete(&user).Error
	 if err != nil {
	 	return errmsg.ERROR
	 }
	 return errmsg.SUCCESE
}

func AdminEditPassword (uid uint,newPassword string,id int)int{
	if checkAdmin(uid) != errmsg.SUCCESE{
		return errmsg.ERROR
	}
	err := db.Model(User{}).Select("password").Where("id =?",id).Update("password",newPassword).Error
	if err != nil {
		 return errmsg.ERROR
	}
return errmsg.SUCCESE
}


func AdminAddUser (uid uint, username string , password string , name string) int{
	fmt.Println(uid,username,password,name)
var user User
	if uid != 0 {
		if password != "" {
			db.Model(User{}).Where("id =? ",uid).Update("password",ScryptPw(password))
		}
		if name != "" {
			db.Model(User{}).Where("id =?",uid).Update("name",name)
		}
	}else {
		if CheckUser(username) != errmsg.SUCCESE{
			fmt.Println(CheckUser(username))
			return errmsg.ERROR_USERNAME_USED
		}
		user.Username = username
		user.Password = ScryptPw(password)
		user.Name = name
		err := db.Create(&user).Error
		if err != nil {
			return errmsg.ERROR
		}
	}
	return errmsg.SUCCESE
}

func AdminItemGetList(uid uint , pageSize int , pageNum int,itemName string,username string)([]ItemListInfo,int,int){
	// var itemMember []ItemMember
	var itemInfo []ItemListInfo
	var result []ItemListInfo
	var totalMember []TotalMember
	var total int

	if checkAdmin(uid) != errmsg.SUCCESE{
		return nil,total,errmsg.ERROR
	}

	if itemName != ""{
		err := db.Table("Item").Select("user.username,item.*").Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Joins("LEFT JOIN user on item.user_id = user.id").Where("title LIKE ? ",itemName + "%").Scan(&itemInfo).Error
		db.Model(Item{}).Where("title LIKE ? ",itemName + "%").Count(&total)
		db.Model(ItemMember{}).Select("item_id , count(uid) as member_num").Group("item_id").Scan(&totalMember)
		for _, i2 := range itemInfo {
			i2.Member_num = getMemberNum(totalMember, i2.ID)
			result = append(result, i2)
		}
		if err != nil && err != gorm.ErrRecordNotFound{
			return nil ,total, 0
		}
		return result,total,errmsg.SUCCESE
	}
	if username != "" {
		err := db.Model(Item{}).Select("user.username,item.*").Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Joins("LEFT JOIN user on item.user_id = user.id").Where("user.username LIKE ? ",username + "%").Scan(&itemInfo).Error
		db.Model(Item{}).Joins("LEFT JOIN user on item.user_id = user.id").Where("User.username LIKE ? ",username + "%").Count(&total)
		db.Model(ItemMember{}).Select("item_id , count(uid) as member_num").Group("item_id").Scan(&totalMember)
		for _, i2 := range itemInfo {
			i2.Member_num = getMemberNum(totalMember, i2.ID)
			result = append(result, i2)
		}
		if err != nil && err != gorm.ErrRecordNotFound{
			return nil ,total, 0
		}
		return result,total,errmsg.SUCCESE
	}

	err := db.Model(Item{}).Select("user.username,item.*").Limit(pageSize).Offset((pageNum - 1) * pageSize).Joins("LEFT JOIN user on item.user_id = user.id").Order("Created_At DESC").Scan(&itemInfo).Error
	db.Model(Item{}).Where("deleted_at is NULL").Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil  ,total, 0
	}
	db.Model(ItemMember{}).Select("item_id , count(uid) as member_num").Group("item_id").Scan(&totalMember)
	for _, i2 := range itemInfo {
		i2.Member_num = getMemberNum(totalMember, i2.ID)
		result = append(result, i2)
	}
return result,total,errmsg.SUCCESE
}

func getMemberNum(totalMember []TotalMember , itemId int)int{
	 var members int
	if totalMember !=nil  {
		for _, member := range totalMember {
			if member.ItemId == itemId {
				members = member.MemberNum
				return members
			} else {
				members = 1
			}
		}
	}
	return members
}

func AdminDeleteItem(uid uint, itemId int)int{
	var item Item
	if checkAdmin(uid) != errmsg.SUCCESE{
		return errmsg.ERROR
	}
	err := db.Model(Item{}).Where("id =?",itemId).Delete(&item).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESE
}

func AdminItemAttorn(uid uint , username string ,itemId int )int{
	var item Item
	var user User
	var maps = make(map[string]interface{})
	if checkAdmin(uid) != errmsg.SUCCESE{
		return errmsg.ERROR
	}
	db.Model(Item{}).Where("id =? ",itemId).Find(&item)
	err = db.Model(User{}).Where("username =? ",username).Find(&user).Error
	if err != nil {
		return errmsg.ERROR_USER_NOT_EXIST
	}

if &user == nil {
	return errmsg.ERROR
}
	maps["user_id"] = user.ID
err := db.Model(Item{}).Where("id =?",itemId).Update(maps).Error
if err != nil {
	 return errmsg.ERROR
}
return errmsg.SUCCESE
}

