package databasepostgres

import (
	"github.com/harlitad/job-board-api/internal/domain"
	ports "github.com/harlitad/job-board-api/internal/port/repository"
	"gorm.io/gorm"
)

type JobRepositoryImpl struct {
	DB *gorm.DB
}

func NewJobRepository(db *gorm.DB) ports.JobRepository {
	return &JobRepositoryImpl{
		DB: db,
	}
}

func (r *JobRepositoryImpl) Create(job *domain.JobPost) error {
	return r.DB.Create(job).Error
}

func (r *JobRepositoryImpl) GetAll() ([]domain.JobPost, error) {
	var jobs []domain.JobPost
	err := r.DB.Find(&jobs).Error
	return jobs, err
}
