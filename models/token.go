package models

import (
	"awesomeProject3/utils/errmsg"
	"crypto/md5"
	"fmt"
	"github.com/jinzhu/gorm"
	"io"
	"strconv"
	"time"
    "math/rand"
)
type Data struct {
	gorm.Model
	Item_id int `gorm:"type:int;not null" json:"item_id"`
	ApiKey string `gorm:"type:varchar(255)" json:"apikey"`
	ApiToken string `gorm:"type:varchar(255)" json:"apitoken"`

}
func CreateToken (itemId int)Data {
    var data Data
	b := []byte("showdoc")
	c := []byte("greypanel")
	rand.Read(b)
	rand.Read(c)
    crutime := time.Now().Unix()
	h := md5.New()
	h.Write(b)
    io.WriteString(h, strconv.FormatInt(crutime, 10))
    apiKey := fmt.Sprintf("%x",h.Sum(b))
    h.Write(c)
	apiToken := fmt.Sprintf("%x",h.Sum(c))
    data.Item_id = itemId
	data.ApiKey = apiKey
	data.ApiToken = apiToken
	err :=db.Model(Data{}).Create(&data).Error
	if err != nil {
		 return data
	}
return data
}

func GetTokenByItemId(itemId int)Data{
	var data Data
	err := db.Where("item_id = ?",itemId).Where("deleted_at is NULL").Find(&data).Error
	if err == gorm.ErrRecordNotFound{
		CreateToken(itemId)
		db.Where("item_id = ?",itemId).Find(&data)
	}
	return data
}

func GetTokenByKey(apikey string)Data{
	var data Data
	db.Where("apitoken = ?",apikey).Find(&data)
	return data

}
func SetLastTime(itemId int)int{
	err := db.Where("item_id =?",itemId).Update("updated_at=?",time.Now()).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESE
}

func Check (apikey string , apitoken string  )int{
	var data Data
	var itemId int
	data = GetTokenByKey(apikey)
	if data.ApiToken == apitoken{
		itemId = data.Item_id
		SetLastTime(itemId)
		return itemId
	}
	return 0
}