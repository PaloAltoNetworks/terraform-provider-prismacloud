package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/cloud/account-v2/org"
	"golang.org/x/net/context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceOrgV2CloudAccount() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceOrgV2CloudAccountRead,

		Schema: map[string]*schema.Schema{
			// Input.
			"cloud_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The cloud type",
				ValidateFunc: validation.StringInSlice(
					[]string{
						org.TypeAwsOrg,
						org.TypeAzureOrg,
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
			//AWS
			org.TypeAwsOrg: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "AWS account type",
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
						"default_account_group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Account group id to which you are assigning this account",
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
							Description: "Account type - organization or account",
						},
						"hierarchy_selection": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of hierarchy selection. Each item has resource id, display name, node type and selection type",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"resource_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Resource ID. Valid values are AWS OU ID, AWS account ID, or AWS Organization ID. Note you must escape any double quotes in the resource ID with a backslash.",
									},
									"display_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Display name for AWS OU, AWS account, or AWS organization",
									},
									"node_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Valid values: OU, ACCOUNT, ORG",
									},
									"selection_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Selection type. Valid values: INCLUDE to include the specified resource to onboard, EXCLUDE to exclude the specified resource and onboard the rest, ALL to onboard all resources in the organization.",
									},
								},
							},
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
							Description: "Cloud type",
						},
						"parent_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Parent id",
						},
						"deleted": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Deleted",
						},
						"protection_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Protection mode",
						},
						"deployment_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Deployment type",
						},
						"customer_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Customer name",
						},
						"created_epoch_millis": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Created epoch millis",
						},
						"last_modified_epoch_millis": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Last modified epoch millis",
						},
						"last_modified_by": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "External id",
						},
						"has_member_role": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Member role",
						},
						"template_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Template url",
						},
						"eventbridge_rule_name_prefix": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "EventbridgeRuleNamePrefix",
						},
						"group_ids": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "List of account IDs to which you are assigning this account",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			//AZURE
			org.TypeAzureOrg: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Azure account type",
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
							Description: "Account type - account or tenant",
						},
						"protection_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Protection mode",
						},
						"cloud_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Cloud Type",
						},
						"default_account_group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Account group id to which you are assigning this account",
						},
						"root_sync_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Azure tenant has children. Must be set to true when azure tenant is onboarded with children",
						},
						"hierarchy_selection": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of hierarchy selection. Each item has resource id, display name, node type and selection type",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"resource_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Resource ID. Valid values are Azure account ID, or AWS tenant ID. Note you must escape any double quotes in the resource ID with a backslash.",
									},
									"display_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Display name for Azure account, or Azure tenant",
									},
									"node_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Valid values: SUBSCRIPTION, TENANT or MANAGEMENT_GROUP",
									},
									"selection_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Selection type. Valid values: INCLUDE to include the specified resource to onboard, EXCLUDE to exclude the specified resource and onboard the rest, ALL to onboard all resources in the organization.",
									},
								},
							},
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
						"parent_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Parent id",
						},
						"customer_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Customer name",
						},
						"created_epoch_millis": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Created epoch millis",
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
						"deleted": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Deleted",
						},
						"template_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Template url",
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
						"member_sync_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Member sync enabled",
						},
					},
				},
			},
		},
	}
}

func dataSourceOrgV2CloudAccountRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	var (
		obj interface{}
		err error
	)

	cloudType := d.Get("cloud_type").(string)
	id := d.Get("account_id").(string)
	name := d.Get("name").(string)

	if id == "" {
		id, err = org.Identify(client, cloudType, name)
		if err != nil {
			if err == pc.ObjectNotFoundError {
				d.SetId("")
				return nil
			}
			return diag.FromErr(err)
		}
	}

	obj, err = org.Get(client, cloudType, id)

	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	if name == "" {
		switch v := obj.(type) {
		case org.AwsOrgV2:
			name = v.Name
		case org.AzureOrgV2:
			name = v.CloudAccountAzureResp.Name
		}
	}

	d.SetId(TwoStringsToId(cloudType, id))
	d.Set("cloud_type", cloudType)
	d.Set("name", name)
	d.Set("account_id", id)

	saveOrgV2CloudAccount(d, cloudType, obj)

	return nil
}
