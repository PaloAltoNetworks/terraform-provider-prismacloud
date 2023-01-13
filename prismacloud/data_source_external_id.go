package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/cloud/account/externalid"
	"golang.org/x/net/context"
)

func dataSourceExternalId() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceExternalIdRead,

		Schema: map[string]*schema.Schema{
			"account_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The aws account type",
				ValidateFunc: validation.StringInSlice(
					[]string{
						"account",
						"organization",
					},
					false,
				),
			},
			"account_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "AWS account ID",
			},
			"aws_partition": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The aws cloud account partition",
				ValidateFunc: validation.StringInSlice(
					[]string{
						"us-gov-west-1",
						"cn-north-1",
						"us-east-1",
					},
					false,
				),
			},
			"features": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Features applicable for aws account",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"external_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "External id",
			},
			"create_stack_link_with_s3_presigned_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "CFT link",
			},
			"event_bridge_rule_name_prefix": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "AWS account eventBridge rule name prefix",
			},
		},
	}
}

func dataSourceExternalIdRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)

	req := externalid.ExternalIdReq{
		AccountType:  d.Get("account_type").(string),
		AccountId:    d.Get("account_id").(string),
		AwsPartition: d.Get("aws_partition").(string),
		Features:     SetToStringSlice(d.Get("features").(*schema.Set)),
	}

	resp, err := externalid.GetExternalId(client, req)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.ExternalId)
	d.Set("external_id", resp.ExternalId)
	d.Set("create_stack_link_with_s3_presigned_url", resp.CreateStackLinkWithS3PresignedUrl)
	d.Set("event_bridge_rule_name_prefix", resp.EventBridgeRuleNamePrefix)
	return nil
}
