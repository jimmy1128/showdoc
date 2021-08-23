package models

import (
	"awesomeProject3/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	PageId uint `json:"page_id"`
	GuestId uint `json:"guest_id"`
	Username string `json:"username"`
	Rate int `json:"rate"`
	Content string `gorm:"type:varchar(500);not null;" json:"content"`
	Status int8 `gorm:"type:tinyint;default:2" json:"status"`
}

func AddComment (data *Comment)int{

	err = db.Create(&data).Error
	if err != nil {
		 return errmsg.ERROR
	}
	return errmsg.SUCCESE
}

func GetComment(id int)(Comment,int){
	var comment Comment
	err = db.Where("id=?",id).Find(&comment).Error
	if err != nil {
		return comment ,errmsg.ERROR
	}
	return comment ,errmsg.SUCCESE

}

func GetCommentList(){

}
func GetCommentCount (id int)int64 {
	var comment Comment
	var total int64
	db.Find(&comment).Where("page_id=?",id).Where("status=?",2).Count(&total)
	return total
}

func GetCommentListFront(id int,pageSize int,pageNum int)([]Comment,int64,int){
	var commentList [] Comment
	var total int64
	db.Model(&Comment{}).Where("page_id =?",id).Where("status = ?",2).Count(&total)
	err = db.Model(&Comment{}).Limit(pageSize).Offset((pageNum-1)*pageSize).Order("Created_At DESC").Where("page_id =?",id).Where("status =?",2).Scan(&commentList).Error
	if err != nil{
		return commentList,0,errmsg.ERROR
	}
	return commentList,total,errmsg.SUCCESE
}
func DeleteComment(id uint) int {
	var comment Comment
	err = db.Where("id = ?",id).Delete(&comment).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESE
}
