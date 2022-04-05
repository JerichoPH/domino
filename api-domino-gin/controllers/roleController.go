package controllers

import (
	"domino-api-gin/errors"
	"domino-api-gin/models"
	"domino-api-gin/tools"
	gin "github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RoleController struct {
	CTX *gin.Context
	DB  *gorm.DB
}

// Index 列表
func (cls *RoleController) Index() {
	roleModel := models.RoleModel{CTX: cls.CTX, DB: cls.DB}
	roles := roleModel.FindManyByQuery()
	cls.CTX.JSON(tools.CorrectIns("").Ok(gin.H{"roles": roles}))
}

// Show 详情
func (cls *RoleController) Show() {
	id := tools.StringToInt(cls.CTX.Param("id"))

	roleModel := models.RoleModel{CTX: cls.CTX, DB: cls.DB}
	role := roleModel.FindOneByID(id)

	cls.CTX.JSON(tools.CorrectIns("").Ok(gin.H{"role": role}))
}

// Store 新建
func (cls *RoleController) Store() {
	role := (&models.RoleModel{CTX: cls.CTX, DB: cls.DB}).Store()
	cls.CTX.JSON(tools.CorrectIns("").Created(gin.H{"role": role}))
}

// Update 编辑
func (cls *RoleController) Update() {
	id := tools.StringToInt(cls.CTX.Param("id"))

	role := (&models.RoleModel{CTX: cls.CTX, DB: cls.DB}).UpdateOneByID(id)
	cls.CTX.JSON(tools.CorrectIns("").Updated(gin.H{"role": role}))
}

// Destroy 删除
func (cls *RoleController) Destroy() {
	id := tools.StringToInt(cls.CTX.Param("id"))

	(&models.RoleModel{CTX: cls.CTX, DB: cls.DB}).DeleteOneByID(id)
	cls.CTX.JSON(tools.CorrectIns("").Deleted())
}

// BindAccounts 绑定用户
func (cls *RoleController) PostBindAccounts() {
	roleID := tools.StringToUint(cls.CTX.Param("id"))
	if accountIDs, err := cls.CTX.GetQueryArray("account_ids"); !err {
		panic(errors.ThrowForbidden("获取参数失败"))
	} else {
		ret := (&models.RoleAccountModel{CTX: cls.CTX, DB: cls.DB}).BindAccountsToRole(roleID, accountIDs)
		if len(ret) > 0 {
			cls.CTX.JSON(tools.CorrectIns("绑定成功").Created(nil))
		} else {
			panic(errors.ThrowForbidden("绑定失败"))
		}
	}
}
