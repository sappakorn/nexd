package productModel

import "my-api/model"

type Product struct {
	model.Base
	ProductName 	string `json:"product_name"`
	Category 		string `json:"category"`
	Price 			float64 `json:"price"`
	Quantity 		int `json:"quantity"`
	Image 			string `json:"image"`
	Description 	string `json:"description"`
}

