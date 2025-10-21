/*
Copyright 2025.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PhilipSpec defines the desired state of Philip
type PhilipSpec struct {
	Replicas int32 `json:"replicas"`
}

// PhilipStatus defines the observed state of Philip.
type PhilipStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`

	// conditions represent the current state of the Philip resource.
	// Each condition has a unique type and reflects the status of a specific aspect of the resource.
	//
	// Standard condition types include:
	// - "Available": the resource is fully functional
	// - "Progressing": the resource is being created or updated
	// - "Degraded": the resource failed to reach or maintain its desired state
	//
	// The status of each condition is one of True, False, or Unknown.
	// +listType=map
	// +listMapKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Philip is the Schema for the philips API
type Philip struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty,omitzero"`

	// spec defines the desired state of Philip
	// +required
	Spec PhilipSpec `json:"spec"`

	// status defines the observed state of Philip
	// +optional
	Status PhilipStatus `json:"status,omitempty,omitzero"`
}

// +kubebuilder:object:root=true

// PhilipList contains a list of Philip
type PhilipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Philip `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Philip{}, &PhilipList{})
}
