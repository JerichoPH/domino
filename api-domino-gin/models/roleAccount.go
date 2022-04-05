package models

import (
	"domino-api-gin/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
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
func (cls *RoleAccountModel) BindAccountsToRole(roleID uint, accountIDs []string) []RoleAccount {
	roleAccounts := make([]RoleAccount, 0)
	for _, accountID := range accountIDs {
		roleAccount := &RoleAccount{RoleID: roleID, AccountID: tools.StringToUint(accountID)}
		roleAccounts = append(roleAccounts, *roleAccount)
	}
	cls.DB.Create(&roleAccounts)

	return roleAccounts
}

// BindRolesToAccount 绑定角色到用户
func (cls *RoleAccountModel) BindRolesToAccount(accountID uint, roleIDs []string) []RoleAccount {
	roleAccounts := make([]RoleAccount, 0)
	for _, roleID := range roleIDs {
		roleAccount := &RoleAccount{AccountID: accountID, RoleID: tools.StringToUint(roleID)}
		roleAccounts = append(roleAccounts, *roleAccount)
	}
	cls.DB.Create(&roleAccounts)

	log.Println(roleAccounts)

	return roleAccounts
}
