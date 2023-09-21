package routes

import (
	"github.com/gofiber/fiber/v2"
	handler "github.com/thulani196/recruits-hub/api"
	"github.com/thulani196/recruits-hub/types"
)

func SetupCompanyRoutes(group fiber.Router) {
	repo := handler.NewMongoCompanyRepository()

	group.Get("/", func(c *fiber.Ctx) error {
		companies, err := repo.GetAll()

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"data":    companies,
		})
	})

	group.Get("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		company, err := repo.GetById(id)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"data":    company,
		})
	})

	group.Post("/create", func(c *fiber.Ctx) error {
		company := new(types.Company)
		err := c.BodyParser(&company)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		}

		createdCompany, err := repo.Create(company)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"data":    createdCompany,
		})
	})
}
