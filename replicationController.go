package kates

import (
	"errors"
	apiv1 "k8s.io/api/core/v1"
)


// ReplicationController
func ReplicationController(input *Input) (*Output, error) {
	output := Output{}
	var err error
	if input.Client == nil {
		return &output, errors.New(" No kubernetes client, cannot connect")
	}
	rc := input.Data.(*apiv1.ReplicationController)
	if rc.Namespace == "" {
		rc.Namespace = "default"
	}
	switch input.Operation {
	case OpCreate:
		rc, err = input.Client.CoreV1().ReplicationControllers(rc.Namespace).Create(rc)
	case OpModify:
		rc, err = input.Client.CoreV1().ReplicationControllers(rc.Namespace).Update(rc)
	}
	output.Result = rc
	if err != nil {
		output.Verified = false
		return &output, err
	}
	return &output, nil
}
