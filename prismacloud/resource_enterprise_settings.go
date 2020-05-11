package prismacloud

import (
	"log"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/settings/enterprise"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceEnterpriseSettings() *schema.Resource {
	return &schema.Resource{
		Create: createUpdateEnterpriseSettings,
		Read:   readEnterpriseSettings,
		Update: createUpdateEnterpriseSettings,
		Delete: deleteEnterpriseSettings,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"session_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Browser session timeout",
			},
			"anomaly_training_model_threshold": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Anomaly training model threshold",
				ValidateFunc: validation.StringInSlice(
					[]string{enterprise.Low, enterprise.Medium, enterprise.High},
					false,
				),
			},
			"anomaly_alert_disposition": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Anomaly alert disposition",
				ValidateFunc: validation.StringInSlice(
					[]string{enterprise.Low, enterprise.Medium, enterprise.High},
					false,
				),
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
		SessionTimeout:                d.Get("session_timeout").(int),
		AnomalyTrainingModelThreshold: d.Get("anomaly_training_model_threshold").(string),
		AnomalyAlertDisposition:       d.Get("anomaly_alert_disposition").(string),
		UserAttributionInNotification: d.Get("user_attribution_in_notification").(bool),
		RequireAlertDismissalNote:     d.Get("require_alert_dismissal_note").(bool),
		DefaultPoliciesEnabled:        dpe,
		ApplyDefaultPoliciesEnabled:   d.Get("apply_default_policies_enabled").(bool),
	}
}

func saveEnterpriseSettings(d *schema.ResourceData, conf enterprise.Config) {
	var err error

	d.Set("session_timeout", conf.SessionTimeout)
	d.Set("anomaly_training_model_threshold", conf.AnomalyTrainingModelThreshold)
	d.Set("anomaly_alert_disposition", conf.AnomalyAlertDisposition)
	d.Set("user_attribution_in_notification", conf.UserAttributionInNotification)
	d.Set("require_alert_dismissal_note", conf.RequireAlertDismissalNote)
	if err = d.Set("default_policies_enabled", conf.DefaultPoliciesEnabled); err != nil {
		log.Printf("[WARN] Error setting 'default_policies_enabled' for %s: %s", d.Id(), err)
	}
	d.Set("apply_default_policies_enabled", conf.ApplyDefaultPoliciesEnabled)
}

func createUpdateEnterpriseSettings(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	conf := parseEnterpriseSettings(d)

	if err := enterprise.Update(client, conf); err != nil {
		return err
	}

	d.SetId("config")

	return readEnterpriseSettings(d, meta)
}

func readEnterpriseSettings(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)

	conf, err := enterprise.Get(client)
	if err != nil {
		return err
	}

	d.SetId("config")
	saveEnterpriseSettings(d, conf)

	return nil
}

func deleteEnterpriseSettings(d *schema.ResourceData, meta interface{}) error {
	return nil
}
