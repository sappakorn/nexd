package productModel

import (
	"gorm.io/gorm"
	"my-api/model"
)

type Keyboard struct {
	model.Base
	Name 			string 	`json:"name" gorm:"unique"`
	Price 			float64 	`json:"price"`
	Quantity 		int 	`json:"quantity"`
	Image 			string 	`json:"image"`
	Description 	string 	`json:"description"`
	Color 			string 	`json:"color"`
	TypeProfile 	string 	`json:"type_profile"` //highProfile, lowProfile 
	SwitchType 		string 	`json:"switch_type"` //linear, tactile, clicky,
	SwitchName 		string 	`json:"switch_name"` //red, blue, brown, black, yellow,
	SwitchColor 	string 	`json:"switch_brand"` //cherry, kailh, etc
	Layout 			string 	`json:"layout"` //60%, 65%, 75%, 80%, 85%, 90%, 95%, 100%
	DeletedAt 		gorm.DeletedAt `gorm:"index"`
}

