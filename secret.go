package kates

import (
	"errors"
	apiv1 "k8s.io/api/core/v1"
)

// CreateSecret new secret
func CreateSecret(input *Input) (*Output, error) {

	output := Output{}
	if input.Client == nil {
		return &output, errors.New(" No kubernetes client, cannot connect")
	}
	secret := input.Data.(*apiv1.Secret)
	if secret.Namespace == "" {
		input.Namespace = "default"
	}
	secret, err := input.Client.CoreV1().Secrets(input.Namespace).Create(secret)
	if err != nil {
		output.Verified = false
		return &output, err
	}
	return &output, nil
}
