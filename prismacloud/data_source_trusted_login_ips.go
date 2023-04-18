package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"golang.org/x/net/context"
	"log"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/ip-address"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTrustedLoginIps() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTrustedLoginIpsRead,

		Schema: map[string]*schema.Schema{
			//Output
			"total": totalSchema("trusted_login_ip"),
			"listing": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of trusted_login_ips",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"trusted_login_ip_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Login IP Allow List ID",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Unique name for CIDR (IP addresses) allow list",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description of CIDR (IP addresses) allow list",
						},
						"cidr": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "List of CIDRs to Allow List for login access. You can include from 1 to 10 CIDRs",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"last_modified_ts": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Timestamp for last modification of CIDR block list",
						},
					},
				},
			},
		},
	}
}

func dataSourceTrustedLoginIpsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)

	items, err := ip_address.List(client)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("trusted_login_ips_list")
	d.Set("total", len(items))

	list := make([]interface{}, 0, len(items))
	for _, i := range items {
		list = append(list, map[string]interface{}{
			"trusted_login_ip_id": i.Id,
			"name":                i.Name,
			"cidr":                i.Cidr,
			"description":         i.Description,
			"last_modified_ts":    i.LastModifiedTs,
		})
	}

	if err := d.Set("listing", list); err != nil {
		log.Printf("[WARN] Error setting 'listing' field for %q: %s", d.Id(), err)
	}

	return nil
}
