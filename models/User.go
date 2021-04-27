package models

import (
	"awesomeProject3/utils/errmsg"
	"encoding/base64"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
	"log"
	"time"
)

type User struct {
	gorm.Model
	Username  string    `gorm:"type:varchar(255);not null" json:"username" validate:"required,min=10,max=50" label:"用户名"`
	Password  string    `gorm:"type:varchar(30);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role      int       `gorm:"type:int;DEFAULT:2" json:"role"`
	Email     string    `gorm:"type:varchar(255)" json:"email"`
	Name      string    `gorm:"type:varchar(255)" json:"name"`
	LastLogin *time.Time    `gorm:"type:datetime" json:"last_login"`
}

//查询用户是否存在
func CheckUser(name string)(code int){
	if RegisterSetting() == "1"{
		return errmsg.ERROR_APPLY_NEW_ACCOUNT
	}
	if checkAdminUser () != errmsg.SUCCESE{
		return errmsg.ERROR_APPLY_NEW_ACCOUNT
	}
	var users User
	db.Select("id").Where("username =?",name).First(&users)
	if users.ID >= 1{
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESE
}

func CheckUpUser(id int , name string)(code int ){
	var user User
	db.Select("id, username").Where("username =?",name).First(&user)
	if user.ID == uint(id) {
		return errmsg.SUCCESE
	}
	if user.ID >=1 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESE
}
//新增用户
func CreateUser(data *User)int{
	err:= db.Create(&data).Error
	if err != nil{
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESE
}
//查询单个用户
func GetUser(id uint,redirect string)(User,int){
	var user User
	if id == 0 && redirect == "false"{
		return user,errmsg.ERROR_TOKEN_RUNTIME
	}
	err = db.Where("id =?",id).First(&user).Error
	if err != nil{
		return user,errmsg.ERROR_USER_NO_RIGHT
	}
	return user,errmsg.SUCCESE
}

type Result struct {
	Username string `gorm:"type:varchar(255)" json:"value"`
}
//查询用户列表
func GetUsers(username string) ([]Result,int){
	var result []Result

	if username != "" {
		db.Raw("select username from `user` where username LIKE ?",username+"%").Find(&result)
		return result,errmsg.SUCCESE
	}
	db.Raw("select username from `user` where 1=1").Find(&result)
	if err == gorm.ErrRecordNotFound{
		return result,0
	}
	return result,errmsg.SUCCESE
}
//编辑用户信息
func EditUser(id int,data *User)int{
	var maps = make(map[string]interface{})
	maps["username"]=data.Username
	maps["role"]=data.Role
	maps["password"]=ScryptPw(data.Password)
	err =db.Model(&User{}).Where("id =?",id ).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESE
}

//修改密码
func ChangePassword(id int,data *User)int{
	err := db.Model(User{}).Select("password").Where("id =?",id).Update(&data).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESE
}
//删除用户
func DeleteUser(id int) int {
	var user User
	err = db.Where("id =?",id).Delete(&user).Error
	if err!= nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESE
}
//密码加密
func(u *User)BeforeSave(){
	u.Password=ScryptPw(u.Password)
}
func (u *User) BeforeUpdate(_ *gorm.DB)(err error){
	u.Password = ScryptPw(u.Password)
	return nil
}
func ScryptPw(password string)string{
	const KeyLen = 10
	salt :=make([]byte,8)
	salt =[]byte{12,32,4,6,66,22,222,11}
	HashPw,err := scrypt.Key([]byte(password),salt,16384,8,1,KeyLen)
	if err != nil{
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw

}
//登陆验证
func CheckLogin(username string ,password string)(User,int){
	var user User


	db.Where("username =?",username).First(&user)
	if ScryptPw(password) != user.Password{
		return user,errmsg.ERROR_PASSWORD_WRONG
	}
	if user.ID==0{
		return user,errmsg.ERROR_USER_NOT_EXIST
	}
	//if user.Role != 3{
	//	return user,errmsg.ERROR_USER_NO_RIGHT
	//}
	currentTime := time.Now()
	db.Model(User{}).Where("username =?",username).Update("last_login",currentTime.Format("2006-01-02 15:04:05"))

	return user,errmsg.SUCCESE
}
// 前台登入
func CheckLoginFront(username string, password string) (User, int) {
	var user User


	db.Where("username = ?", username).First(&user)


	if user.ID == 0 {
		return user, errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPw(password) != user.Password{
		return user,errmsg.ERROR_PASSWORD_WRONG
	}
	return user, errmsg.SUCCESE
}
func checkAdminUser ()int{
	var user User
	err = db.Model(User{}).Where("username = ?","greypanel@gmail.com").Find(&user).Error
	if err != gorm.ErrRecordNotFound{
		return errmsg.SUCCESE
	}
	user.Username = "greypanel@gmail.com"
	user.Password = "123456789"
	user.Name="admin"
	user.Role = 1

    err := db.Model(User{}).Create(&user).Error
    if err != nil {
    	return errmsg.ERROR
	}
	return errmsg.SUCCESE
}