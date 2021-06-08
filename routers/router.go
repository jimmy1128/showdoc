package routers

import (
	v1 "awesomeProject3/api/v1"
	"awesomeProject3/middleware"
	"awesomeProject3/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	//"github.com/gin-gonic/contrib/sessions"
)

func IniRouter(){
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())
	r.Use(middleware.CORSMiddleware())
	r.Use(v1.EnableCookieSession())
	r.LoadHTMLGlob("static/index.html")
	r.Static("static","static/static")
	r.Static("upload","upload")
	r.GET("/doc",func(c *gin.Context){
		c.HTML(200,"index.html",nil)
	})
	r.NoRoute(v1.NoResponse)

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())

	{
		//用户模块
		auth.POST("admin/user",v1.GetUsers)
		auth.PUT("user/:id",v1.EditUser)
		auth.DELETE("user/:id",v1.DeleteUser)
		auth.POST("admin/resetPassword",v1.ChangeUserPassword)
		//template 模块
		auth.POST("template/add",v1.AddTemplate)
		auth.POST("/admin/template",v1.GetList)
		auth.DELETE("template/:id",v1.DeleteTemplate)
		auth.PUT("template/:id",v1.EditTemplate)
		auth.POST("template/save",v1.SaveTemp)
		//page 模块
		auth.POST("page/add",v1.AddPage)
		auth.PUT("page/:id",v1.EditPage)
		auth.POST("page/delete",v1.DeletePage)
		auth.POST("page/save",v1.SavePage)
		auth.POST("page/setLock",v1.SetLock)
		auth.POST("page/isLock",v1.IsLock)
		auth.POST("page/history",v1.History)
		auth.POST("page/sort",v1.Sort)

		//catalogs 模块
		auth.POST("delcatalogs/:id",v1.DeleteCatalogs)
		auth.PUT("catalogs/:id",v1.EditCat)
		auth.POST("catalogs/add",v1.AddCat)
		auth.POST("catListGroup",v1.CatListGroup)
		auth.POST("secondCatList",v1.SecondCatList)
		auth.POST("childCatList",v1.ChildCatList)
		auth.POST("getDefaultCat",v1.GetDefaultCat)
		auth.POST("catalogs/save",v1.SaveCatalogs)
		auth.POST("catalog/getPagesBycat",v1.GetPagesByCat)
		auth.POST("catalog/batUpdate",v1.BatUpdates)
		//item 模块

		auth.POST("item/delete",v1.DeleteItem)
		auth.POST("item/add",v1.AddItem)
		auth.POST("item/update",v1.UpdateItem)
		auth.POST("item/exitItem",v1.ExitItem)
		auth.POST("item/archive",v1.ArchiveItem)
		auth.POST("item/attorn",v1.AttornItem)
		auth.POST("item/sort",v1.SortByItem)

		//team
		auth.POST("team/save",v1.SaveTeam)
		auth.POST("team/list",v1.GetTeamList)
		auth.POST("team/delete",v1.DeleteTeam)
		auth.POST("team/attorn",v1.AttornTeam)

		//teamMember
		auth.POST("member/save",v1.SaveTeamMember)
		auth.POST("member/list",v1.GetTeamMemberList)
		auth.POST("member/delete",v1.DeleteTeamMember)
		//TeamItemMember
		auth.POST("teamitemmember/save",v1.SaveTeamItemMember)
		auth.POST("teamitemmember/list",v1.GetTeamItemMemberList)
		//TeamItem
		auth.POST("teamitem/save",v1.SaveTeamItem)
		auth.POST("teamitem/list",v1.GetTeamItemList)
		auth.POST("teamitem/listbyteam",v1.GetListByTeam)
		auth.POST("teamitem/delete",v1.DeleteTeamItem)

		// api
		auth.POST("item/getKey",v1.GetKey)
		auth.POST("item/resetKey",v1.GetReset)

		//Member
		auth.POST("itemmember/getList",v1.GetMemberList)
		auth.POST("itemmember/save",v1.SaveMember)
		auth.POST("itemmember/delete",v1.DeleteMember)

		//Recycle
		auth.POST("recycle/getList",v1.GetRecycleList)
		auth.POST("recycle/recover",v1.Recover)

		//admin
		auth.POST("adminUser/deleteUser",v1.AdminDeleteUser)
		auth.POST("adminUser/addUser",v1.AdminAddUser)
		auth.POST("adminUser/getList",v1.AdminGetList)
		auth.POST("adminItem/getList",v1.AdminItemGetList)
		auth.POST("adminItem/attorn",v1.AdminItemAttorn)
		auth.POST("adminItem/deleteItem",v1.AdminDeleteItem)

		// adminConfig
		auth.POST("adminSetting/saveConfig",v1.SaveConfig)
		auth.GET("adminSetting/loadConfig",v1.LoadConfig)
		auth.POST("adminSetting/saveLangConfig",v1.SaveLangConfig)


		// LangConfig
		auth.POST("lang/add",v1.AddLang)
		auth.GET("admin/lang",v1.GetLang)
		auth.POST("lang/edit",v1.EditLang)
		auth.POST("lang/delete",v1.DeleteLang)



	}
	router := r.Group("api/v1")

	{
		//用户信息模块
		router.POST("user/info",v1.GetUserInfo)
		//router.GET("users",v1.GetUsers)
		router.POST("admin/check_token",v1.CheckToken)
		router.POST("login",v1.Login)
		router.GET("exit", v1.ExitGet)
		router.POST("item/info",v1.ItemsInfo)

		// 上传模块
		router.POST("upload",v1.UpLoad)
		router.POST("user/register",v1.AddUser)
		//item 模块
		router.GET("admin/list",v1.MyList)
		router.POST("item/link",v1.LinkItem)
		router.POST("item/detail",v1.ItemDetail)
		router.POST("item/pwd",v1.Pwd)
		//page 模块
		router.POST("admin/page",v1.GetInfo)
		//router.GET("cookie", v1.Cookie)
		router.GET("lang/info",v1.GetLangInfo)
		router.GET("lang",v1.GetLang)
		// config
		router.GET("adminSetting/loadLangConfig",v1.LoadLangConfig)
		//ImportItem
		router.POST("import/auto",v1.ImportItem)

	}
	fmt.Println("Visit : http://localhost"+utils.HttpPort+"/doc")
	_ = r.Run(utils.HttpPort)

}

