// config-updater-provider is a Terraform provider for config-updater.
//
// Currently, only implements secrets_export.
//
// https://labs.docs.magenta.dk/config-updater.html

package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return &schema.Provider{
				Schema: map[string]*schema.Schema{
					"url": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
						Default:  "https://config-updater.magentahosted.dk/",
					},
				},
				ResourcesMap: map[string]*schema.Resource{
					"configupdater_secret": secretsExport(),
				},
				ConfigureFunc: providerConfigure,
			}
		},
	})
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	url := d.Get("url").(string)
	return url, nil
}
