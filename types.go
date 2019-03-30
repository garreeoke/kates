package kates

import (
	"k8s.io/client-go/kubernetes"
)

type Input struct {
	// Operation will be either Create, Modify, Delete, and Verify
	Operation string `json:"operation,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Data interface{} `json:"data,omitempty"`
	Client *kubernetes.Clientset
}

type Output struct {
	// Resulting Object from Kubernetes
	Result interface{} `json:"result,omitempty"`
	// Was the object able to be verified
	Verified bool `json:"verified,omitempty"`
	// Get events for the object
	Events []string `json:"events,omitempty"`
}
