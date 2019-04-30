package kates

import (
	"errors"
	netv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	case OpUpdate:
		np, err = input.Client.NetworkingV1().NetworkPolicies(np.Namespace).Update(np)
	case OpDynamic:
		_, err := input.Client.NetworkingV1().NetworkPolicies(np.Namespace).Get(np.Name, metav1.GetOptions{})
		if err != nil {
			np, err = input.Client.NetworkingV1().NetworkPolicies(np.Namespace).Create(np)
		} else {
			np, err = input.Client.NetworkingV1().NetworkPolicies(np.Namespace).Update(np)
		}
	}
	output.Result = np
	if err != nil {
		output.Verified = false
		return &output, err
	}
	return &output, nil
}