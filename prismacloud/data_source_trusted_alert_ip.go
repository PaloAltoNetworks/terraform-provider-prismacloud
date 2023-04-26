package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/trusted-alert-ip"
	"golang.org/x/net/context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTrustedAlertIp() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTrustedAlertIpRead,

		Schema: map[string]*schema.Schema{
			"uuid": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "Trusted alert ip ID",
				AtLeastOneOf: []string{"name", "uuid"},
			},
			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "Trusted alert ip name",
				AtLeastOneOf: []string{"name", "uuid"},
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
	}
}

func dataSourceTrustedAlertIpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)

	var err error
	id := d.Get("uuid").(string)

	if id == "" {
		name := d.Get("name").(string)
		id, err = trustedalertip.Identify(client, name)
		if err != nil {
			if err == pc.ObjectNotFoundError {
				d.SetId("")
				return nil
			}
			return diag.FromErr(err)
		}
	}

	obj, err := trustedalertip.Get(client, id)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	d.SetId(obj.UUID)
	saveTrustedAlertIp(d, obj)

	return nil
}
