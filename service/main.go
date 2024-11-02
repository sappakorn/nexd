package main

import (
	"my-api/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	routesController "my-api/controller"
	"log"
	"strconv"
	"github.com/joho/godotenv"
	
)

func main() {
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
	cfg := config.GetConfig()
	db := config.DatabasePostgres(&cfg)

	//สร้าง app จาก fiber ที่มีการกำหนดขนาดของ body ที่ส่งได้ 15mb
	app := fiber.New(fiber.Config{
		BodyLimit: 15 * 1024 * 1024,
	})

	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		// AllowOrigins: "https://app.dentalliances.com",
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, X-Access-Token",
	}))

	app.Static("/static", "./public")

	api := app.Group("/my-api")
	routesController.RoutesAuth(api, db)
	routesController.RoutesProduct(api, db)
	routesController.RoutesAdmin(api, db)

	if err := app.Listen(":" + strconv.Itoa(cfg.App.Port)); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
