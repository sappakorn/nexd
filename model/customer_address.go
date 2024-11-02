package model
import (
	"gorm.io/gorm"
)

type CustomerAddress struct {
	Base
	CustomerId    string  `json:"customer_id" gorm:"type:varchar(255);not null"`
	Address       string  `json:"address" gorm:"type:varchar(255);not null"`
	Telephone     string  `json:"telephone" gorm:"type:varchar(255);default: null"`
	ProvinceId    *string `json:"province_id" gorm:"type:varchar(255)"`
	Province      string  `json:"province"`
	DistrictId    *string `json:"district_id" gorm:"type:varchar(255)"`
	District      string  `json:"district"`
	SubDistrictId *string `json:"sub_district_id" gorm:"type:varchar(255)"`
	SubDistrict   string  `json:"sub_district"`
	ZipCode       string  `json:"zip_code" gorm:"type:varchar(20);not null"`
	Latitude      string  `json:"latitude" gorm:"type:varchar(255);default: null"`
	Longtitude    string  `json:"longtitude" gorm:"type:varchar(255);default: null"`
	Type          string  `json:"type" gorm:"type:varchar(255);default: HOME"`
	Default       *bool   `json:"default" gorm:"default: true"`
	DeletedAt     gorm.DeletedAt
}
