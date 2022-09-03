package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/mitchellh/mapstructure"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/anomalySettings/anomalyTrustedList"
	"golang.org/x/net/context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"sort"
)

func dataSourceAnomalyTrustedLists() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAnomalyTrustedListsRead,

		Schema: map[string]*schema.Schema{
			// Output.
			"total": totalSchema("anomaly trusted list"),
			"listing": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of accounts",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"atl_id": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Anomaly Trusted List ID",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Anomaly Trusted List name",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Reason for trusted listing",
						},
						"trusted_list_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Anomaly Trusted List type",
						},
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Anomaly Trusted List account id",
						},
						"applicable_policies": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Applicable Policies",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"vpc": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "VPC",
						},
						"created_by": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Created by",
						},
						"created_on": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Created on",
						},
						"trusted_list_entries": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of network anomalies in the trusted list",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"image_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Image ID",
									},
									"tag_key": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Tag key",
									},
									"tag_value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Tag value",
									},
									"ip_cidr": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Ip CIDR",
									},
									"port": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Port",
									},
									"resource_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Resource ID",
									},
									"service": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Service",
									},
									"subject": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Subject",
									},
									"domain": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Domain",
									},
									"protocol": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Protocol",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceAnomalyTrustedListsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var err error
	client := meta.(*pc.Client)

	items, err := anomalyTrustedList.List(client)
	if err != nil {
		return diag.FromErr(err)
	}

	keys := make([]int, 0, len(items))
	for k := range items {
		keys = append(keys, k)
	}
	sort.Ints(keys) //keys are sorted now

	ans := make([]interface{}, 0, len(items))

	for key := range keys {
		tles := make([]interface{}, 0, len(items[key].TrustedListEntries))
		v := items[key]

		result := anomalyTrustedList.AnomalyTrustedList{}
		mapstructure.Decode(v, &result)

		for _, tle := range v.TrustedListEntries {
			tles = append(tles, map[string]interface{}{
				"tag_value":   tle.TagValue,
				"image_id":    tle.ImageID,
				"ip_cidr":     tle.IpCIDR,
				"port":        tle.Port,
				"tag_key":     tle.TagKey,
				"resource_id": tle.ResourceID,
				"service":     tle.Service,
				"subject":     tle.Subject,
				"domain":      tle.Domain,
				"protocol":    tle.Protocol,
			})
		}
		ans = append(ans, map[string]interface{}{
			"atl_id":               result.Atl_Id,
			"name":                 result.Name,
			"description":          result.Description,
			"trusted_list_type":    result.TrustedListType,
			"account_id":           result.AccountId,
			"vpc":                  result.VPC,
			"created_by":           result.CreatedBy,
			"created_on":           result.CreatedOn,
			"trusted_list_entries": tles,
			"applicable_policies":  StringSliceToSet(result.ApplicablePolicies),
		})
	}

	d.SetId("anomaly_trusted_list")
	d.Set("total", len(items))

	if err = d.Set("listing", ans); err != nil {
		log.Printf("[WARN] Error setting 'listing' field for %q: %s", d.Id(), err)
	}
	return nil
}
