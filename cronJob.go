package kates

import (
	"errors"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
)

// CreateCronJob new cronjob
func CreateCronJob(input *Input) (*Output, error) {

	output := Output{}
	if input.Client == nil {
		return &output, errors.New(" No kubernetes client, cannot connect")
	}
	cronJob := input.Data.(*batchv1beta1.CronJob)
	if cronJob.Namespace == "" {
		input.Namespace = "default"
	}
	cronJob, err := input.Client.BatchV1beta1().CronJobs(input.Namespace).Create(cronJob)
	if err != nil {
		output.Verified = false
		return &output, err
	}
	return &output, nil
}
