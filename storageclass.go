package kates

import (
	"errors"
	storagev1 "k8s.io/api/storage/v1"
)

// StorageClass
func StorageClass(input *Input) (*Output, error) {
	output := Output{}
	var err error
	if input.Client == nil {
		return &output, errors.New(" No kubernetes client, cannot connect")
	}
	sc := input.Data.(*storagev1.StorageClass)
	switch input.Operation {
	case OpCreate:
		sc, err = input.Client.StorageV1().StorageClasses().Create(sc)
	case OpModify:
		sc, err = input.Client.StorageV1().StorageClasses().Update(sc)
	}
	output.Result = sc
	if err != nil {
		output.Verified = false
		return &output, err
	}
	return &output, nil
}