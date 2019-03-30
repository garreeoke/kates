package kates

import (
	"errors"
	storagev1 "k8s.io/api/storage/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)
func CreateStorageClass(input *Input) (Output, error) {

	var output Output
	if input.Client == nil {
		return output, errors.New(" No kubernetes client, cannot connect")
	}

	sc := input.Data.(*storagev1.StorageClass)
	sc, err := input.Client.StorageV1().StorageClasses().Create(sc)
	if err != nil {
		output.Verified = false
		err = eventMessages(input, &output, meta_v1.ListOptions{
			FieldSelector: "involvedObject.name="+sc.Name,
		})
		return output, err
	}
	return output, nil
}
