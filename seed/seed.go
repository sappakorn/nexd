package seed

import (
	"my-api/config"
)
func init() {
	config.SetupEnv()
}

func Run() {
	// ตั้งค่าการเชื่อมต่อกับฐานข้อมูล
	cfg := config.GetConfig()
	db := config.DatabasePostgres(&cfg)
	seed := SeedHandler{Db: db}
	seed.TruncateTable([]string{
		"keyboards", 
		"admins",
		"keycaps",
	})
	seed.KeyboardSeeder() 
	seed.AdminSeeder()
	seed.KeycapsSeeder()
}