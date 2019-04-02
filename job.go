package kates

import (
	"errors"
	batchv1 "k8s.io/api/batch/v1"
)

// CreateJob new job
func CreateJob(input *Input) (*Output, error) {

	output := Output{}
	if input.Client == nil {
		return &output, errors.New(" No kubernetes client, cannot connect")
	}
	job := input.Data.(*batchv1.Job)
	if job.Namespace == "" {
		input.Namespace = "default"
	}
	job, err := input.Client.BatchV1().Jobs(input.Namespace).Create(job)
	if err != nil {
		output.Verified = false
		return &output, err
	}
	return &output, nil
}
