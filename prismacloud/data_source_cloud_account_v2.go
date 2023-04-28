package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/cloud/account-v2"
	"golang.org/x/net/context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceV2CloudAccount() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceV2CloudAccountRead,

		Schema: map[string]*schema.Schema{
			// Input.
			"cloud_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The cloud type",
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
			accountv2.TypeAws: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "AWS account type",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "AWS account ID",
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
						"role_arn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Unique identifier for an AWS resource (ARN)",
						},
						"storage_scan_config": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Storage scan config",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"scan_option": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Scan option",
									},
									"sns_topic_arn": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Scan topic arn",
									},
									"buckets": {
										Type:        schema.TypeSet,
										Description: "Buckets",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"backward": {
													Type:        schema.TypeSet,
													Computed:    true,
													Description: "Backward",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"forward": {
													Type:        schema.TypeSet,
													Computed:    true,
													Description: "Foward",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
								},
							},
						},
						"storage_vvid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Storage VVID",
						},
						"account_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Account Type",
						},
						"features": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Aws account features",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Feature name",
									},
									"state": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Feature state",
									},
								},
							},
						},
						"cloud_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "",
						},
						"parent_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "",
						},
						"deleted": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "",
						},
						"protection_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "",
						},
						"deployment_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "",
						},
						"customer_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "",
						},
						"created_epoch_millis": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "",
						},
						"last_modified_epoch_millis": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "",
						},
						"last_modified_by": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "",
						},
						"external_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "",
						},
						"has_member_role": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "",
						},
						"template_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "",
						},
						"eventbridge_rule_name_prefix": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "EventbridgeRuleNamePrefix",
						},
					},
				},
			},
			//Azure type.
			accountv2.TypeAzure: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "AWS account type",
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
							Sensitive:   true,
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
							Description: "Unique ID of the service principle object associated with the Prisma Cloud application that you create",
						},
						"account_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Account Type",
						},
						"protection_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "",
						},
						"features": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Azure account features",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Feature name",
									},
									"state": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Feature state",
									},
								},
							},
						},
						"environment_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Environment type",
						},
						"cloud_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "",
						},
						"parent_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "",
						},
						"customer_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "",
						},
						"created_epoch_millis": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "",
						},
						"last_modified_epoch_millis": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "",
						},
						"last_modified_by": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "",
						},
						"deleted": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "",
						},
						"template_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "",
						},
						"deployment_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "",
						},
						"deployment_type_description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "",
						},
					},
				},
			},
			//Gcp type
			accountv2.TypeGcp: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Gcp account type",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "GCP project ID",
						},
						"account_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Account type - organization or account",
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
							Description: "Enable or disable compressed network flow log generation",
						},
						"credentials": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Content of the JSON credentials file",
							Sensitive:   true,
						},
						"dataflow_enabled_project": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Project ID where the Dataflow API is enabled",
						},
						"features": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Gcp account features",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Feature name",
									},
									"state": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Feature state",
									},
								},
							},
						},
						"flow_log_storage_bucket": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Cloud Storage Bucket name that is used store the flow logs",
						},
						"protection_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "",
						},
						"cloud_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "",
						},
						"parent_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "",
						},
						"customer_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "",
						},
						"created_epoch_millis": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "",
						},
						"last_modified_epoch_millis": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "",
						},
						"last_modified_by": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "",
						},
						"deleted": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "",
						},
						"storage_scan_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "",
						},
						"added_on_ts": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "",
						},
						"deployment_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "",
						},
						"deployment_type_description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "",
						},
						"project_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "GCP project ID",
						},
						"service_account_email": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "",
						},
						"authentication_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "",
						},
						"account_group_creation_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "",
						},
						"default_account_group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Account group id to which you are assigning this account. Must be provided for gcp `masterServiceAccount`",
						},
					},
				},
			},
			//IBM type
			accountv2.TypeIbm: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "IBM account type",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IBM account ID",
						},
						"account_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Account type -  account",
						},
						"api_key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IBM service API key",
							Sensitive:   true,
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
						"svc_id_iam_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IBM service ID",
						},
						"added_on_ts": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Added on time stamp",
						},
						"cloud_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Cloud type",
						},
						"created_epoch_millis": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Created epoch millis",
						},
						"customer_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Customer name",
						},
						"deleted": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Deleted",
						},
						"deployment_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Deployment type",
						},
						"deployment_type_description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Deployment type description. Valid values : Commercial or Government",
						},
						"features": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "IBM account features",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Feature name",
									},
									"state": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Feature state",
									},
								},
							},
						},
						"last_modified_epoch_millis": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Last modified epoch millis",
						},
						"last_modified_by": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Last modified by",
						},
						"parent_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "",
						},
						"protection_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Protection mode",
						},
						"storage_scan_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether or not the storage scan is enabled",
						},
					},
				},
			},
		},
	}
}

func dataSourceV2CloudAccountRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	var (
		obj interface{}
		err error
	)

	cloudType := d.Get("cloud_type").(string)
	id := d.Get("account_id").(string)
	name := d.Get("name").(string)

	if id == "" {
		id, err = accountv2.Identify(client, cloudType, name)
		if err != nil {
			if err == pc.ObjectNotFoundError {
				d.SetId("")
				return nil
			}
			return diag.FromErr(err)
		}
	}

	obj, err = accountv2.Get(client, cloudType, id)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	if name == "" {
		switch v := obj.(type) {
		case accountv2.AwsV2:
			name = v.Name
		case accountv2.AzureV2:
			name = v.CloudAccountAzureResp.Name
		case accountv2.GcpV2:
			name = v.CloudAccountGcpResp.Name
		case accountv2.IbmV2:
			name = v.CloudAccountIbmResp.Name
		}
	}

	d.SetId(TwoStringsToId(cloudType, id))
	d.Set("cloud_type", cloudType)
	d.Set("name", name)
	d.Set("account_id", id)

	saveV2CloudAccount(d, cloudType, obj)

	return nil
}
