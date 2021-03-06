package v1

import (
	"domino-api-gin/controllers"
	"domino-api-gin/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AccountRouter struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func (cls *AccountRouter) Load() {
	r := cls.Router.Group("/v1", middlewares.JwtCheck(cls.DB))
	{
		// 用户列表
		r.GET("/account", func(ctx *gin.Context) {
			(&controllers.AccountController{CTX: ctx, DB: cls.DB}).Index()
		})

		// 根据id获取用户详情
		r.GET("/account/:id", func(ctx *gin.Context) {
			(&controllers.AccountController{CTX: ctx, DB: cls.DB}).Show()
		})

		// 绑定角色到用户
		r.POST("/account/:id/bindRoles", func(ctx *gin.Context) {
			(&controllers.AccountController{CTX: ctx, DB: cls.DB}).PostBindRoles()
		})
	}
}
