# Fiber REST API

โปรเจ็กต์ REST API ที่สร้างด้วย Go และ Fiber Framework สำหรับการจัดการผู้ใช้ (User Management System)

## 📋 ความต้องการระบบ

- Go 1.24.2 หรือสูงกว่า
- Fiber v2.52.8

## 🚀 การติดตั้งและรันโปรเจ็กต์

### 1. Clone โปรเจ็กต์

```bash
git clone <repository-url>
cd fiber-restapi
```

### 2. ติดตั้ง Dependencies

```bash
go mod download
```

### 3. รันแอปพลิเคชัน

```bash
go run main.go
```

เซิร์ฟเวอร์จะรันที่ `http://localhost:3000`

### 4. Build สำหรับ Production

```bash
go build -o app main.go
./app
```

## 📁 โครงสร้างโปรเจ็กต์

```
fiber-restapi/
├── main.go              # Entry point ของแอปพลิเคชัน
├── go.mod               # Go module dependencies
├── go.sum               # Go module checksums
├── README.md            # เอกสารโปรเจ็กต์
├── controllers/         # Controller layer สำหรับจัดการ HTTP requests
│   └── user_controller.go
├── models/              # Data models และ structs
│   └── user.go
├── services/            # Business logic layer
│   └── user_service.go
└── tmp/                 # ไฟล์ชั่วคราว
    └── main.exe
```

## 🎯 คุณสมบัติ

- ✅ ระบบลงทะเบียนผู้ใช้ (User Registration)
- ✅ ระบบเข้าสู่ระบบ (User Login)
- ✅ การจัดการผู้ใช้ในหน่วยความจำ (In-Memory User Management)
- ✅ การสร้าง Token แบบง่าย (Simple Token Generation)
- ✅ การตรวจสอบผู้ใช้ซ้ำ (Duplicate User Validation)

## 📌 API Endpoints

| Method | Endpoint    | Description           | Body                                                         |
|--------|-------------|-----------------------|--------------------------------------------------------------|
| GET    | `/`         | หน้าแรก               | -                                                            |
| POST   | `/register` | ลงทะเบียนผู้ใช้ใหม่   | `{"username": "string", "password": "string", "email": "string", "full_name": "string"}` |
| POST   | `/login`    | เข้าสู่ระบบ           | `{"username": "string", "password": "string"}`              |

## 📝 ตัวอย่างการใช้งาน

### ลงทะเบียนผู้ใช้ใหม่

```bash
curl -X POST http://localhost:3000/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john_doe",
    "password": "password123",
    "email": "john@example.com",
    "full_name": "John Doe"
  }'
```

**Response:**
```json
{
  "success": true,
  "message": "ลงทะเบียนสำเร็จ",
  "user": {
    "id": 1,
    "username": "john_doe",
    "email": "john@example.com",
    "full_name": "John Doe",
    "role": "user",
    "created_at": "2025-05-31T10:30:00Z",
    "updated_at": "2025-05-31T10:30:00Z"
  }
}
```

### เข้าสู่ระบบ

```bash
curl -X POST http://localhost:3000/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john_doe",
    "password": "password123"
  }'
```

**Response:**
```json
{
  "success": true,
  "message": "เข้าสู่ระบบสำเร็จ",
  "token": "token_john_doe_1735646200",
  "user": {
    "id": 1,
    "username": "john_doe",
    "email": "john@example.com",
    "full_name": "John Doe",
    "role": "user"
  }
}
```

## 🏗️ สถาปัตยกรรม

โปรเจ็กต์นี้ใช้หลักการ **Clean Architecture** แบบง่าย:

1. **Controllers Layer**: จัดการ HTTP requests และ responses
2. **Services Layer**: ประมวลผล business logic
3. **Models Layer**: กำหนด data structures

## 🔧 การพัฒนาต่อ

### เพิ่มฟีเจอร์ใหม่

1. สร้าง model ใหม่ใน `models/`
2. เพิ่ม business logic ใน `services/`
3. สร้าง controller สำหรับ HTTP handling ใน `controllers/`
4. เพิ่ม routes ใน `main.go`

### ข้อแนะนำสำหรับ Production

- [ ] เพิ่มฐานข้อมูลจริง (PostgreSQL, MySQL, MongoDB)
- [ ] ใช้ JWT สำหรับ authentication
- [ ] เพิ่ม password hashing (bcrypt)
- [ ] เพิ่ม input validation
- [ ] เพิ่ม logging และ monitoring
- [ ] เพิ่ม CORS middleware
- [ ] เพิ่ม rate limiting
- [ ] เพิ่ม unit tests

## 🛠️ Dependencies

- [Fiber v2](https://gofiber.io/) - Web framework สำหรับ Go
- [Go 1.24.2](https://golang.org/) - ภาษา Go

## 📄 License

ใช้ได้เพื่อการศึกษาและพัฒนา

## 👥 การสนับสนุน

หากมีคำถามหรือปัญหาในการใช้งาน สามารถ:
- เปิด Issue ในระบบ Git
- ติดต่อทีมพัฒนา

---

**หมายเหตุ**: โปรเจ็กต์นี้สร้างขึ้นเพื่อการศึกษา ไม่แนะนำให้ใช้ใน production โดยตรงโดยไม่มีการปรับปรุงด้านความปลอดภัย
