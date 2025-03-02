package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harlitad/job-board-api/internal/adapter/rest"
	"github.com/harlitad/job-board-api/internal/port/service"
)

func SetupJobRoutes(api fiber.Router, jobService service.JobService) {
	jobRoutes := api.Group("/jobs")
	jobRoutes.Post("/", rest.CreateJobPostHandler(jobService))
	jobRoutes.Get("/", rest.GetJobPostsHandler(jobService))
}
