package serviceimpl

import (
	"github.com/harlitad/job-board-api/internal/domain"
	"github.com/harlitad/job-board-api/internal/port/repository"
	"github.com/harlitad/job-board-api/internal/port/service"
)

type JobServiceImpl struct {
	JobRepo repository.JobRepository
}

func NewJobService(repo repository.JobRepository) service.JobService {
	return &JobServiceImpl{
		JobRepo: repo,
	}
}

func (s *JobServiceImpl) CreateJob(title, description string) (*domain.JobPost, error) {
	job := &domain.JobPost{
		Title:       title,
		Description: description,
	}

	err := s.JobRepo.Create(job)
	if err != nil {
		return nil, err
	}

	return job, nil
}

func (s *JobServiceImpl) GetAllJobs() ([]domain.JobPost, error) {
	return s.JobRepo.GetAll()
}

func (s *JobServiceImpl) GetJobByID(id int) (*domain.JobPost, error) {
	return s.JobRepo.GetByID(id)
}

func (s *JobServiceImpl) UpdateJob(id int, title, description string) (*domain.JobPost, error) {
	job, err := s.JobRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	job.Title = title
	job.Description = description

	err = s.JobRepo.Update(job)
	if err != nil {
		return nil, err
	}

	return job, nil
}

func (s *JobServiceImpl) DeleteJob(id int) error {
	return s.JobRepo.Delete(id)
}
