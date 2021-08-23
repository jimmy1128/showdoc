package models
import(
	"fmt"
	"awesomeProject3/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var db *gorm.DB
var err error

func IniDb(){
	db,err = gorm.Open(utils.Db,fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",utils.DbUser,utils.DbPassWord,utils.DbHost,utils.DbPort,utils.DbName))
	if err != nil{
		fmt.Println("链接数据库失败")
		fmt.Println(err)
	}
	db.SingularTable(true)
	db.AutoMigrate(&User{},&Catalogs{},&Page{},&Item{},&Template{},&Lang{},&Team{},ItemAvatar{},&PageLock{},&ItemSort{},&TeamMember{},&PageHistory{},&TeamItem{},ItemHeader{},&TeamItemMember{},&Result{},&Data{},&ItemMember{},&Options{},&Comment{},&Guest{})
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(10 *time.Second)


}