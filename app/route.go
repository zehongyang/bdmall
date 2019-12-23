package app

import (
	admin "bdmall/app/admin/controller"
	"bdmall/app/bis/controller"
	"bdmall/app/bis/middleware"
	home "bdmall/app/index/controller"
	"bdmall/utils"
	"github.com/gin-gonic/gin"
	"html/template"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"TimeFormat":utils.TimeFormat,
	})
	router.Static("/public","./public/static")
	router.LoadHTMLGlob("public/view/**/**/*")
	router.GET("/",home.Index)
	adminRouter := router.Group("/admin")
	{
		adminRouter.GET("/",admin.Index)
		adminRouter.GET("/welcome",admin.Welcome)
		adminRouter.GET("/category/index",admin.CategoryIndex)
		adminRouter.GET("/category/create",admin.CategoryCreate)
		adminRouter.POST("/category/store",admin.CategoryStore)
		adminRouter.GET("/category/edit",admin.CategoryEdit)
		adminRouter.POST("/category/update",admin.CategoryUpdate)
		adminRouter.POST("/cities",admin.GetCities)
		adminRouter.POST("/categories",admin.GetCategories)
		adminRouter.POST("/upload",admin.UploadImage)
	}
	bisRouter := router.Group("/bis")
	bisRouter.Use(middleware.LoginCheck())
	{
		bisRouter.GET("/login",controller.BisLoginIndex)
		bisRouter.GET("/register",controller.BisRegister)
		bisRouter.POST("/add",controller.BisAdd)
		bisRouter.GET("/index",controller.BisIndex)
		bisRouter.GET("/logout",controller.BisLogout)
		bisRouter.POST("/login",controller.BisLogin)
	}
	userRouter := router.Group("/user")
	{
		userRouter.GET("/captcha",home.GeneratorCaptcha)
		userRouter.GET("/register",home.UserRegisterIndex)
		userRouter.POST("/register",home.UserRegister)
		userRouter.GET("/login",home.UserLoginIndex)
		userRouter.POST("/login",home.UserLogin)
		userRouter.GET("/logout",home.UserLogOut)
	}
	return router
}
