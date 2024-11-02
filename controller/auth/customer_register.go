package authController

import (
	"my-api/constant"
	"my-api/model"
	"my-api/utils"
	"github.com/gofiber/fiber/v2"
	
)



type CustomerInput struct {
	Telephone string `json:"telephone"`
	TypeVerify string `json:"type_verify"`
}


func (h *Handler) CustomerRegister(c *fiber.Ctx) error {
	var input CustomerInput
	if err := c.BodyParser(&input); err != nil {
		return utils.ResponseError(c, err.Error())
	}
	//เช็คว่ามีหมายเลขโทรศัพท์นี้หรือไม่

	if errFindTelePhone := h.Db.Where("telephone = ?", input.Telephone).First(&model.Customer{}).Error; errFindTelePhone == nil {
		return utils.ResponseError(c, "TELEPHONE_ALREADY_EXISTS")
	}
	
	if errUpdateCustomerVerify := h.Db.Where("telephone = ? AND type_verify = ?", input.Telephone, input.TypeVerify).Updates(&model.CustomerVerify{
		Status: constant.STATUS_VERIFY_FAIL,
	}).Error; errUpdateCustomerVerify != nil {
		return utils.ResponseError(c, "ERROR_UPDATE_CUSTOMER_VERIFY")
	}

	var customerVerify model.CustomerVerify
	customerVerify.Telephone = input.Telephone
	customerVerify.Reference = utils.GenerateCodeNumberAndString()
	customerVerify.Code      = utils.GenerateOTP()
	customerVerify.TypeVerify = input.TypeVerify
	customerVerify.TypeSystem = constant.TYPE_SYSTEM_SMS

	if errCreate := h.Db.Create(&customerVerify).Error; errCreate != nil {
		return utils.ResponseError(c, "ERROR_CREATE_CUSTOMER_VERIFY")
	}

	// responseSuccess ต้องการ 2 ค่า คือ message กับ resultส่งแบบMap JSON กลับไป
	return utils.ResponseSuccess(c, "VERIFY_SUCCESS", map[string]interface{}{
		"id": 			customerVerify.Id,
		"telephone": 	customerVerify.Telephone,
		"email": 		customerVerify.Email,
		"reference": 	customerVerify.Reference,
		"code": 		customerVerify.Code,
		"type_verify": 	customerVerify.TypeVerify,
		"type_system": 	customerVerify.TypeSystem,
		"status": 		constant.STATUS_VERIFY_PENDING,//รอผู้ใช้กรอกรหัส
	})

}

type CustomerSubmitInput struct {
	Name string `json:"name"`
	Password string `json:"password"`
	Telephone string `json:"telephone"`
}

func (h *Handler) CustomerSubmit(c *fiber.Ctx) error {
	var input CustomerSubmitInput

	//เช็คerr การส่งข้อมูลมา
	if err := c.BodyParser(&input); err != nil {
		return utils.ResponseError(c, err.Error())
	}
	TelephoneVerify := true 
	var customer model.Customer
	customer.Name = input.Name
	customer.Telephone = input.Telephone

	customer.TelephoneVerify = &TelephoneVerify // ชี้ไปที่ตัวแปรที่สร้างขึ้น
	customer.Password = utils.HashPassword(input.Password)
	customer.PasswordView = input.Password
	customer.StatusRegister = constant.STATUS_COMPLETE
	customer.TypeRegister = constant.TYPE_REGISTER_TELEPHONE

	if errCreate := h.Db.Create(&customer).Error; errCreate != nil {
		return utils.ResponseError(c, "ERROR_CREATE_CUSTOMER")
	}
	token := GenerateAccessToken(customer.Id.String())
	return utils.ResponseSuccess(c,"Register Success",map[string]interface{}{
		"customer": customer,
		"token": token,
	})
}

	
	
