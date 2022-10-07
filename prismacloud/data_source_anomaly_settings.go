package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/anomalySettings"
	"golang.org/x/net/context"
	"log"
	"sort"
)

func dataSourceAnomalySettings() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAnomalySettingsRead,

		Schema: map[string]*schema.Schema{
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "type",
			},
			"total": totalSchema("cloud accounts"),
			"listing": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of accounts",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"policy_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Policy id",
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
										Description: "low",
									},
									"medium": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "medium",
									},
									"high": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "high",
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
				},
			},
		},
	}
}

func dataSourceAnomalySettingsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	t := d.Get("type").(string)
	ans, err := anomalySettings.List(client, t)
	if err != nil {
		return diag.FromErr(err)
	}

	keys := make([]string, 0, len(ans))
	for key := range ans {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	data := make([]interface{}, 0, len(ans))

	for _, k := range keys {
		i := ans[k]
		var adp1 = i.(map[string]interface{})["alertDispositionDescription"]
		adp := make(map[string]interface{})
		if adp1 != nil {
			adp = map[string]interface{}{
				"aggressive":   adp1.(map[string]interface{})["aggressive"],
				"moderate":     adp1.(map[string]interface{})["moderate"],
				"conservative": adp1.(map[string]interface{})["conservative"],
			}
		}

		var tmd1 = i.(map[string]interface{})["trainingModelDescription"]
		tmd := make(map[string]interface{})
		if tmd1 != nil {
			tmd = map[string]interface{}{
				"low":    tmd1.(map[string]interface{})["low"],
				"medium": tmd1.(map[string]interface{})["medium"],
				"high":   tmd1.(map[string]interface{})["high"],
			}
		}

		item := map[string]interface{}{
			"policy_id":                     k,
			"alert_disposition":             i.(map[string]interface{})["alertDisposition"],
			"alert_disposition_description": []interface{}{adp},
			"policy_description":            i.(map[string]interface{})["policyDescription"],
			"policy_name":                   i.(map[string]interface{})["policyName"],
			"training_model_description":    []interface{}{tmd},
			"training_model_threshold":      i.(map[string]interface{})["trainingModelThreshold"],
		}
		data = append(data, item)
	}

	d.SetId("anomaly_settings")
	d.Set("total", len(ans))

	if err := d.Set("listing", data); err != nil {
		log.Printf("[WARN] Error setting 'anomaly_settings' field for %q: %s", d.Id(), err)
	}
	return nil
}
