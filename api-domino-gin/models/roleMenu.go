package models

import (
	"domino-api-gin/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RoleMenu struct {
	RoleID uint
	MenuID uint
}

type RoleMenuModel struct {
	CTX *gin.Context
	DB  *gorm.DB
}

// Store 新建
func (cls *RoleMenuModel) Store(menuID, roleID uint) (roleMenu RoleMenu) {
	cls.DB.Create(map[string]interface{}{"role_id": roleID, "menu_id": menuID}).FirstOrCreate(roleMenu)
	return
}

// BindRoles 绑定角色到菜单
func (cls *RoleMenuModel) BindRolesToMenu(menuID uint, roleIDs []string) (roleMenus []RoleMenu) {
	for _, roleID := range roleIDs {
		roleMenus = append(roleMenus, cls.Store(menuID, tools.StringToUint(roleID)))
	}
	return
}

// BindMenusToRole 绑定菜单到角色
func (cls *RoleMenuModel) BindMenusToRole(roleID uint, menuIDs []string) (roleMenus []RoleMenu) {
	for _, menuID := range menuIDs {
		roleMenus = append(roleMenus, cls.Store(tools.StringToUint(menuID), roleID))
	}
	return
}
