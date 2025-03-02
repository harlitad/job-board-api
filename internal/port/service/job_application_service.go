package service

import "github.com/harlitad/job-board-api/internal/domain"

type JobApplicationService interface {
	ApplyToJob(applicantID, jobPostID int) (*domain.JobApplication, error)
}
