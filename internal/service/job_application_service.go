package serviceimpl

import (
	"github.com/harlitad/job-board-api/internal/domain"
	"github.com/harlitad/job-board-api/internal/port/repository"
)

// ApplicantService mengelola logika bisnis terkait aplikasi pekerjaan
type ApplicantService struct {
	ApplicationRepository repository.JobApplicationRepository
}

func (s *ApplicantService) ApplyToJob(applicantID, jobPostID int) (*domain.JobApplication, error) {
	application := domain.NewJobApplication(applicantID, jobPostID)
	err := s.ApplicationRepository.Create(application)
	if err != nil {
		return nil, err
	}

	return application, nil
}
