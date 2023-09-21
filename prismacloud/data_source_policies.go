package prismacloud

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"golang.org/x/net/context"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/policy"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePolicies() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePoliciesRead,

		Schema: map[string]*schema.Schema{
			// Input.
			"filters": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Filter policy results",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Output.
			"total": totalSchema("policies"),
			"listing": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of policies",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"policy_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Policy ID",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Policy name",
						},
						"policy_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Policy type",
						},
						"system_default": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "If the policy is a system default for Prisma Cloud",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description",
						},
						"severity": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Severity",
						},
						"recommendation": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Remediation recommendation",
						},
						"cloud_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Cloud type",
						},
						"labels": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Labels",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enabled",
						},
						"overridden": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Overridden",
						},
						"deleted": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Deleted",
						},
						"open_alerts_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Open alerts count",
						},
						"policy_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Policy mode",
						},
						"remediable": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Remediable",
						},
					},
				},
			},
		},
	}
}

func dataSourcePoliciesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var buf bytes.Buffer
	client := meta.(*pc.Client)

	filters := d.Get("filters").(map[string]interface{})
	query := make(map[string]string)
	for key := range filters {
		if buf.Len() > 0 {
			buf.WriteString("&")
		}
		query[key] = filters[key].(string)
		buf.WriteString(fmt.Sprintf("%s=%s", key, query[key]))
	}

	items, err := policy.List(client, query)
	if err != nil {
		return diag.FromErr(err)
	}

	if buf.Len() == 0 {
		d.SetId("all")
	} else {
		d.SetId(base64.StdEncoding.EncodeToString(buf.Bytes()))
	}
	d.Set("total", len(items))

	list := make([]interface{}, 0, len(items))
	for _, i := range items {
		list = append(list, map[string]interface{}{
			"policy_id":         i.PolicyId,
			"name":              i.Name,
			"policy_type":       i.PolicyType,
			"system_default":    i.SystemDefault,
			"description":       i.Description,
			"severity":          i.Severity,
			"recommendation":    i.Recommendation,
			"cloud_type":        i.CloudType,
			"labels":            StringSliceToSet(i.Labels),
			"enabled":           i.Enabled,
			"overridden":        i.Overridden,
			"deleted":           i.Deleted,
			"open_alerts_count": i.OpenAlertsCount,
			"policy_mode":       i.PolicyMode,
			"remediable":        i.Remediable,
		})
	}

	if err := d.Set("listing", list); err != nil {
		log.Printf("[WARN] Error setting 'listing' field for %q: %s", d.Id(), err)
	}

	return nil
}
