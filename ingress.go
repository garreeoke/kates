package kates

import (
	"errors"
	netv1beta "k8s.io/api/networking/v1beta1"
)

// Ingress
func Ingress(input *Input) (*Output, error) {
	output := Output{}
	var err error
	if input.Client == nil {
		return &output, errors.New(" No kubernetes client, cannot connect")
	}
	ingress := input.Data.(*netv1beta.Ingress)
	if ingress.Namespace == "" {
		ingress.Namespace = "default"
	}
	switch input.Operation {
	case OpCreate:
		ingress, err = input.Client.NetworkingV1beta1().Ingresses(ingress.Namespace).Create(ingress)
	case OpModify:
		ingress, err = input.Client.NetworkingV1beta1().Ingresses(ingress.Namespace).Update(ingress)
	}
	output.Result = ingress
	if err != nil {
		output.Verified = false
		return &output, err
	}
	return &output, nil
}