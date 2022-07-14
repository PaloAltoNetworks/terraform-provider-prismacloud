package prismacloud

import (
	"fmt"
	"testing"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/alert/rule"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccAlertRule(t *testing.T) {
	var o rule.Rule
	grp := fmt.Sprintf("tf%s", acctest.RandString(6))
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	tags := []rule.Tag{
		{
			Key:    fmt.Sprintf("tf%s", acctest.RandString(6)),
			Values: []string{fmt.Sprintf("tf%s", acctest.RandString(6)), fmt.Sprintf("tf%s", acctest.RandString(6))},
		},
		{
			Key:    fmt.Sprintf("tf%s", acctest.RandString(6)),
			Values: []string{fmt.Sprintf("tf%s", acctest.RandString(6)), fmt.Sprintf("tf%s", acctest.RandString(6))},
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccAlertRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAlertRuleConfig(grp, name, "alert rule desc1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAlertRuleExists("prismacloud_alert_rule.test", &o),
					testAccCheckAlertRuleAttributes(&o, name, "alert rule desc1", nil),
				),
			},
			{
				Config: testAccAlertRuleConfig(grp, name, "alert rule desc2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAlertRuleExists("prismacloud_alert_rule.test", &o),
					testAccCheckAlertRuleAttributes(&o, name, "alert rule desc2", nil),
				),
			},
			{
				Config: testAccAlertRuleConfigTags(grp, name, "alert rule tags", tags),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAlertRuleExists("prismacloud_alert_rule.test", &o),
					testAccCheckAlertRuleAttributes(&o, name, "alert rule with tags", tags),
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

func testAccCheckAlertRuleAttributes(o *rule.Rule, name, desc string, tags []rule.Tag) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if o.Name != name {
			return fmt.Errorf("Name is %s, expected %s", o.Name, name)
		}

		if o.Description != desc {
			return fmt.Errorf("Description is %s, expected %s", o.Description, desc)
		}

		if tags != nil {
			for i := 0; i < len(tags); i++ {
				if o.Target.Tags[i].Key != tags[i].Key {
					return fmt.Errorf("Tags is %+v, expected %+v", o.Target.Tags, tags)
				}

				for j := 0; j < len(tags[i].Values); j++ {
					if o.Target.Tags[i].Values[j] != tags[i].Values[j] {
						return fmt.Errorf("Tags is %+v, expected %+v", o.Target.Tags, tags)
					}
				}
			}

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

func testAccAlertRuleConfigTags(grp, name, desc string, tags []rule.Tag) string {
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
			tags {
				key = "%s"
				values = ["%s", "%s"]
			}

			tags {
				key = "%s"
				values = ["%s", "%s"]
			}
		}
	}
	`, grp, name, desc, tags[0].Key, tags[0].Values[0], tags[0].Values[1], tags[1].Key, tags[1].Values[0], tags[1].Values[1])
}
