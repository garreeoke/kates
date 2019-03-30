package kates

import (
	"errors"
	apiv1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Create new config
func CreateSecret(input *Input) (Output, error) {

	var output Output
	if input.Client == nil {
		return output, errors.New(" No kubernetes client, cannot connect")
	}
	secret := input.Data.(*apiv1.Secret)
	if secret.Namespace == "" {
		input.Namespace = "default"
	}
	secret, err := input.Client.CoreV1().Secrets(input.Namespace).Create(secret)
	if err != nil {
		output.Verified = false
		err = eventMessages(input, &output, meta_v1.ListOptions{
			FieldSelector: "involvedObject.name="+ secret.Name,
		})
		return output, err
	}
	return output, nil
}
