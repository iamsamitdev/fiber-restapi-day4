package models

// กำหนด Struct สำหรับ User
// คล้ายการสร้าง model ในภาษาอื่น ๆ
// โดยมีฟิลด์ต่าง ๆ ที่จำเป็นสำหรับผู้ใช้
type User struct {
	ID        int    `json:"id"`         // รหัสผู้ใช้
	Username  string `json:"username"`   // ชื่อผู้ใช้
	Password  string `json:"password"`   // รหัสผ่าน
	Email     string `json:"email"`      // อีเมล
	FullName  string `json:"full_name"`  // ชื่อเต็ม
	Role      string `json:"role"`       // บทบาทของผู้ใช้ (เช่น admin, user)
	CreatedAt string `json:"created_at"` // วันที่สร้างผู้ใช้
	UpdatedAt string `json:"updated_at"` // วันที่ปรับปรุงข้อมูลผู้ใช้
}
