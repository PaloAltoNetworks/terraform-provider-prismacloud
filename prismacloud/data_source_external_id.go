package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/cloud/account/externalid"
)

func dataSourceExternalId() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceExternalIdRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The cloud account name",
			},
			"aws_partition": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The aws cloud account partition",
			},
			"account_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "AWS account ID",
			},
			"protection_mode": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Monitor or Monitor and Protect",
				ValidateFunc: validation.StringInSlice(
					[]string{
						"MONITOR",
						"MONITOR_AND_PROTECT",
					},
					false,
				),
			},
			"storage_scan_enabled": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Enables storage scan for AWS account",
			},
			"external_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "AWS account external ID",
				Sensitive:   true,
			},
			"cft_path": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "AWS account cft path",
			},
			"cloud_formation_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "AWS account cloud formation url",
			},
		},
	}
}

func dataSourceExternalIdRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)

	req := externalid.ExternalIdReq{
		Name:               d.Get("name").(string),
		AccountId:          d.Get("account_id").(string),
		ProtectionMode:     d.Get("protection_mode").(string),
		AwsPartition:       d.Get("aws_partition").(string),
		StorageScanEnabled: d.Get("storage_scan_enabled").(bool),
	}

	resp, err := externalid.GetExternalId(client, req)
	if err != nil {
		return err
	}

	d.SetId(resp.ExternalId)
	d.Set("external_id", resp.ExternalId)
	d.Set("cft_path", resp.CftPath)
	d.Set("cloud_formation_url", resp.CloudFormationUrl)
	return nil
}
