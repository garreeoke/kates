package kates

import (
	"errors"
	apiv1 "k8s.io/api/core/v1"
)

// CreateConfigMap new config
func CreateConfigMap(input *Input) (*Output, error) {

	output := Output{}
	if input.Client == nil {
		return &output, errors.New(" No kubernetes client, cannot connect")
	}
	cm := input.Data.(*apiv1.ConfigMap)
	if cm.Namespace == "" {
		input.Namespace = "default"
	}
	cm, err := input.Client.CoreV1().ConfigMaps(input.Namespace).Create(cm)
	if err != nil {
		output.Verified = false
		return &output, err
	}
	return &output, nil
}
