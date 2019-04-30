package kates

import (
	"errors"
	appsv1 "k8s.io/api/apps/v1"
)

// DaemonSet
func DaemonSet(input *Input) (*Output, error) {
	output := Output{}
	var err error
	if input.Client == nil {
		return &output, errors.New(" No kubernetes client, cannot connect")
	}
	daemonSet := input.Data.(*appsv1.DaemonSet)
	if daemonSet.Namespace == "" {
		daemonSet.Namespace = "default"
	}
	switch input.Operation {
	case OpCreate:
		daemonSet, err = input.Client.AppsV1().DaemonSets(daemonSet.Namespace).Create(daemonSet)
	case OpModify:
		daemonSet, err = input.Client.AppsV1().DaemonSets(daemonSet.Namespace).Update(daemonSet)
	}
	output.Result = daemonSet
	if err != nil {
		output.Verified = false
		return &output, err
	}
	return &output, nil
}