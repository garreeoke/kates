package kates

import (
	"errors"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ConfigMap - ConfigMap operations
func ConfigMap(input *Input) (*Output, error) {
	output := Output{}
	var err error
	if input.Client == nil {
		return &output, errors.New(" No kubernetes client, cannot connect")
	}
	cm := input.Data.(*apiv1.ConfigMap)
	if cm.Namespace == "" {
		cm.Namespace = "default"
	}
	switch input.Operation {
	case OpCreate:
		cm, err = input.Client.CoreV1().ConfigMaps(cm.Namespace).Create(cm)
	case OpUpdate:
		cm, err = input.Client.CoreV1().ConfigMaps(cm.Namespace).Update(cm)
	case OpDynamic:
		_, err := input.Client.CoreV1().ConfigMaps(cm.Namespace).Get(cm.Name, metav1.GetOptions{})
		if err != nil {
			cm, err = input.Client.CoreV1().ConfigMaps(cm.Namespace).Create(cm)
		} else {
			cm, err = input.Client.CoreV1().ConfigMaps(cm.Namespace).Update(cm)
		}
	}
	output.Result = cm
	if err != nil {
		output.Verified = false
		return &output, err
	}
	return &output, nil
}