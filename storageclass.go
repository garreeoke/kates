package kates

import (
	"errors"
	storagev1 "k8s.io/api/storage/v1"
)

//CreateStorageClass new storage class
func CreateStorageClass(input *Input) (*Output, error) {

	output := Output{}
	if input.Client == nil {
		return &output, errors.New(" No kubernetes client, cannot connect")
	}
	sc := input.Data.(*storagev1.StorageClass)
	sc, err := input.Client.StorageV1().StorageClasses().Create(sc)
	output.Result = sc
	if err != nil {
		output.Verified = false
		return &output, err
	}
	return &output, nil
}
