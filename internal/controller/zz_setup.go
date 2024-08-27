// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	blockdevice "github.com/therandombyte/provider-vra8/internal/controller/blockdevice/blockdevice"
	blockdevicesnapshot "github.com/therandombyte/provider-vra8/internal/controller/blockdevice/blockdevicesnapshot"
	blueprint "github.com/therandombyte/provider-vra8/internal/controller/blueprint/blueprint"
	version "github.com/therandombyte/provider-vra8/internal/controller/blueprint/version"
	catalogitementitlement "github.com/therandombyte/provider-vra8/internal/controller/catalogitementitlement/catalogitementitlement"
	catalogsourceblueprint "github.com/therandombyte/provider-vra8/internal/controller/catalogsourceblueprint/catalogsourceblueprint"
	catalogsourceentitlement "github.com/therandombyte/provider-vra8/internal/controller/catalogsourceentitlement/catalogsourceentitlement"
	accountaws "github.com/therandombyte/provider-vra8/internal/controller/cloudaccount/accountaws"
	accountazure "github.com/therandombyte/provider-vra8/internal/controller/cloudaccount/accountazure"
	accountgcp "github.com/therandombyte/provider-vra8/internal/controller/cloudaccount/accountgcp"
	accountnsxt "github.com/therandombyte/provider-vra8/internal/controller/cloudaccount/accountnsxt"
	accountvmc "github.com/therandombyte/provider-vra8/internal/controller/cloudaccount/accountvmc"
	accountvsphere "github.com/therandombyte/provider-vra8/internal/controller/cloudaccount/accountvsphere"
	contentsource "github.com/therandombyte/provider-vra8/internal/controller/contentsource/contentsource"
	deployment "github.com/therandombyte/provider-vra8/internal/controller/deployment/deployment"
	compute "github.com/therandombyte/provider-vra8/internal/controller/fabric/compute"
	datastorevsphere "github.com/therandombyte/provider-vra8/internal/controller/fabric/datastorevsphere"
	networkvsphere "github.com/therandombyte/provider-vra8/internal/controller/fabric/networkvsphere"
	profile "github.com/therandombyte/provider-vra8/internal/controller/flavorprofile/profile"
	profileimageprofile "github.com/therandombyte/provider-vra8/internal/controller/imageprofile/profile"
	integration "github.com/therandombyte/provider-vra8/internal/controller/integration/integration"
	loadbalancer "github.com/therandombyte/provider-vra8/internal/controller/loadbalancer/loadbalancer"
	machine "github.com/therandombyte/provider-vra8/internal/controller/machine/machine"
	network "github.com/therandombyte/provider-vra8/internal/controller/network/network"
	networkiprange "github.com/therandombyte/provider-vra8/internal/controller/network/networkiprange"
	networkprofile "github.com/therandombyte/provider-vra8/internal/controller/network/networkprofile"
	project "github.com/therandombyte/provider-vra8/internal/controller/project/project"
	providerconfig "github.com/therandombyte/provider-vra8/internal/controller/providerconfig"
	profilestorage "github.com/therandombyte/provider-vra8/internal/controller/storage/profile"
	profileaws "github.com/therandombyte/provider-vra8/internal/controller/storage/profileaws"
	profileazure "github.com/therandombyte/provider-vra8/internal/controller/storage/profileazure"
	profilevsphere "github.com/therandombyte/provider-vra8/internal/controller/storage/profilevsphere"
	zone "github.com/therandombyte/provider-vra8/internal/controller/zone/zone"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		blockdevice.Setup,
		blockdevicesnapshot.Setup,
		blueprint.Setup,
		version.Setup,
		catalogitementitlement.Setup,
		catalogsourceblueprint.Setup,
		catalogsourceentitlement.Setup,
		accountaws.Setup,
		accountazure.Setup,
		accountgcp.Setup,
		accountnsxt.Setup,
		accountvmc.Setup,
		accountvsphere.Setup,
		contentsource.Setup,
		deployment.Setup,
		compute.Setup,
		datastorevsphere.Setup,
		networkvsphere.Setup,
		profile.Setup,
		profileimageprofile.Setup,
		integration.Setup,
		loadbalancer.Setup,
		machine.Setup,
		network.Setup,
		networkiprange.Setup,
		networkprofile.Setup,
		project.Setup,
		providerconfig.Setup,
		profilestorage.Setup,
		profileaws.Setup,
		profileazure.Setup,
		profilevsphere.Setup,
		zone.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
