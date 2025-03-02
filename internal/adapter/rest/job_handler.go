package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harlitad/job-board-api/internal/port/service"
)

func CreateJobPostHandler(jobService service.JobService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		}

		if err := c.BodyParser(&input); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}

		job, err := jobService.CreateJob(input.Title, input.Description)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to create job"})
		}

		return c.Status(201).JSON(job)
	}
}

func GetJobPostsHandler(jobService service.JobService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jobs, err := jobService.GetAllJobs()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch jobs"})
		}

		return c.Status(200).JSON(jobs)
	}
}
