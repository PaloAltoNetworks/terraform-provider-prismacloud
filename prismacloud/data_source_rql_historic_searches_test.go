package prismacloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDsRqlHistoricSearches(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDsRqlHistoricSearchesConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloud_rql_historic_searches.test", "total"),
				),
			},
		},
	})
}

func testAccDsRqlHistoricSearchesConfig() string {
	return `
data "prismacloud_rql_historic_searches" "test" {}
`
}
