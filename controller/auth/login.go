package authController

import (
	"my-api/model"
	"my-api/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type CustomerLoginInput struct {
	Telephone 	string 	`json:"telephone"`
	Password 	string 	`json:"password"`
}

func (h *Handler) CustomerLogin(c *fiber.Ctx) error {
	var input CustomerLoginInput
	var customer model.Customer
	//เช็ค input ถ้ามี error 
	if err := c.BodyParser(&input); err != nil {
		return utils.ResponseError(c, "ERROR_CUSTOMER_LOGIN")
	}
	//เช็คว่ามีหมายเลขโทรศัพท์นี้หรือไม่
	if errFindCustomer := h.Db.Where("telephone = ?", input.Telephone).First(&customer).Error; errFindCustomer != nil {
		return utils.ResponseError(c, "CUSTOMER_NOT_FOUND")
	}
	
	if errPassword := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(input.Password)); errPassword != nil {
		return utils.ResponseError(c, "PASSWORD_INCORRECT")
	} 
	
								//ดึงข้อมูล customerId แล้วส่งไปให้ GenerateAccessToken เพื่อสร้าง token
	token := GenerateAccessToken(customer.Id.String())
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "CUSTOMER_LOGIN_SUCCESS",
		"result": fiber.Map{
			"customer" : customer,
			"token" : token,
		},
	})
}	

