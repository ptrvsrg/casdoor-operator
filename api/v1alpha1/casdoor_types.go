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

	// OrganizationName is the name of Casdoor organization
	// +kubebuilder:validation:Required
	OrganizationName string `json:"organizationName"`

	// ApplicationName is the name of Casdoor application
	// +kubebuilder:validation:Required
	ApplicationName string `json:"applicationName"`

	// ClientID is the client ID of Casdoor
	// +kubebuilder:validation:Required
	ClientID string `json:"clientID"`

	// ClientSecret is the secret selector of Casdoor client secret
	// +kubebuilder:validation:Required
	ClientSecret corev1.SecretKeySelector `json:"clientSecret"`

	// JwtCertificate is the secret selector of Casdoor JWT certificate
	// +kubebuilder:validation:Required
	JwtCertificate corev1.SecretKeySelector `json:"jwtCertificate"`
}

// +kubebuilder:object:root=true
// +kubebuilder:printcolumn:name="URL",type=string,JSONPath=`.spec.url`
// +kubebuilder:printcolumn:name="Age",type=string,JSONPath=`.metadata.creationTimestamp`
// +kubebuilder:resource:shortName={"casdoor"},categories=all;casdoor
// +kubebuilder:k8s:openapi-gen=true

// Casdoor is the Schema for the casdoors API
type Casdoor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec CasdoorSpec `json:"spec,omitempty"`
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
