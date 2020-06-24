package prismacloud

import (
	"fmt"
	"testing"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/alert/rule"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccAlertRule(t *testing.T) {
	var o rule.Rule
	grp := fmt.Sprintf("tf%s", acctest.RandString(6))
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccAlertRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAlertRuleConfig(grp, name, "alert rule desc1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAlertRuleExists("prismacloud_alert_rule.test", &o),
					testAccCheckAlertRuleAttributes(&o, name, "alert rule desc1"),
				),
			},
			{
				Config: testAccAlertRuleConfig(grp, name, "alert rule desc2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAlertRuleExists("prismacloud_alert_rule.test", &o),
					testAccCheckAlertRuleAttributes(&o, name, "alert rule desc2"),
				),
			},
		},
	})
}

func testAccCheckAlertRuleExists(n string, o *rule.Rule) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Resource not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Object label ID is not set")
		}

		client := testAccProvider.Meta().(*pc.Client)
		id := rs.Primary.ID
		lo, err := rule.Get(client, id)
		if err != nil {
			return fmt.Errorf("Error in get: %s", err)
		}

		*o = lo

		return nil
	}
}

func testAccCheckAlertRuleAttributes(o *rule.Rule, name, desc string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if o.Name != name {
			return fmt.Errorf("Name is %s, expected %s", o.Name, name)
		}

		if o.Description != desc {
			return fmt.Errorf("Description is %s, expected %s", o.Description, desc)
		}

		return nil
	}
}

func testAccAlertRuleDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*pc.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "prismacloud_alert_rule" {
			continue
		}

		if rs.Primary.ID != "" {
			id := rs.Primary.ID
			if _, err := rule.Get(client, id); err == nil {
				return fmt.Errorf("Object %q still exists", rs.Primary.ID)
			}
		}
		return nil
	}

	return nil
}

func testAccAlertRuleConfig(grp, name, desc string) string {
	return fmt.Sprintf(`
resource "prismacloud_account_group" "x" {
    name = %q
    description = "for alert rule acctest"
}

resource "prismacloud_alert_rule" "test" {
    name = %q
    description = %q
    enabled = false
    scan_all = true
    target {
        account_groups = [
            prismacloud_account_group.x.group_id,
        ]
    }
}
`, grp, name, desc)
}
