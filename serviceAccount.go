package kates

import (
	"errors"
	apiv1 "k8s.io/api/core/v1"
)

// ServiceAccount
func ServiceAccount(input *Input) (*Output, error) {
	output := Output{}
	var err error
	if input.Client == nil {
		return &output, errors.New(" No kubernetes client, cannot connect")
	}
	sa := input.Data.(*apiv1.ServiceAccount)
	if sa.Namespace == "" {
		sa.Namespace = "default"
	}
	switch input.Operation {
	case OpCreate:
		sa, err = input.Client.CoreV1().ServiceAccounts(sa.Namespace).Create(sa)
	case OpModify:
		sa, err = input.Client.CoreV1().ServiceAccounts(sa.Namespace).Update(sa)
	}
	output.Result = sa
	if err != nil {
		output.Verified = false
		return &output, err
	}
	return &output, nil
}
