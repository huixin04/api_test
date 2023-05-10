package purchase

import (
	"net/http"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	preset "eirc.app/internal/v1/presenter"
	"eirc.app/internal/v1/structure/purchases"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Created
// @Summary 新增公司行號
// @description 新增公司行號
// @Tags purchases
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /v1.0/authority/purchases [post]
func (p *presenter) Created(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	create_by := util.GenerateUUID()
	input := &purchases.Created{}
	input.Createdby = create_by
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.PurchaseResolver.Created(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// List
// @Summary 條件搜尋公司行號
// @description 條件公司行號
// @Tags purchases
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=purchases.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /v1.0/authority/purchases [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &purchases.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.PurchaseResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary 取得單一公司行號
// @description 取得單一公司行號
// @Tags purchases
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param purchaseID path string true "公司編號"
// @success 200 object code.SuccessfulMessage{body=purchases.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /v1.0/authority/purchases/{purchaseID} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	purchaseID := ctx.Param("purchaseID")
	input := &purchases.Field{}
	input.PurchaseID = purchaseID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.PurchaseResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary 刪除公司
// @description 刪除公司
// @Tags purchases
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param purchaseID path string true "公司編號"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /v1.0/authority/purchases/{purchaseID} [delete]
func (p *presenter) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	updatedBy := util.GenerateUUID()
	purchaseID := ctx.Param("purchaseID")
	input := &purchases.Updated{}
	input.PurchaseID = purchaseID
	input.Updatedby = util.PointerString(updatedBy)
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.PurchaseResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Updated
// @Summary 更新單一使用者
// @description 更新單一使用者
// @Tags purchases
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param purchaseID path string true "公司編號"
// @param * body purchases.Updated true "更新公司"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /v1.0/authority/purchases/{purchaseID} [patch]
func (p *presenter) Updated(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	updatedBy := util.GenerateUUID()
	purchaseID := ctx.Param("purchaseID")
	input := &purchases.Updated{}
	input.PurchaseID = purchaseID
	input.Updatedby = util.PointerString(updatedBy)
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.PurchaseResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
