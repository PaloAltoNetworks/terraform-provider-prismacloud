package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/cloud/account-v2/azureTemplate"
	"golang.org/x/net/context"
)

func dataSourceAzureTemplate() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAzureTemplateRead,

		Schema: map[string]*schema.Schema{
			"account_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The azure account type",
				ValidateFunc: validation.StringInSlice(
					[]string{
						"account",
						"tenant",
					},
					false,
				),
			},
			"subscription_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Azure subscription ID",
			},
			"features": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Features applicable for azure account",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"tenant_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Tenant id",
			},
			"root_sync_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Azure tenant has children. Must be set to true when azure tenant is onboarded with children",
			},
			"deployment_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Deployment type",
				Default:     "azure",
				ValidateFunc: validation.StringInSlice(
					[]string{
						"azure",
						"azure_gov",
					},
					false,
				),
			},
			"file_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "File name to store azure template",
			},
		},
	}
}

func dataSourceAzureTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)

	req := azureTemplate.AzureTemplateReq{
		AccountType:     d.Get("account_type").(string),
		SubscriptionId:  d.Get("subscription_id").(string),
		TenantId:        d.Get("tenant_id").(string),
		DeploymentType:  d.Get("deployment_type").(string),
		RootSyncEnabled: d.Get("root_sync_enabled").(bool),
		Features:        SetToStringSlice(d.Get("features").(*schema.Set)),
		FileName:        d.Get("file_name").(string),
	}

	err := azureTemplate.GetAzureTemplate(client, req)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(d.Get("tenant_id").(string))
	d.Set("subscription_id", d.Get("subscription_id").(string))
	d.Set("deployment_type", d.Get("deployment_type").(string))
	d.Set("account_type", d.Get("account_type").(string))
	d.Set("tenant_id", d.Get("tenant_id").(string))
	d.Set("file_name", d.Get("file_name").(string))
	d.Set("root_sync_enabled", d.Get("root_sync_enabled").(bool))
	d.Set("features", SetToStringSlice(d.Get("features").(*schema.Set)))

	return nil
}
