package routers

import (
	"account/controller"
	_ "account/docs" // 千万不要忘了导入把你上一步生成的docs
	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var R *gin.Engine

func NewAPI() *gin.Engine {
	gin.ForceConsoleColor()
	R = gin.Default()
	R.Static("/static", "E:/accountProject/account/img")
	// user操作
	R.POST("/xqy/user/register", controller.RegisterHandler)
	R.POST("/xqy/user", controller.AuthHandler)
	user := R.Group("/user")
	{
		user.POST("/auth", controller.AuthHandler) //验证token + 登录
		user.POST("/register", controller.RegisterHandler)
	}

	// 连接关系
	link := R.Group("/link", controller.JWTAuthMiddleware())
	{
		link.POST("/:id1/:id2", controller.SetLinkHandler)
		link.GET("/:id", controller.GetLinkHandler)
		link.DELETE("/:id1/:id2", controller.DeleteLinkHandler)
	}

	// images
	img := R.Group("/img", controller.JWTAuthMiddleware())
	{
		img.GET("", controller.GetImageHandler)
		img.POST("", controller.PostImageHandler)
		img.DELETE("/:name", controller.DeleteImageHandler)
	}

	// account
	account := R.Group("/account", controller.JWTAuthMiddleware())
	{
		account.GET("", controller.GetAccountHandler)
		account.POST("", controller.PostAccountHandler)
		account.DELETE("/:id", controller.DeleteAccountHandler)
	}
	// gin-swagger
	R.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	return R
}
