package productModel

import "my-api/model"

type Switch struct {
	model.Base
	Name 		string `json:"name"`
	Price 		int `json:"price"`
	Quantity 	int `json:"quantity"`
	Image 		string `json:"image"`
	Description string `json:"description"` //คำอธิบาย
	Color 		string `json:"color"` //สี
	Type 		string `json:"type"` //linear, tactile, clicky
}

