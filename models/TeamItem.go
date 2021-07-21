package models

import (
	"awesomeProject3/utils/errmsg"
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
)
type TeamItem struct {
	gorm.Model
	Item_id uint `gorm:"type:int;not null" json:"item_id"`
	Team_id uint `gorm:"type:int;not null" json:"team_id"`

}

type TeamItemInfo struct {
	TeamItem TeamItem `gorm:"foreignkey:Cid"`
	Id int `gorm:"type:int;not null" json:"item_id"`
	ID uint `gorm:"type:int" json:"cid"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	Title string `gorm:"type:varchar(255)" json:"titlename"`
	Password string `gorm:"type:varchar(255)" json:"password"`
	UserId uint `gorm:"type:int;not null" json:"userid"`
	Team_id int `gorm:"type:int;not null" json:"team_id"`
    Types int `gorm:"type:int;not null" json:"types"`
    CreatedAt string `gorm:"type:date" json:"created_at"`
	DeletedAt string `gorm:"type:date" json:"deleted_at"`
}

type TeamItemInfo2 struct {
	TeamItem TeamItem `gorm:"foreignkey:Cid"`
	Cid uint `gorm:"type:int" json:"cid"`
	Team_id  int `gorm:"type:int;not null" json:"team_id"`
	Item_id int `gorm:"type:int;not null" json:"item_id"`
	Uid uint `gorm:"type:int;not null" json:"uid"`
	CreatedAt string `gorm:"type:date" json:"created_at"`
	UpdateAt string `gorm:"type:date" json:"update_at"`
	Teamname   string `gorm:"type:varchar(255)" json:"teamname"`

}

func SaveTeamItem(uid uint , id string,teamId int)(TeamItem,int){
	var teamInfo Team
	var itemIdArray []string
	var teamItem TeamItem
	var teamMember []TeamMember
	var teamitemMember TeamItemMember
	err = db.Model(Team{}).Where("id=?",teamId).Where("uid=?",uid).Find(&teamInfo).Error
	if err != nil {
		 return teamItem,errmsg.ERROR
	}
	itemIdArray = strings.Split(id,",")
	for _,v := range itemIdArray {
		i, _ := strconv.Atoi(v)
		if checkItemCreator(uid, i) != errmsg.SUCCESE {
			return teamItem, errmsg.ERROR
		}
		if db.Model(TeamItem{}). Where("team_id= ?",teamId).Where("item_id=?",i).Find(&teamItem) ==nil {
			continue
		}
		teamItem.Item_id = uint(i)
		teamItem.Team_id = uint(teamId)
		db.Model(TeamMember{}).Where("team_id =?",teamId).Find(&teamMember)
		for _,v := range teamMember {
			teamitemMember.Team_id = v.Team_id
			teamitemMember.Member_uid = v.Member_uid
			teamitemMember.Member_username =v.Member_username
			teamitemMember.Item_id= uint(i)
			teamitemMember.Member_group_id=1
	}
		db.Create(&teamitemMember)
		err := db.Create(&teamItem).Error
		if err != nil {
			return teamItem , errmsg.ERROR_TEAM_ITEM_EXIT
		}

	}
	return teamItem , errmsg.SUCCESE
}

func checkItemCreator (uid uint, item_id int)int{
	var item Item
	var user User
	db.Model(User{}).Where("id = ?",uid).Find(&user)
	if user.Role == 1 {
		 return errmsg.SUCCESE
	}
	if uid == 0 {
		return errmsg.ERROR
	}
	db.Model(Item{}).Where("id=?",item_id).Find(&item)
	if item.UserId == uid{
		return errmsg.SUCCESE
	}
return errmsg.ERROR
}

func GetTeamItemList(uid uint, itemId int)([]TeamItemInfo2,int){
	var teamItem []TeamItemInfo2
	if checkItemCreator(uid,itemId) != errmsg.SUCCESE{
		return teamItem, errmsg.ERROR
	}
	db.Model(Team{}).Where("team_item.item_id=?",itemId).Where("team_item.deleted_at is NULL").Joins("LEFT JOIN team_item on team.id = team_item.team_id").Select("team.*,team_item.team_id,team_item.id as cid,item_id").Scan(&teamItem)
	return teamItem , errmsg.SUCCESE
}

func GetListByTeam (teamId int ,uid uint )([]TeamItemInfo,int) {
	var team [] Team
	var teamItemInfo []TeamItemInfo
	err =db.Model(Team{}).Where("id =?",teamId). Where("uid=?",uid).Find(&team).Error
	if err != nil {
		return teamItemInfo,errmsg.ERROR_USER_NO_RIGHT
	}
	err=db.Model(Item{}).Select("item.*,team_item.team_id,team_item.id").Where("team_item.team_id=?",teamId).Where("team_item.deleted_at is NULL").Joins("LEFT JOIN team_item on item.id = team_item.item_id").Scan(&teamItemInfo).Error

return teamItemInfo , errmsg.SUCCESE
}

func DeleteTeamItem(id int , uid  uint)int{
	var teamItem TeamItem

	var itemId int
	var teamId int
	var teamItemMember TeamItemMember
	err = db.Model(TeamItem{}).Where("id=?",id).Find(&teamItem).Error
	if err == gorm.ErrRecordNotFound{
		return errmsg.ERROR
	}
	itemId = int(teamItem.Item_id)
	teamId = int(teamItem.Team_id)

	if checkItemCreator(uid,itemId) != errmsg.SUCCESE{
		return errmsg.ERROR_USER_NO_RIGHT
	}

	err = db.Model(&TeamItemMember{}).Where("item_id=?",itemId).Where("team_id=?",teamId).Delete(&teamItemMember).Error
	err = db.Model(&TeamItem{}).Where("id=?",id).Delete(&teamItem).Error
	if err != nil {
		 return errmsg.ERROR
	}
	return errmsg.SUCCESE
}