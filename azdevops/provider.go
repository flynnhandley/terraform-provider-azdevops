package azdevops

import (
	"net/http"

	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"azdevops_github_service_connection":  resourceGithubConnection(),
			"azdevops_azurerm_service_connection": resourceAzureRMConnection(),
			"azdevops_build_definition":           resourceBuildDefinition(),
		},
		ConfigureFunc: providerConfigure,
		Schema: map[string]*schema.Schema{
			"account": {
				Required:    true,
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("AZDEVOPS_ACCOUNT", nil),
			},
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("AZDEVOPS_TOKEN", nil),
			},
		},
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	client := &AzDevOpsClient{
		Account:    d.Get("account").(string),
		Token:      d.Get("token").(string),
		HTTPClient: &http.Client{},
	}

	return client, nil
}
