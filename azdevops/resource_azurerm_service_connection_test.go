package azdevops

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAzureRmServiceConnection_Success(t *testing.T) {

	resource.Test(t, resource.TestCase{
		IsUnitTest: true,
		Providers: map[string]terraform.ResourceProvider{
			"azdevops": Provider(),
		},
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testHelperLoadString(t, "azurerm_service_connection.tf"),
			},
		},
	})
}
