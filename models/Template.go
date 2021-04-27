package models

import (
	"awesomeProject3/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Template struct {
	gorm.Model
	UserId uint `gorm:"type:int;not null" json:"userid"`
	Title string `gorm:"type:varchar(255);not null" json:"title"`
	Content string `gorm:"type:varchar(255)" json:"content"`
	Username string `gorm:"type:varchar(30);not null" json:"username"`
	CatId uint `gorm:"type:int;not null" json:"catid"`

}

func CreateTemplate(data *Template)int {
	err := db.Create(&data).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESE
}

func EditTemplate(id int , data *Template)int{
	var temp Template
	var maps = make(map[string]interface{})
	maps["userid"] = data.UserId
	maps["title"] = data.Title
	maps["content"] = data.Content
	err = db.Model(&temp).Where("id =?",id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESE
}

func GetTemplateByUid(id uint)([]Template,int){
	var templateList []Template

	err := db.Where("user_id =?",id).Find(&templateList).Error
	if err != nil {
		return templateList,errmsg.ERROR
	}
	return templateList,errmsg.SUCCESE
}

func GetTemplate(id int)(Template,int){
	var temp Template
	err := db.Preload("Template").Where("id =?",id).First(&temp).Error
	if err != nil{
		return temp,errmsg.ERROR_ART_NO_EXIST
	}
	return temp ,errmsg.SUCCESE
}

func DeleteTemplate(id int)int{
var temp Template
err = db.Where("id =?",id).Delete(&temp).Error
if err != nil{
	return errmsg.ERROR
}
return errmsg.SUCCESE
}

func (data *Template)Save()(*Template,int){

	if data.ID == 0 {
		if CreateTemplate(data) == 200{
			return data,errmsg.SUCCESE
		}else {
			return data,errmsg.ERROR
		}

	}else {
		err = db.Table("Template").Where("id =?",data.ID).Update(data).Error
		if err != nil {
			return nil,errmsg.ERROR
		}

	}
	return data,errmsg.SUCCESE
}

