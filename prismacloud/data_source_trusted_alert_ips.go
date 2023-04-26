package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/paloaltonetworks/prisma-cloud-go/trusted-alert-ip"
	"golang.org/x/net/context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
)

func dataSourceTrustedAlertIps() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTrustedAlertIpsRead,

		Schema: map[string]*schema.Schema{
			// Output.
			"total": totalSchema("trusted alert ips"),
			"listing": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of trusted alert ips",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Trusted alert ip ID",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Trusted alert ip name",
						},
						"cidrs": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "CIDRs",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cidr": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "CIDR",
									},
									"uuid": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "UUID",
									},
									"created_on": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Created on",
									},
									"description": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Description",
									},
								},
							},
						},
						"cidr_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Associated cloud account type",
						},
					},
				},
			},
		},
	}
}

func dataSourceTrustedAlertIpsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)

	items, err := trustedalertip.List(client)
	if err != nil {
		return diag.FromErr(err)
	}

	ans := make([]interface{}, 0, len(items))
	for _, v := range items {
		cidrInfo := make([]interface{}, 0, len(v.CIDRS))
		for _, cidr := range v.CIDRS {
			cidrInfo = append(cidrInfo, map[string]interface{}{
				"uuid":        cidr.UUID,
				"description": cidr.Description,
				"created_on":  cidr.CreatedOn,
				"cidr":        cidr.CIDR,
			})
		}
		ans = append(ans, map[string]interface{}{
			"uuid":       v.UUID,
			"name":       v.Name,
			"cidrs":      cidrInfo,
			"cidr_count": v.CidrCount,
		})
	}

	d.SetId("trustedAlertIps")
	d.Set("total", len(items))
	if err = d.Set("listing", ans); err != nil {
		log.Printf("[WARN] Error setting 'listing' field for %q: %s", d.Id(), err)
	}

	return nil
}
