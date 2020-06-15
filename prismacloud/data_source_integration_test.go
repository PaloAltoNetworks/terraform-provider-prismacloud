package prismacloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDsIntegration(t *testing.T) {
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDsIntegrationConfig(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloud_integration.test", "integration_id"),
					resource.TestCheckResourceAttrSet("data.prismacloud_integration.test", "name"),
				),
			},
		},
	})
}

func testAccDsIntegrationConfig(name string) string {
	return fmt.Sprintf(`
data "prismacloud_integration" "test" {
    integration_id = prismacloud_integration.x.integration_id
}

resource "prismacloud_integration" "x" {
    name = %q
    description = "integration ds acctest"
    enabled = false
    integration_type = "pager_duty"
    integration_config {
        integration_key = "mySecretKey"
        auth_token = "my-secret-auth-token"
    }
}
`, name)
}
