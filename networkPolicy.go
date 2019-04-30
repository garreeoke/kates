package kates

import (
	"errors"
	netv1 "k8s.io/api/networking/v1"
)

// NetworkPolicy
func NetworkPolicy(input *Input) (*Output, error) {
	output := Output{}
	var err error
	if input.Client == nil {
		return &output, errors.New(" No kubernetes client, cannot connect")
	}
	np := input.Data.(*netv1.NetworkPolicy)
	if np.Namespace == "" {
		np.Namespace = "default"
	}
	switch input.Operation {
	case OpCreate:
		np, err = input.Client.NetworkingV1().NetworkPolicies(np.Namespace).Create(np)
	case OpModify:
		np, err = input.Client.NetworkingV1().NetworkPolicies(np.Namespace).Update(np)
	}
	output.Result = np
	if err != nil {
		output.Verified = false
		return &output, err
	}
	return &output, nil
}