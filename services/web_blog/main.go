package main

import (
	"fmt"
	"log"

	"github.com/beego/beego/v2/client/orm"
	"github.com/gofiber/fiber/v2"
	"github.com/joewilson27/microservices-go-fiber/database"
	"github.com/joewilson27/microservices-go-fiber/models"
)

func main() {
	port := 3069 // ini connect ke port image docker
	// tapi nanti project yg dari docker di jalanin di 3062 karena panggil port public 3062:3069
	// karena project dipanggil dari sana
	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello Mars!")
	})

	// register model to database, right before we call ConnectDB() function.
	// jika ingin di register pada file database juga bisa, caranya register setelah baris code
	// orm.RegisterDataBase
	orm.RegisterModel(new(models.Authors))
	// connect db
	database.ConnectDB()

	addr := fmt.Sprintf(":%d", port)
	log.Fatal(app.Listen(addr))
}
