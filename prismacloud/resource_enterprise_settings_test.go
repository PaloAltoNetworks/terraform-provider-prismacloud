package prismacloud

import (
	"bytes"
	"fmt"
	"testing"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/settings/enterprise"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccEnterpriseSettings(t *testing.T) {
	if originalEnterpriseSettings == nil {
		t.Skip("Failed to retrieve initial enterprise settings")
	}

	times := []int{180, 240}
	tos := make([]int, 0, 2)
	for _, t := range times {
		if t == originalEnterpriseSettings.SessionTimeout {
			continue
		}
		tos = append(tos, t)
		if len(tos) == 1 {
			tos = append(tos, originalEnterpriseSettings.SessionTimeout)
			break
		}
	}

	var o enterprise.Config

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccEnterpriseSettingsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccEnterpriseSettingsConfig(tos[0]),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEnterpriseSettingsExists("prismacloud_enterprise_settings.test", &o),
					testAccCheckEnterpriseSettingsAttributes(&o, tos[0]),
				),
			},
			{
				Config: testAccEnterpriseSettingsConfig(tos[1]),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEnterpriseSettingsExists("prismacloud_enterprise_settings.test", &o),
					testAccCheckEnterpriseSettingsAttributes(&o, tos[1]),
				),
			},
		},
	})
}

func testAccCheckEnterpriseSettingsExists(n string, o *enterprise.Config) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Resource not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Object label ID is not set")
		}

		client := testAccProvider.Meta().(*pc.Client)
		lo, err := enterprise.Get(client)
		if err != nil {
			return fmt.Errorf("Error in get: %s", err)
		}

		*o = lo

		return nil
	}
}

func testAccCheckEnterpriseSettingsAttributes(o *enterprise.Config, tout int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if o.SessionTimeout != tout {
			return fmt.Errorf("Session timeout is %d, expected %d", o.SessionTimeout, tout)
		}

		if o.AnomalyTrainingModelThreshold != originalEnterpriseSettings.AnomalyTrainingModelThreshold {
			return fmt.Errorf("Anomaly training model threshold is %q, expected %q", o.AnomalyTrainingModelThreshold, originalEnterpriseSettings.AnomalyTrainingModelThreshold)
		}

		if o.AnomalyAlertDisposition != originalEnterpriseSettings.AnomalyAlertDisposition {
			return fmt.Errorf("Anomaly alert disposition is %q, expected %q", o.AnomalyAlertDisposition, originalEnterpriseSettings.AnomalyAlertDisposition)
		}

		if o.UserAttributionInNotification != originalEnterpriseSettings.UserAttributionInNotification {
			return fmt.Errorf("User attribution in notification is %t, expected %t", o.UserAttributionInNotification, originalEnterpriseSettings.UserAttributionInNotification)
		}

		if o.RequireAlertDismissalNote != originalEnterpriseSettings.RequireAlertDismissalNote {
			return fmt.Errorf("Require alert dismissal note is %t, expected %t", o.RequireAlertDismissalNote, originalEnterpriseSettings.RequireAlertDismissalNote)
		}

		if o.ApplyDefaultPoliciesEnabled != originalEnterpriseSettings.ApplyDefaultPoliciesEnabled {
			return fmt.Errorf("Apply default policies enabled is %t, expected %t", o.ApplyDefaultPoliciesEnabled, originalEnterpriseSettings.ApplyDefaultPoliciesEnabled)
		}

		if len(o.DefaultPoliciesEnabled) != len(originalEnterpriseSettings.DefaultPoliciesEnabled) {
			return fmt.Errorf("Len of default policies enabled is %d, expected %d", len(o.DefaultPoliciesEnabled), len(originalEnterpriseSettings.DefaultPoliciesEnabled))
		}

		for key := range o.DefaultPoliciesEnabled {
			if o.DefaultPoliciesEnabled[key] != originalEnterpriseSettings.DefaultPoliciesEnabled[key] {
				return fmt.Errorf("Default policies enabled %q is %t, expected %t", key, o.DefaultPoliciesEnabled[key], originalEnterpriseSettings.DefaultPoliciesEnabled[key])
			}
		}

		return nil
	}
}

func testAccEnterpriseSettingsDestroy(s *terraform.State) error {
	return nil
}

func testAccEnterpriseSettingsConfig(tout int) string {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf(`
resource "prismacloud_enterprise_settings" "test" {
    session_timeout = %d
    anomaly_training_model_threshold = %q
    anomaly_alert_disposition = %q
    user_attribution_in_notification = %t
    require_alert_dismissal_note = %t
    apply_default_policies_enabled = %t
    default_policies_enabled = {`,
		tout,
		originalEnterpriseSettings.AnomalyTrainingModelThreshold,
		originalEnterpriseSettings.AnomalyAlertDisposition,
		originalEnterpriseSettings.UserAttributionInNotification,
		originalEnterpriseSettings.RequireAlertDismissalNote,
		originalEnterpriseSettings.ApplyDefaultPoliciesEnabled,
	))

	for key, val := range originalEnterpriseSettings.DefaultPoliciesEnabled {
		buf.WriteString(fmt.Sprintf(`
        %q: %t,`, key, val))
	}
	buf.WriteString(`
    }
}
`)

	return buf.String()
}
