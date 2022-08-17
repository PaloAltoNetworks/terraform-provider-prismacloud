package prismacloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDsEnterpriseSettings(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDsEnterpriseSettingsConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloud_enterprise_settings.test", "access_key_max_validity"),
					resource.TestCheckResourceAttrSet("data.prismacloud_enterprise_settings.test", "session_timeout"),
					resource.TestCheckResourceAttrSet("data.prismacloud_enterprise_settings.test", "user_attribution_in_notification"),
					resource.TestCheckResourceAttrSet("data.prismacloud_enterprise_settings.test", "require_alert_dismissal_note"),
					resource.TestCheckResourceAttrSet("data.prismacloud_enterprise_settings.test", "apply_default_policies_enabled"),
					resource.TestCheckResourceAttrSet("data.prismacloud_enterprise_settings.test", "alarm_enabled"),
				),
			},
		},
	})
}

func testAccDsEnterpriseSettingsConfig() string {
	return `
data "prismacloud_enterprise_settings" "test" {}
`
}
