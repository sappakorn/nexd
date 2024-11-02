package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	
)
type UserExtract struct {
	UserId 		string 
	Username 	string 
}



func CustomerExtractToken(c *fiber.Ctx) UserExtract {
	user := c.Locals("User")
	if user == nil {
		return UserExtract{}
	}
	token := user.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	user_id := claims["user_id"].(string)
	return UserExtract{
		UserId: user_id,
	}

}

func AdminExtractToken(c *fiber.Ctx) UserExtract  {
	admin := c.Locals("admin")
	if admin == nil {
		return UserExtract{}
	}
	token := admin.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	admin_id := claims["admin_id"].(string)
	return UserExtract{
		UserId: admin_id,
	}
}