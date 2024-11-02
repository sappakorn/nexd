package model

import "gorm.io/gorm"

type Customer struct {
	Base
	Username          string            `json:"username" gorm:"default: null"`
	Password          string            `json:"password"`
	PasswordView      string            `json:"password_view" gorm:"default: null"`
	ImageProfile      string            `json:"image_profile" gorm:"default: null"`
	ImageCover        string            `json:"image_cover" gorm:"default: null"`
	Name              string            `json:"name" gorm:"default: null"`
	Telephone         string            `json:"telephone"`
	TelephoneVerify   *bool             `json:"telephone_verify" gorm:"default: false"`
	Email             string            `json:"email" gorm:"default: null"`
	EmailVerify       *bool             `json:"email_verify" gorm:"default: false"`
	TypeRegister      string            `json:"type_register" gorm:"default: 1"` // 1:telephone // 2:facebook login // 3:line login // 4:gmail login
	Gender            string            `json:"gender" gorm:"default: MALE"`
	BirthDate         string            `json:"birth_date" gorm:"default: null"`
	Age               string            `json:"age" gorm:"default: null"`
	FacebookId        string            `json:"facebook_id" gorm:"default: null"`
	GmailId           string            `json:"gmail_id" gorm:"default: null"`
	LineId            string            `json:"line_id" gorm:"default: null"`
	Status            *bool             `json:"status" gorm:"default: true"`
	StatusRegister    string            `json:"status_register" gorm:"default: TELEPHONE"` // TELEPHONE // VERIFY //INFORMATION // COMPLETE
	DeletedAt         gorm.DeletedAt    `json:"deleted_at"`
	CustomerAddresses []CustomerAddress `json:"customer_addresses"`
}

type CustomerInput struct {
	Username          string            `json:"username" gorm:"default: null"`
	Password          string            `json:"password"`
	PasswordView      string            `json:"password_view" gorm:"default: null"`
	ImageProfile      string            `json:"image_profile" gorm:"default: null"`
	Name              string            `json:"name" gorm:"default: null"`
	Telephone         string            `json:"telephone" gorm:"default: null"`
	TelephoneVerify   *bool             `json:"telephone_verify" gorm:"default: false"`
	Email             string            `json:"email" gorm:"default: null"`
	EmailVerify       *bool             `json:"email_verify" gorm:"default: false"`
	TypeRegister      int8              `json:"type_register" gorm:"default: 1"` // 1:telephone // 2:facebook login // 3:line login // 4:gmail login
	Gender            string            `json:"gender" gorm:"default: male"`
	BirthDate         string            `json:"birth_date" gorm:"default: null"`
	FacebookId        string            `json:"facebook_id" gorm:"default: null"`
	GoogleId          string            `json:"google_id" gorm:"default: null"`
	LineId            string            `json:"line_id" gorm:"default: null"`
	Status            *bool             `json:"status" gorm:"default: true"`
	DeletedAt         gorm.DeletedAt    `json:"deleted_at"`
	CustomerAddresses []CustomerAddress `json:"customer_addresses"`
	TypeSocial        string            `json:"type_social"` // "facebook // "line"" // "gmail"
}
