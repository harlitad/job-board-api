package domain

import "time"

type JobApplication struct {
	ID          int       `json:"id"`
	JobPostID   int       `json:"job_post_id"`
	ApplicantID int       `json:"applicant_id"`
	Status      string    `json:"status"`
	AppliedAt   time.Time `json:"applied_at"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func NewJobApplication(applicantID, jobPostID int) *JobApplication {
	return &JobApplication{
		JobPostID:   jobPostID,
		ApplicantID: applicantID,
		Status:      "pending", // status default
		AppliedAt:   time.Now(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
