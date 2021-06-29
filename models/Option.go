package models

import (
	"awesomeProject3/utils/errmsg"
	"github.com/jinzhu/gorm"
	"strings"
)

type Options struct {
	OptionId uint `gorm:"primaryKey;type int autoIncrement:false" json:"optionId"`
	OptionName string `gorm:"type:varchar(255)" json:"optionName"`
	OptionValue string `gorm:"type:varchar(255)" json:"optionValue"`
}
func SaveConfig (uid uint,id int)int{
	if checkAdmin(uid) != errmsg.SUCCESE{
		return errmsg.ERROR
	}
err =db.Model(Options{}).Where("option_name = ?","register_open").Update("option_value",id).Error
if err != nil {
	return errmsg.ERROR
}
return errmsg.SUCCESE
}
func RegisterSetting ()string {
	var option Options
	err = db.Model(Options{}).Where("option_name = ? ","register_open").Find(&option).Error
	if err ==  gorm.ErrRecordNotFound{
		option.OptionId= 1
		option.OptionName= "register_open"
		option.OptionValue = "0"
		err = db.Create(&option).Error
		if err != nil {
			return ""
		}
	}
	return option.OptionValue
}

func LoadConfig() ([]Options,int){
	var option []Options
	var result []Options
	err = db.Model(Options{}).Where("option_name = ?","register_open").Find(&option).Error
	for _, options := range option {
		if options.OptionValue == "1"{
			options.OptionValue = "false"
		}else if options.OptionValue == "0"{
			options.OptionValue = "true"
		}
		result= append(result , options)
	}
	if err != nil {
		 return result,errmsg.ERROR
	}
return result,errmsg.SUCCESE
}

func SaveLangConfig(uid uint , id int) int{
	if checkAdmin(uid) != errmsg.SUCCESE{
		return errmsg.ERROR
	}
	err =db.Model(Options{}).Where("option_name = ?","lang").Update("option_value",id).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESE
}
func LoadLangConfig(itemId int) (Lang,int){
	var option Options
	var lang Lang
	var item Item
	err = db.Model(Options{}).Where("option_name = ?","lang").Find(&option).Error
	db.Model(Item{}).Where("id = ?",itemId).Find(&item)
	s := strings.Split(item.LangList, ",")
	for _, i2 := range s  {
		if i2 == option.OptionValue{
			db.Model(Lang{}).Where("id =?",option.OptionValue).Find(&lang)
			return lang,errmsg.SUCCESE
		}
	}
	db.Model(Lang{}).Where("id =?",s[0]).Find(&lang)
	if err != nil {
		return lang,errmsg.SUCCESE
	}
	return lang,errmsg.SUCCESE
}
