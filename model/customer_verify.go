package model

import "time"

type CustomerVerify struct {
	Base
	Telephone  string    `json:"telephone" gorm:"default: null"`
	Email      string    `json:"email" gorm:"default: null"`
	Reference  string    `json:"reference"`
	Code       int       `json:"code"`
	TypeVerify string    `json:"type_verify"`                    // 1:register // 2:reset_password
	TypeSystem string    `json:"type_system"`                    // 1:sms // 2:email
	Status     string    `json:"status" gorm:"default: PENDING"` // PENDING // SUCCESS // FAIL
	ExpireAt   time.Time `json:"expire_at" gorm:"default: null"`
}

