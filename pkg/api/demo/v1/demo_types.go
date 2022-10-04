package v1

import (
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DemoSpec struct {
    // Message is a simple string
	Message string `json:"message"`
}

type DemoStatus struct {
	Phase string `json:"phase"`
}

// Demo is a sample demo CRD
// +genclient
//+kubebuilder:object:root=true
//+kubebuilder:resource:path=demos,shortName=dm,categories=all
//+kubebuilder:subresource:status
type Demo struct {
    metav1.TypeMeta   `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`

    Spec   DemoSpec   `json:"spec,omitempty"`
    Status DemoStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true
type DemoList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata,omitempty"`
    Items           []Demo `json:"items"`
}
