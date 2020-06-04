package prismacloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDsEnterpriseSettings(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDsEnterpriseSettingsConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloud_enterprise_settings.test", "session_timeout"),
					resource.TestCheckResourceAttrSet("data.prismacloud_enterprise_settings.test", "anomaly_training_model_threshold"),
					resource.TestCheckResourceAttrSet("data.prismacloud_enterprise_settings.test", "anomaly_alert_disposition"),
					resource.TestCheckResourceAttrSet("data.prismacloud_enterprise_settings.test", "user_attribution_in_notification"),
					resource.TestCheckResourceAttrSet("data.prismacloud_enterprise_settings.test", "require_alert_dismissal_note"),
					resource.TestCheckResourceAttrSet("data.prismacloud_enterprise_settings.test", "apply_default_policies_enabled"),
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
