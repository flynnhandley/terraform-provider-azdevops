package azdevops

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAzureRMConnection() *schema.Resource {
	return &schema.Resource{
		Create: resourceAzureRMConnectionCreate,
		Read:   resourceAzureRMConnectionRead,
		Update: resourceAzureRMConnectionUpdate,
		Delete: resourceAzureRMConnectionDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"project": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"authorization_parameters": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service_principal_id": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validateUUID(),
						},
						"tenant_id": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validateUUID(),
						},
						"service_principal_key": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"data": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subscription_id": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validateUUID(),
						},
						"subscription_name": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func newAzureConnectionFromResourceData(d *schema.ResourceData) *azureEndpoint {

	e := &azureEndpoint{}
	e.Name = d.Get("name").(string)
	e.Type = "azurerm"
	e.URL = "https://management.azure.com/"

	data := d.Get("data").([]interface{})[0].(map[string]interface{})
	authParams := d.Get("authorization_parameters").([]interface{})[0].(map[string]interface{})

	e.Data.SubscriptionID = data["subscription_id"].(string)
	e.Data.SubscriptionName = data["subscription_name"].(string)
	e.Data.CreationMode = "Manual"

	e.Authorization.Scheme = "ServicePrincipal"

	e.Authorization.Parameters.ServicePrincipalID = authParams["service_principal_id"].(string)
	e.Authorization.Parameters.ServicePrincipalKey = authParams["service_principal_key"].(string)
	e.Authorization.Parameters.TenantID = authParams["tenant_id"].(string)

	return e
}

func resourceAzureRMConnectionCreate(d *schema.ResourceData, m interface{}) error {

	var reqBody []byte
	var aze *azureEndpoint
	var project = d.Get("project").(string)
	var c *AzDevOpsClient
	var err error
	var resp *http.Response
	var respBodyBytes []byte

	c = m.(*AzDevOpsClient)
	aze = newAzureConnectionFromResourceData(d)

	for reqBody, err = json.Marshal(aze); err != nil; {

		return err
	}

	for resp, err = c.Post(
		getServiceEndpointsURL(project),
		bytes.NewBuffer(reqBody)); err != nil; {
		return err
	}

	for respBodyBytes, err = ioutil.ReadAll(resp.Body); err != nil; {
		return err
	}

	for err = json.Unmarshal(respBodyBytes, &aze); err != nil; {
		return err
	}

	d.SetId(aze.ID)
	return nil
}

func resourceAzureRMConnectionRead(d *schema.ResourceData, m interface{}) error {

	var aze *githubEndpoint
	var project = d.Get("project").(string)
	var c *AzDevOpsClient
	var err error
	var resp *http.Response
	var respBodyBytes []byte

	project = d.Get("project").(string)

	c = m.(*AzDevOpsClient)

	for resp, err = c.Get(getServiceEndpointURL(project, d.Id())); err != nil; {
		return err
	}

	for respBodyBytes, err = ioutil.ReadAll(resp.Body); err != nil; {
		return err
	}

	for err = json.Unmarshal(respBodyBytes, &aze); err != nil; {
		return err
	}

	d.Set("name", aze.Name)
	d.SetId(aze.ID)

	return nil
}

func resourceAzureRMConnectionUpdate(d *schema.ResourceData, m interface{}) error {
	var reqBody []byte
	var aze *azureEndpoint
	var project = d.Get("project").(string)
	var c *AzDevOpsClient
	var err error

	c = m.(*AzDevOpsClient)
	aze = newAzureConnectionFromResourceData(d)

	for reqBody, err = json.Marshal(aze); err != nil; {
		return err
	}

	for _, err = c.Put(
		getServiceEndpointURL(project, d.Id()),
		bytes.NewBuffer(reqBody)); err != nil; {
		return err
	}

	return nil
}

func resourceAzureRMConnectionDelete(d *schema.ResourceData, m interface{}) error {
	var c *AzDevOpsClient
	var err error
	var project = d.Get("project").(string)

	c = m.(*AzDevOpsClient)

	for _, err = c.Delete(getServiceEndpointURL(project, d.Id())); err != nil; {
		return err
	}

	return nil
}
