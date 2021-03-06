package models

import (
	"domino-api-gin/errors"
	"domino-api-gin/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
	"time"
)

// Account 用户表
type Account struct {
	gorm.Model
	Username         string    `form:"username" gorm:"type:VARCHAR(64);UNIQUE;NOT NULL;COMMENT:'用户名';"`
	Password         string    `form:"password" gorm:"type:VARCHAR(128);NOT NULL;comment:'密码';"`
	Nickname         string    `form:"nickname" gorm:"type:VARCHAR(64);NOT NULL;DEFAULT '';INDEX:users__nickname__idx;COMMENT:'昵称';"`
	Email            string    `form:"email" gorm:"type:VARCHAR(64);NOT NULL;UNIQUE;COMMENT:'邮箱';"`
	ActivatedAt      time.Time `gorm:"type:DATETIME;COMMENT:'激活时间';"`
	StatusUniqueCode string    `gorm:"type:VARCHAR(64);COMMENT:'状态代码';"`
	Status           Status    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:StatusUniqueCode;references:UniqueCode;"`
	Roles            []*Role   `gorm:"many2many:role_accounts;"`
}

// AccountModel 用户模型
type AccountModel struct {
	CTX      *gin.Context
	DB       *gorm.DB
	Preloads map[string][]string
}

// ScopeIsActivated 已经激活的
func (cls *AccountModel) ScopeIsActivated(db *gorm.DB) *gorm.DB {
	return db.Not(map[string]interface{}{"activated_at": nil})
}

// ScopeCanLogin 可以登录的
func (cls *AccountModel) ScopeCanLogin(db *gorm.DB) *gorm.DB {
	return db.Where(map[string]interface{}{"status_unique_code": "NORMAL"})
}

// ScopeIsNotActivated 未激活的
func (cls *AccountModel) ScopeIsNotActivated(db *gorm.DB) *gorm.DB {
	return db.Where(map[string]interface{}{"activated_at": nil})
}

// FindOneById 根据id获取用户
func (cls *AccountModel) FindOneById(id int) Account {
	var account Account
	cls.DB.Preload("status").Preload("Roles").Where(map[string]interface{}{"id": id}).First(&account)
	return account
}

// FindOneByUsername 根据用户名读取一条
func (cls *AccountModel) FindOneByUsername(username string) (account Account) {
	cls.DB.Scopes(cls.ScopeIsActivated, cls.ScopeCanLogin).Preload(clause.Associations).Where(map[string]interface{}{"username": username}).First(&account)
	return
}

// FindOneByEmail 根据邮箱读取一条
func (cls *AccountModel) FindOneByEmail(email string) (account Account) {
	cls.DB.Scopes(cls.ScopeIsActivated, cls.ScopeCanLogin).Preload(clause.Associations).Where(map[string]interface{}{"email": email}).First(&account)
	return
}

// FindManyByQuery 根据Query读取用户列表
func (cls *AccountModel) FindManyByQuery() []Account {
	var accounts []Account

	// 搜索条件
	w := make(map[string]interface{})
	n := make(map[string]interface{})
	if username := cls.CTX.Query("username"); username != "" {
		w["username"] = username
	}
	if nickname := cls.CTX.Query("nickname"); nickname != "" {
		w["nickname"] = nickname
	}
	if email := cls.CTX.Query("email"); email != "" {
		w["email"] = email
	}
	if activatedAt := cls.CTX.Query("activated_at"); activatedAt != "" {
		switch activatedAt {
		case "1":
			n["activated_at"] = nil
		case "0":
			w["activated_at"] = nil
		}
	}

	query := (&tools.QueryBuilder{CTX: cls.CTX, DB: cls.DB}).Init(w, n)
	if activatedAtBetween := cls.CTX.Query("activated_at_between"); activatedAtBetween != "" {
		query.Where("activated_at BETWEEN ? AND ?", strings.Split(activatedAtBetween, "~"))
	}
	query.Find(&accounts)

	if len(accounts) == 0 {
		panic(errors.ThrowEmpty("用户不存在"))
	}

	return accounts
}
