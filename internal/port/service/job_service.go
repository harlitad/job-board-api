package service

import "github.com/harlitad/job-board-api/internal/domain"

type JobService interface {
	CreateJob(title, description string) (*domain.JobPost, error)
	GetAllJobs() ([]domain.JobPost, error)
	GetJobByID(id int) (*domain.JobPost, error)
	UpdateJob(id int, title, description string) (*domain.JobPost, error)
	DeleteJob(id int) error
}
