package prismacloud

import (
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/cloud/account/org"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceOrgCloudAccount() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceOrgCloudAccountRead,

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
						org.TypeGcpOrg,
						org.TypeOci,
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
							Description: "Account type - organization or account",
						},
						"member_role_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "AWS Member account role name",
						},
						"member_external_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "AWS Member account role's external ID",
						},
						"member_role_status": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "true = The member role created using stack set exists in all the member accounts. All the Org accounts will be added.\nfalse = Only the master account will be added.",
						},
						"protection_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Monitor or Monitor and Protect",
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
					},
				},
			},

			org.TypeAzureOrg: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Azure account type",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Application ID registered with Active Directory",
						},
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
						"account_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Account type - account or tenant",
						},
						"protection_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Monitor or Monitor and Protect",
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
						"monitor_flow_logs": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Automatically ingest flow logs",
						},
						"key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Application ID key",
							Sensitive:   true,
						},
						"root_sync_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Azure tenant has children. Must be set to true when azure tenant is onboarded with children",
						},
						"hierarchy_selection": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of subscriptions and/or management groups to onboard",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"resource_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Resource ID. Management group ID or subscription ID.\nNote you must escape any double quotes in the resource ID with a backslash.",
									},
									"display_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Display name for management group or subscription",
									},
									"node_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Node type. Valid values: SUBSCRIPTION, TENANT, MANAGEMENT_GROUP",
									},
									"selection_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Selection type. Valid values: INCLUDE to include the specified resource to onboard, EXCLUDE to exclude the specified resource and onboard the rest, ALL to onboard all resources in the tenant.",
									},
								},
							},
						},
					},
				},
			},

			org.TypeGcpOrg: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "GCP account type",
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
							Description: "Name to be used for the account on the Prisma Cloud platform",
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
							Description: "Account type - organization or account",
						},
						"protection_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Protection Mode - Monitor or Monitor and Protect",
						},
						"organization_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "GCP organization name",
						},
						"account_group_creation_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Cloud account group creation mode - manual, auto or recursive",
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
										Description: "Resource ID. For folders, format is folders/{folder ID}. For projects, format is {project number}. For orgs, format is organizations/{org ID}",
									},
									"display_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Display name for folder, project, or organization",
									},
									"node_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Node type - folder, project, org",
									},
									"selection_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Selection type - INCLUDE, EXCLUDE, ALL",
									},
								},
							},
						},
					},
				},
			},

			org.TypeOci: {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Oci account type",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Oci account ID",
						},
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether or not the account is enabled",
						},
						"group_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "OCI identity group name that you define. Can be an existing group",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name to be used for the account on the Prisma Cloud platform (must be unique)",
						},
						"account_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Account type - account or tenant",
						},
						"default_account_group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Account group id for this account",
						},
						"home_region": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "OCI tenancy home region",
						},
						"policy_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "OCI identity policy name that you define. Can be an existing policy that has the right policy statements",
						},
						"user_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "OCI identity user name that you define. Can be an existing user that has the right privileges",
						},
						"user_ocid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "OCI identity user ocid that you define. Can be an existing user that has the right privileges",
						},
					},
				},
			},
		},
	}
}

func dataSourceOrgCloudAccountRead(d *schema.ResourceData, meta interface{}) error {
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
			return err
		}
	}

	obj, err = org.Get(client, cloudType, id)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return err
	}

	if name == "" {
		switch v := obj.(type) {
		case org.AwsOrg:
			name = v.Name
		case org.AzureOrg:
			name = v.Account.Name
		case org.GcpOrg:
			name = v.Account.Name
		case org.Oci:
			name = v.Name
		}
	}

	d.SetId(TwoStringsToId(cloudType, id))
	d.Set("cloud_type", cloudType)
	d.Set("name", name)
	d.Set("account_id", id)

	saveOrgCloudAccount(d, cloudType, obj)

	return nil
}
