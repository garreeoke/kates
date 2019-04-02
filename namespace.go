package kates

import (
	"errors"
	"time"

	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiv1 "k8s.io/api/core/v1"
)

// CreateNamespace new namespace
func CreateNamespace(input *Input) (*Output, error) {

	output := Output{}
	//var nameSpace apiv1.Namespace
	if input.Client == nil {
		return &output, errors.New(" No kubernetes client, cannot connect")
	}

	nameSpace := input.Data.(*apiv1.Namespace)
	nameSpace, err := input.Client.CoreV1().Namespaces().Create(nameSpace)
	if err != nil {
		return &output, err
	}

	// Verify
	for i := 1; i <= 10 && nameSpace.Status.Phase != "Active"; i++ {
		time.Sleep(1000 *time.Millisecond)
		nameSpace, err = input.Client.CoreV1().Namespaces().Get(nameSpace.Name, meta_v1.GetOptions{})
		if err != nil {
			return &output, nil
		}
	}
	if nameSpace.Status.Phase == "Active" {
		output.Verified = true
	} else {
		output.Verified = false
		// Get the events
		err = eventMessages(input, &output, meta_v1.ListOptions{
			FieldSelector: "involvedObject.name="+nameSpace.Name,
		})

	}
	output.Result = nameSpace

	return &output, nil

}
