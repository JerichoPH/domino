package controllers

import (
	"domino-api-gin/errors"
	"domino-api-gin/models"
	"domino-api-gin/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"reflect"
)

// AccountController 用户控制器
type AccountController struct {
	CTX *gin.Context
	DB  *gorm.DB
}

// Index 列表
func (cls *AccountController) Index() {
	accountModel := &models.AccountModel{CTX: cls.CTX, DB: cls.DB}
	accounts := accountModel.FindManyByQuery()
	cls.CTX.JSON(tools.CorrectIns("").Ok(gin.H{"accounts": accounts}))
}

// Show 详情
func (cls *AccountController) Show() {
	id := tools.StringToInt(cls.CTX.Param("id"))

	accountModel := &models.AccountModel{CTX: cls.CTX, DB: cls.DB}
	account := accountModel.FindOneById(id)
	if reflect.DeepEqual(account, models.Account{}) {
		panic(errors.ThrowEmpty("用户不存在"))
	}

	cls.CTX.JSON(tools.CorrectIns("").Ok(gin.H{"account": account}))
}

// PostBindRoles 绑定角色到用户
func (cls *AccountController) PostBindRoles() {
	accountID := tools.StringToUint(cls.CTX.Param("id"))

	if roleIDs, err := cls.CTX.GetQueryArray("role_ids"); !err {
		panic(err)
	} else {
		ret := (&models.RoleAccountModel{CTX: cls.CTX, DB: cls.DB}).BindRolesToAccount(accountID, roleIDs)
		if len(ret) > 0 {
			cls.CTX.JSON(tools.CorrectIns("绑定成功").Created(nil))
		} else {
			panic(errors.ThrowForbidden("绑定失败"))
		}
	}
}
