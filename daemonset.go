package kates

import (
	"errors"
	appsv1 "k8s.io/api/apps/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)
func CreateDaemonSet(input *Input) (Output, error) {

	var output Output
	if input.Client == nil {
		return output, errors.New(" No kubernetes client, cannot connect")
	}

	ds := input.Data.(*appsv1.DaemonSet)
	if ds.Namespace == "" {
		input.Namespace = "default"
	}
	ds, err := input.Client.AppsV1().DaemonSets(input.Namespace).Create(ds)
	if err != nil {
		output.Verified = false
		err = eventMessages(input, &output, meta_v1.ListOptions{
			FieldSelector: "involvedObject.name="+ds.Name,
		})
		return output, err
	}
	return output, nil
}