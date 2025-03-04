package repository

import "github.com/harlitad/job-board-api/internal/domain"

type JobRepository interface {
	Create(job *domain.JobPost) error
	GetAll() ([]domain.JobPost, error)
	GetByID(id int) (*domain.JobPost, error)
	Update(job *domain.JobPost) error
	Delete(id int) error
}
