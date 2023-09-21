package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/thulani196/recruits-hub/database"
	"github.com/thulani196/recruits-hub/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file could not be found.")
	}

	database.ConnectDB()

	app := fiber.New()
	// Group the routes with a common prefix
	jobGroup := app.Group("/api/jobs")
	userGroup := app.Group("/api/users")
	companyGroup := app.Group("/api/company")

	routes.SetupUserRoutes(userGroup)
	routes.SetupJobRoutes(jobGroup)
	routes.SetupCompanyRoutes(companyGroup)

	log.Fatal(app.Listen(":8080"))
}
