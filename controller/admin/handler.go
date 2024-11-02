package adminController

import (
	"gorm.io/gorm"
	"github.com/golang-jwt/jwt"
	"time"
	"os"
	"fmt"
)

type Handler struct {
	Db *gorm.DB
}

						//รับ adminId  return token = string
func GenerateAccessToken(adminId string) string {
	token := jwt.New(jwt.SigningMethodHS256) //สร้าง token โดยใช้ jwt.SigningMethodHS256 เป็นประเภทการสร้าง token
	claims := token.Claims.(jwt.MapClaims) //กำหนด claims เป็น jwt.MapClaims
	claims["admin_id"] = adminId
	fmt.Println(claims["admin_id"])
	claims["exp"] = time.Now().Add(time.Hour * 12).Unix() 
	t, _ := token.SignedString([]byte(os.Getenv("ADMIN_JWT_SECRET")))
	return t
}

