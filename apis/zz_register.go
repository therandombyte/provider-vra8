// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/therandombyte/provider-vra8/apis/blockdevice/v1alpha1"
	v1alpha1blueprint "github.com/therandombyte/provider-vra8/apis/blueprint/v1alpha1"
	v1alpha1catalogitementitlement "github.com/therandombyte/provider-vra8/apis/catalogitementitlement/v1alpha1"
	v1alpha1catalogsourceblueprint "github.com/therandombyte/provider-vra8/apis/catalogsourceblueprint/v1alpha1"
	v1alpha1catalogsourceentitlement "github.com/therandombyte/provider-vra8/apis/catalogsourceentitlement/v1alpha1"
	v1alpha1cloudaccount "github.com/therandombyte/provider-vra8/apis/cloudaccount/v1alpha1"
	v1alpha1contentsource "github.com/therandombyte/provider-vra8/apis/contentsource/v1alpha1"
	v1alpha1deployment "github.com/therandombyte/provider-vra8/apis/deployment/v1alpha1"
	v1alpha1fabric "github.com/therandombyte/provider-vra8/apis/fabric/v1alpha1"
	v1alpha1flavorprofile "github.com/therandombyte/provider-vra8/apis/flavorprofile/v1alpha1"
	v1alpha1imageprofile "github.com/therandombyte/provider-vra8/apis/imageprofile/v1alpha1"
	v1alpha1integration "github.com/therandombyte/provider-vra8/apis/integration/v1alpha1"
	v1alpha1loadbalancer "github.com/therandombyte/provider-vra8/apis/loadbalancer/v1alpha1"
	v1alpha1machine "github.com/therandombyte/provider-vra8/apis/machine/v1alpha1"
	v1alpha1network "github.com/therandombyte/provider-vra8/apis/network/v1alpha1"
	v1alpha1project "github.com/therandombyte/provider-vra8/apis/project/v1alpha1"
	v1alpha1storage "github.com/therandombyte/provider-vra8/apis/storage/v1alpha1"
	v1alpha1apis "github.com/therandombyte/provider-vra8/apis/v1alpha1"
	v1beta1 "github.com/therandombyte/provider-vra8/apis/v1beta1"
	v1alpha1zone "github.com/therandombyte/provider-vra8/apis/zone/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1blueprint.SchemeBuilder.AddToScheme,
		v1alpha1catalogitementitlement.SchemeBuilder.AddToScheme,
		v1alpha1catalogsourceblueprint.SchemeBuilder.AddToScheme,
		v1alpha1catalogsourceentitlement.SchemeBuilder.AddToScheme,
		v1alpha1cloudaccount.SchemeBuilder.AddToScheme,
		v1alpha1contentsource.SchemeBuilder.AddToScheme,
		v1alpha1deployment.SchemeBuilder.AddToScheme,
		v1alpha1fabric.SchemeBuilder.AddToScheme,
		v1alpha1flavorprofile.SchemeBuilder.AddToScheme,
		v1alpha1imageprofile.SchemeBuilder.AddToScheme,
		v1alpha1integration.SchemeBuilder.AddToScheme,
		v1alpha1loadbalancer.SchemeBuilder.AddToScheme,
		v1alpha1machine.SchemeBuilder.AddToScheme,
		v1alpha1network.SchemeBuilder.AddToScheme,
		v1alpha1project.SchemeBuilder.AddToScheme,
		v1alpha1storage.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
		v1beta1.SchemeBuilder.AddToScheme,
		v1alpha1zone.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
