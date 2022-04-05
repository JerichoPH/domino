package models

import (
	"domino-api-gin/errors"
	"domino-api-gin/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Permission struct {
	gorm.Model
	Name   string `form:"name" binding:"required" gorm:"type:VARCHAR(64);NOT NULL;UNIQUE;COMMENT:'权限名称';"`
	Url    string `form:"url" binding:"required" gorm:"type:TEXT;NOT NULL;COMMENT:'路由';"`
	RoleID uint   `form:"role_id" gorm:"type:UINT;COMMENT:'所属角色';"`
	Role   Role   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type PermissionModel struct {
	CTX *gin.Context
	DB  *gorm.DB
}

// Store 新建
func (cls *PermissionModel) Store(permissionForm Permission) (permission Permission) {
	var repeatPermission Permission
	cls.DB.Omit(clause.Associations).Where(map[string]interface{}{"role_id": permissionForm.RoleID, "name": permissionForm.Name}).First(&repeatPermission)
	tools.ThrowErrorWhenIsRepeat(repeatPermission, Permission{}, "权限")

	permission.Name = permissionForm.Name
	permission.Url = permissionForm.Url
	permission.RoleID = 0

	cls.DB.Omit(clause.Associations).Create(&permission)
	return
}

// DeleteByID 根据ID删除
func (cls *PermissionModel) DeleteByID(id uint) {
	cls.DB.Delete(Permission{}, id)
}

// UpdateByID 根据ID更新
func (cls *PermissionModel) UpdateByID(id uint) (permission Permission) {
	var permissionForm Permission
	if err := cls.CTX.ShouldBind(&permissionForm); err != nil {
		panic(err)
	}

	cls.DB.Where(map[string]interface{}{"id": id}).First(&permission)
	tools.ThrowErrorWhenIsEmpty(permission, Permission{}, "权限")

	permission.Name = permissionForm.Name
	permission.Url = permissionForm.Url
	permission.RoleID = permissionForm.RoleID
	cls.DB.Where(map[string]interface{}{
		"role_id": permissionForm.RoleID,
		"name":    permissionForm.Name,
	}).
		Not(map[string]interface{}{"id": id}).
		FirstOrCreate(&permission)

	return
}

// FindOneByID 根据ID读取一条
func (cls *PermissionModel) FindOneByID(id uint) (permission Permission) {
	cls.DB.Preload("Role").Where(map[string]interface{}{"id": id}).First(&permission)
	tools.ThrowErrorWhenIsEmpty(permission, Permission{}, "权限")
	return
}

// FindManyByRoleID 根据角色ID读取多条
func (cls *PermissionModel) FindManyByRoleID(roleID uint) (permissions []Permission) {
	cls.DB.Preload("Role").Where(map[string]interface{}{"role_id": roleID}).Find(&permissions)
	return
}

// FindManyByRoleIDs 根据角色ID组读取多条
func (cls *PermissionModel) FindManyByRoleIDs(roleIDs []string) (permissions []Permission) {
	cls.DB.Preload("Role").Where(map[string]interface{}{"role_id": roleIDs}).Find(&permissions)
	return
}

// FindManyByQuery 根据Query读取多条
func (cls *PermissionModel) FindManyByQuery() (permissions []Permission) {
	w := make(map[string]interface{})
	n := make(map[string]interface{})

	if names, ok := cls.CTX.GetQueryArray("names"); ok {
		w["name"] = names
	}
	if roleID := cls.CTX.Query("role_id"); roleID != "" {
		w["role_id"] = roleID
	}
	if roleIDs, ok := cls.CTX.GetQueryArray("role_ids"); ok {
		w["role_id"] = roleIDs
	}
	if url := cls.CTX.Query("url"); url != "" {
		w["url"] = url
	}
	if urls, ok := cls.CTX.GetQueryArray("urls"); !ok {
		w["url"] = urls
	}

	query := (&tools.QueryBuilder{CTX: cls.CTX, DB: cls.DB}).Init(w, n)
	if name := cls.CTX.Query("name"); name != "" {
		query.Where("`name` LIKE '%?%'", name)
	}
	query.Preload("Role").Find(&permissions)

	if len(permissions) == 0 {
		panic(errors.ThrowEmpty("权限不存在"))
	}

	return
}
