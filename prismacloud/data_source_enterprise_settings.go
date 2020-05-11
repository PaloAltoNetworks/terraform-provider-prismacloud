package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceEnterpriseSettings() *schema.Resource {
	return &schema.Resource{
		Read: readEnterpriseSettings,

		Schema: map[string]*schema.Schema{
			// Output.
			"session_timeout": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Browser session timeout",
			},
			"anomaly_training_model_threshold": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Anomaly training model threshold",
			},
			"anomaly_alert_disposition": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Anomaly alert disposition",
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
		},
	}
}
