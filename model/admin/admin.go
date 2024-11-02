package adminModel

import (
	"gorm.io/gorm"
	"my-api/model"
)

type Admin struct {
	model.Base
	Username 		string 			`json:"username"`
	Password 		string 			`json:"password"`
	PasswordView 	string 			`json:"password_view"`
	Email   		string 			`json:"email"`
	Phone   		string 			`json:"phone"`
	AdminRole 		string 			`json:"admin_role"`
	DeletedAt 		gorm.DeletedAt 	`json:"deleted_at"`
}
