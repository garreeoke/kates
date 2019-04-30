package kates

import (
	"errors"
	apiv1 "k8s.io/api/core/v1"
)

// Secret
func Secret(input *Input) (*Output, error) {
	output := Output{}
	var err error
	if input.Client == nil {
		return &output, errors.New(" No kubernetes client, cannot connect")
	}
	secret := input.Data.(*apiv1.Secret)
	if secret.Namespace == "" {
		secret.Namespace = "default"
	}
	switch input.Operation {
	case OpCreate:
		secret, err = input.Client.CoreV1().Secrets(secret.Namespace).Create(secret)
	case OpModify:
		secret, err = input.Client.CoreV1().Secrets(secret.Namespace).Update(secret)
	}
	output.Result = secret
	if err != nil {
		output.Verified = false
		return &output, err
	}
	return &output, nil
}
