package blueprint

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vra_blueprint", func(r *config.Resource) {

		r.ShortGroup = "blueprint"
		r.Version = "v1alpha1"
		r.References["project_id"] = config.Reference{
			Type: "github.com/therandombyte/provider-vra8/apis/project/v1alpha1.Project",
		}
	})

	p.AddResourceConfigurator("vra_blueprint_version", func(r *config.Resource) {

		r.ShortGroup = "blueprint"
		r.Version = "v1alpha1"
		r.References["blueprint_id"] = config.Reference{
			Type: "Blueprint", // github.com/crossplane-contrib/provider-vra/apis/blueprint/v1alpha1.Blueprint
		}
	})
}
