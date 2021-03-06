package models

import (
	"domino-api-gin/errors"
	"domino-api-gin/tools"
	"gorm.io/gorm/clause"
	"reflect"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Status struct {
	gorm.Model
	UniqueCode string    `gorm:"type:VARCHAR(64);unique;NOT NULL;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;COMMENT:'状态代码';"`
	Name       string    `form:"name" binding:"required" gorm:"type:VARCHAR(64);unique;NOT NULL;COMMENT:'状态名称';"`
	Accounts   []Account `gorm:"foreignKey:StatusUniqueCode;references:UniqueCode;"`
}

type StatusModel struct {
	CTX *gin.Context
	DB  *gorm.DB
}

// Store 新建
func (cls *StatusModel) Store() Status {
	var status Status
	if err := cls.CTX.ShouldBind(&status); err != nil {
		panic(err)
	}

	var repeatStatus Status
	cls.DB.Where(map[string]interface{}{"name": status.Name}).First(&repeatStatus)
	tools.ThrowErrorWhenIsRepeat(repeatStatus, Status{}, "状态名称")

	cls.DB.Omit(clause.Associations).Create(status)
	return status
}

// DeleteByID 根据ID删除
func (cls *StatusModel) DeleteByID(id int) *StatusModel {
	cls.DB.Delete(&Status{}, id)

	return cls
}

// UpdateById 根据ID编辑
func (cls *StatusModel) UpdateById(id int) Status {
	status := cls.FindOneByID(id)

	var statusForm Status
	if err := cls.CTX.ShouldBind(&statusForm); err != nil {
		panic(err)
	}

	var repeatStatus Status
	cls.DB.Where(map[string]interface{}{"name": statusForm.Name}).Not(map[string]interface{}{"id": id}).First(&repeatStatus)
	if !reflect.DeepEqual(repeatStatus, Status{}) {
		panic(errors.ThrowForbidden("状态名称重复"))
	}

	status.Name = statusForm.Name
	cls.DB.Omit(clause.Associations).Save(&status)

	return status
}

// FindOneByID 根据ID读取一条
func (cls *StatusModel) FindOneByID(id int, ) (status Status) {
	cls.DB.Preload("Accounts").Where(map[string]interface{}{"id": id}).First(&status)
	tools.ThrowErrorWhenIsEmpty(status, Status{}, "状态")
	return
}

// FindManyByQuery 根据Query读取多条
func (cls *StatusModel) FindManyByQuery() (statuses []Status) {
	w := make(map[string]interface{})
	n := make(map[string]interface{})

	query := (&tools.QueryBuilder{CTX: cls.CTX, DB: cls.DB}).Init(w, n)
	if name := cls.CTX.Query("name"); name != "" {
		query.Where("`name` like '%?%'", name)
	}
	query.Find(&statuses)

	if len(statuses) == 0{
		panic(errors.ThrowEmpty("状态不存在"))
	}

	return
}
