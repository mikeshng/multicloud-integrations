/*
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

package v1beta1

import (
	"crypto/tls"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

const (
	// TLS minimum version as an integer
	TLSMinVersionInt = tls.VersionTLS12
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope="Namespaced"

// The GitOpsCluster uses placement to import selected managed clusters into the Argo CD.
type GitOpsCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   GitOpsClusterSpec   `json:"spec"`
	Status GitOpsClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GitOpsClusterSpec defines the desired state of GitOpsCluster.
type GitOpsClusterSpec struct {
	ArgoServer ArgoServerSpec `json:"argoServer"`

	PlacementRef *corev1.ObjectReference `json:"placementRef"`

	// ManagedServiceAccountRef defines managed service account in the managed cluster namespace used to create the ArgoCD cluster secret.
	ManagedServiceAccountRef string `json:"managedServiceAccountRef,omitempty"`

	// Internally used.
	CreateBlankClusterSecrets *bool `json:"createBlankClusterSecrets,omitempty"`

	// Create default policy template if it is true.
	CreatePolicyTemplate *bool `json:"createPolicyTemplate,omitempty"`
}

// ArgoServerSpec specifies the location of the Argo CD server.
type ArgoServerSpec struct {
	// Not used and reserved for defining a managed cluster name.
	Cluster string `json:"cluster,omitempty"`

	// ArgoNamespace is the namespace in which the Argo CD server is installed.
	ArgoNamespace string `json:"argoNamespace"`
}

// +kubebuilder:object:root=true

// GitOpsClusterStatus defines the observed state of GitOpsCluster.
type GitOpsClusterStatus struct {
	// LastUpdateTime provides the last updated timestamp of the gitOpsCluster status
	LastUpdateTime metav1.Time `json:"lastUpdateTime,omitempty"`

	// Message provides the detailed message of the GitOpsCluster status.
	Message string `json:"message,omitempty"`

	// Phase provides the overall phase of the GitOpsCluster status. Valid values include failed or successful.
	Phase string `json:"phase,omitempty"`
}

// +kubebuilder:object:root=true

// GitOpsClusterList providess a list of GitOpsClusters.
type GitOpsClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GitOpsCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GitOpsCluster{}, &GitOpsClusterList{})
}
