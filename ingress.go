package kates

import (
	"errors"
	netv1beta "k8s.io/api/networking/v1beta1"
)

// CreateIngress new Ingress
func CreateIngress(input *Input) (*Output, error) {

	output := Output{}
	if input.Client == nil {
		return &output, errors.New(" No kubernetes client, cannot connect")
	}
	ing := input.Data.(*netv1beta.Ingress)
	if ing.Namespace == "" {
		input.Namespace = "default"
	}
	ing, err := input.Client.NetworkingV1beta1().Ingresses(input.Namespace).Create(ing)
	if err != nil {
		output.Verified = false
		return &output, err
	}
	return &output, nil
}