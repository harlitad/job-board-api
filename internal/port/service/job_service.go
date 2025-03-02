package service

import "github.com/harlitad/job-board-api/internal/domain"

type JobService interface {
	CreateJob(title, description string) (*domain.JobPost, error)
	GetAllJobs() ([]domain.JobPost, error)
}
