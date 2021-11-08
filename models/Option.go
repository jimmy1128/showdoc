package models

import (
	"awesomeProject3/utils/errmsg"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"gopkg.in/ldap.v2"
	"strconv"
	"strings"
	"time"
)

type Options struct {
	OptionId    uint   `gorm:"primaryKey;type int autoIncrement:false" json:"optionId"`
	OptionName  string `gorm:"type:varchar(255)" json:"optionName"`
	OptionValue string `gorm:"type:varchar(255)" json:"optionValue"`
}
type Ldap struct {
	Url      string
	Port     int
	Username string
	Password string
	BaseDN   string
	filter   string
}

//url        = "b8f40ada2504.sn.mynetname.net"
//port       = 389
//userName   = "cn=administrator,cn=Users,dc=greypanel,dc=com"
//password   = "Grey@support!666"
//baseDN     = "ou=department,dc=greypanel,dc=com"
//filter = "(sAMAccountName=%(user)s)"

func SaveConfig(uid uint, id int) int {
	if checkAdmin(uid) != errmsg.SUCCESE {
		return errmsg.ERROR
	}
	err = db.Model(Options{}).Where("option_name = ?", "register_open").Update("option_value", id).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESE
}
func RegisterSetting() string {
	var option Options
	var ldap Options
	err = db.Model(Options{}).Where("option_name = ? ", "register_open").Find(&option).Error
	if err == gorm.ErrRecordNotFound {
		option.OptionId = 1
		option.OptionName = "register_open"
		option.OptionValue = "0"
		err = db.Create(&option).Error
		if err != nil {
			return ""
		}
	}
	err = db.Model(Options{}).Where("option_name = ? ", "lang").Find(&option).Error
	if err == gorm.ErrRecordNotFound {
		option.OptionId = 2
		option.OptionName = "lang"
		option.OptionValue = "0"
		err = db.Create(&option).Error
		if err != nil {
			return ""
		}
	}
	err = db.Model(Options{}).Where("option_name = ? ", "ldap_open").Find(&ldap).Error
	if err == gorm.ErrRecordNotFound {
		ldap.OptionId = 3
		ldap.OptionName = "ldap_open"
		ldap.OptionValue = "1"
		err = db.Create(&ldap).Error
		if err != nil {
			return ""
		}
	}
	err = db.Model(Options{}).Where("option_name = ? ", "ldap_setting").Find(&ldap).Error
	if err == gorm.ErrRecordNotFound {
		ldap.OptionId = 4
		ldap.OptionName = "ldap_setting"
		ldap.OptionValue = ""
		err = db.Create(&ldap).Error
		if err != nil {
			return ""
		}
	}
	return option.OptionValue
}

func LoadConfig() ([]Options, int) {
	var option []Options
	var result []Options
	var ldapload Options
	var ldaps Ldap
	//var maps = make(map[string]string)
	db.Model(Options{}).Where("option_name =?", "ldap_setting").Find(&ldapload)
	json.Unmarshal([]byte(ldapload.OptionValue), &ldaps)

	err = db.Model(Options{}).Find(&option).Error
	for _, options := range option {
		if options.OptionValue == "1" {
			options.OptionValue = "false"
		} else if options.OptionValue == "0" {
			options.OptionValue = "true"
		}
		result = append(result, options)
	}
	if err != nil {
		return result, errmsg.ERROR
	}
	return result, errmsg.SUCCESE
}

func SaveLangConfig(uid uint, id int) int {
	if checkAdmin(uid) != errmsg.SUCCESE {
		return errmsg.ERROR
	}
	err = db.Model(Options{}).Where("option_name = ?", "lang").Update("option_value", id).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESE
}
func LoadLangConfig(itemId int) (Lang, int) {
	var option Options
	var lang Lang
	var item Item
	err = db.Model(Options{}).Where("option_name = ?", "lang").Find(&option).Error
	db.Model(Item{}).Where("id = ?", itemId).Find(&item)
	s := strings.Split(item.LangList, ",")
	for _, i2 := range s {
		if i2 == option.OptionValue {
			db.Model(Lang{}).Where("id =?", option.OptionValue).Find(&lang)
			return lang, errmsg.SUCCESE
		}
	}
	db.Model(Lang{}).Where("id =?", s[0]).Find(&lang)
	if itemId == 0 {
		db.Model(Lang{}).Where("id =?", option.OptionValue).Find(&lang)
	}

	if err != nil {
		return lang, errmsg.SUCCESE
	}
	return lang, errmsg.SUCCESE
}

func CheckLdapLogin(loginUser string, password string) (User, int) {
	var user User
	var ldapSetting Ldap
	var option Options
	var value Options

	db.Model(Options{}).Where("option_name =?", "ldap_open").Find(&option)
	if option.OptionValue == "1" {
		return user, errmsg.SUCCESE
	}
	var maps = make(map[string]string)

	db.Model(Options{}).Where("option_name =?", "ldap_setting").Find(&value)
	json.Unmarshal([]byte(value.OptionValue), &maps)
	ldapSetting.Url = maps["host"]
	ldapSetting.BaseDN = maps["base_dn"]
	ldapSetting.Username = maps["bind_dn"]
	ldapSetting.Password = maps["bind_password"]
	ldapSetting.filter = maps["user_field"]
	t, _ := strconv.Atoi(maps["port"])
	ldapSetting.Port = t
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldapSetting.Url, ldapSetting.Port))
	if err != nil {
		return user, errmsg.ERROR_CONNECTING_SERVER
	}
	l.SetTimeout(5 * time.Second)
	defer l.Close()
	//err = l.StartTLS(&tls.Config{InsecureSkipVerify: false})
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	err = l.Bind(ldapSetting.Username, ldapSetting.Password)
	if err != nil {
		return user, errmsg.ERROR_USER_NO_RIGHT
	}
	srsql := ldap.NewSearchRequest(ldapSetting.BaseDN, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false, "(cn=*)", []string{ldapSetting.filter, "cn", "mail"}, nil)

	cur, err := l.Search(srsql)
	if err != nil {
		return user, errmsg.ERROR_ARG
	}
	if len(cur.Entries) > 0 {
		for _, entry := range cur.Entries {
			username := entry.GetAttributeValue(ldapSetting.filter)
			name := entry.GetAttributeValue("cn")
			email := entry.GetAttributeValue("mail")

			if loginUser == username {
				err = l.Bind(entry.DN, password)
				if err != nil {
					continue
				}
				if err == nil {
					err = db.Model(User{}).Where("username =?", loginUser).Find(&user).Error

					if err == gorm.ErrRecordNotFound {
						user.Name = name
						user.Username = username
						user.Email = email
						err = db.Model(User{}).Create(&user).Error
						if err != nil {
							return user, errmsg.ERROR
						}
					}
					return user, errmsg.SUCCESE
				}
			}
		}
	}
	return user, errmsg.ERROR_USER_NOT_EXIST
}

func SaveLDapConfig(id int, ldapForm string) int {
	var ldapSetting Ldap
	var user User
	var maps = make(map[string]string)
	json.Unmarshal([]byte(ldapForm), &maps)
	ldapSetting.Url = maps["host"]
	ldapSetting.BaseDN = maps["base_dn"]
	ldapSetting.Username = maps["bind_dn"]
	ldapSetting.Password = maps["bind_password"]
	ldapSetting.filter = maps["user_field"]
	t, _ := strconv.Atoi(maps["port"])
	ldapSetting.Port = t
	if id == 0 {
		l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldapSetting.Url, ldapSetting.Port))
		if err != nil {
			return errmsg.ERROR_CONNECTING_SERVER
		}
		//err = l.StartTLS(&tls.Config{InsecureSkipVerify: true})
		//if err != nil {
		//	return errmsg.ERROR_CONNECTING_SERVER
		//}
		l.SetTimeout(5 * time.Second)
		defer l.Close()
		err = l.Bind(ldapSetting.Username, ldapSetting.Password)
		if err != nil {
			return errmsg.ERROR_USER_NOT_EXIST
		}
		srsql := ldap.NewSearchRequest(ldapSetting.BaseDN, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false, "(cn=*)", []string{ldapSetting.filter, "cn", "mail"}, nil)
		cur, err := l.Search(srsql)
		if err != nil {
			return errmsg.ERROR_ARG
		}
		if len(cur.Entries) > 0 {
			for _, entry := range cur.Entries {
				username := entry.GetAttributeValue(ldapSetting.filter)
				name := entry.GetAttributeValue("cn")
				email := entry.GetAttributeValue("mail")

				err = db.Model(User{}).Where("username =?", username).Find(&user).Error

				if err == gorm.ErrRecordNotFound {
					var adduser User
					adduser.Name = name
					adduser.Username = username
					adduser.Email = email
					err = db.Model(User{}).Create(&adduser).Error
					if err != nil {
						return errmsg.ERROR
					}

				}
			}
			err = db.Model(Options{}).Where("option_name = ?", "ldap_setting").Update("option_value", ldapForm).Error
		}

	}
	db.Model(Options{}).Where("option_name =?", "ldap_open").Update("option_value", id)
	return errmsg.SUCCESE
}
