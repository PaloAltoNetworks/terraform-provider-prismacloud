package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/cloud/account-v2/supportedFeatures"
	"golang.org/x/net/context"
)

func dataSourceCloudAccountSupportedFeatures() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCloudAccountSupportedFeaturesRead,

		Schema: map[string]*schema.Schema{
			"account_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Cloud Account Type",
				ValidateFunc: validation.StringInSlice(
					[]string{
						"account",
						"organization",
						"masterServiceAccount",
						"tenant",
						"workspace_domain",
					},
					false,
				),
			},
			"deployment_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "*Applicable only for cloud_type: **azure**.*\n\n * **azure** -  Account type is commercial\n\n * **azure_gov** - Account type is Government on Prisma Commercial and Government stacks.\n\n * **azure_china** - Prisma China Stack.",
				ValidateFunc: validation.StringInSlice(
					[]string{
						"azure",
						"azure_gov",
						"azure_china",
					},
					false,
				),
			},
			"aws_partition": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "*Applicable only for cloud_type: **aws** on Prisma Government Stack(**app.gov.prismacloud.io**) given if the Cloud account Global Deployment option is enabled*\n\n * **us-east-1** -  AWS Commercial/Global account\n\n * **us-gov-west-1** - AWS GovCloud account.",
				ValidateFunc: validation.StringInSlice(
					[]string{
						"us-east-1",
						"us-gov-west-1",
						"cn-north-1",
					},
					false,
				),
			},
			"root_sync_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "AWS account ID",
			},
			"cloud_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Cloud Type",
			},
			"license_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Customer License type.",
			},
			"supported_features": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "List of supported feature names.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"supported_features_all": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "List of all supported features including the default 'Cloud Visibility Compliance and Governance' feature",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceCloudAccountSupportedFeaturesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)

	req := supportedFeatures.SupportedFeaturesReq{
		CloudType:       d.Get("cloud_type").(string),
		AccountType:     d.Get("account_type").(string),
		DeploymentType:  d.Get("deployment_type").(string),
		AwsPartition:    d.Get("aws_partition").(string),
		RootSyncEnabled: d.Get("root_sync_enabled").(bool),
	}

	resp, err := supportedFeatures.GetSupportedFeatures(client, req)
	if err != nil {
		return diag.FromErr(err)
	}
	supported_features_all := resp.SupportedFeatures
	d.Set("supported_features_all", resp.SupportedFeatures)
	// Remove default feature 'Cloud Visibility Compliance and Governance'. Since its not required to be passed in any api.
	index_of_default_feature := indexOf("Cloud Visibility Compliance and Governance", supported_features_all)
	if index_of_default_feature != -1 {
		supported_features := RemoveIndex(supported_features_all, index_of_default_feature)
		d.Set("supported_features", supported_features)
	} else {
		d.Set("supported_features", resp.SupportedFeatures)
	}

	d.SetId(resp.AccountType)
	d.Set("cloud_type", resp.CloudType)
	d.Set("deployment_type", resp.DeploymentType)
	d.Set("account_type", resp.AccountType)
	d.Set("license_type", resp.LicenseType)
	return nil
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
