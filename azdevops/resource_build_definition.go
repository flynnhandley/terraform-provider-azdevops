package azdevops

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
)

type AzError struct {
	ID             string      `json:"$id"`
	InnerException interface{} `json:"innerException"`
	Message        string      `json:"message"`
	TypeName       string      `json:"typeName"`
	TypeKey        string      `json:"typeKey"`
	ErrorCode      int         `json:"errorCode"`
	EventID        int         `json:"eventId"`
}

func resourceBuildDefinition() *schema.Resource {
	return &schema.Resource{
		Create: resourceBuildDefinitionCreate,
		Read:   resourceBuildDefinitionRead,
		Update: resourceBuildDefinitionUpdate,
		Delete: resourceBuildDefinitionDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"project": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"build_queue_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"yaml_file": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"github_repository": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"endpoint_id": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validateUUID(),
						},
						"repo": {
							Type:     schema.TypeString,
							Required: true,
						},
						"url": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			}, "variables": {
				Type:     schema.TypeList,
				Optional: true,
				Required: false,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
						"is_secret": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"allow_override": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
					},
				},
			},
		},
	}
}
func newBuildDefinitionFromResourceData(d *schema.ResourceData) *buildDefinition {

	ghr := d.Get("github_repository").([]interface{})[0].(map[string]interface{})

	b := &buildDefinition{}

	b.Name = d.Get("name").(string)
	b.Repository.Properties.LabelSources = "0"
	b.Repository.Properties.ReportBuildStatus = "true"
	b.Repository.Properties.ConnectedServiceID = ghr["endpoint_id"].(string)
	b.Repository.Clean = "false"
	b.Repository.DefaultBranch = "master"
	b.Repository.Type = "GitHub"
	b.Repository.Name = ghr["repo"].(string)
	b.Repository.URL = ghr["url"].(string)
	b.Process.YamlFilename = d.Get("yaml_file").(string)
	b.Process.Type = 2
	b.Queue.ID = d.Get("build_queue_id").(int)
	b.Variables = expandVariables(d)

	return b
}

func expandVariables(*schema.ResourceData) []variable {
	return []variable{}
}

func resourceBuildDefinitionCreate(d *schema.ResourceData, m interface{}) error {
	var reqBody []byte
	var bdef *buildDefinition
	var project = d.Get("project").(string)
	var c *AzDevOpsClient
	var err error
	var resp *http.Response
	var respBodyBytes []byte

	c = m.(*AzDevOpsClient)
	bdef = newBuildDefinitionFromResourceData(d)

	for reqBody, err = json.Marshal(bdef); err != nil; {

		return err
	}

	print(string(reqBody[:]))

	for resp, err = c.Post(
		getBuildDefinitionURL(project),
		bytes.NewBuffer(reqBody)); err != nil; {
		print(err.Error())
		var azerr AzError

		resp.Body.Read(respBodyBytes)

		ddddd := json.NewDecoder(resp.Body).Decode(&azerr)
		respBodyBytes, err = ioutil.ReadAll(resp.Body)
		asdf := json.Unmarshal(respBodyBytes, &azerr)

		if ddddd != nil {
			return asdf
		}
		if asdf != nil {
			return asdf
		}

		return fmt.Errorf(azerr.Message)
	}

	for respBodyBytes, err = ioutil.ReadAll(resp.Body); err != nil; {
		return err
	}

	for err = json.Unmarshal(respBodyBytes, &bdef); err != nil; {
		return err
	}

	d.SetId(strconv.Itoa(bdef.ID))
	return nil
}

func resourceBuildDefinitionRead(d *schema.ResourceData, m interface{}) error {

	return nil
}

func resourceBuildDefinitionUpdate(d *schema.ResourceData, m interface{}) error {

	return nil
}

func resourceBuildDefinitionDelete(d *schema.ResourceData, m interface{}) error {

	var c *AzDevOpsClient
	var err error
	var project = d.Get("project").(string)
	c = m.(*AzDevOpsClient)

	url := fmt.Sprintf("/%s/_apis/build/definitions/%s?api-version=4.1", project, d.Id())

	for _, err = c.Delete(url); err != nil; {
		return err
	}

	return nil
}
