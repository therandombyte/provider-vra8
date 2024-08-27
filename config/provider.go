/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"
	blockdevice "github.com/therandombyte/provider-vra8/config/block_device"
	blueprint "github.com/therandombyte/provider-vra8/config/blueprint"
	catalogitementitlement "github.com/therandombyte/provider-vra8/config/catalog_item"
	catalogsource "github.com/therandombyte/provider-vra8/config/catalog_source"
	cloudaccount "github.com/therandombyte/provider-vra8/config/cloud_account"
	contentsource "github.com/therandombyte/provider-vra8/config/content_source"
	deployment "github.com/therandombyte/provider-vra8/config/deployment"
	fabric "github.com/therandombyte/provider-vra8/config/fabric"
	flavorprofile "github.com/therandombyte/provider-vra8/config/flavor_profile"
	imageprofile "github.com/therandombyte/provider-vra8/config/image_profile"
	integration "github.com/therandombyte/provider-vra8/config/integration"
	loadbalancer "github.com/therandombyte/provider-vra8/config/load_balancer"
	machine "github.com/therandombyte/provider-vra8/config/machine"
	network "github.com/therandombyte/provider-vra8/config/network"
	project "github.com/therandombyte/provider-vra8/config/project"
	storage "github.com/therandombyte/provider-vra8/config/storage"
	zone "github.com/therandombyte/provider-vra8/config/zone"
)

const (
	resourcePrefix = "vra8"
	modulePath     = "github.com/therandombyte/provider-vra8"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("vmware.vra"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		project.Configure,
		blueprint.Configure,
		deployment.Configure,
		fabric.Configure,
		blockdevice.Configure,
		flavorprofile.Configure,
		imageprofile.Configure,
		storage.Configure,
		catalogsource.Configure,
		catalogitementitlement.Configure,
		cloudaccount.Configure,
		contentsource.Configure,
		integration.Configure,
		loadbalancer.Configure,
		machine.Configure,
		network.Configure,
		zone.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
