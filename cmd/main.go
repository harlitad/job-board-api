package main

import (
	"github.com/gofiber/fiber/v2"
	databaspostgres "github.com/harlitad/job-board-api/internal/adapter/database/postgres"
	"github.com/harlitad/job-board-api/internal/domain"
	"github.com/harlitad/job-board-api/internal/router"
	serviceimpl "github.com/harlitad/job-board-api/internal/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	db, err := gorm.Open(postgres.Open("host=localhost user=postgres dbname=jobboard password=mysecretpassword port=5432"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}

	db.AutoMigrate(&domain.JobPost{})

	jobRepo := databaspostgres.NewJobRepository(db)
	jobService := serviceimpl.NewJobService(jobRepo)

	router.SetupRoutes(app, jobService)
	app.Listen(":8080")
}
