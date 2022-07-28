package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"golang.org/x/net/context"
	"log"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/settings/enterprise"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEnterpriseSettings() *schema.Resource {
	return &schema.Resource{
		CreateContext: createUpdateEnterpriseSettings,
		ReadContext:   readEnterpriseSettings,
		UpdateContext: createUpdateEnterpriseSettings,
		DeleteContext: deleteEnterpriseSettings,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"access_key_max_validity": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Access Keys maximum validity in days",
			},
			"session_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Browser session timeout",
			},
			"user_attribution_in_notification": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "User attribution in notification",
			},
			"require_alert_dismissal_note": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Require alert dismissal note",
			},
			"default_policies_enabled": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Default policies enabled",
				Elem: &schema.Schema{
					Type: schema.TypeBool,
				},
			},
			"apply_default_policies_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply default policies enabled",
			},
			"alarm_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Alarms enabled",
				Default:     true,
			},
		},
	}
}

func parseEnterpriseSettings(d *schema.ResourceData) enterprise.Config {
	dpe := make(map[string]bool)
	dpec := d.Get("default_policies_enabled").(map[string]interface{})
	for key := range dpec {
		dpe[key] = dpec[key].(bool)
	}

	return enterprise.Config{
		AccessKeyMaxValidity:          d.Get("access_key_max_validity").(int),
		SessionTimeout:                d.Get("session_timeout").(int),
		UserAttributionInNotification: d.Get("user_attribution_in_notification").(bool),
		RequireAlertDismissalNote:     d.Get("require_alert_dismissal_note").(bool),
		DefaultPoliciesEnabled:        dpe,
		ApplyDefaultPoliciesEnabled:   d.Get("apply_default_policies_enabled").(bool),
		AlarmEnabled:                  d.Get("alarm_enabled").(bool),
	}
}

func saveEnterpriseSettings(d *schema.ResourceData, conf enterprise.Config) {
	var err error

	d.Set("access_key_max_validity", conf.AccessKeyMaxValidity)
	d.Set("session_timeout", conf.SessionTimeout)
	d.Set("user_attribution_in_notification", conf.UserAttributionInNotification)
	d.Set("require_alert_dismissal_note", conf.RequireAlertDismissalNote)
	if err = d.Set("default_policies_enabled", conf.DefaultPoliciesEnabled); err != nil {
		log.Printf("[WARN] Error setting 'default_policies_enabled' for %s: %s", d.Id(), err)
	}
	d.Set("apply_default_policies_enabled", conf.ApplyDefaultPoliciesEnabled)
	d.Set("alarm_enabled", conf.AlarmEnabled)
}

func createUpdateEnterpriseSettings(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	conf := parseEnterpriseSettings(d)

	if err := enterprise.Update(client, conf); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("config")

	return readEnterpriseSettings(ctx, d, meta)
}

func readEnterpriseSettings(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)

	conf, err := enterprise.Get(client)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("config")
	saveEnterpriseSettings(d, conf)

	return nil
}

func deleteEnterpriseSettings(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
