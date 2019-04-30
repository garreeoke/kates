package kates

import (
	"errors"
	netv1beta "k8s.io/api/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	case OpUpdate:
		ingress, err = input.Client.NetworkingV1beta1().Ingresses(ingress.Namespace).Update(ingress)
	case OpDynamic:
		_, err := input.Client.NetworkingV1beta1().Ingresses(ingress.Namespace).Get(ingress.Name, metav1.GetOptions{})
		if err != nil {
			ingress, err = input.Client.NetworkingV1beta1().Ingresses(ingress.Namespace).Create(ingress)
		} else {
			ingress, err = input.Client.NetworkingV1beta1().Ingresses(ingress.Namespace).Update(ingress)
		}
	}
	output.Result = ingress
	if err != nil {
		output.Verified = false
		return &output, err
	}
	return &output, nil
}