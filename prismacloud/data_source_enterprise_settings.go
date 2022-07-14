package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEnterpriseSettings() *schema.Resource {
	return &schema.Resource{
		Read: readEnterpriseSettings,

		Schema: map[string]*schema.Schema{
			// Output.
			"access_key_max_validity": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Access Keys maximum validity in days",
			},
			"session_timeout": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Browser session timeout",
			},
			"user_attribution_in_notification": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "User attribution in notification",
			},
			"require_alert_dismissal_note": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Require alert dismissal note",
			},
			"default_policies_enabled": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Default policies enabled",
				Elem: &schema.Schema{
					Type: schema.TypeBool,
				},
			},
			"apply_default_policies_enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Apply default policies enabled",
			},
			"alarm_enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Alarms enabled",
			},
		},
	}
}
