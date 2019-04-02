package kates

import (
	"errors"
	apiv1 "k8s.io/api/core/v1"
)

// CreateService new service
func CreateService(input *Input) (*Output, error) {

	output := Output{}
	if input.Client == nil {
		return &output, errors.New(" No kubernetes client, cannot connect")
	}
	svc := input.Data.(*apiv1.Service)
	if svc.Namespace == "" {
		input.Namespace = "default"
	}
	svc, err := input.Client.CoreV1().Services(input.Namespace).Create(svc)
	if err != nil {
		output.Verified = false
		return &output, err
	}
	return &output, nil
}
