package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEnterpriseSettings() *schema.Resource {
	return &schema.Resource{
		ReadContext: readEnterpriseSettings,

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
			"named_users_access_keys_expiry_notifications_enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Named users access keys expiry notifications enabled",
			},
			"service_users_access_keys_expiry_notifications_enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Service users access keys expiry notifications enabled",
			},
			"notification_threshold_access_keys_expiry": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Notification threshold access keys expiry",
			},
		},
	}
}
