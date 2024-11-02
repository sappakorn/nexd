package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"

)

func CheckMiddleware(c *fiber.Ctx) error {
	accessToken := c.Get("X-Access-Token") //ดึง token จาก header
	secretKey := os.Getenv("JWT_SECRET") 	//ดึง secret key จาก env
	token, err := jwt.ParseWithClaims(accessToken, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil //ตรวจสอบ token ว่าถูกต้องหรือไม่
	}); if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized) //ถ้า token ไม่ถูกต้อง ส่ง status 401
	}

	c.Locals("user", token) //เก็บ token ไว้ใน locals
	return c.Next() //ส่งต่อไปยัง route ต่อไป
}  

func CheckAdminMiddleware(c *fiber.Ctx) error {
	accessToken := c.Get("X-Access-Token") 
	secretKey := os.Getenv("ADMIN_JWT_SECRET") 	
	token, err := jwt.ParseWithClaims(accessToken, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil 
	})
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized) 
	}
	c.Locals("admin", token) //เก็บ token ไว้ใน locals
	return c.Next() //ส่งต่อไปยัง route ต่อไป
}
