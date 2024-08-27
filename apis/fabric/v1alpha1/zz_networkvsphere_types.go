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

type NetworkVsphereInitParameters struct {
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	DNSSearchDomains []*string `json:"dnsSearchDomains,omitempty" tf:"dns_search_domains,omitempty"`

	DNSServerAddresses []*string `json:"dnsServerAddresses,omitempty" tf:"dns_server_addresses,omitempty"`

	DefaultGateway *string `json:"defaultGateway,omitempty" tf:"default_gateway,omitempty"`

	DefaultIPv6Gateway *string `json:"defaultIpv6Gateway,omitempty" tf:"default_ipv6_gateway,omitempty"`

	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	IPv6Cidr *string `json:"ipv6Cidr,omitempty" tf:"ipv6_cidr,omitempty"`

	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default,omitempty"`

	IsPublic *bool `json:"isPublic,omitempty" tf:"is_public,omitempty"`

	Tags []NetworkVsphereTagsInitParameters `json:"tags,omitempty" tf:"tags,omitempty"`
}

type NetworkVsphereLinksInitParameters struct {
}

type NetworkVsphereLinksObservation struct {
	Href *string `json:"href,omitempty" tf:"href,omitempty"`

	// +listType=set
	Hrefs []*string `json:"hrefs,omitempty" tf:"hrefs,omitempty"`

	Rel *string `json:"rel,omitempty" tf:"rel,omitempty"`
}

type NetworkVsphereLinksParameters struct {
}

type NetworkVsphereObservation struct {
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// +listType=set
	CloudAccountIds []*string `json:"cloudAccountIds,omitempty" tf:"cloud_account_ids,omitempty"`

	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// +mapType=granular
	CustomProperties map[string]*string `json:"customProperties,omitempty" tf:"custom_properties,omitempty"`

	DNSSearchDomains []*string `json:"dnsSearchDomains,omitempty" tf:"dns_search_domains,omitempty"`

	DNSServerAddresses []*string `json:"dnsServerAddresses,omitempty" tf:"dns_server_addresses,omitempty"`

	DefaultGateway *string `json:"defaultGateway,omitempty" tf:"default_gateway,omitempty"`

	DefaultIPv6Gateway *string `json:"defaultIpv6Gateway,omitempty" tf:"default_ipv6_gateway,omitempty"`

	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	ExternalID *string `json:"externalId,omitempty" tf:"external_id,omitempty"`

	ExternalRegionID *string `json:"externalRegionId,omitempty" tf:"external_region_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IPv6Cidr *string `json:"ipv6Cidr,omitempty" tf:"ipv6_cidr,omitempty"`

	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default,omitempty"`

	IsPublic *bool `json:"isPublic,omitempty" tf:"is_public,omitempty"`

	Links []NetworkVsphereLinksObservation `json:"links,omitempty" tf:"links,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	Tags []NetworkVsphereTagsObservation `json:"tags,omitempty" tf:"tags,omitempty"`

	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type NetworkVsphereParameters struct {

	// +kubebuilder:validation:Optional
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// +kubebuilder:validation:Optional
	DNSSearchDomains []*string `json:"dnsSearchDomains,omitempty" tf:"dns_search_domains,omitempty"`

	// +kubebuilder:validation:Optional
	DNSServerAddresses []*string `json:"dnsServerAddresses,omitempty" tf:"dns_server_addresses,omitempty"`

	// +kubebuilder:validation:Optional
	DefaultGateway *string `json:"defaultGateway,omitempty" tf:"default_gateway,omitempty"`

	// +kubebuilder:validation:Optional
	DefaultIPv6Gateway *string `json:"defaultIpv6Gateway,omitempty" tf:"default_ipv6_gateway,omitempty"`

	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// +kubebuilder:validation:Optional
	IPv6Cidr *string `json:"ipv6Cidr,omitempty" tf:"ipv6_cidr,omitempty"`

	// +kubebuilder:validation:Optional
	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default,omitempty"`

	// +kubebuilder:validation:Optional
	IsPublic *bool `json:"isPublic,omitempty" tf:"is_public,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []NetworkVsphereTagsParameters `json:"tags,omitempty" tf:"tags,omitempty"`
}

type NetworkVsphereTagsInitParameters struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type NetworkVsphereTagsObservation struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type NetworkVsphereTagsParameters struct {

	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

// NetworkVsphereSpec defines the desired state of NetworkVsphere
type NetworkVsphereSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkVsphereParameters `json:"forProvider"`
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
	InitProvider NetworkVsphereInitParameters `json:"initProvider,omitempty"`
}

// NetworkVsphereStatus defines the observed state of NetworkVsphere.
type NetworkVsphereStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkVsphereObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// NetworkVsphere is the Schema for the NetworkVspheres API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vra8}
type NetworkVsphere struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkVsphereSpec   `json:"spec"`
	Status            NetworkVsphereStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkVsphereList contains a list of NetworkVspheres
type NetworkVsphereList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkVsphere `json:"items"`
}

// Repository type metadata.
var (
	NetworkVsphere_Kind             = "NetworkVsphere"
	NetworkVsphere_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NetworkVsphere_Kind}.String()
	NetworkVsphere_KindAPIVersion   = NetworkVsphere_Kind + "." + CRDGroupVersion.String()
	NetworkVsphere_GroupVersionKind = CRDGroupVersion.WithKind(NetworkVsphere_Kind)
)

func init() {
	SchemeBuilder.Register(&NetworkVsphere{}, &NetworkVsphereList{})
}
