package v1

import (
	"domino-api-gin/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PermissionRouter struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func (cls *PermissionRouter) Load() {
	r := cls.Router.Group("/v1")
	{
		// 列表
		r.GET("/permission", func(ctx *gin.Context) {
			(&controllers.PermissionController{CTX: ctx, DB: cls.DB}).Index()
		})

		// 详情
		r.GET("/permission/:id", func(ctx *gin.Context) {
			(&controllers.PermissionController{CTX: ctx, DB: cls.DB}).Show()
		})

		// 新建
		r.POST("/permission",func(ctx *gin.Context){
			(&controllers.PermissionController{CTX:ctx,DB:cls.DB}).Store()
		})
	}
}
