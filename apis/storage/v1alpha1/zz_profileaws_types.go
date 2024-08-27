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

type ProfileAwsInitParameters struct {
	DefaultItem *bool `json:"defaultItem,omitempty" tf:"default_item,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	DeviceType *string `json:"deviceType,omitempty" tf:"device_type,omitempty"`

	Iops *string `json:"iops,omitempty" tf:"iops,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	RegionID *string `json:"regionId,omitempty" tf:"region_id,omitempty"`

	SupportsEncryption *bool `json:"supportsEncryption,omitempty" tf:"supports_encryption,omitempty"`

	Tags []ProfileAwsTagsInitParameters `json:"tags,omitempty" tf:"tags,omitempty"`

	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type ProfileAwsLinksInitParameters struct {
}

type ProfileAwsLinksObservation struct {
	Href *string `json:"href,omitempty" tf:"href,omitempty"`

	// +listType=set
	Hrefs []*string `json:"hrefs,omitempty" tf:"hrefs,omitempty"`

	Rel *string `json:"rel,omitempty" tf:"rel,omitempty"`
}

type ProfileAwsLinksParameters struct {
}

type ProfileAwsObservation struct {
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	DefaultItem *bool `json:"defaultItem,omitempty" tf:"default_item,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	DeviceType *string `json:"deviceType,omitempty" tf:"device_type,omitempty"`

	ExternalRegionID *string `json:"externalRegionId,omitempty" tf:"external_region_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Iops *string `json:"iops,omitempty" tf:"iops,omitempty"`

	Links []ProfileAwsLinksObservation `json:"links,omitempty" tf:"links,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	RegionID *string `json:"regionId,omitempty" tf:"region_id,omitempty"`

	SupportsEncryption *bool `json:"supportsEncryption,omitempty" tf:"supports_encryption,omitempty"`

	Tags []ProfileAwsTagsObservation `json:"tags,omitempty" tf:"tags,omitempty"`

	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`

	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type ProfileAwsParameters struct {

	// +kubebuilder:validation:Optional
	DefaultItem *bool `json:"defaultItem,omitempty" tf:"default_item,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	DeviceType *string `json:"deviceType,omitempty" tf:"device_type,omitempty"`

	// +kubebuilder:validation:Optional
	Iops *string `json:"iops,omitempty" tf:"iops,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	RegionID *string `json:"regionId,omitempty" tf:"region_id,omitempty"`

	// +kubebuilder:validation:Optional
	SupportsEncryption *bool `json:"supportsEncryption,omitempty" tf:"supports_encryption,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []ProfileAwsTagsParameters `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type ProfileAwsTagsInitParameters struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ProfileAwsTagsObservation struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ProfileAwsTagsParameters struct {

	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

// ProfileAwsSpec defines the desired state of ProfileAws
type ProfileAwsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProfileAwsParameters `json:"forProvider"`
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
	InitProvider ProfileAwsInitParameters `json:"initProvider,omitempty"`
}

// ProfileAwsStatus defines the observed state of ProfileAws.
type ProfileAwsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProfileAwsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ProfileAws is the Schema for the ProfileAwss API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vra8}
type ProfileAws struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.defaultItem) || (has(self.initProvider) && has(self.initProvider.defaultItem))",message="spec.forProvider.defaultItem is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.regionId) || (has(self.initProvider) && has(self.initProvider.regionId))",message="spec.forProvider.regionId is a required parameter"
	Spec   ProfileAwsSpec   `json:"spec"`
	Status ProfileAwsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProfileAwsList contains a list of ProfileAwss
type ProfileAwsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProfileAws `json:"items"`
}

// Repository type metadata.
var (
	ProfileAws_Kind             = "ProfileAws"
	ProfileAws_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProfileAws_Kind}.String()
	ProfileAws_KindAPIVersion   = ProfileAws_Kind + "." + CRDGroupVersion.String()
	ProfileAws_GroupVersionKind = CRDGroupVersion.WithKind(ProfileAws_Kind)
)

func init() {
	SchemeBuilder.Register(&ProfileAws{}, &ProfileAwsList{})
}
