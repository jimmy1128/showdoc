package models

import (
	"awesomeProject3/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Lang struct {
	//gorm.Model
	ID uint `gorm:"primaryKey;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
	Icon string `gorm:"type:varchar(255);not null" json:"icon"`
}

func CheckLang(name string)(code int){
	var cate Lang
	var options Options
	options.OptionId = 2
	options.OptionName = "lang"
	options.OptionValue = "0"
	if options.OptionName != "lang"{
		db.Model(Options{}).Create(&options)
	}

	db.Select("id").Where("name =?",name).First(&cate)

	if cate.ID > 0 {
		return errmsg.ERROR_LANG_USED //1001
	}
	return errmsg.SUCCESE
}
//新增分类
func CreateLang(data *Lang)int{
	err:= db.Create(&data).Error
	if err != nil{
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESE
}
//todo 查询分类下的所有分类
func GetLangInfo(id int) (Lang,int){
	var cate Lang
	db.Where("id = ?",id).First(&cate)
	return cate,errmsg.SUCCESE
}

//查询分类列表
func GetLang(pageSize int ,pageNum int) ([]Lang,int){
	var cate []Lang
	var data [] Lang
	var total int
	err = db.Find(&cate).Count(&total).Limit(pageSize).Offset((pageNum-1) * pageSize).Error
	for _, lang := range cate {
		if lang.Icon != "" {
			lang.Icon = `<use xlink:href="` + lang.Icon + `"></use>`
		}
		data = append(data, lang)
	}

	if err != nil && err != gorm.ErrRecordNotFound{
		return nil,0
	}
	return data,total
}
//编辑分类信息
func EditLang(id int,data *Lang)int{
	var cate Lang
	var maps = make(map[string]interface{})
	maps["name"]= data.Name
	maps["icon"]= data.Icon
	err =db.Model(&cate).Where("id =?",id ).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESE
}
//删除用户
func DeleteLang(id int) int {
	var cate Lang
	err = db.Where("id =?",id).Delete(&cate).Error
	if err!= nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESE
}