/*
Copyright 2025 ptrvsrg.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/ptrvsrg/casdoor-operator/api"
)

// CasdoorSpec defines the desired state of Casdoor
type CasdoorSpec struct {
	// URL of Casdoor instance
	// +kubebuilder:validation:Required
	URL string `json:"url"`

	// AdminCredentialsSecret is the secret name of Casdoor admin credentials
	// +kubebuilder:validation:Required
	AdminCredentialsSecret corev1.LocalObjectReference `json:"adminCredentialsSecret"`

	// Healthcheck is the configuration of healthcheck
	// +kubebuilder:validation:Optional
	Healthcheck *CasdoorHealthcheckSpec `json:"healthcheck,omitempty"`
}

type CasdoorHealthcheckSpec struct {
	// Enabled is the flag to enable healthcheck
	// +kubebuilder:default=false
	Enabled bool `json:"enabled"`

	// Method of healthcheck
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=GET;POST;PUT;PATCH;DELETE;CONNECT;OPTIONS;TRACE;HEAD
	// +kubebuilder:default="GET"
	Method string `json:"method"`

	// Path of healthcheck
	// +kubebuilder:validation:Optional
	// +kubebuilder:default="/"
	Path string `json:"path"`

	// Interval of healthcheck
	// +kubebuilder:validation:Optional
	// +kubebuilder:default="1m"
	Interval metav1.Duration `json:"interval"`

	// Timeout of healthcheck
	// +kubebuilder:validation:Optional
	// +kubebuilder:default="1m"
	Timeout metav1.Duration `json:"timeout"`

	// Retries of healthcheck
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:default=3
	Retries int `json:"retries"`
}

type CasdoorStatusCode string

const (
	CasdoorStatusReady  CasdoorStatusCode = "Ready"
	CasdoorStatusFailed CasdoorStatusCode = "Failed"
)

// CasdoorStatus defines the observed state of Casdoor
type CasdoorStatus struct {
	// Code of Casdoor instance status
	Code CasdoorStatusCode `json:"code"`

	// Reason of Casdoor instance failure
	Reason string `json:"reason,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type=string,JSONPath=`.metadata.creationTimestamp`
// +kubebuilder:printcolumn:name="Status",type=string,JSONPath=`.status.code`
// +kubebuilder:resource:shortName={"casdoor"},categories=all;casdoor
// +kubebuilder:k8s:openapi-gen=true

// Casdoor is the Schema for the casdoors API
type Casdoor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CasdoorSpec   `json:"spec,omitempty"`
	Status CasdoorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CasdoorList contains a list of Casdoor
type CasdoorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Casdoor `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Casdoor{}, &CasdoorList{})
}

func (c *Casdoor) GetGroup() string {
	return GroupName
}

func (c *Casdoor) GetVersion() string {
	return VersionName
}

func (c *Casdoor) GetResourceKind() api.ResourceKind {
	return CasdoorKind
}
