package migration

import (
	"my-api/model"
	"gorm.io/gorm"
	"my-api/config"
	productModel "my-api/model/product"
	adminModel "my-api/model/admin"
)

func init() {
	config.SetupEnv()
}

func Run() {
	cfg := config.GetConfig()
	db := config.DatabasePostgres(&cfg)
	migrate(db)
}

func migrate(db *gorm.DB) {

	// ลบตารางที่มีอยู่
	if err := db.Migrator().DropTable(
		&model.Customer{},
		&model.CustomerVerify{},
		&productModel.Keyboard{},
		&productModel.Keycaps{},
		&productModel.Switch{},
		&productModel.Accessories{},
		&adminModel.Admin{},
	); err != nil {
		panic("Failed to drop tables: " + err.Error())
	}



	// ทำการสร้างตารางใหม่
	if err := db.AutoMigrate(
		&model.Customer{},
		&model.CustomerVerify{},
		&productModel.Keyboard{},
		&productModel.Keycaps{},
		&productModel.Switch{},
		&productModel.Accessories{},
		&adminModel.Admin{},
	); err != nil {
		panic("Failed to migrate tables: " + err.Error())
	}
}

