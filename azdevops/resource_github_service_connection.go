package azdevops

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGithubConnection() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubConnectionCreate,
		Read:   resourceGithubConnectionRead,
		Update: resourceGithubConnectionUpdate,
		Delete: resourceGithubConnectionDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"token": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"project": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func newGithubConnectionFromResourceData(d *schema.ResourceData) *githubEndpoint {

	e := &githubEndpoint{}

	e.Name = d.Get("name").(string)
	e.Type = "github"
	e.URL = "https://github.com"
	e.Authorization.Parameters.Token = d.Get("token").(string)
	e.Authorization.Scheme = "PersonalAccessToken"

	return e
}

func resourceGithubConnectionCreate(d *schema.ResourceData, m interface{}) error {

	var reqBody []byte
	var ghe *githubEndpoint
	var project = d.Get("project").(string)
	var c *AzDevOpsClient
	var err error
	var resp *http.Response
	var respBodyBytes []byte

	c = m.(*AzDevOpsClient)
	ghe = newGithubConnectionFromResourceData(d)

	for reqBody, err = json.Marshal(ghe); err != nil; {
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

	for err = json.Unmarshal(respBodyBytes, &ghe); err != nil; {
		return err
	}

	d.SetId(ghe.ID)
	return nil
}

func resourceGithubConnectionRead(d *schema.ResourceData, m interface{}) error {

	var ghe *githubEndpoint
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

	for err = json.Unmarshal(respBodyBytes, &ghe); err != nil; {
		return err
	}

	d.Set("name", ghe.Name)

	return nil
}

func resourceGithubConnectionUpdate(d *schema.ResourceData, m interface{}) error {
	var reqBody []byte
	var ghe *githubEndpoint
	var project = d.Get("project").(string)
	var c *AzDevOpsClient
	var err error

	c = m.(*AzDevOpsClient)
	ghe = newGithubConnectionFromResourceData(d)

	for reqBody, err = json.Marshal(ghe); err != nil; {
		return err
	}

	for _, err = c.Put(
		getServiceEndpointURL(project, d.Id()),
		bytes.NewBuffer(reqBody)); err != nil; {
		return err
	}

	return nil
}

func resourceGithubConnectionDelete(d *schema.ResourceData, m interface{}) error {
	var c *AzDevOpsClient
	var err error
	var project = d.Get("project").(string)

	c = m.(*AzDevOpsClient)

	for _, err = c.Delete(getServiceEndpointURL(project, d.Id())); err != nil; {
		return err
	}

	return nil
}
