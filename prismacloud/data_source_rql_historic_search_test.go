package prismacloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDsRqlHistoricSearch(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDsRqlHistoricSearchConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloud_rql_historic_search.audit", "search_id"),
					resource.TestCheckResourceAttrSet("data.prismacloud_rql_historic_search.audit", "name"),
					resource.TestCheckResourceAttrSet("data.prismacloud_rql_historic_search.config", "search_id"),
					resource.TestCheckResourceAttrSet("data.prismacloud_rql_historic_search.config", "name"),
					resource.TestCheckResourceAttrSet("data.prismacloud_rql_historic_search.network", "search_id"),
					resource.TestCheckResourceAttrSet("data.prismacloud_rql_historic_search.network", "name"),
				),
			},
		},
	})
}

func testAccDsRqlHistoricSearchConfig() string {
	return `
data "prismacloud_rql_historic_searches" "x" {}

locals {
    audit_event_ids = [
        for inst in data.prismacloud_rql_historic_searches.x.listing :
        inst.search_id
        if inst.search_type == "audit_event"
    ]
    config_ids = [
        for inst in data.prismacloud_rql_historic_searches.x.listing :
        inst.search_id
        if inst.search_type == "config"
    ]
    network_ids = [
        for inst in data.prismacloud_rql_historic_searches.x.listing :
        inst.search_id
        if inst.search_type == "network"
    ]
}

data "prismacloud_rql_historic_search" "audit" {
    search_id = local.audit_event_ids[0]
}

data "prismacloud_rql_historic_search" "config" {
    search_id = local.config_ids[0]
}

data "prismacloud_rql_historic_search" "network" {
    search_id = local.network_ids[0]
}
`
}
