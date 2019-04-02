package kates

import (
	"errors"
	apiv1 "k8s.io/api/core/v1"
)

// CreateReplicationController new RC
func CreateReplicationController(input *Input) (*Output, error) {

	output := Output{}
	if input.Client == nil {
		return &output, errors.New(" No kubernetes client, cannot connect")
	}
	rc := input.Data.(*apiv1.ReplicationController)
	if rc.Namespace == "" {
		input.Namespace = "default"
	}
	rc, err := input.Client.CoreV1().ReplicationControllers(input.Namespace).Create(rc)
	if err != nil {
		output.Verified = false
		return &output, err
	}
	return &output, nil
}
