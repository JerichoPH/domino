package controllers

import (
	"domino-api-gin/models"
	"domino-api-gin/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PermissionController struct {
	CTX *gin.Context
	DB  *gorm.DB
}

// Index 权限列表
func (cls *PermissionController) Index() {
	cls.CTX.JSON(tools.CorrectIns().Ok(gin.H{"permissions": (&models.PermissionModel{CTX: cls.CTX, DB: cls.DB}).FindManyByQuery()}, ""))
}

func (cls *PermissionController) Show() {
	cls.CTX.JSON(tools.CorrectIns().Ok(gin.H{"permission": (&models.PermissionModel{CTX: cls.CTX, DB: cls.DB}).FindOneByID(tools.StringToUint(cls.CTX.Param("id")))}, ""))
}

func (cls *PermissionController) Store() {
	var permissionForm models.Permission
	if err := cls.CTX.ShouldBind(models.Permission{}); err != nil {
		panic(err)
	}

	permission := (&models.PermissionModel{CTX: cls.CTX, DB: cls.DB}).Store(permissionForm)
	cls.CTX.JSON(tools.CorrectIns().Created(gin.H{"permission": permission}, ""))
}
