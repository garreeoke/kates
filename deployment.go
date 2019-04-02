package kates

import (
	"errors"
	appsv1 "k8s.io/api/apps/v1"
)

// CreateDeployment new deployment
func CreateDeployment(input *Input) (*Output, error) {

	output := Output{}
	if input.Client == nil {
		return &output, errors.New(" No kubernetes client, cannot connect")
	}
	deployment := input.Data.(*appsv1.Deployment)
	if deployment.Namespace == "" {
		input.Namespace = "default"
	}
	deployment, err := input.Client.AppsV1().Deployments(input.Namespace).Create(deployment)
	if err != nil {
		output.Verified = false
		/*
		err = eventMessages(input, &output, meta_v1.ListOptions{
			FieldSelector: "involvedObject.name="+deployment.Name,
		})
		*/
		return &output, err
	}
	return &output, nil
}

/*
// Verify Deployment
func verifyDeployment(deployment *appsv1.Deployment, client *kubernetes.Clientset) (bool,error) {

	// Query to find pods
	query := EquityLabelSelector(deployment.Spec.Selector.MatchLabels)
	verifyPods()

	tries := 1
	for tries <= 30 {
		deployment, err := client.AppsV1().Deployments(deployment.Namespace).Get(deployment.Name, meta_v1.GetOptions{})
		if err != nil {
			return false, err
		}
		if *deployment.Spec.Replicas != deployment.Status.AvailableReplicas && tries < 30 {
			time.Sleep(1000 * time.Millisecond)
			tries++
		} else if *deployment.Spec.Replicas != deployment.Status.AvailableReplicas && tries == 30 {
			return false, nil
			tries++
		} else if *deployment.Spec.Replicas == deployment.Status.AvailableReplicas {
			tries = 31
		}
	}

	return true, nil
}
*/