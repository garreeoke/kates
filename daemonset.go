package kates

import (
	"errors"
	appsv1 "k8s.io/api/apps/v1"
)

//CreateDaemonSet new daemonset
func CreateDaemonSet(input *Input) (*Output, error) {

	output := Output{}
	if input.Client == nil {
		return &output, errors.New(" No kubernetes client, cannot connect")
	}

	ds := input.Data.(*appsv1.DaemonSet)
	if ds.Namespace == "" {
		ds.Namespace = "default"
	}
	ds, err := input.Client.AppsV1().DaemonSets(ds.Namespace).Create(ds)
	if err != nil {
		output.Verified = false
		return &output, err
	}
	return &output, nil
}