package kates

import (
	"errors"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Create new deployment
func CreateCronJob(input *Input) (Output, error) {

	var output Output
	if input.Client == nil {
		return output, errors.New(" No kubernetes client, cannot connect")
	}
	cronJob := input.Data.(*batchv1beta1.CronJob)
	if cronJob.Namespace == "" {
		input.Namespace = "default"
	}
	cronJob, err := input.Client.BatchV1beta1().CronJobs(cronJob.Namespace).Create(cronJob)
	if err != nil {
		output.Verified = false
		err = eventMessages(input, &output, meta_v1.ListOptions{
			FieldSelector: "involvedObject.name="+ cronJob.Name,
		})
		return output, err
	}
	return output, nil
}
