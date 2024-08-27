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

type LinksInitParameters struct {
}

type LinksObservation struct {
	Href *string `json:"href,omitempty" tf:"href,omitempty"`

	// +listType=set
	Hrefs []*string `json:"hrefs,omitempty" tf:"hrefs,omitempty"`

	Rel *string `json:"rel,omitempty" tf:"rel,omitempty"`
}

type LinksParameters struct {
}

type TagsInitParameters struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TagsObservation struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TagsParameters struct {

	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type TagsToMatchInitParameters struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TagsToMatchObservation struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TagsToMatchParameters struct {

	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type ZoneInitParameters struct {

	// The ids of the compute resources that will be explicitly assigned to this zone.
	// +listType=set
	ComputeIds []*string `json:"computeIds,omitempty" tf:"compute_ids,omitempty"`

	// A list of key value pair of properties that will be used.
	// +mapType=granular
	CustomProperties map[string]*string `json:"customProperties,omitempty" tf:"custom_properties,omitempty"`

	// A human-friendly description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The folder relative path to the datacenter where resources are deployed to (only applicable for vSphere cloud zones).
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// A human-friendly name used as an identifier for the zone resource instance.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The placement policy for the zone. One of DEFAULT, SPREAD or BINPACK.
	PlacementPolicy *string `json:"placementPolicy,omitempty" tf:"placement_policy,omitempty"`

	// The id of the region for which this zone is created.
	RegionID *string `json:"regionId,omitempty" tf:"region_id,omitempty"`

	Tags []TagsInitParameters `json:"tags,omitempty" tf:"tags,omitempty"`

	TagsToMatch []TagsToMatchInitParameters `json:"tagsToMatch,omitempty" tf:"tags_to_match,omitempty"`
}

type ZoneObservation struct {

	// The ID of the cloud account this zone belongs to.
	CloudAccountID *string `json:"cloudAccountId,omitempty" tf:"cloud_account_id,omitempty"`

	// The ids of the compute resources that will be explicitly assigned to this zone.
	// +listType=set
	ComputeIds []*string `json:"computeIds,omitempty" tf:"compute_ids,omitempty"`

	// Date when the entity was created. The date is in ISO 8601 and UTC.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// A list of key value pair of properties that will be used.
	// +mapType=granular
	CustomProperties map[string]*string `json:"customProperties,omitempty" tf:"custom_properties,omitempty"`

	// A human-friendly description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The id of the region for which this zone is defined.
	ExternalRegionID *string `json:"externalRegionId,omitempty" tf:"external_region_id,omitempty"`

	// The folder relative path to the datacenter where resources are deployed to (only applicable for vSphere cloud zones).
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Links []LinksObservation `json:"links,omitempty" tf:"links,omitempty"`

	// A human-friendly name used as an identifier for the zone resource instance.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The id of the organization this entity belongs to.
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// Email of the user that owns the entity.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// The placement policy for the zone. One of DEFAULT, SPREAD or BINPACK.
	PlacementPolicy *string `json:"placementPolicy,omitempty" tf:"placement_policy,omitempty"`

	// The id of the region for which this zone is created.
	RegionID *string `json:"regionId,omitempty" tf:"region_id,omitempty"`

	Tags []TagsObservation `json:"tags,omitempty" tf:"tags,omitempty"`

	TagsToMatch []TagsToMatchObservation `json:"tagsToMatch,omitempty" tf:"tags_to_match,omitempty"`

	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type ZoneParameters struct {

	// The ids of the compute resources that will be explicitly assigned to this zone.
	// +kubebuilder:validation:Optional
	// +listType=set
	ComputeIds []*string `json:"computeIds,omitempty" tf:"compute_ids,omitempty"`

	// A list of key value pair of properties that will be used.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	CustomProperties map[string]*string `json:"customProperties,omitempty" tf:"custom_properties,omitempty"`

	// A human-friendly description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The folder relative path to the datacenter where resources are deployed to (only applicable for vSphere cloud zones).
	// +kubebuilder:validation:Optional
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// A human-friendly name used as an identifier for the zone resource instance.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The placement policy for the zone. One of DEFAULT, SPREAD or BINPACK.
	// +kubebuilder:validation:Optional
	PlacementPolicy *string `json:"placementPolicy,omitempty" tf:"placement_policy,omitempty"`

	// The id of the region for which this zone is created.
	// +kubebuilder:validation:Optional
	RegionID *string `json:"regionId,omitempty" tf:"region_id,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []TagsParameters `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	TagsToMatch []TagsToMatchParameters `json:"tagsToMatch,omitempty" tf:"tags_to_match,omitempty"`
}

// ZoneSpec defines the desired state of Zone
type ZoneSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ZoneParameters `json:"forProvider"`
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
	InitProvider ZoneInitParameters `json:"initProvider,omitempty"`
}

// ZoneStatus defines the observed state of Zone.
type ZoneStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ZoneObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Zone is the Schema for the Zones API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vra8}
type Zone struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.regionId) || (has(self.initProvider) && has(self.initProvider.regionId))",message="spec.forProvider.regionId is a required parameter"
	Spec   ZoneSpec   `json:"spec"`
	Status ZoneStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ZoneList contains a list of Zones
type ZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Zone `json:"items"`
}

// Repository type metadata.
var (
	Zone_Kind             = "Zone"
	Zone_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Zone_Kind}.String()
	Zone_KindAPIVersion   = Zone_Kind + "." + CRDGroupVersion.String()
	Zone_GroupVersionKind = CRDGroupVersion.WithKind(Zone_Kind)
)

func init() {
	SchemeBuilder.Register(&Zone{}, &ZoneList{})
}
