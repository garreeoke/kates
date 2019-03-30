package kates

import (
	"errors"
	appsv1 "k8s.io/api/apps/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Create new StatefulSet
func CreateStatefulSet(input *Input) (Output, error) {

	var output Output
	if input.Client == nil {
		return output, errors.New(" No kubernetes client, cannot connect")
	}
	ss := input.Data.(*appsv1.StatefulSet)
	if ss.Namespace == "" {
		input.Namespace = "default"
	}
	ss, err := input.Client.AppsV1().StatefulSets(input.Namespace).Create(ss)
	if err != nil {
		output.Verified = false
		err = eventMessages(input, &output, meta_v1.ListOptions{
			FieldSelector: "involvedObject.name="+ ss.Name,
		})
		return output, err
	}
	return output, nil
}
