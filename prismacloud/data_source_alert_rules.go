package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"golang.org/x/net/context"
	"log"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/alert/rule"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAlertRules() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAlertRulesRead,

		Schema: map[string]*schema.Schema{
			// Output.
			"total": totalSchema("alert rules"),
			"listing": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of alert rules",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"policy_scan_config_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Policy scan config ID",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Rule/Scan name",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description",
						},
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enabled",
						},
						"scan_all": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Scan all policies",
						},
						"policies": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "List of specific policies to scan",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"owner": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Owner",
						},
						"open_alerts_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Open alerts count",
						},
						"read_only": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Read only",
						},
						"deleted": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Deleted",
						},
					},
				},
			},
		},
	}
}

func dataSourceAlertRulesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var err error
	client := meta.(*pc.Client)

	items, err := rule.List(client)
	if err != nil {
		return diag.FromErr(err)
	}

	ans := make([]interface{}, 0, len(items))
	for _, v := range items {
		ans = append(ans, map[string]interface{}{
			"policy_scan_config_id": v.PolicyScanConfigId,
			"name":                  v.Name,
			"description":           v.Description,
			"enabled":               v.Enabled,
			"scan_all":              v.ScanAll,
			"policies":              StringSliceToSet(v.Policies),
			"owner":                 v.Owner,
			"open_alerts_count":     v.OpenAlertsCount,
			"read_only":             v.ReadOnly,
			"deleted":               v.Deleted,
		})
	}

	d.SetId("alert_rules")
	d.Set("total", len(items))
	if err = d.Set("listing", ans); err != nil {
		log.Printf("[WARN] Error setting 'listing' field for %q: %s", d.Id(), err)
	}

	return nil
}
