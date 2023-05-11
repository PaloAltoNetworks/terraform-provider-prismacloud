package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/cloud/account-v2/externalid"
	"golang.org/x/net/context"
)

func dataSourceStorageUUID() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStorageUUIDRead,

		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "AWS account ID",
			},
			"external_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "External id",
			},
			"role_arn": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "AWS Role ARN",
			},
			"storage_uuid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Storage UUID",
			},
		},
	}
}

func dataSourceStorageUUIDRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)

	req := externalid.StorageUUID{
		AccountId:  d.Get("account_id").(string),
		ExternalId: d.Get("external_id").(string),
		RoleArn:    d.Get("role_arn").(string),
	}

	resp, err := externalid.GetStorageUUID(client, req)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.StorageUUID)
	d.Set("storage_uuid", resp.StorageUUID)
	return nil
}
