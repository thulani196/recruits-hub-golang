package routes

import (
	"github.com/gofiber/fiber/v2"
	handler "github.com/thulani196/recruits-hub/api"
	"github.com/thulani196/recruits-hub/types"
)

func SetupUserRoutes(group fiber.Router) {
	repo := handler.NewMongoUserRepository()

	group.Post("/login", func(c *fiber.Ctx) error {
		user := new(types.User)
		err := c.BodyParser(&user)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": err,
			})
		}

		res, err := repo.LoginHandler(user)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(res)
	})

	group.Post("/register", func(c *fiber.Ctx) error {
		user := new(types.User)
		err := c.BodyParser(&user)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Cannot parse JSON",
				"error":   err,
			})
		}

		err = repo.RegisterHandler(user)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": err,
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "Registeration successfully.",
		})
	})
}
