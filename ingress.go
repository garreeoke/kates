package kates

import (
	"errors"
	netv1beta "k8s.io/api/networking/v1beta1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Create new Ingress
func CreateIngress(input *Input) (Output, error) {

	var output Output
	if input.Client == nil {
		return output, errors.New(" No kubernetes client, cannot connect")
	}
	ing := input.Data.(*netv1beta.Ingress)
	if ing.Namespace == "" {
		input.Namespace = "default"
	}
	ing, err := input.Client.NetworkingV1beta1().Ingresses(input.Namespace).Create(ing)
	if err != nil {
		output.Verified = false
		err = eventMessages(input, &output, meta_v1.ListOptions{
			FieldSelector: "involvedObject.name="+ ing.Name,
		})
		return output, err
	}
	return output, nil
}