package databasepostgres

import (
	"github.com/harlitad/job-board-api/internal/domain"
	ports "github.com/harlitad/job-board-api/internal/port/repository"
	"gorm.io/gorm"
)

type JobApplicationRepositoryImpl struct {
	DB *gorm.DB
}

func NewJobApplicationRepository(db *gorm.DB) ports.JobApplicationRepository {
	return &JobApplicationRepositoryImpl{
		DB: db,
	}
}

func (r *JobApplicationRepositoryImpl) Create(job *domain.JobApplication) error {
	return r.DB.Create(job).Error
}

func (r *JobApplicationRepositoryImpl) GetAll() ([]domain.JobApplication, error) {
	var jobApplications []domain.JobApplication
	err := r.DB.Find(&jobApplications).Error
	return jobApplications, err
}
