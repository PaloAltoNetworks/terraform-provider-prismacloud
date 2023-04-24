package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/cloud/account-v2/gcpTemplate"
	"golang.org/x/net/context"
)

func dataSourceGcpTemplate() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGcpTemplateRead,

		Schema: map[string]*schema.Schema{
			"account_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The gcp account type",
				ValidateFunc: validation.StringInSlice(
					[]string{
						"account",
						"organization",
						"masterServiceAccount",
					},
					false,
				),
			},
			"authentication_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Authentication type",
				Default:     "service_account",
				ValidateFunc: validation.StringInSlice(
					[]string{
						"service_account",
					},
					false,
				),
			},
			"project_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Gcp Project ID",
			},
			"features": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Features applicable for gcp account",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"org_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Gcp organization ID",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name to be used for the account on the Prisma Cloud platform (must be unique)",
			},
			"file_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "File name to store gcp template",
			},
			"flow_log_storage_bucket": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Cloud Storage Bucket name that is used store the flow logs",
			},
		},
	}
}
func dataSourceGcpTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)

	req := gcpTemplate.GcpTemplateReq{
		AccountType:          d.Get("account_type").(string),
		ProjectId:            d.Get("project_id").(string),
		OrgId:                d.Get("org_id").(string),
		FileName:             d.Get("file_name").(string),
		Name:                 d.Get("name").(string),
		AuthenticationType:   d.Get("authentication_type").(string),
		FlowLogStorageBucket: d.Get("flow_log_storage_bucket").(string),
		Features:             SetToStringSlice(d.Get("features").(*schema.Set)),
	}

	err := gcpTemplate.GetGcpTemplate(client, req)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(d.Get("account_type").(string))
	d.Set("project_id", d.Get("project_id").(string))
	d.Set("authentication_type", d.Get("authentication_type").(string))
	d.Set("account_type", d.Get("account_type").(string))
	d.Set("org_id", d.Get("org_id").(string))
	d.Set("name", d.Get("name").(string))
	d.Set("file_name", d.Get("file_name").(string))
	d.Set("flow_log_storage_bucket", d.Get("flow_log_storage_bucket").(string))
	d.Set("features", SetToStringSlice(d.Get("features").(*schema.Set)))

	return nil
}
