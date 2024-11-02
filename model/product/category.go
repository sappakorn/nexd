package productModel

import (
	"my-api/model"
	"gorm.io/gorm"
)

type Category struct {
	model.Base
	Name    string     		`json:"name" gorm:"type:varchar(255);not null"`
	Icon      string     	`json:"icon" gorm:"type:varchar(255);default:null"`
	IconUrl   string     	`json:"icon_url" gorm:"default:null"`
	Parent    *Category  	`json:"category_parent" gorm:"foreignkey:ParentId"`
	Children  []Category 	`json:"category_childrens" gorm:"foreignkey:ParentId"`
	DeletedAt gorm.DeletedAt
}
