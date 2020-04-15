package prismacloud

import (
	"log"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/alert/rule"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceAlertRules() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAlertRulesRead,

		Schema: map[string]*schema.Schema{
			// Output.
			"rules": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of alert rules",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"policy_scan_config_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Policy ID",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Policy name",
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

func dataSourceAlertRulesRead(d *schema.ResourceData, meta interface{}) error {
	var err error
	client := meta.(*pc.Client)

	items, err := rule.List(client)
	if err != nil {
		return err
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
	if err = d.Set("rules", ans); err != nil {
		log.Printf("[WARN] Error setting 'rules' field for %q: %s", d.Id(), err)
	}

	return nil
}
