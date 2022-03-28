package models

import (
	"awesomeProject3/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Team struct {
	gorm.Model
	Uid         uint   `gorm:"type:int;not null" json:"uid"`
	Username    string `gorm:"type:varchar(255)" json:"username"`
	Teamname    string `gorm:"type:varchar(255)" json:"teamname"`
	MemberCount int    `gorm:"type:int" json:"memberCount"`
	ItemCount   int    `gorm:"type:int" json:"itemCount"`
}

//添加和编辑

func (team *Team)SaveTeam()(*Team,int) {
	if team.ID == 0 {
		err := db.Create(&team).Error
		if err != nil {
			return team ,errmsg.ERROR
		}
		return team,errmsg.SUCCESE
	}else {
		err := db.Model(Team{}).Where("id=?",team.ID).Update(&team).Error
		if err!= nil{
			return team,errmsg.ERROR
		}
		return team,errmsg.SUCCESE
	}
}
func List(uid int)([]Team,int){
	var team []Team
	var memberCount int
	var itemCount int
	var maps = make(map[string]interface{})
	db.Model(Team{}).Where("uid =?",uid).Order("created_at").Find(&team)
	for _, t := range team {
		db.Model(TeamMember{}).Where("team_id=?",t.ID).Count(&memberCount)
		db.Model(TeamItem{}).Where("team_id =?",t.ID).Where("item.deleted_at is NULL").Joins("LEFT JOIN item ON item.id = team_item.item_id").Count(&itemCount)
		maps["member_count"]=memberCount
		maps["item_count"]=itemCount
		db.Model(Team{}).Where("id=?",t.ID).Update(maps)
	}
	if err != nil {
		if err == gorm.ErrRecordNotFound{
			return team ,errmsg.SUCCESE
		}
		return team,errmsg.ERROR
	}
	return team,errmsg.SUCCESE
}

func DeleteTeam(v int,id int )int{
	var team Team
	err = db.Where("id=?",id).Where("uid=?",v).Delete(&team).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESE
}
//转让团队
func Attorn(id int ,uid int,username string,password string)int {
	var team Team
	var teamItem TeamItem
	var user User

	err = db.Where("id =?",id ).Where("uid=?",uid).Find(&team).Error
	if err != nil {
		return errmsg.ERROR
	}
	db.Model(User{}).Select("username").Where("id =?",uid).Find(&user)
	_,code,_ := CheckLogin(user.Username,password)
	if code != 200 {
		return  errmsg.ERROR
	}
	db.Model(User{}).Where("username =?", username).Find(&user)

	team.Username = user.Username
	team.Uid = user.ID
	db.Model(Team{}).Where("id=?",id).Update(team)


	//读取出该团队的所有项目，准备转上
	db.Model(TeamItem{}).Where("team_id =?",id).Find(&teamItem)
	db.Model(Item{}).Where("id =? ",teamItem.Item_id).Update(team)


return errmsg.SUCCESE
}
