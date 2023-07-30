package main

import (
	"fmt"
	"log"

	"github.com/beego/beego/v2/client/orm"
	"github.com/gofiber/fiber/v2"
	"github.com/joewilson27/microservices-go-fiber/database"
	"github.com/joewilson27/microservices-go-fiber/models"
	"github.com/joewilson27/microservices-go-fiber/routes"
	"github.com/joho/godotenv"
)

func main() {
	port := 3069 // ini connect ke port image docker
	// tapi nanti project yg dari docker di jalanin di 3062 karena panggil port public 3062:3069
	// karena project dipanggil dari sana

	// load .env
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
	}
	// lets verify that .env printing correctly on terminal
	// fmt.Println("print db host", os.Getenv("DB_HOST"))

	// register model to database, right before we call ConnectDB() function.
	// jika ingin di register pada file database juga bisa, caranya register setelah baris code
	// orm.RegisterDataBase
	orm.RegisterModel(new(models.Authors))
	// connect db
	database.ConnectDB()

	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello Mars!")
	})
	// assign routes
	SetupRoutes(app)

	addr := fmt.Sprintf(":%d", port)
	log.Fatal(app.Listen(addr))
}

func SetupRoutes(app *fiber.App) {
	app.Post("/author", routes.AddAuthor)
	app.Get("/authors", routes.GetAllAuthors)
	app.Get("/author/:id", routes.GetSingleAuthor)
	app.Delete("/author", routes.DeleteAuthor)
	app.Put("/author", routes.UpdateAuthor)
}
