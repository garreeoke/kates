package kates

import (
	"errors"
	apiv1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Create new deployment
func CreateReplicationController(input *Input) (Output, error) {

	var output Output
	if input.Client == nil {
		return output, errors.New(" No kubernetes client, cannot connect")
	}
	rc := input.Data.(*apiv1.ReplicationController)
	if rc.Namespace == "" {
		input.Namespace = "default"
	}
	rc, err := input.Client.CoreV1().ReplicationControllers(rc.Namespace).Create(rc)
	if err != nil {
		output.Verified = false
		err = eventMessages(input, &output, meta_v1.ListOptions{
			FieldSelector: "involvedObject.name="+ rc.Name,
		})
		return output, err
	}
	return output, nil
}
