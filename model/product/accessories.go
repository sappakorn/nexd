package productModel

import "my-api/model"

type Accessories struct {
	model.Base
	Name 			string 	`json:"name"` //ชื่อสินค้า
	Price 			int 	`json:"price"` //ราคา
	Quantity 		int 	`json:"quantity"` //จำนวนคงเหลือ
	Image 			string 	`json:"image"` //รูปภาพไอคอน
	Description 	string 	`json:"description"` //คำอธิบาย
	AccessoriesType string 	`json:"accessories_type"` //ประเภทสินค้า
	Color 			string 	`json:"color"` //สี
}

