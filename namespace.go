package kates

import (
	"errors"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiv1 "k8s.io/api/core/v1"
)

// Namespace
func Namespace(input *Input) (*Output, error) {
	output := Output{}
	var err error
	if input.Client == nil {
		return &output, errors.New(" No kubernetes client, cannot connect")
	}
	nameSpace := input.Data.(*apiv1.Namespace)
	switch input.Operation {
	case OpCreate:
		nameSpace, err = input.Client.CoreV1().Namespaces().Create(nameSpace)
	case OpUpdate:
		nameSpace, err = input.Client.CoreV1().Namespaces().Update(nameSpace)
	case OpDynamic:
		_, err := input.Client.CoreV1().Namespaces().Get(nameSpace.Name, metav1.GetOptions{})
		if err != nil {
			nameSpace, err = input.Client.CoreV1().Namespaces().Create(nameSpace)
		} else {
			nameSpace, err = input.Client.CoreV1().Namespaces().Update(nameSpace)
		}
	}
	output.Result = nameSpace
	if err != nil {
		output.Verified = false
		return &output, err
	}
	return &output, nil
}

// VerifyNamespace
func VerifyNamespace(input *Input) (*Output, error) {

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
		nameSpace, err = input.Client.CoreV1().Namespaces().Get(nameSpace.Name, metav1.GetOptions{})
		if err != nil {
			return &output, nil
		}
	}
	if nameSpace.Status.Phase == "Active" {
		output.Verified = true
	} else {
		output.Verified = false
		// Get the events
		err = eventMessages(input, &output, metav1.ListOptions{
			FieldSelector: "involvedObject.name="+nameSpace.Name,
		})

	}
	output.Result = nameSpace

	return &output, nil

}
