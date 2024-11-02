package seed

import (
	adminModel "my-api/model/admin"

	"my-api/utils" // นำเข้าฟังก์ชัน HashPassword
)

func (sh *SeedHandler) AdminSeeder() []adminModel.Admin {
	admins := []adminModel.Admin{

		{
			Username:  "admin",
			Password:  utils.HashPassword("123456789"), // แฮชรหัสผ่าน
			PasswordView: "123456789",
			Email:     "admin@gmail.com",
			Phone:     "0610180599",
			AdminRole: "admin",
		},
	}

	for _, admin := range admins {
		sh.Db.Create(&admin)

	}
	return admins
}