// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ProfileAzureInitParameters struct {
	DataDiskCaching *string `json:"dataDiskCaching,omitempty" tf:"data_disk_caching,omitempty"`

	DefaultItem *bool `json:"defaultItem,omitempty" tf:"default_item,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	DiskType *string `json:"diskType,omitempty" tf:"disk_type,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	OsDiskCaching *string `json:"osDiskCaching,omitempty" tf:"os_disk_caching,omitempty"`

	RegionID *string `json:"regionId,omitempty" tf:"region_id,omitempty"`

	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`

	SupportsEncryption *bool `json:"supportsEncryption,omitempty" tf:"supports_encryption,omitempty"`

	Tags []ProfileAzureTagsInitParameters `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ProfileAzureLinksInitParameters struct {
}

type ProfileAzureLinksObservation struct {
	Href *string `json:"href,omitempty" tf:"href,omitempty"`

	// +listType=set
	Hrefs []*string `json:"hrefs,omitempty" tf:"hrefs,omitempty"`

	Rel *string `json:"rel,omitempty" tf:"rel,omitempty"`
}

type ProfileAzureLinksParameters struct {
}

type ProfileAzureObservation struct {
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	DataDiskCaching *string `json:"dataDiskCaching,omitempty" tf:"data_disk_caching,omitempty"`

	DefaultItem *bool `json:"defaultItem,omitempty" tf:"default_item,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	DiskType *string `json:"diskType,omitempty" tf:"disk_type,omitempty"`

	ExternalRegionID *string `json:"externalRegionId,omitempty" tf:"external_region_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Links []ProfileAzureLinksObservation `json:"links,omitempty" tf:"links,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	OsDiskCaching *string `json:"osDiskCaching,omitempty" tf:"os_disk_caching,omitempty"`

	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	RegionID *string `json:"regionId,omitempty" tf:"region_id,omitempty"`

	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`

	SupportsEncryption *bool `json:"supportsEncryption,omitempty" tf:"supports_encryption,omitempty"`

	Tags []ProfileAzureTagsObservation `json:"tags,omitempty" tf:"tags,omitempty"`

	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type ProfileAzureParameters struct {

	// +kubebuilder:validation:Optional
	DataDiskCaching *string `json:"dataDiskCaching,omitempty" tf:"data_disk_caching,omitempty"`

	// +kubebuilder:validation:Optional
	DefaultItem *bool `json:"defaultItem,omitempty" tf:"default_item,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	DiskType *string `json:"diskType,omitempty" tf:"disk_type,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	OsDiskCaching *string `json:"osDiskCaching,omitempty" tf:"os_disk_caching,omitempty"`

	// +kubebuilder:validation:Optional
	RegionID *string `json:"regionId,omitempty" tf:"region_id,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`

	// +kubebuilder:validation:Optional
	SupportsEncryption *bool `json:"supportsEncryption,omitempty" tf:"supports_encryption,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []ProfileAzureTagsParameters `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ProfileAzureTagsInitParameters struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ProfileAzureTagsObservation struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ProfileAzureTagsParameters struct {

	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

// ProfileAzureSpec defines the desired state of ProfileAzure
type ProfileAzureSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProfileAzureParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ProfileAzureInitParameters `json:"initProvider,omitempty"`
}

// ProfileAzureStatus defines the observed state of ProfileAzure.
type ProfileAzureStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProfileAzureObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ProfileAzure is the Schema for the ProfileAzures API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vra8}
type ProfileAzure struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.defaultItem) || (has(self.initProvider) && has(self.initProvider.defaultItem))",message="spec.forProvider.defaultItem is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.regionId) || (has(self.initProvider) && has(self.initProvider.regionId))",message="spec.forProvider.regionId is a required parameter"
	Spec   ProfileAzureSpec   `json:"spec"`
	Status ProfileAzureStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProfileAzureList contains a list of ProfileAzures
type ProfileAzureList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProfileAzure `json:"items"`
}

// Repository type metadata.
var (
	ProfileAzure_Kind             = "ProfileAzure"
	ProfileAzure_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProfileAzure_Kind}.String()
	ProfileAzure_KindAPIVersion   = ProfileAzure_Kind + "." + CRDGroupVersion.String()
	ProfileAzure_GroupVersionKind = CRDGroupVersion.WithKind(ProfileAzure_Kind)
)

func init() {
	SchemeBuilder.Register(&ProfileAzure{}, &ProfileAzureList{})
}
