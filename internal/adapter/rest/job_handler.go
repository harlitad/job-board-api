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

func GetJobByIDHandler(jobService service.JobService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid job ID"})
		}

		job, err := jobService.GetJobByID(id)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Job not found"})
		}

		return c.Status(200).JSON(job)
	}
}

func UpdateJobHandler(jobService service.JobService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid job ID"})
		}

		var input struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		}

		if err := c.BodyParser(&input); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}

		updatedJob, err := jobService.UpdateJob(id, input.Title, input.Description)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to update job"})
		}

		return c.Status(200).JSON(updatedJob)
	}
}

func DeleteJobHandler(jobService service.JobService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid job ID"})
		}

		err = jobService.DeleteJob(id)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to delete job"})
		}

		return c.Status(200).JSON(fiber.Map{"message": "Job deleted successfully"})
	}
}
