package main

import (
	"fmt"
	"log"

	routes "github.com/Sirawithon11/Golang_Project/route_and_controller"
	"github.com/gofiber/fiber/v2" //ใน go.mod มี แล้ว ไม่รุ้ต้องใส่อีกไหม ??????
)

func add(a int, b int) int {
	return a + b
}

func main() {

	fmt.Println("Hell World")
	app := fiber.New() // สร้าง instance server
	routes.SetupUserRoutes(app)
	log.Fatal(app.Listen(":4000")) // log นำเข้าจาก log package มีหน้าที่จะหยุดการทำงานของโปรแกรม เมื่อการทำงาน error

}

// เพิ่มเติมความเข้าใจ
// ก่อนแปลงเป็น JSON แล้วส่งไปทาง HTTP จะสามารถใช้เป็น Data structure อะไรก็ได้จาก logic หรือ condition ใดๆ
// พอส่งไปถึง frontend ก็จะนำข้อมูล JSON ไปแปลง เป็นข้อมูลที่ใช้ Data Struct ที่ต้องการตามภาษานั้น ฝั้ง frontend
