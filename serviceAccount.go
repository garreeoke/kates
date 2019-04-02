package kates

import (
	"errors"
	apiv1 "k8s.io/api/core/v1"
)

// CreateServiceAccount new service account
func CreateServiceAccount(input *Input) (*Output, error) {

	output := Output{}
	if input.Client == nil {
		return &output, errors.New(" No kubernetes client, cannot connect")
	}
	sa := input.Data.(*apiv1.ServiceAccount)
	if sa.Namespace == "" {
		input.Namespace = "default"
	}
	sa, err := input.Client.CoreV1().ServiceAccounts(input.Namespace).Create(sa)
	if err != nil {
		output.Verified = false
		return &output, err
	}
	return &output, nil
}
