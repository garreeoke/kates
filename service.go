package kates

import (
	"errors"
	apiv1 "k8s.io/api/core/v1"
)

// Service
func Service(input *Input) (*Output, error) {
	output := Output{}
	var err error
	if input.Client == nil {
		return &output, errors.New(" No kubernetes client, cannot connect")
	}
	service := input.Data.(*apiv1.Service)
	if service.Namespace == "" {
		service.Namespace = "default"
	}
	switch input.Operation {
	case OpCreate:
		service, err = input.Client.CoreV1().Services(service.Namespace).Create(service)
	case OpModify:
		service, err = input.Client.CoreV1().Services(service.Namespace).Update(service)
	}
	output.Result = service
	if err != nil {
		output.Verified = false
		return &output, err
	}
	return &output, nil
}