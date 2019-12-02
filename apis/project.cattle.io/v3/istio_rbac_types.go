package v3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// Disable Istio RBAC completely, Istio RBAC policies will not be enforced.
	OFF RbacConfigMode = 0
	// Enable Istio RBAC for all services and namespaces. Note Istio RBAC is deny-by-default
	// which means all requests will be denied if it's not allowed by RBAC rules.
	ON RbacConfigMode = 1
	// Enable Istio RBAC only for services and namespaces specified in the inclusion field. Any other
	// services and namespaces not in the inclusion field will not be enforced by Istio RBAC policies.
	ON_WITH_INCLUSION RbacConfigMode = 2
	// Enable Istio RBAC for all services and namespaces except those specified in the exclusion field. Any other
	// services and namespaces not in the exclusion field will be enforced by Istio RBAC policies.
	ON_WITH_EXCLUSION RbacConfigMode = 3
)

type RbacConfigMode int32

type ClusterRbacConfig struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec RbacConfigSpec `json:"spec,omitempty"`
}

type RbacConfigSpec struct {
	// Istio RBAC mode.
	Mode RbacConfigMode `json:"mode,omitempty"`
	// A list of services or namespaces that should be enforced by Istio RBAC policies. Note: This field have
	// effect only when mode is ON_WITH_INCLUSION and will be ignored for any other modes.
	Inclusion RbacConfigTarget `json:"inclusion,omitempty"`
	// A list of services or namespaces that should not be enforced by Istio RBAC policies. Note: This field have
	// effect only when mode is ON_WITH_EXCLUSION and will be ignored for any other modes.
	Exclusion RbacConfigTarget `json:"exclusion,omitempty"`

	EnforcementMode EnforcementMode `json:"enforcementMode,omitempty"`
}

// Target defines a list of services or namespaces.
type RbacConfigTarget struct {
	// A list of services.
	Services []string `json:"services,omitempty"`
	// A list of namespaces.
	Namespaces []string `json:"namespaces,omitempty"`
}

type EnforcementMode int32

const (
	// Policy in ENFORCED mode has impact on user experience.
	// Policy is in ENFORCED mode by default.
	ENFORCED EnforcementMode = 0
	// Policy in PERMISSIVE mode isn't enforced and has no impact on users.
	// RBAC engine run policies in PERMISSIVE mode and logs stats.
	PERMISSIVE EnforcementMode = 1
)

