package models

import (
	"domino-api-gin/tools"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Role struct {
	gorm.Model
	Name     string     `gorm:"type:VARCHAR(64);unique;NOT NULL;COMMENT:'角色名称';"`
	Accounts []*Account `gorm:"many2many:role_accounts;"`
	Menus    []*Menu    `gorm:"many2many:role_menus;"`
}

type RoleModel struct {
	CTX *gin.Context
	DB  *gorm.DB
}

// Store 新建
func (cls *RoleModel) Store() Role {
	var role Role
	if err := cls.CTX.ShouldBind(&role); err != nil {
		panic(err)
	}

	var repeatRole Role
	cls.DB.Where(map[string]interface{}{"name": role.Name}).First(&repeatRole)
	tools.ThrowErrorWhenIsRepeat(repeatRole, Role{}, "角色名称")

	cls.DB.Omit(clause.Associations).Create(&role)
	return role
}

// DeleteOneByID 根据id删除
func (cls *RoleModel) DeleteOneByID(id int) *RoleModel {
	role := cls.FindOneByID(id)
	cls.DB.Omit(clause.Associations).Delete(&role)

	return cls
}

// UpdateOneByID 根据id编辑
func (cls *RoleModel) UpdateOneByID(id int) Role {
	role := cls.FindOneByID(id)

	var roleForm Role
	if err := cls.CTX.ShouldBind(&roleForm); err != nil {
		panic(err)
	}

	var repeatRole Role
	cls.DB.Where(map[string]interface{}{"name": roleForm.Name}).Not(map[string]interface{}{"id": id}).First(&repeatRole)
	tools.ThrowErrorWhenIsRepeat(repeatRole, Role{}, "角色名称")

	role.Name = roleForm.Name
	cls.DB.Omit(clause.Associations).Save(&role)

	return role
}

// FindOneByID 根据编号查询
func (cls *RoleModel) FindOneByID(id int) Role {
	var role Role
	cls.DB.Preload("Accounts").Preload("Account.status").Where(map[string]interface{}{"id": id}).First(&role)

	tools.ThrowErrorWhenIsEmpty(role, Role{}, "角色")

	return role
}

// FindManyByQuery 根据query参数获取列表
func (cls *RoleModel) FindManyByQuery() []Role {
	var roles []Role
	w := make(map[string]interface{})
	n := make(map[string]interface{})

	if name := cls.CTX.Query("name"); name != "" {
		w["name"] = name
	}

	query := (&tools.QueryBuilder{CTX: cls.CTX, DB: cls.DB}).Init(w, n)
	if name := cls.CTX.Query("name"); name != "" {
		query.Where("`name` like '%?%'", name)
	}
	query.Preload("Accounts").Preload("Accounts.status").Find(&roles)

	return roles
}
