package kates

import (
	"errors"
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Job
func Job(input *Input) (*Output, error) {
	output := Output{}
	var err error
	if input.Client == nil {
		return &output, errors.New(" No kubernetes client, cannot connect")
	}
	job := input.Data.(*batchv1.Job)
	if job.Namespace == "" {
		job.Namespace = "default"
	}
	switch input.Operation {
	case OpCreate:
		job, err = input.Client.BatchV1().Jobs(job.Namespace).Create(job)
	case OpUpdate:
		job, err = input.Client.BatchV1().Jobs(job.Namespace).Update(job)
	case OpDynamic:
		_, err := input.Client.BatchV1().Jobs(job.Namespace).Get(job.Name, metav1.GetOptions{})
		if err != nil {
			job, err = input.Client.BatchV1().Jobs(job.Namespace).Create(job)
		} else {
			job, err = input.Client.BatchV1().Jobs(job.Namespace).Update(job)
		}
	}
	output.Result = job
	if err != nil {
		output.Verified = false
		return &output, err
	}
	return &output, nil
}
