package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2" //ใน go.mod มี แล้ว ไม่รุ้ต้องใส่อีกไหม ??????
)

type Todo struct { //เป็นการสร้าง แบบแปลน ยังไม่ได้สร้างพื้นที่ใน memory
	Id      int    `json:"id"`
	Success bool   `json:"success"`
	Body    string `json:"body"`
}

func main() {

	fmt.Println("Hell World")
	app := fiber.New() // สร้าง instance server
	todos := []Todo{}  //เก็บค่ากรณีไม่ได้ทำงาน กับ

	app.Get("/", func(c *fiber.Ctx) error { //ทุกค่าที่ return เป็น data type error หมด แต่ค่าที่ ทำงานถูกต้องจะreturn error datatype เป็น nil
		// c เป็น pointer สำหรับการจัดการทุกสิ่งสำหรับส่งไป หรือ รับ ผ่าน Http request ไม่ว่าจะการเข้าถึงข้อมูลใดๆที่มากับ Http request
		err := c.Status(200).JSON(todos) // ใน golang มาแปลงเป็น JSON สำหรับการส่งข้อมูล ทาง http
		if err != nil {
			// ตรวจจับ error แล้วส่ง response ตามต้องการ
			return c.Status(500).JSON(fiber.Map{
				"error": "ไม่สามารถแปลงข้อมูลเป็น JSON ได้",
			})
		}
		return err
	})

	app.Post("/", func(c *fiber.Ctx) error {
		todo := &Todo{}                            // Todo{} เป็นการประกาศสร้าง struct จริงๆ โดยอ้างจาก แปลน Todo ค่าจะกำหนดให้เป็น default และ todo เก็น addrs ของ Todo struct ที่พึ่งสร้าง (เป็น pointer)
		if err := c.BodyParser(todo); err != nil { //เมื่อทำงานถึง bodyParser จะตรวจสอบ ข้อมูลที่ส่งเข้ามา กับ struct ที่ todo ชี้อยู่ตรงกันไหม ถ้าตรงกันเอาค่าที่ส่งมาไปใส่ใน struct(ไม่ได้ใส่ใน todo) แล้วเรียกใช้ค่าโดย todo.Body ได้เลย
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{ // fiber.Map เป็นการนำข้อมูลใน {} มาสร้างเป็น map struct หรือประกาศ map ใน golang แล้วค่อยมาแปลงเป็น JSON สำหรับการส่งข้อมูล ทาง http
				"Success": " false",
				"Error":   "Body is required",
			})
		}

		todo.Id = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(201).JSON(*todo)
	})

	type PartialBody struct {
		Body string `json:"Body"` // `json:"Body"` เป็นการบอกว่า field นี้ focus ไปที่ field ไหน ของ json
	}

	app.Patch("/:id", func(c *fiber.Ctx) error {
		pb := &PartialBody{}
		id := c.Params("id")
		c.BodyParser(pb)
		for i, todo := range todos {
			if fmt.Sprint(todo.Id) == id {
				todos[i].Success = true
				todos[i].Body = pb.Body
				return c.Status(200).JSON(todos[i])
			}
		}

		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	})

	app.Delete("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.Id) == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"success": "truetrue"})
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "todos not found"})
	})

	log.Fatal(app.Listen(":4000")) // log นำเข้าจาก log package มีหน้าที่จะหยุดการทำงานของโปรแกรม เมื่อการทำงาน error

}

// เพิ่มเติมความเข้าใจ
// ก่อนแปลงเป็น JSON แล้วส่งไปทาง HTTP จะสามารถใช้เป็น Data structure อะไรก็ได้จาก logic หรือ condition ใดๆ
// พอส่งไปถึง frontend ก็จะนำข้อมูล JSON ไปแปลง เป็นข้อมูลที่ใช้ Data Struct ที่ต้องการตามภาษานั้น ฝั้ง frontend
