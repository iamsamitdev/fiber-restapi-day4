package main

import (
	"fiber-restapi/controllers"
	"fiber-restapi/services"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	// Define a simple route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// สร้าง service และ controller สำหรับผู้ใช้
	userService := services.NewUserService()
	userController := controllers.NewUserController(userService)

	// กำหนด route สำหรับการลงทะเบียนผู้ใช้
	app.Post("/register", userController.Register)
	// กำหนด route สำหรับการเข้าสู่ระบบผู้ใช้
	app.Post("/login", userController.Login)

	// Start the server on port 3000
	app.Listen(":3000")

}
