package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	Name     string  `gorm:"type:VARCHAR(64);unique;NOT NULL;COMMENT:'菜单名称';"`
	Sort     uint    `gorm:"type:UINT;NOT NULL;DEFAULT: 0;COMMENT:'排序';"`
	IsShow   bool    `gorm:"type:BOOLEAN;NOT NULL;DEFAULT: 0;COMMENT:'是否显示';"`
	ParentID *uint   `gorm:"type:UINT;COMMENT:'父级ID';"`
	Parents  []Menu  `gorm:"foreignKey:ParentID"`
	Roles    []*Role `gorm:"many2many:role_menus;"`
}

type MenuModel struct {
	CTX *gin.Context
	DB  *gorm.DB
}
