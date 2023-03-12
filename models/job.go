package models

import "time"

// Job model
type Job struct {
	ID                    string    `json:"id"`
	AssignedTruck         string    `json:"assigned_truck"`
	AssignedDriver        string    `json:"assigned_driver"`
	CreatedDate           time.Time `json:"created_date"`
	Status                JobStatus `json:"status" gorm:"type:job_status"`
	OriginLocationID      string    `json:"origin_location_id"`
	DestinationLocationID string    `json:"destination_location_id"`
}

type JobStatus string

const (
	JobInProgress      JobStatus = "Job In Progress"
	JobFinished        JobStatus = "Job Finished"
	JobNotYetStarted   JobStatus = "Not Yet Started"
	JobReceivingInProg JobStatus = "Job Receiving In Progress"
	JobFailedToReceive JobStatus = "Failed To Receive"
	JobFailed          JobStatus = "Job Failed"
)

func NewJob(assignedTruck string, assignedDriver string, createdDate time.Time, status JobStatus, originLocationID string, destinationLocationID string) Job {
	return Job{
		ID:                    randomString(16),
		AssignedTruck:         assignedTruck,
		AssignedDriver:        assignedDriver,
		CreatedDate:           createdDate,
		Status:                status,
		OriginLocationID:      originLocationID,
		DestinationLocationID: destinationLocationID,
	}
}
