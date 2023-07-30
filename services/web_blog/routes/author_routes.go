package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/joewilson27/microservices-go-fiber/database"
	"github.com/joewilson27/microservices-go-fiber/models"
	"github.com/joewilson27/microservices-go-fiber/utilities"
	"github.com/joewilson27/microservices-go-fiber/validators"
)

func AddAuthor(c *fiber.Ctx) error {
	// Get a base response object
	response := utilities.GetBaseResponseObject()
	postBody := &validators.AuthorAddPostBody{}

	// Verify if the post body is in proper JSON format. Check errors if there are none move forward.
	if err := c.BodyParser(postBody); err != nil {
		fmt.Println("Errrrrrorrr 111")
		response["error"] = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	} else {
		// Validate the post body. Check errors if there are none move forward.
		if err := validators.ValidateStruct(postBody); err != nil {
			fmt.Println("Errrrrrorrr 222")
			return c.Status(fiber.StatusInternalServerError).JSON(err)
		} else {
			author := models.Authors{Title: postBody.Title}
			if _, err := database.Database.Orm.Insert(&author); err != nil {
				fmt.Println("Errrrrrorrr 333")
				response["error"] = err.Error()
				return c.Status(fiber.StatusInternalServerError).JSON(response)
			} else {
				fmt.Println("Errrrrrorrr 444")
				response["message"] = "Author successfully added"
				response["status"] = "pass"
				return c.Status(fiber.StatusCreated).JSON(response)
			}
		}
	}
}

func GetAllAuthors(c *fiber.Ctx) error {
	return c.SendString("All Author")
}

func GetSingleAuthor(c *fiber.Ctx) error {
	return c.SendString("Single Author")
}

func DeleteAuthor(c *fiber.Ctx) error {
	return c.SendString("Delete Author")
}

func UpdateAuthor(c *fiber.Ctx) error {
	return c.SendString("Update Author")
}
