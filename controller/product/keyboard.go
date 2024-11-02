package productController

import (
	"github.com/gofiber/fiber/v2"
	productModel "my-api/model/product"
	"my-api/utils"
)

func (h *Handler) GetAllKeyboard(c *fiber.Ctx) error {
	var keyboards []productModel.Keyboard

	if err := h.Db.Find(&keyboards).Error; err != nil {
		return utils.ResponseError(c, err.Error())
	}

	return utils.ResponseSuccess(c, "GET_ALL_KEYBOARD_SUCCESS", keyboards)
}

func (h *Handler) GetKeyboardById(c *fiber.Ctx) error {
	var id = c.Params("id")
	var keyboard productModel.Keyboard

	if errParse := c.BodyParser(&keyboard); errParse != nil {
		return utils.ResponseError(
			c,
			errParse.Error())
	}
	//ค้าหาข้อมูลจาก id ที่ส่งมาและเก็บค่าเข้าไปใน ตัวแปร keyboard
	if err := h.Db.Where("id = ?", id).First(&keyboard).Error; err != nil {
		return utils.ResponseError(
			c,
			err.Error(),
		)
	}

	return utils.ResponseSuccess(
		c,
		"GET_KEYBOARD_BY_ID_SUCCESS",
		keyboard,
	)
}
