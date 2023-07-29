package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	port := 3069 // ini connect ke port image docker
	// tapi nanti project yg dari docker di jalanin di 3062 karena panggil port public 3062:3069
	// karena project dipanggil dari sana
	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello Mars!")
	})
	addr := fmt.Sprintf(":%d", port)
	log.Fatal(app.Listen(addr))
}
