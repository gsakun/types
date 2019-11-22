package v3

import (
	"github.com/rancher/norman/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	managementv3 "github.com/rancher/types/apis/management.cattle.io/v3"
)

type Application struct {
	types.Namespaced
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	
	Spec ApplicationSpec `json:"spec,omitempty"`
	Status ApplicationStatus `json:"spec,omitempty"`
}

type ApplicationSpec struct {
	Components []managementv3.Component `json:"components"`
}

type ApplicationStatus struct{}