package routes

import (
	"github.com/gofiber/fiber/v2"
)

type Todo struct { //เป็นการสร้าง แบบแปลน ยังไม่ได้สร้างพื้นที่ใน memory
	Id int `json:"id"` // `json:"id"` เป็นการบอกว่า field นี้ เมื่อเปลี่ยนค่ากลับไปกลับมาระหว่าง Json จะบอกว่า field บน GO จะแปลง เป็น field
	// ชื่อ อะไรใน JSON และ ชื่อ field ใน struct ควรมีตัวแรกเป็นตัวใหญ่ จะได้มีการตรวจจับการเข้าถึง ได้ง่าย
	Success bool   `json:"success"`
	Body    string `json:"body"`
}

type user struct {
	todos []Todo //เก็บค่ากรณีไม่ได้ทำงาน กับ
}

// SetupUserRoutes ตั้งค่า route ที่เกี่ยวกับผู้ใช้
func SetupUserRoutes(app *fiber.App) {
	userGroup := app.Group("/users")
	var isUser user
	userGroup.Get("/", isUser.getAllUsers)
	userGroup.Post("/", isUser.createUser)
}

func (u *user) getAllUsers(c *fiber.Ctx) error {
	err := c.Status(200).JSON(u.todos) // ใน golang มาแปลงเป็น JSON สำหรับการส่งข้อมูล ทาง http
	if err != nil {
		// ตรวจจับ error แล้วส่ง response ตามต้องการ
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ไม่สามารถแปลงข้อมูลเป็น JSON ได้",
		})
	}
	return err
}

func (u *user) createUser(c *fiber.Ctx) error {
	todo := &Todo{}                            // Todo{} เป็นการประกาศสร้าง struct จริงๆ โดยอ้างจาก แปลน Todo ค่าจะกำหนดให้เป็น default และ todo เก็น addrs ของ Todo struct ที่พึ่งสร้าง (เป็น pointer)
	if err := c.BodyParser(todo); err != nil { //เมื่อทำงานถึง bodyParser จะตรวจสอบ ข้อมูลที่ส่งเข้ามา กับ struct ที่ todo ชี้อยู่ตรงกันไหม ถ้าตรงกันเอาค่าที่ส่งมาไปใส่ใน struct(ไม่ได้ใส่ใน todo) แล้วเรียกใช้ค่าโดย todo.Body ได้เลย
		return err
	}

	if todo.Body == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{ // fiber.Map เป็นการนำข้อมูลใน {} มาสร้างเป็น map struct หรือประกาศ map ใน golang แล้วค่อยมาแปลงเป็น JSON สำหรับการส่งข้อมูล ทาง http
			"Success": " false",
			"Error":   "Body is required",
		})
	}

	todo.Id = len(u.todos) + 1
	todo.Success = true
	u.todos = append(u.todos, *todo)

	return c.Status(fiber.StatusOK).JSON(*todo)
}
