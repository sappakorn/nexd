package seed

import (
	productModel "my-api/model/product"
)

func (sh *SeedHandler) KeycapsSeeder() []productModel.Keycaps {
	keycaps := []productModel.Keycaps{
		{
			Name:        "Keychron Double Shot 141 Keys Keycap (EN/TH) White",
			Price:       1390.00,
			Quantity:    20,
			Image:       "https://mercular.s3.ap-southeast-1.amazonaws.com/images/products/2024/02/Product/keychron-double-shot-pbt-osa-141-keys-keycap-white-top-view.jpg",
			Description: "Keychron Double Shot PBT OSA 141 Keys Keycap (EN/TH) White",
			Color:       "White",
			Type:        "OSA",
			Profile:     "High Profile",
		},
		{
			Name:        "(Keycaps) Maya Bay Low profile Keycaps 128 Keys ABS",
			Price:       1350.00,
			Quantity:    35,
			Image:       "https://www.keychron.co.th/cdn/shop/files/Screenshot2024-01-17181131.png?v=1705490225&width=800",
			Description: `ปิดตัวชุดคีย์แคป Low profile ที่สามารถใส่ได้กับคีย์บอร์ดทุก LayoutConcept หลักของ Design Keycaps set ชุดนี้คือ เราได้รับ Inspiration มาจาก อ่าวมาหยา เกาะพีพี ประเทศไทย 
						  มาเป็นจุดเริ่มต้นในการดีไซน์ ด้วยการใช้สีเป็นสีหลัก CI ของ Sillicons Studio`,
			Color:       "light blue", 
			Type:        "ABS",
			Profile:     "Low Profile",
		},
		{
			Name:        "Keychron Everest Low-profile Keycap Set",
			Price:       1350.00,
			Quantity:    20,
			Image:       "https://www.keychron.co.th/cdn/shop/files/Everest-Keycap-Set-1.jpg?v=1719396189&width=800",
			Description: "สามารถใช้ได้ร่วมกับรุ่น : Keychron K1 version5, K1 SE,  K7, K1 Pro, K3 Pro,  K7 Pro, K9 Pro, K13 Pro, and S1",
			Color:       "White Blue Black",
			Type:        "OSA",
			Profile:     "Low Profile",
		},
		{
			Name:        "Blue - OEM Dye-Sub PBT",
			Price:       1355.00,
			Quantity:    20,
			Image:       "https://www.keychron.co.th/cdn/shop/products/Blue.jpg?v=1653911748&width=800",
			Description: "Keychron Double Shot PBT OSA 141 Keys Keycap (EN/TH) White",
			Color:       "Blue White",
			Type:        "OEM",
			Profile:     "High Profile",
		},
		{
			Name:        "NuPhy® Gem",
			Price:       1152.70,
			Quantity:    20,
			Image:       "https://nuphy.com/cdn/shop/files/ObsidianBlackv2_1080x.png?v=1705565310",
			Description: "Keychron Double Shot PBT OSA 141 Keys Keycap (EN/TH) White",
			Color:       "Black",
			Type:        "mSA",
			Profile:     "High Profile",
		},
		{
			Name:        "NuPhy® x Suda Carmine Cloud",
			Price:       1661.59,
			Quantity:    20,
			Image:       "https://nuphy.com/cdn/shop/files/CarmineClound_Main01_230323_1080x.jpg?v=1683599677",
			Description: "Keychron Double Shot PBT OSA 141 Keys Keycap (EN/TH) White",
			Color:       "Carmine Cloud",
			Type:        "Cherry",
			Profile:     "High Profile",
		},
		
	}

	for _, keycaps := range keycaps {
		sh.Db.Create(&keycaps)
	}
	return keycaps
}
