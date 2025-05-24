package services

import (
	"fiber-restapi/models"
	"fmt"
	"time"
)

// UserService struct สำหรับจัดการผู้ใช้
type UserService struct {
	users  []models.User // จำลองฐานข้อมูลผู้ใช้ในหน่วยความจำ โดยใช้ slice
	nextID int           // ID ถัดไปที่จะกำหนดให้ผู้ใช้ใหม่
}

// สร้างฟังก์ชันใหม่สำหรับ UserService
func NewUserService() *UserService {
	return &UserService{
		users:  []models.User{},
		nextID: 1, // เริ่มต้น ID ถัดไปที่ 1
	}
}

// สร้างฟังก์ชัน Register สำหรับการลงทะเบียนผู้ใช้ใหม่
func (s *UserService) Register(user models.User) (*models.User, error) {
	// ตรวจสอบว่ามีผู้ใช้อยู่แล้วหรือไม่
	for _, u := range s.users {
		if u.Username == user.Username {
			return nil, fmt.Errorf("username %s already exists", user.Username)
		}
		if u.Email == user.Email {
			return nil, fmt.Errorf("email %s already exists", user.Email)
		}
	}

	// กำหนด ID และ timestamps ให้ผู้ใช้ใหม่
	user.ID = s.nextID
	s.nextID++
	now := time.Now().Format(time.RFC3339) // ใช้รูปแบบ ISO 8601 สำหรับวันเวลา เช่น "2023-10-01T12:00:00Z"
	user.CreatedAt = now
	user.UpdatedAt = now

	// กำหนดค่าเริ่มต้นให้กับ Role ถ้าไม่ระบุ
	if user.Role == "" {
		user.Role = "user" // กำหนดบทบาทเริ่มต้นเป็น "user"
	}

	// เพิ่มผู้ใช้ใหม่เข้าไปใน slice
	s.users = append(s.users, user)
	return &user, nil
}

// สร้างฟังก์ชัน Login
func (s *UserService) Login(username, password string) (*models.User, string, error) {
	// ค้นหาผู้ใช้ตามชื่อผู้ใช้
	for _, u := range s.users {
		if u.Username == username && u.Password == password {
			// สร้าง token แบบง่าย (ใน production ควรใช้ JWT)
			token := fmt.Sprintf("token_%s_%d", u.Username, time.Now().Unix())
			return &u, token, nil // คืนค่าผู้ใช้และ token
		}
	}

	return nil, "", fmt.Errorf("invalid username or password")
}
