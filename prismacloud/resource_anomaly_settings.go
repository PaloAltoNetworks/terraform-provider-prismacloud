package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"golang.org/x/net/context"
	"log"

	pc "github.com/paloaltonetworks/prisma-cloud-go"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/paloaltonetworks/prisma-cloud-go/anomalySettings"
)

func resourceAnomalySettings() *schema.Resource {
	return &schema.Resource{
		CreateContext: createAnomalySettings,
		ReadContext:   readAnomalySettings,
		UpdateContext: updateAnomalySettings,
		DeleteContext: deleteAnomalySettings,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"policy_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Policy ID",
			},
			"alert_disposition": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Alert disposition",
				Computed:    true,
				ValidateFunc: validation.StringInSlice(
					[]string{
						"aggressive",
						"moderate",
						"conservative",
					},
					false,
				),
			},
			"alert_disposition_description": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Alert disposition info",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"aggressive": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Aggressive",
						},
						"moderate": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Moderate",
						},
						"conservative": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Conservative",
						},
					},
				},
			},
			"policy_description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Policy description",
			},
			"policy_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Policy name",
			},
			"training_model_description": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Training model info",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"low": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Low",
						},
						"medium": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Medium",
						},
						"high": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "High",
						},
					},
				},
			},
			"training_model_threshold": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Training model threshold info",
				ValidateFunc: validation.StringInSlice(
					[]string{
						"low",
						"medium",
						"high",
					},
					false,
				),
			},
		},
	}
}

func parseAnomalySettings(d *schema.ResourceData, id string) anomalySettings.AnomalySettings {
	ans := anomalySettings.AnomalySettings{
		PolicyId:               d.Get("policy_id").(string),
		AlertDisposition:       d.Get("alert_disposition").(string),
		TrainingModelThreshold: d.Get("training_model_threshold").(string),
	}

	return ans
}

func saveAnomalySettings(d *schema.ResourceData, obj anomalySettings.AnomalySettings) {
	d.Set("alert_disposition", obj.AlertDisposition)
	d.Set("policy_description", obj.PolicyDescription)
	d.Set("policy_name", obj.PolicyName)
	d.Set("training_model_threshold", obj.TrainingModelThreshold)
	d.Set("type", obj.Type)
	d.Set("policy_id", d.Id())

	adp := map[string]interface{}{
		"aggressive":   obj.AlertDispositionDescription.Aggressive,
		"moderate":     obj.AlertDispositionDescription.Moderate,
		"conservative": obj.AlertDispositionDescription.Conservative,
	}

	if err := d.Set("alert_disposition_description", []interface{}{adp}); err != nil {
		log.Printf("[WARN] Error setting 'alert_disposition_description' for %q: %s", d.Id(), err)
	}

	tmd := map[string]interface{}{
		"low":    obj.TrainingModelDescription.Low,
		"medium": obj.TrainingModelDescription.Medium,
		"high":   obj.TrainingModelDescription.High,
	}

	if err := d.Set("training_model_description", []interface{}{tmd}); err != nil {
		log.Printf("[WARN] Error setting 'training_model_description' for %q: %s", d.Id(), err)
	}
}

func createAnomalySettings(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	obj := parseAnomalySettings(d, "")

	if err := anomalySettings.Create(client, obj); err != nil {
		return diag.FromErr(err)
	}

	id := d.Get("policy_id").(string)

	PollApiUntilSuccess(func() error {
		_, err := anomalySettings.Get(client, id)
		return err
	})

	d.SetId(id)
	return readAnomalySettings(ctx, d, meta)
}

func readAnomalySettings(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	id := d.Id()

	obj, err := anomalySettings.Get(client, id)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}
	saveAnomalySettings(d, obj)
	return nil
}

func updateAnomalySettings(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	id := d.Id()
	obj := parseAnomalySettings(d, id)

	if err := anomalySettings.Update(client, obj); err != nil {
		return diag.FromErr(err)
	}
	return readAnomalySettings(ctx, d, meta)
}

func deleteAnomalySettings(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// There is no way to delete anomaly settings
	return nil
}
