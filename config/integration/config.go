package integration

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vra_integration", func(r *config.Resource) {
		r.ShortGroup = "integration"
		r.Kind = "Integration"
		r.Version = "v1alpha1"
		r.References["project_id"] = config.Reference{
			Type: "github.com/therandombyte/provider-vra8/apis/project/v1alpha1.Project",
		}
	})
}
