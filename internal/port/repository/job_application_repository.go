package repository

import "github.com/harlitad/job-board-api/internal/domain"

type JobApplicationRepository interface {
	Create(job *domain.JobApplication) error
	GetAll() ([]domain.JobApplication, error)
}
