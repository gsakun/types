// Code generated by kubetype-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	// Package-wide variables from generator "register".
	SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1alpha1"}
	//SchemeBuilder      = runtime.NewSchemeBuilder(addKnownTypes)
	//localSchemeBuilder = &SchemeBuilder
	//AddToScheme        = localSchemeBuilder.AddToScheme
)

const (
	// Package-wide consts from generator "register".
	GroupName = "rbac.istio.io"
)

func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&ClusterRbacConfig{},
		&ClusterRbacConfigList{},
		&RbacConfig{},
		&RbacConfigList{},
		&ServiceRole{},
		&ServiceRoleList{},
		&ServiceRoleBinding{},
		&ServiceRoleBindingList{},
	)
	v1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
