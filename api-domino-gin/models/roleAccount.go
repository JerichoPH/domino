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

func (cls *RoleAccountModel) StoreBatch(roleID uint, accountIDs []string) []RoleAccount {
	roleAccounts := make([]RoleAccount, 0)
	for _, accountID := range accountIDs {
		roleAccount := &RoleAccount{RoleID: roleID, AccountID: tools.StringToUint(accountID)}
		roleAccounts = append(roleAccounts, *roleAccount)
	}
	cls.DB.Create(&roleAccounts)

	return roleAccounts
}
