package azdevops

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestBuildDefinition_Success(t *testing.T) {

	resource.Test(t, resource.TestCase{
		IsUnitTest: true,
		Providers: map[string]terraform.ResourceProvider{
			"azdevops": Provider(),
		},
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testHelperLoadString(t, "build_definition.tf"),
			},
		},
	})
}
