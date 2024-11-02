package authController

import (
	"my-api/utils"
	"my-api/model"
	"my-api/constant"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type VerifyInput struct {
	Reference string `json:"reference"`
	Code      int    `json:"code"`	
}

func (h *Handler) Verify(c *fiber.Ctx) error {
	var customerVerifyId = c.Params("customerVerifyId")
	var body VerifyInput
	if errInput := c.BodyParser(&body); errInput != nil {
		return utils.ResponseError(c, errInput.Error())
	}
	
	errCheckVerify := h.Db.Where("id = ? AND reference = ? AND code = ?", customerVerifyId, body.Reference, body.Code).First(&model.CustomerVerify{}).Error
	if errCheckVerify == gorm.ErrRecordNotFound {
		return utils.ResponseError(c, "CODE_NOT_MATCH")
	}

	if errUpdateVerify := h.Db.Model(&model.CustomerVerify{}).Where("id = ?", customerVerifyId).Updates(&model.CustomerVerify{
		Status: constant.STATUS_VERIFY_SUCCESS,
	}).Error; errUpdateVerify != nil {
		return utils.ResponseError(c, "ERROR_UPDATE_CUSTOMER_VERIFY")
	}

	return utils.ResponseSuccess(c, "VERIFY_SUCCESS", nil)
}