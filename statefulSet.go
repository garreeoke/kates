package kates

import (
	"errors"
	appsv1 "k8s.io/api/apps/v1"
)

// CreateStatefulSet new statefulset
func CreateStatefulSet(input *Input) (*Output, error) {

	output := Output{}
	if input.Client == nil {
		return &output, errors.New(" No kubernetes client, cannot connect")
	}
	ss := input.Data.(*appsv1.StatefulSet)
	if ss.Namespace == "" {
		input.Namespace = "default"
	}
	ss, err := input.Client.AppsV1().StatefulSets(input.Namespace).Create(ss)
	if err != nil {
		output.Verified = false
		return &output, err
	}
	return &output, nil
}
