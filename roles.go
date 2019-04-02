package kates

import (
	"errors"
	rbacv1 "k8s.io/api/rbac/v1"
)

//ClusterRoles
//ClusterRoleBinding
//Roles
//RoleBindings

// CreateClusterRoles new cluster roles
func CreateClusterRoles(input *Input) (*Output, error) {

	output := Output{}
	if input.Client == nil {
		return &output, errors.New(" No kubernetes client, cannot connect")
	}
	crs := input.Data.(*rbacv1.ClusterRole)
	crs, err := input.Client.RbacV1().ClusterRoles().Create(crs)
	if err != nil {
		output.Verified = false
		return &output, err
	}
	return &output, nil
}

// CreateClusterRoleBindings new ClusterRoleBindings
func CreateClusterRoleBindings(input *Input) (*Output, error) {

	output := Output{}
	if input.Client == nil {
		return &output, errors.New(" No kubernetes client, cannot connect")
	}
	crb := input.Data.(*rbacv1.ClusterRoleBinding)
	crb, err := input.Client.RbacV1().ClusterRoleBindings().Create(crb)
	if err != nil {
		//output.Verified = false
		return &output, err
	}
	return &output, nil
}

//CreateRole new role
func CreateRole(input *Input) (*Output, error) {

	output := Output{}
	if input.Client == nil {
		return &output, errors.New(" No kubernetes client, cannot connect")
	}
	role := input.Data.(*rbacv1.Role)
	if role.Namespace == "" {
		input.Namespace = "default"
	}
	role, err := input.Client.RbacV1().Roles(input.Namespace).Create(role)
	if err != nil {
		output.Verified = false
		return &output, err
	}
	return &output, nil
}

// CreateRoleBindings new RoleBindings
func CreateRoleBindings(input *Input) (*Output, error) {

	var output *Output
	if input.Client == nil {
		return output, errors.New(" No kubernetes client, cannot connect")
	}
	roleBinding := input.Data.(*rbacv1.RoleBinding)
	if roleBinding.Namespace == "" {
		input.Namespace = "default"
	}
	roleBinding, err := input.Client.RbacV1().RoleBindings(input.Namespace).Create(roleBinding)
	if err != nil {
		output.Verified = false
		return output, err
	}
	return output, nil
}
