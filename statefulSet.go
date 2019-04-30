package kates

import (
	"errors"
	appsv1 "k8s.io/api/apps/v1"
)

// StatefulSet
func StatefulSet(input *Input) (*Output, error) {
	output := Output{}
	var err error
	if input.Client == nil {
		return &output, errors.New(" No kubernetes client, cannot connect")
	}
	ss := input.Data.(*appsv1.StatefulSet)
	if ss.Namespace == "" {
		ss.Namespace = "default"
	}
	switch input.Operation {
	case OpCreate:
		ss, err = input.Client.AppsV1().StatefulSets(ss.Namespace).Create(ss)
	case OpModify:
		ss, err = input.Client.AppsV1().StatefulSets(ss.Namespace).Update(ss)
	}
	output.Result = ss
	if err != nil {
		output.Verified = false
		return &output, err
	}
	return &output, nil
}
