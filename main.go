package main

import (
	"fiber-restapi/controllers"
	"fiber-restapi/services"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// สร้างแอปพลิเคชัน Fiber ใหม่
	app := fiber.New()

	// กำหนด route สำหรับหน้าแรก
	// เมื่อผู้ใช้เข้าถึง root URL จะส่งข้อความ "Hello, World!"
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// สร้าง service และ controller สำหรับผู้ใช้
	// โดยใช้โครงสร้างที่กำหนดไว้ใน services และ controllers
	// คล้ายการสร้าง instance ของ class ในภาษาอื่น ๆ
	userService := services.NewUserService()
	userController := controllers.NewUserController(userService)

	// กำหนด route สำหรับการลงทะเบียนผู้ใช้
	app.Post("/register", userController.Register)
	// กำหนด route สำหรับการเข้าสู่ระบบผู้ใช้
	app.Post("/login", userController.Login)

	// Start the server on port 3000
	app.Listen(":3000")

}
