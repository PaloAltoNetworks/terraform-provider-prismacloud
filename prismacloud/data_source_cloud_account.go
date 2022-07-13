package prismacloud

import (
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/cloud/account"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceCloudAccount() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCloudAccountRead,

		Schema: map[string]*schema.Schema{
			// Input.
			"cloud_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The cloud type",
				ValidateFunc: validation.StringInSlice(
					[]string{
						account.TypeAws,
						account.TypeAzure,
						account.TypeGcp,
						account.TypeAlibaba,
						account.TypeAwsEventBridge,
					},
					false,
				),
			},
			"account_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "The cloud account ID",
				AtLeastOneOf: []string{"account_id", "name"},
			},
			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "The cloud account name",
				AtLeastOneOf: []string{"account_id", "name"},
			},

			// Output.
			// AWS type.
			account.TypeAws: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "AWS account type",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "AWS account ID",
						},
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether or not the account is enabled",
						},
						"external_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "AWS account external ID",
							Sensitive:   true,
						},
						"group_ids": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "List of account IDs to which you are assigning this account",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name to be used for the account on the Prisma Cloud platform (must be unique)",
						},
						"role_arn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Unique identifier for an AWS resource (ARN)",
						},
						"account_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Account Type",
						},
						"protection_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Protection Mode setting",
						},
						"eb_rule_name_prefix": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "EventBridge Rule Name Prefix",
						},
					},
				},
			},

			// Aws Event bridge enabled account
			account.TypeAwsEventBridge: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "AWS account type",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "AWS account ID",
						},
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether or not the account is enabled",
						},
						"external_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "AWS account external ID",
							Sensitive:   true,
						},
						"group_ids": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "List of account IDs to which you are assigning this account",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name to be used for the account on the Prisma Cloud platform (must be unique)",
						},
						"role_arn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Unique identifier for an AWS resource (ARN)",
						},
						"account_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Account Type",
						},
						"protection_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Protection Mode setting",
						},
						"eb_rule_name_prefix": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "EventBridge Rule Name Prefix",
						},
					},
				},
			},
			// Azure type.
			account.TypeAzure: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Azure account type",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Azure account ID",
						},
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether or not the account is enabled",
						},
						"group_ids": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "List of account IDs to which you are assigning this account",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name to be used for the account on the Prisma Cloud platform (must be unique)",
						},
						"client_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Application ID registered with Active Directory",
						},
						"key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Application ID key",
						},
						"monitor_flow_logs": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Automatically ingest flow logs",
						},
						"tenant_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Active Directory ID associated with Azure",
						},
						"service_principal_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Unique ID of the service principal object associated with the Prisma Cloud application that you create",
						},
						"account_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Account Type",
						},
						"protection_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Protection Mode setting",
						},
					},
				},
			},

			// GCP type.
			account.TypeGcp: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "GCP account type",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "GCP project ID",
						},
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether or not the account is enabled",
						},
						"group_ids": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "List of account IDs to which you are assigning this account",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name to be used for the account on the Prisma Cloud platform (must be unique)",
						},
						"compression_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable flow log compression",
						},
						"dataflow_enabled_project": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "GCP project for flow log compression",
						},
						"flow_log_storage_bucket": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "GCP flow logs storage bucket",
						},
						// Use a json string until this feature is added:
						// https://github.com/hashicorp/terraform-plugin-sdk/issues/248
						"credentials_json": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Content of the JSON credentials file",
							Sensitive:   true,
						},
						"account_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Account Type",
						},
						"protection_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Protection Mode setting",
						},
					},
				},
			},

			// Alibaba type.
			account.TypeAlibaba: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Alibaba account type",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Alibaba account ID",
						},
						"group_ids": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "List of account IDs to which you are assigning this account",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name to be used for the account on the Prisma Cloud platform (must be unique)",
						},
						"ram_arn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Unique identifier for an Alibaba RAM role resource",
						},
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether or not the account is enabled",
						},
					},
				},
			},
		},
	}
}

func dataSourceCloudAccountRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	var (
		obj interface{}
		err error
	)

	cloudType := d.Get("cloud_type").(string)
	id := d.Get("account_id").(string)
	name := d.Get("name").(string)

	if id == "" {
		id, err = account.Identify(client, cloudType, name)
		if err != nil {
			if err == pc.ObjectNotFoundError {
				d.SetId("")
				return nil
			}
			return err
		}
	}

	obj, err = account.Get(client, cloudType, id)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return err
	}

	if name == "" {
		switch v := obj.(type) {
		case account.Aws:
			name = v.Name
		case account.AwsEventBridge:
			name = v.Name
		case account.Azure:
			name = v.Account.Name
		case account.Gcp:
			name = v.Account.Name
		case account.Alibaba:
			name = v.Name
		}
	}

	d.SetId(TwoStringsToId(cloudType, id))
	d.Set("cloud_type", cloudType)
	d.Set("name", name)
	d.Set("account_id", id)

	saveCloudAccount(d, cloudType, obj)

	return nil
}
