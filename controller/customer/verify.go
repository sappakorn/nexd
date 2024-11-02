package customerController

import (
	"my-api/constant"
	"my-api/model"
	"my-api/utils"
	"github.com/gofiber/fiber/v2"
)

type VerifyInput struct {
	Telephone string `json:"telephone"`
	TypeVerify string `json:"type_verify"`
}

func (h *Handler) CustomerVerify(c *fiber.Ctx) error {
	var input VerifyInput
	if errParser := c.BodyParser(&input); errParser != nil {
		return utils.ResponseError(c, errParser.Error())
	}
	//sql = SELECT * FROM customers WHERE telephone = ? 
	if errFindTelePhone := h.Db.Where("telephone = ?", input.Telephone ).First(&model.Customer{}).Error; errFindTelePhone != nil {
		return utils.ResponseError(c, "TELEPHONE_ALREADY_EXISTS")
	}

	if errUpdateCustomerVerify := h.Db.Where("telephone = ? AND type_verify = ?", input.Telephone, input.TypeVerify).Updates(&model.CustomerVerify{
		Status: constant.STATUS_VERIFY_PENDING,
	}).Error; errUpdateCustomerVerify != nil {
		return utils.ResponseError(c, "ERROR_UPDATE_CUSTOMER_VERIFY")
	}
	var customerVerify model.CustomerVerify
	customerVerify.Telephone = input.Telephone
	customerVerify.Reference = utils.GenerateCodeNumberAndString()
	customerVerify.Code = utils.GenerateOTP()
	customerVerify.TypeVerify = input.TypeVerify
	customerVerify.TypeSystem = constant.TYPE_SYSTEM_SMS

	if errCreate := h.Db.Create(&customerVerify).Error; errCreate != nil {
		return utils.ResponseError(c, "ERROR_CREATE_CUSTOMER_VERIFY")
	}

															// แมพ json ออกไป
	return utils.ResponseSuccess(c, "TELEPHONE_AVAILABLE", map[string]interface{}{
		"id" : customerVerify.Id,
		"telephone" : customerVerify.Telephone,
		"email" : customerVerify.Email,
		"reference" : customerVerify.Reference,
		"code" : customerVerify.Code,
		"type_verify" : customerVerify.TypeVerify,
		"type_system" : customerVerify.TypeSystem,
		"status" : customerVerify.Status,
	})

}

