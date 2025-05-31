package controllers

import (
	"fiber-restapi/models"
	"fiber-restapi/services"

	"github.com/gofiber/fiber/v2"
)

// UserController struct สำหรับจัดการผู้ใช้
// คล้ายการสร้าง class ในภาษาอื่น ๆ
type UserController struct {
	service *services.UserService
}

// สร้าง NewUserController สำหรับการสร้าง UserController ใหม่
// คล้ายการสร้าง constructor ในภาษาอื่น ๆ
// โดยรับพารามิเตอร์เป็น service ที่ใช้จัดการผู้ใช้
func NewUserController(service *services.UserService) *UserController {
	return &UserController{service: service}
}

// RegisterUser สำหรับการลงทะเบียนผู้ใช้ใหม่
// คล้ายการสร้าง method ในภาษาอื่น ๆ
// โดยรับข้อมูลผู้ใช้จาก request body และส่งไปยัง service เพื่อจัดการการลงทะเบียน
func (ctrl *UserController) Register(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "ข้อมูลไม่ถูกต้อง",
			"error":   "Invalid request",
		})
	}

	registeredUser, err := ctrl.service.Register(*user)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "ลงทะเบียนไม่สำเร็จ",
			"error":   err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "ลงทะเบียนสำเร็จ",
		"user": fiber.Map{
			"id":         registeredUser.ID,
			"username":   registeredUser.Username,
			"email":      registeredUser.Email,
			"full_name":  registeredUser.FullName,
			"role":       registeredUser.Role,
			"created_at": registeredUser.CreatedAt,
			"updated_at": registeredUser.UpdatedAt,
		},
	})
}

// LoginUser สำหรับการเข้าสู่ระบบผู้ใช้
// คล้ายการสร้าง method ในภาษาอื่น ๆ
// โดยรับข้อมูลผู้ใช้จาก request body และส่งไปยัง service เพื่อจัดการการเข้าสู่ระบบ
func (ctrl *UserController) Login(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "ข้อมูลไม่ถูกต้อง",
			"error":   "Invalid request",
		})
	}

	loggedInUser, token, err := ctrl.service.Login(user.Username, user.Password)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "เข้าสู่ระบบไม่สำเร็จ",
			"error":   err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "เข้าสู่ระบบสำเร็จ",
		"user": fiber.Map{
			"id":         loggedInUser.ID,
			"username":   loggedInUser.Username,
			"email":      loggedInUser.Email,
			"full_name":  loggedInUser.FullName,
			"role":       loggedInUser.Role,
			"created_at": loggedInUser.CreatedAt,
			"updated_at": loggedInUser.UpdatedAt,
		},
		"token": token,
	})
}
