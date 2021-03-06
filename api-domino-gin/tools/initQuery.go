package tools

import (
	"domino-api-gin/errors"
	"gorm.io/gorm/clause"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type QueryBuilder struct {
	CTX *gin.Context
	DB  *gorm.DB
}

// Init 初始化查询器
func (cls *QueryBuilder) Init(w, n map[string]interface{}) *gorm.DB {
	tx := cls.DB.Where(w).Not(n)

	// 排序
	if order := cls.CTX.Query("order"); order != "" {
		tx.Order(order)
	}

	// offset
	if offset := cls.CTX.Query("offset"); offset != "" {
		if offset, err := strconv.Atoi(offset); err != nil {
			panic(errors.ThrowForbidden("offset 参数只能填写数字"))
		} else {
			tx.Offset(offset)
		}
	}

	// limit
	if limit := cls.CTX.Query("limit"); limit != "" {
		if limit, err := strconv.Atoi(limit); err != nil {
			panic(errors.ThrowForbidden("limit 参数只能填写数字"))
		} else {
			tx.Limit(limit)
		}
	}

	return tx.Preload(clause.Associations)
}
