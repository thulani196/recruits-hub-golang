package routes

import (
	"github.com/gofiber/fiber/v2"
	handler "github.com/thulani196/recruits-hub/api"
	"github.com/thulani196/recruits-hub/types"
)

func SetupJobRoutes(group fiber.Router) {
	repo := handler.NewMongoJobRepository()

	group.Get("/", func(c *fiber.Ctx) error {
		jobs, err := repo.GetAll()

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": "an error occurred while fetching jobs",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"data":    jobs,
		})
	})

	group.Get("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		job, err := repo.GetJobByID(id)

		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"success": false,
				"message": err,
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"data":    job,
		})
	})

	group.Post("/create", func(c *fiber.Ctx) error {
		job := new(types.Job)
		err := c.BodyParser(&job)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Cannot parse JSON",
				"error":   err,
			})
		}

		createdJob, err := repo.CreateJob(job)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": err,
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"data":    createdJob,
		})
	})

	group.Post("/update/:id", func(c *fiber.Ctx) error {
		job := new(types.Job)
		id := c.Params("id")
		err := c.BodyParser(&job)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Cannot parse JSON",
				"error":   err,
			})
		}

		err = repo.UpdateJob(id, job)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": err,
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "Update successful.",
		})
	})
}
