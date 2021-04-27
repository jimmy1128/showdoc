package main

import (
	"awesomeProject3/models"
	"awesomeProject3/routers"
)

func main(){
	models.IniDb()
	routers.IniRouter()
}
