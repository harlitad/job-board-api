package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harlitad/job-board-api/internal/port/service"
)

func SetupRoutes(app *fiber.App, jobService service.JobService) {
	api := app.Group("/api")

	SetupJobRoutes(api, jobService)
	// SetupJobApplicationRoutes(api, jobAppService)
}
