package authController

import (
	"github.com/gofiber/fiber/v2"
	"my-api/utils"
	"my-api/model"
)

func (h *Handler) FindCustomerById(c *fiber.Ctx) error {
	var id = c.Params("customerId")
	var customer model.Customer
	err := h.Db.Where("id = ?", id).First(&customer).Error
	if err != nil {
		return utils.ResponseError(c, "CUSTOMER_NOT_FOUND")
	}
	return utils.ResponseSuccess(c, "SUCCESS_GET_CUSTOMER", customer)
}
