package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/anomalySettings"
	"golang.org/x/net/context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAnomalySetting() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAnomalySettingRead,

		Schema: map[string]*schema.Schema{
			"policy_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Policy ID",
			},
			"alert_disposition": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Alert disposition",
			},
			"alert_disposition_description": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Alert Disposition info",
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
				Description: "Policy Name",
			},
			"training_model_description": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Training Model Threshold info",
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
				Computed:    true,
				Description: "Training model threshold",
			},
		},
	}
}

func dataSourceAnomalySettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	id := d.Get("policy_id").(string)

	obj, err := anomalySettings.Get(client, id)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}
	d.SetId(id)
	saveAnomalySettings(d, obj)

	return nil
}
