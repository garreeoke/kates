package kates

import (
	"errors"
	apiv1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Create new deployment
func CreateServiceAccount(input *Input) (Output, error) {

	var output Output
	if input.Client == nil {
		return output, errors.New(" No kubernetes client, cannot connect")
	}
	sa := input.Data.(*apiv1.ServiceAccount)
	if sa.Namespace == "" {
		input.Namespace = "default"
	}
	sa, err := input.Client.CoreV1().ServiceAccounts(sa.Namespace).Create(sa)
	if err != nil {
		output.Verified = false
		err = eventMessages(input, &output, meta_v1.ListOptions{
			FieldSelector: "involvedObject.name="+ sa.Name,
		})
		return output, err
	}
	return output, nil
}
