package productModel

import "my-api/model"

type Keycaps struct {
	model.Base
	Name 		string 	`json:"name"`
	Price 		float64 	`json:"price"`
	Quantity 	int 	`json:"quantity"`
	Image 		string 	`json:"image"`
	Description string 	`json:"description"` //คำอธิบาย
	Color 		string 	`json:"color"` 		//สี
	Type 		string 	`json:"type"` 		//cherry, kailh, etc
	Profile 	string 	`json:"profile"` 	//standard, highProfile, lowProfile
}

