package kates

import (
	"errors"
	netv1 "k8s.io/api/networking/v1"
)

// CreateNewtorkPolicy new NetworkPolicy
func CreateNetworkPolicy(input *Input) (*Output, error) {

	output := Output{}
	if input.Client == nil {
		return &output, errors.New(" No kubernetes client, cannot connect")
	}
	np := input.Data.(*netv1.NetworkPolicy)
	if np.Namespace == "" {
		np.Namespace = "default"
	}
	np, err := input.Client.NetworkingV1().NetworkPolicies(np.Namespace).Create(np)
	if err != nil {
		output.Verified = false
		return &output, err
	}
	return &output, nil
}