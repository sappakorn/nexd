package authController

import (
	"os"
	"time"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type Handler struct {
	Db *gorm.DB
}
						//รับ customerId  return token = string
func GenerateAccessToken(customerId string) string {
	token := jwt.New(jwt.SigningMethodHS256) //สร้าง token โดยใช้ jwt.SigningMethodHS256 เป็นประเภทการสร้าง token
	claims := token.Claims.(jwt.MapClaims) //กำหนด claims เป็น jwt.MapClaims
	claims["user_id"] = customerId
	claims["exp"] = time.Now().Add(time.Hour * 12).Unix() 
	t, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return t
}

