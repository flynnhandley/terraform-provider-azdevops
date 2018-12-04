package azdevops

import (
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform/helper/schema"
)

func getServiceEndpointsURL(project string) string {
	return fmt.Sprintf("%s/_apis/distributedtask/serviceendpoints/?api-version=4.1-preview", project)
}

func getBuildDefinitionURL(project string) string {
	return fmt.Sprintf("%s/_apis/build/definitions/?api-version=4.1-preview", project)
}

func getServiceEndpointURL(project, id string) string {
	return fmt.Sprintf("%s/_apis/distributedtask/serviceendpoints/%s/?api-version=4.1-preview", project, id)
}

func validateUUID() schema.SchemaValidateFunc {

	return func(val interface{}, key string) (warns []string, errs []error) {
		v := val.(string)
		r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")

		if r.MatchString(v) == false {
			errs = append(errs, fmt.Errorf("%s is not a valid UUID ", key))
		}
		return
	}

}
