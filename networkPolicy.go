package kates

import (
	"errors"
	netv1 "k8s.io/api/networking/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Create new NetworkPolicy
func CreateNetworkPolicy(input *Input) (Output, error) {

	var output Output
	if input.Client == nil {
		return output, errors.New(" No kubernetes client, cannot connect")
	}
	np := input.Data.(*netv1.NetworkPolicy)
	if np.Namespace == "" {
		input.Namespace = "default"
	}
	np, err := input.Client.NetworkingV1().NetworkPolicies(input.Namespace).Create(np)
	if err != nil {
		output.Verified = false
		err = eventMessages(input, &output, meta_v1.ListOptions{
			FieldSelector: "involvedObject.name="+ np.Name,
		})
		return output, err
	}
	return output, nil
}