package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harlitad/job-board-api/internal/port/service"
)

func SetupJobApplicationRoutes(api fiber.Router, jobAppService service.JobApplicationService) {
	// jobAppRoutes := api.Group("/job-applications")
	// jobAppRoutes.Post("/", rest.CreateJobApplicationHandler(jobAppService))
	// jobAppRoutes.Get("/", rest.GetJobApplicationsHandler(jobAppService))
}
