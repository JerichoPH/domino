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
	cls.CTX.JSON(tools.CorrectIns("").OK(gin.H{"permissions": (&models.PermissionModel{CTX: cls.CTX, DB: cls.DB}).FindManyByQuery()}))
}

// Show 详情
func (cls *PermissionController) Show() {
	cls.CTX.JSON(tools.CorrectIns("").OK(gin.H{"permission": (&models.PermissionModel{CTX: cls.CTX, DB: cls.DB}).FindOneByID(tools.StringToUint(cls.CTX.Param("id")))}))
}

// Store 新建
func (cls *PermissionController) Store() {
	var permissionForm models.Permission
	if err := cls.CTX.ShouldBind(&permissionForm); err != nil {
		panic(err)
	}

	permission := (&models.PermissionModel{CTX: cls.CTX, DB: cls.DB}).Store(permissionForm)

	cls.CTX.JSON(tools.CorrectIns("").Created(gin.H{"permission": permission}))
}

// Update 编辑
func (cls *PermissionController) Update() {
	var permissionForm models.Permission
	if err := cls.CTX.ShouldBind(permissionForm); err != nil {
		panic(err)
	}

	permission := (&models.PermissionModel{CTX: cls.CTX, DB: cls.DB}).UpdateByID(tools.StringToUint(cls.CTX.Param("id")))
	cls.CTX.JSON(tools.CorrectIns("").Updated(gin.H{"permission": permission}))
}

// Destroy 删除
func (cls *PermissionController) Destroy() {
	(&models.PermissionModel{CTX: cls.CTX, DB: cls.DB}).DeleteByID(tools.StringToUint(cls.CTX.Param("id")))
	cls.CTX.JSON(tools.CorrectIns("").Deleted())
}
