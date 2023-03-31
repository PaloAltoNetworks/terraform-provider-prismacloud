package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/ip-address"
	"golang.org/x/net/context"
)

func dataSourceTrustedLoginIp() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTrustedLoginIpRead,

		Schema: map[string]*schema.Schema{
			"trusted_login_ip_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "Login IP Allow List ID",
				AtLeastOneOf: []string{"trusted_login_ip_id", "name"},
			},
			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "Unique name for CIDR (IP addresses) allow list",
				AtLeastOneOf: []string{"trusted_login_ip_id", "name"},
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
	}
}

func dataSourceTrustedLoginIpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)

	var err error
	id := d.Get("trusted_login_ip_id").(string)

	if id == "" {
		name := d.Get("name").(string)
		id, err = ip_address.Identify(client, name)
		if err != nil {
			if err == pc.ObjectNotFoundError {
				d.SetId("")
				return nil
			}
			return diag.FromErr(err)
		}
	}

	obj, err := ip_address.Get(client, id)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	d.SetId(obj.Id)
	saveTrustedLoginIpList(d, obj)

	return nil
}
