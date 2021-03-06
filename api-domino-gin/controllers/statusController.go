package controllers

import (
	"domino-api-gin/errors"
	"domino-api-gin/models"
	"domino-api-gin/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

// StatusController 状态控制器
type StatusController struct {
	CTX *gin.Context
	DB  *gorm.DB
}

// Index 列表
func (cls *StatusController) Index() {
	statusModel := &models.StatusModel{CTX: cls.CTX, DB: cls.DB}
	statuses := statusModel.FindManyByQuery()
	cls.CTX.JSON(tools.CorrectIns("").OK(gin.H{"statuses": statuses}))
}

// Show 详情
func (cls *StatusController) Show() {
	statusModel := &models.StatusModel{CTX: cls.CTX, DB: cls.DB}

	id := tools.StringToInt(cls.CTX.Param("id"))

	status := statusModel.FindOneByID(id)
	cls.CTX.JSON(tools.CorrectIns("").OK(gin.H{"status": status}))
}

// Store 新建
func (cls *StatusController) Store() {
	status := (&models.StatusModel{CTX: cls.CTX, DB: cls.DB}).Store()
	cls.CTX.JSON(tools.CorrectIns("").Created(gin.H{"status": status}))
}

// Update 更新
func (cls *StatusController) Update() {
	statusModel := &models.StatusModel{CTX: cls.CTX, DB: cls.DB}

	id := tools.StringToInt(cls.CTX.Param("id"))

	status := statusModel.UpdateById(id)
	cls.CTX.JSON(tools.CorrectIns("").Updated(gin.H{"status": status}))
}

// Destroy 删除
func (cls *StatusController) Destroy() {
	statusModel := &models.StatusModel{CTX: cls.CTX, DB: cls.DB}

	id, err := strconv.Atoi(cls.CTX.Param("id"))
	if err != nil {
		panic(errors.ThrowForbidden("id必须是数字"))
	}

	statusModel.DeleteByID(id)
	cls.CTX.JSON(tools.CorrectIns("").Deleted())
}
