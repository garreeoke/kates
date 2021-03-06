package kates

import (
	"errors"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//ClusterRoles
//ClusterRoleBinding
//Roles
//RoleBindings

// ClusterRoles new cluster roles
func ClusterRoles(input *Input) (*Output, error) {

	output := Output{}
	var err error
	if input.Client == nil {
		return &output, errors.New(" No kubernetes client, cannot connect")
	}
	cr := input.Data.(*rbacv1.ClusterRole)
	if cr.Namespace == "" {
		cr.Namespace = "default"
	}
	switch input.Operation {
	case OpCreate:
		cr, err = input.Client.RbacV1().ClusterRoles().Create(cr)
	case OpUpdate:
		cr, err = input.Client.RbacV1().ClusterRoles().Update(cr)
	case OpDynamic:
		_, err := input.Client.RbacV1().ClusterRoles().Get(cr.Name, metav1.GetOptions{})
		if err != nil {
			cr, err = input.Client.RbacV1().ClusterRoles().Create(cr)
		} else {
			cr, err = input.Client.RbacV1().ClusterRoles().Update(cr)
		}
	}
	output.Result = cr
	if err != nil {
		output.Verified = false
		return &output, err
	}
	return &output, nil
}

// CreateClusterRoleBindings new ClusterRoleBindings
func ClusterRoleBindings(input *Input) (*Output, error) {

	output := Output{}
	var err error
	if input.Client == nil {
		return &output, errors.New(" No kubernetes client, cannot connect")
	}
	crb := input.Data.(*rbacv1.ClusterRoleBinding)
	if crb.Namespace == "" {
		crb.Namespace = "default"
	}
	switch input.Operation {
	case OpCreate:
		crb, err = input.Client.RbacV1().ClusterRoleBindings().Create(crb)
	case OpUpdate:
		crb, err = input.Client.RbacV1().ClusterRoleBindings().Update(crb)
	case OpDynamic:
		_, err := input.Client.RbacV1().ClusterRoleBindings().Get(crb.Name, metav1.GetOptions{})
		if err != nil {
			crb, err = input.Client.RbacV1().ClusterRoleBindings().Create(crb)
		} else {
			crb, err = input.Client.RbacV1().ClusterRoleBindings().Update(crb)
		}
	}
	output.Result = crb
	if err != nil {
		output.Verified = false
		return &output, err
	}
	return &output, nil
}

//Role new role
func Role(input *Input) (*Output, error) {

	output := Output{}
	var err error
	if input.Client == nil {
		return &output, errors.New(" No kubernetes client, cannot connect")
	}
	role := input.Data.(*rbacv1.Role)
	if role.Namespace == "" {
		role.Namespace = "default"
	}
	switch input.Operation {
	case OpCreate:
		role, err = input.Client.RbacV1().Roles(role.Namespace).Create(role)
	case OpUpdate:
		role, err = input.Client.RbacV1().Roles(role.Namespace).Update(role)
	case OpDynamic:
		_, err := input.Client.RbacV1().Roles(role.Namespace).Get(role.Name, metav1.GetOptions{})
		if err != nil {
			role, err = input.Client.RbacV1().Roles(role.Namespace).Create(role)
		} else {
			role, err = input.Client.RbacV1().Roles(role.Namespace).Update(role)
		}
	}
	output.Result = role
	if err != nil {
		output.Verified = false
		return &output, err
	}
	return &output, nil
}

// RoleBindings new RoleBindings
func RoleBindings(input *Input) (*Output, error) {

	output := Output{}
	var err error
	if input.Client == nil {
		return &output, errors.New(" No kubernetes client, cannot connect")
	}
	rb := input.Data.(*rbacv1.RoleBinding)
	if rb.Namespace == "" {
		rb.Namespace = "default"
	}
	switch input.Operation {
	case OpCreate:
		rb, err = input.Client.RbacV1().RoleBindings(rb.Namespace).Create(rb)
	case OpUpdate:
		rb, err = input.Client.RbacV1().RoleBindings(rb.Namespace).Update(rb)
	case OpDynamic:
		_, err := input.Client.RbacV1().RoleBindings(rb.Namespace).Get(rb.Name, metav1.GetOptions{})
		if err != nil {
			rb, err = input.Client.RbacV1().RoleBindings(rb.Namespace).Create(rb)
		} else {
			rb, err = input.Client.RbacV1().RoleBindings(rb.Namespace).Update(rb)
		}
	}
	output.Result = rb
	if err != nil {
		output.Verified = false
		return &output, err
	}
	return &output, nil
}
