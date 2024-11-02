package routesController

import (
	authController "my-api/controller/auth"
	"my-api/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	productController "my-api/controller/product"
	adminController "my-api/controller/admin"
)

// ตรวจสอบสิทธิเข้าถึง สมาชิกทั่วไป
func RoutesAuth(api fiber.Router, db *gorm.DB) {
	h := &authController.Handler{
		Db: db,
	}
	r := api.Group("/auth")
	r.Post("/register", h.CustomerRegister) // กรอกข้อมูลสมัครสมาชิก
	r.Patch("/register/verify/:customerVerifyId", h.Verify) // ตรวจสอบสมัครสมาชิก
	r.Post("/register/submit", h.CustomerSubmit) // ส่งข้อมูลสมัครสมาชิก
	r.Post("/login", h.CustomerLogin) // เข้าสู่ระบบ
	r.Use(middleware.CheckMiddleware) // ตรวจสอบสิทธิเข้าถึง
	r.Get("/customer/:customerId", h.FindCustomerById) // ดึงข้อมูลสมาชิก ด้วย id สมาชิก


}


//ใช้แสดงสินค้าอย่างหรือเพิ่ม สินค้าได้ในตะกร้า ไม่สามารถทำการ สร้าง ลบ แก้ไข ได้
func RoutesProduct(api fiber.Router, db *gorm.DB) {
	h := &productController.Handler{
		Db: db,
	}
	r := api.Group("/product")
	r.Use(middleware.CheckMiddleware)
	r.Get("/keyboard", h.GetAllKeyboard)
	r.Get("/keyboard/:id", h.GetKeyboardById)
	r.Get("/keycaps", h.GetAllKeycaps)
	r.Get("/keycaps/:id", h.GetKeycapsById)
}


// check middleware คนละแบบกับ user ทั่วไป 
// สิทธิเข้าถึงสินค้า สร้าง ลบ แก้ไข สินค้า มีแค่แอดมิน เท่านั้น
func RoutesAdmin(api fiber.Router, db *gorm.DB) {
	h := &adminController.Handler{
		Db: db,
	}
	r := api.Group("/admin")
	r.Post("/login", h.AdminLogin)
	r.Use(middleware.CheckAdminMiddleware)
	r.Post("/create-keyboard", h.CreateKeyboard) // สร้างข้อมูลคีย์บอร์ด
	r.Put("/update-keyboard/:id", h.UpdateKeyboard) // อัพเดตข้อมูลคีย์บอร์ด
	r.Delete("/delete-keyboard/:id", h.DeleteKeyboard) // ลบข้อมูลคีย์บอร์ด ด้วย id คีย์บอร์ด

	
	//function สำหรับตรวจสอบสิทธิเข้าถึง ให้ใส่token ใน header X-Access-Token
	r.Get("/hello-admin", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello Admin",
		})
	})
}

