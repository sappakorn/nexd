package adminController

import (
	"my-api/utils"
	productModel "my-api/model/product"
	"github.com/gofiber/fiber/v2"
)



func (h *Handler) CreateKeyboard(c *fiber.Ctx) error {
	var keyboard productModel.Keyboard
	if errParser := c.BodyParser(&keyboard); errParser != nil{
		return utils.ResponseError(c, "Error Parse Body")
	}

	if errFindKeyboard := h.Db.Where("name = ?", keyboard.Name).First(&productModel.Keyboard{}).Error; errFindKeyboard == nil {
		return utils.ResponseError(c, "KEYBOARD_ALREADY_EXISTS")
	}

	if errCreate := h.Db.Create(&keyboard).Error; errCreate != nil {
		return utils.ResponseError(c, "ERROR_CREATE_KEYBOARD")
	}

	return utils.ResponseSuccess(c, "Create Keyboard Success", keyboard)
}



type UpdateKeyboardInput struct {
	Name 		string 	`json:"name"`
	Price 		float64 	`json:"price"`
	Quantity 	int 	`json:"quantity"`
	Image 		string 	`json:"image"`
	Description string 	`json:"description"`
}

func (h *Handler) UpdateKeyboard(c *fiber.Ctx) error {
	
	id := c.Params("id")
	var input UpdateKeyboardInput
	if id == "" {
		return utils.ResponseError(c, "ID_CANNOT_BE_EMPTY")
	}
	if errParse := c.BodyParser(&input); errParse != nil {
		return utils.ResponseError(c, errParse.Error())
	}
	var keyboard productModel.Keyboard
	keyboard.Name = input.Name
	keyboard.Price = input.Price
	keyboard.Quantity = input.Quantity
	keyboard.Image = input.Image
	keyboard.Description = input.Description
	
	if errUpdate := h.Db.Where("id = ?", id).Updates(&keyboard).Error; errUpdate != nil {
		return utils.ResponseError(c, errUpdate.Error())
	}

	return utils.ResponseSuccess(c, "Update Keyboard Success", keyboard)
}



func (h *Handler) DeleteKeyboard(c *fiber.Ctx) error {
	id := c.Params("id")
	var keyboard productModel.Keyboard


	if errfind := h.Db.Where("id = ?", id).First(&keyboard).Error; errfind != nil {
		return utils.ResponseError(c, "KEYBOARD_ID_NOT_FOUND")
	}

	if errDelete := h.Db.Delete(&keyboard).Error; errDelete != nil {
		return utils.ResponseError(c, errDelete.Error())
	}

	return utils.ResponseSuccess(c, "Delete Keyboard Success", nil)
}
