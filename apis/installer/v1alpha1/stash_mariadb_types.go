/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Free Trial License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Free-Trial-1.0.0.md

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

const (
	ResourceKindStashMariadb = "StashMariadb"
	ResourceStashMariadb     = "stashmariadb"
	ResourceStashMariadbs    = "stashmariadbs"
)

// StashMariadb defines the schama for Stash MariaDB Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=stashmariadbs,singular=stashmariadb,categories={stash,appscode}
type StashMariadb struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StashMariadbSpec `json:"spec,omitempty"`
}

// StashMariadbSpec is the schema for Stash MariaDB values file
type StashMariadbSpec struct {
	// +optional
	NameOverride string `json:"nameOverride"`
	// +optional
	FullnameOverride string         `json:"fullnameOverride"`
	Image            ImageRef       `json:"image"`
	Backup           MariaDBBackup  `json:"backup"`
	Restore          MariaDBRestore `json:"restore"`
	WaitTimeout      int64          `json:"waitTimeout"`
}

type ImageRef struct {
	Registry   string `json:"registry"`
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
}

type MariaDBBackup struct {
	// +optional
	Args string `json:"args"`
}

type MariaDBRestore struct {
	// +optional
	Args string `json:"args"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StashMariadbList is a list of StashMariadbs
type StashMariadbList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StashMariadb CRD objects
	Items []StashMariadb `json:"items,omitempty"`
}
