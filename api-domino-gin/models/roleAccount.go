package models

import (
	"domino-api-gin/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RoleAccount struct {
	RoleID    uint
	AccountID uint
}

type RoleAccountModel struct {
	CTX *gin.Context
	DB  *gorm.DB
}

// Store 新建
func (cls *RoleAccountModel) Store(roleID, accountID uint) (roleAccount RoleAccount) {
	cls.DB.Where(map[string]interface{}{"role_id": roleID, "account_id": accountID}).FirstOrCreate(&roleAccount)
	return
}

// BindAccountsToRole 绑定用户到角色
func (cls *RoleAccountModel) BindAccountsToRole(roleID uint, accountIDs []string) (roleAccounts []RoleAccount) {
	for _, accountID := range accountIDs {
		roleAccounts = append(roleAccounts,cls.Store(roleID,tools.StringToUint(accountID)))
	}
	return roleAccounts
}

// BindRolesToAccount 绑定角色到用户
func (cls *RoleAccountModel) BindRolesToAccount(accountID uint, roleIDs []string) (roleAccounts []RoleAccount) {
	for _, roleID := range roleIDs {
		roleAccounts = append(roleAccounts, cls.Store(tools.StringToUint(roleID),accountID))
	}

	return roleAccounts
}
