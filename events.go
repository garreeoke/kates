package kates

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func eventMessages(input *Input, output *Output, options meta_v1.ListOptions) error {

	eventList, err := input.Client.CoreV1().Events(input.Namespace).List(options)
	if err != nil {
		return err
	}
	for _,e := range eventList.Items {
		// Potential crash here
		output.Events = append(output.Events, e.Message)
	}
	return nil
}
