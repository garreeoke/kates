package kates

import (
	"errors"
	batchv1 "k8s.io/api/batch/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Create new deployment
func CreateJob(input *Input) (Output, error) {

	var output Output
	if input.Client == nil {
		return output, errors.New(" No kubernetes client, cannot connect")
	}
	job := input.Data.(*batchv1.Job)
	if job.Namespace == "" {
		input.Namespace = "default"
	}
	job, err := input.Client.BatchV1().Jobs(job.Namespace).Create(job)
	if err != nil {
		output.Verified = false
		err = eventMessages(input, &output, meta_v1.ListOptions{
			FieldSelector: "involvedObject.name="+ job.Name,
		})
		return output, err
	}
	return output, nil
}
