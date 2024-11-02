package adminController

import (
	adminModel "my-api/model/admin"
	"my-api/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

type AdminInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) AdminLogin(c *fiber.Ctx) error {
	var input AdminInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var admin adminModel.Admin

	//ฟังก์ชันหา admin ที่มี username และ admin_role เท่ากับ admin
	if errFind := h.Db.Where("username = ? and admin_role = ?", input.Username, "admin").First(&admin).Error; errFind != nil {
		return utils.ResponseError(c, "YOU_ARE_NOT_ADMIN")
	}
	
	if errPassword := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(input.Password)); errPassword != nil {
		return utils.ResponseError(c, "PASSWORD_INCORRECT")
	}

	fmt.Println(admin.Id.String())
	token := GenerateAccessToken(admin.Id.String())
	adminExtract := utils.AdminExtractToken(c)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ADMIN_LOGIN_SUCCESS",
		"token_admin": token, //token ของ admin
		"admin_id": adminExtract.UserId, //id ของ admin
	})
}
