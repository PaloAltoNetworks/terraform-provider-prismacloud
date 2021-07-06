package prismacloud

import (
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/cloud/account/org"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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
				Optional:    true,
				Description: "AWS account type",
				MaxItems:    1,
				ConflictsWith: []string{
					org.TypeGcpOrg,
					org.TypeAzureOrg,
					org.TypeOci,
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "AWS account ID",
						},
						"enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether or not the account is enabled",
							Default:     true,
						},
						"external_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "AWS account external ID",
							Sensitive:   true,
						},
						"group_ids": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of account IDs to which you are assigning this account",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Name to be used for the account on the Prisma Cloud platform (must be unique)",
						},
						"role_arn": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique identifier for an AWS resource (ARN)",
						},
						"account_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "organization",
							Description: "Account type - organization or account",
						},
						"member_role_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "AWS Member account role name",
						},
						"member_external_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "AWS Member account role's external ID",
						},
						"member_role_status": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "true = The member role created using stack set exists in all the member accounts. All the Org accounts will be added.\nfalse = Only the master account will be added.",
						},
						"protection_mode": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "MONITOR",
							Description: "Monitor or Monitor and Protect",
						},
					},
				},
			},

			org.TypeAzureOrg: {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Azure account type",
				MaxItems:    1,
				ConflictsWith: []string{
					org.TypeGcpOrg,
					org.TypeAwsOrg,
					org.TypeOci,
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Application ID registered with Active Directory",
						},
						"account_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Azure account ID",
						},
						"enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether or not the account is enabled",
							Default:     true,
						},
						"group_ids": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of account IDs to which you are assigning this account",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Name to be used for the account on the Prisma Cloud platform (must be unique)",
						},
						"account_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "tenant",
							Description: "Account type - account or tenant",
						},
						"protection_mode": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "MONITOR",
							Description: "Monitor or Monitor and Protect",
						},
						"tenant_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Active Directory ID associated with Azure",
						},
						"service_principal_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique ID of the service principal object associated with the Prisma Cloud application that you create",
						},
						"monitor_flow_logs": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Automatically ingest flow logs",
						},
						"key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Application ID key",
							Sensitive:   true,
						},
					},
				},
			},

			org.TypeGcpOrg: {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "GCP account type",
				MaxItems:    1,
				ConflictsWith: []string{
					org.TypeAwsOrg,
					org.TypeAzureOrg,
					org.TypeOci,
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "GCP project ID",
						},
						"enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether or not the account is enabled",
							Default:     true,
						},
						"group_ids": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of account IDs to which you are assigning this account",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Name to be used for the account on the Prisma Cloud platform (must be unique)",
						},
						"compression_enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable flow log compression",
						},
						"dataflow_enabled_project": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "GCP project for flow log compression",
						},
						"flow_log_storage_bucket": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "GCP flow logs storage bucket",
						},
						// Use a json string until this feature is added:
						// https://github.com/hashicorp/terraform-plugin-sdk/issues/248
						"credentials_json": {
							Type:             schema.TypeString,
							Required:         true,
							Description:      "Content of the JSON credentials file",
							Sensitive:        true,
							DiffSuppressFunc: gcpOrgCredentialsMatch,
						},
						"account_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "organization",
							Description: "Account type - organization or account",
						},
						"protection_mode": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "MONITOR",
							Description: "Monitor or Monitor and Protect",
						},
						"organization_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "GCP organization name",
						},
						"account_group_creation_mode": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "MANUAL",
							Description: "Cloud account group creation mode - manual, auto or recursive",
						},
						"hierarchy_selection": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"resource_id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Resource ID. For folders, format is folders/{folder ID}. For projects, format is {project number}. For orgs, format is organizations/{org ID}",
									},
									"display_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Display name for folder, project, or organization",
									},
									"node_type": {
										Type:        schema.TypeString,
										Optional:    true,
										Default:     "FOLDER",
										Description: "Valid values - folder, project, org",
									},
									"selection_type": {
										Type:        schema.TypeString,
										Optional:    true,
										Default:     "EXCLUDE",
										Description: "Valid values: INCLUDE, EXCLUDE, INCLUDE ALL. If hierarchySelection.nodeType is PROJECT or FOLDER, then a valid value is either INCLUDE or EXCLUDE",
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
				MaxItems:    1,
				ConflictsWith: []string{
					org.TypeAwsOrg,
					org.TypeAzureOrg,
					org.TypeGcpOrg,
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Azure account ID",
						},
						"enabled": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Whether or not the account is enabled",
						},
						"group_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "OCI identity group name that you define. Can be an existing group",
						},
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Name to be used for the account on the Prisma Cloud platform (must be unique)",
						},
						"account_type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Account type - account or tenant",
						},
						"default_account_group_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Account group id for this account",
						},
						"home_region": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "OCI tenancy home region",
						},
						"policy_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "OCI identity policy name that you define. Can be an existing policy that has the right policy statements",
						},
						"user_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "OCI identity user name that you define. Can be an existing user that has the right privileges",
						},
						"user_ocid": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "OCI identity user name that you define. Can be an existing user that has the right privileges",
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
