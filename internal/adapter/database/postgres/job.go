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
	if err != nil {
		return nil, err
	}
	return jobs, nil
}

func (r *JobRepositoryImpl) GetByID(id int) (*domain.JobPost, error) {
	var job domain.JobPost
	err := r.DB.First(&job, id).Error
	if err != nil {
		return nil, err
	}
	return &job, nil
}

func (r *JobRepositoryImpl) Update(job *domain.JobPost) error {
	return r.DB.Save(job).Error
}

func (r *JobRepositoryImpl) Delete(id int) error {
	return r.DB.Delete(&domain.JobPost{}, id).Error
}
