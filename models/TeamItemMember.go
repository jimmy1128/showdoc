package models

import (
	"awesomeProject3/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type TeamItemMember struct {
	gorm.Model
	Item_id uint `gorm:"type:int;not null" json:"item_id"`
	Member_group_id uint `gorm:"type:int;default:1" json:"member_group_id"`
	Cat_id int `gorm:"type:int;not null" json:"cat_id"`
	Team_id uint `gorm:"type:int;not null" json:"team_id"`
	Name string `gorm:"type:varchar(255)" json:"catname"`
	Member_uid uint `gorm:"type:int;default:1" json:"member_uid"`
	Member_username string `gorm:"type:varchar(255)" json:"member_username"`
	Catalogs Catalogs `gorm:"foreignkey:Cid"`


}

func SaveTeamItemMember(id int , catId int , uid uint , memberGroupId int) (TeamItemMember,int){
	var teamItemMemberInfo TeamItemMember
	var maps = make(map[string]interface{})
	var item_id int
	var team_id int
	db.Model(TeamItemMember{}).Where("id=?",id).Find(&teamItemMemberInfo)
	item_id = int(teamItemMemberInfo.Item_id)
	team_id = int(teamItemMemberInfo.Team_id)
	if checkItemCreator(uid,item_id) != errmsg.SUCCESE{
		return teamItemMemberInfo,errmsg.ERROR
	}
	err =db.Model(Team{}).Where("id =?",team_id).Where("uid =?",uid).Error
	if err != nil {
		 return teamItemMemberInfo,errmsg.ERROR
	}
	maps["member_group_id"] = memberGroupId
	maps["cat_id"]=catId
	db.Model(TeamItemMember{}).Where("id=?",id).Update(maps)
	if err != nil {
		 return teamItemMemberInfo,errmsg.ERROR
	}
	return teamItemMemberInfo,errmsg.SUCCESE
}

func GetTeamItemMemberList(uid uint , itemId int , teamId int) ([]TeamItemMember,int){
	var teamItemMember []TeamItemMember
	var catalogs Catalogs
	if checkItemCreator(uid,itemId) != errmsg.SUCCESE{
		return teamItemMember,errmsg.ERROR_USER_NO_RIGHT
	}
	db.Model(TeamItemMember{}).Where("item_id=?",itemId).Where("team_id=?",teamId).Find(&teamItemMember)
	for _, member := range teamItemMember {
		if member.Cat_id == 0 {
			member.Name = "所有目录"
		}

		if member.Cat_id >0{
			db.Model(Catalogs{}).Where("id = ?",member.Cat_id).Find(&catalogs)
			member.Name = catalogs.Name
		}
	}
	return teamItemMember,errmsg.SUCCESE
}