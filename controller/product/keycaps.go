package productController

import (
	productModel "my-api/model/product"
	"github.com/gofiber/fiber/v2"
	"my-api/utils"
)

func (h *Handler) GetAllKeycaps(c *fiber.Ctx) error {
	var keycaps []productModel.Keycaps

	if err := h.Db.Find(&keycaps).Error; err != nil {
		return utils.ResponseError(c, err.Error())
	}

	return utils.ResponseSuccess(
		c,
		"GET_ALL_KEYCAPS_SUCCESS",
		keycaps,
	)
}

func (h *Handler) GetKeycapsById(c *fiber.Ctx) error {
	var id = c.Params("id")
	var keycaps productModel.Keycaps

	if errParse := c.BodyParser(&keycaps); errParse != nil {
		return utils.ResponseError(
			c,
			errParse.Error(),
		)
	}

	if err := h.Db.Where("id = ?", id).First(&keycaps).Error; err != nil {
		return utils.ResponseError(c, err.Error())
	}

	return utils.ResponseSuccess(
		c,
		"GET_KEYCAPS_BY_ID_SUCCESS",
		keycaps,
	)
}

