package models

func auto(filename string){
	//filenameall := path.Base(filename)
	//filesuffix := path.Ext(filename)
	//fileprefix := filenameall[0:len(filenameall) - len(filesuffix)]


}

//func ImportItem (jsonData string,uid uint , itemName string , itemDescription string , itemPassowrd string , itemDomain string)int{
//	var itemInfo ItemInfo
//	var result map[string]string
//	db.Model(User{}).Where("id = ?",uid).Find(&itemInfo)
//	err = json.Unmarshal([]byte(jsonData),&result )
//	if err != nil {
//		return errmsg.ERROR
//	}
//	if result != nil {
//		if result["item_domain"] != "" {
//			db.Model(Item{}).Where("item_domain in ?",result["item_domain"]).Find(&itemInfo)
//			if &itemInfo != nil{
//				return false
//			}
//			if
//		}
//	}
//
//
//
//}