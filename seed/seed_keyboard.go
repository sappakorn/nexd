package seed

import (
	productModel "my-api/model/product"
)

func (sh *SeedHandler) KeyboardSeeder() []productModel.Keyboard {
	keyboards := []productModel.Keyboard{
		{
			Name:        "IQUNIX ZX75 Dark Side Wireless Mechanical Keyboard",
			Price:       6042.79,
			Quantity:    10,
			Image:       "https://iqunix.store/cdn/shop/products/iqunix-zx75-dark-side-wireless-mechanical-keyboard-283507_1800x1800.jpg?v=1686823812",
			Description: "Inspired by the dark and mysterious world of gaming, where players immerse themselves in thrilling battles and epic adventures. The sleek black color with red RGB accents exudes a sense of power and intensity, reflecting the dark and daring nature of gaming. The dark-themed keycaps are not only visually striking, but they also convey a sense of mystery, inviting users to explore the depths of their creativity and productivity.",
			Color:       "Black",
			TypeProfile: "High Profile",
			SwitchType:  "linear",
			SwitchName:  "Gold Pink",
			SwitchColor: "Gold Pink",
			Layout:      "75%",
		},
		{
			Name:        "Lofree Flow 84",
			Price:       5500,
			Quantity:    50,
			Image:       "https://www.lofree.co/cdn/shop/files/Flow84_white_PI01_d336e44a-a51d-45c6-b203-fc4c2e577994.webp?v=1714526288&width=800",
			Description: "Embrace the Flow 84, the ultra-slim mechanical keyboard, where industrial design aesthetic meets smooth typing feel. You can ready to complement your professional and tasteful lifestyle at a moment's notice. An all-aluminum base and frame with the low profile switches, perfectly combines professional craftsmanship with sleek aesthetics. Choose bright white for a modern and organized office atmosphere or sophisticated black for a touch of luxury.",
			Color:       "White",
			TypeProfile: "Low Profile",
			SwitchType:  "linear",
			SwitchName:  "ghost",
			SwitchColor: "green",
			Layout:      "84%",
		},
		{
			Name:        "Keychron Q1 HE QMK Wireless Custom Keyboard",
			Price:       7800,
			Quantity:    50,
			Image:       "https://www.keychron.com/cdn/shop/files/Keychron-Q1-HE-QMK-Wireless-Custom-Keyboard-Carbon-Black.jpg?v=1704881599&width=900",
			Description: "In addition to the Hall Effect magnetic switches, the Q1 HE introduces a 2.4G connection mode with an impressive 1000Hz polling rate, further ensuring a lightning-fast gaming experience. But that's not all; it also boasts premium features such as Bluetooth connectivity, a double-gasket design, an all-aluminum body, PBT keycaps, and more",
			Color:       "Carbon Black",
			TypeProfile: "Low Profile",
			SwitchType:  "linear",
			SwitchName:  "Gateron Double-Rail Magnetic Hall Effect Switches",
			SwitchColor: "purple",
			Layout:      "75%",
		},
		{
			Name:        "IQUNIX x Little Prince ZX75 Sunset Ponder Mechanical Keyboard",
			Price:       6717.96,
			Quantity:    10,
			Image:       "https://iqunix.store/cdn/shop/files/ZX75-_-1_540x.jpg?v=1705371799",
			Description: "Tenkeyless mechanical gaming keyboard designed for FPS gamers.",
			Color:       "CreamYellow",
			TypeProfile: "High Profile",
			SwitchType:  "Linear",
			SwitchName:  "Little Prince",
			SwitchColor: "yellow",
			Layout:      "75%",
		},
		{
			Name:        "IQUNIX OG80 Dark Side Wireless Mechanical Keyboard",
			Price:       6042.79,
			Quantity:    25,
			Image:       "https://iqunix.store/cdn/shop/products/iqunix-og80-dark-side-wireless-mechanical-keyboard-216359_1800x1800.jpg?v=1686823791",
			Description: "Customizable mechanical keyboard with adjustable actuation points.",
			Color:       "Black",
			TypeProfile: "High Profile",
			SwitchType:  "mechanical",
			SwitchName:  "ACE",
			SwitchColor: "White",
			Layout:      "80%",
		},
	}

	for _, keyboard := range keyboards {
		sh.Db.Create(&keyboard)
	}
	return keyboards
}
