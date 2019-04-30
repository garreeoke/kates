package kates

import (
	"errors"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
)

// CronJob
func CronJob(input *Input) (*Output, error) {
	output := Output{}
	var err error
	if input.Client == nil {
		return &output, errors.New(" No kubernetes client, cannot connect")
	}
	cronJob := input.Data.(*batchv1beta1.CronJob)
	if cronJob.Namespace == "" {
		cronJob.Namespace = "default"
	}
	switch input.Operation {
	case OpCreate:
		cronJob, err = input.Client.BatchV1beta1().CronJobs(cronJob.Namespace).Create(cronJob)
	case OpModify:
		cronJob, err = input.Client.BatchV1beta1().CronJobs(cronJob.Namespace).Update(cronJob)
	}
	output.Result = cronJob
	if err != nil {
		output.Verified = false
		return &output, err
	}
	return &output, nil
}
