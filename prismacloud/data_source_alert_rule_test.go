package prismacloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDsAlertRule(t *testing.T) {
	name := fmt.Sprintf("tf%s", acctest.RandString(6))
	group := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDsAlertRuleConfig(name, group),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloud_alert_rule.test", "policy_scan_config_id"),
					resource.TestCheckResourceAttrSet("data.prismacloud_alert_rule.test", "last_modified_on"),
				),
			},
		},
	})
}

func testAccDsAlertRuleConfig(name, group string) string {
	return fmt.Sprintf(`
resource "prismacloud_alert_rule" "x" {
    name = %q
    description = "data source alert rule acctest"
    scan_all = true
    notify_on_open = true
    target {
        account_groups = [
            prismacloud_account_group.x.group_id,
        ]
    }
}

resource "prismacloud_account_group" "x" {
    name = %q
    description = "data source alert rule acctest"
}

data "prismacloud_alert_rule" "test" {
    policy_scan_config_id = prismacloud_alert_rule.x.policy_scan_config_id
}
`, name, group)
}
