package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"golang.org/x/net/context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/paloaltonetworks/prisma-cloud-go/anomalySettings/anomalyTrustedList"
	"strconv"
)

func dataSourceAnomalyTrustedList() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAnomalyTrustedListRead,

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
				Description: "Anomaly Trusted List account id.",
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
				Description: "Anomaly Trusted List name",
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
	}
}

func dataSourceAnomalyTrustedListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)

	var err error
	id := d.Get("atl_id").(int)

	obj, err := anomalyTrustedList.Get(client, strconv.Itoa(id))
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(id))
	saveAnomalyTrustedList(d, obj)
	return nil
}
