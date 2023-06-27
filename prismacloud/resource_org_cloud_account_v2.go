package prismacloud

import (
	"context"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	accountv2 "github.com/paloaltonetworks/prisma-cloud-go/cloud/account-v2"
	"log"
	"strings"
	"time"

	"github.com/paloaltonetworks/prisma-cloud-go/cloud/account-v2/org"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceOrgV2CloudAccount() *schema.Resource {
	return &schema.Resource{
		CreateContext: createOrgV2CloudAccount,
		ReadContext:   readOrgV2CloudAccount,
		UpdateContext: updateOrgV2CloudAccount,
		DeleteContext: deleteOrgV2CloudAccount,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"disable_on_destroy": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "to disable cloud account instead of deleting on calling destroy",
				Default:     false,
			},

			org.TypeAwsOrg: {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "AWS account type",
				MaxItems:    1,
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
						"default_account_group_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Account group id to which you are assigning this account",
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
							ValidateFunc: validation.StringInSlice(
								[]string{
									"organization",
									"account",
								},
								true,
							),
						},
						"hierarchy_selection": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "List of hierarchy selection. Each item has resource id, display name, node type and selection type",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"resource_id": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Resource ID. Valid values are AWS OU ID, AWS account ID, or AWS Organization ID. Note you must escape any double quotes in the resource ID with a backslash.",
									},
									"display_name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Display name for AWS OU, AWS account, or AWS organization",
									},
									"node_type": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Valid values: OU, ACCOUNT, ORG",
										ValidateFunc: validation.StringInSlice(
											[]string{
												"OU",
												"ACCOUNT",
												"ORG",
											},
											false,
										),
									},
									"selection_type": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Selection type. Valid values: INCLUDE to include the specified resource to onboard, EXCLUDE to exclude the specified resource and onboard the rest, ALL to onboard all resources in the organization.",
										ValidateFunc: validation.StringInSlice(
											[]string{
												"INCLUDE",
												"EXCLUDE",
												"ALL",
											},
											false,
										),
									},
								},
							},
						},
						"features": {
							Type:        schema.TypeSet,
							Optional:    true,
							Computed:    true,
							Description: "Features applicable for aws account",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Feature name",
									},
									"state": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Feature state, one of enabled and disabled",
										ValidateFunc: validation.StringInSlice(
											[]string{
												"enabled",
												"disabled",
											},
											false,
										),
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
						"external_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "",
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
							Optional:    true,
							Description: "List of account IDs to which you are assigning this account",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			org.TypeAzureOrg: {
				Type:        schema.TypeList,
				Optional:    true,
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
							Optional:    true,
							Description: "Whether or not the account is enabled",
							Default:     true,
						},
						"group_ids": {
							Type:        schema.TypeSet,
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
						"client_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Application ID registered with Active Directory",
						},
						"key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Application ID key",
							Sensitive:   true,
						},
						"monitor_flow_logs": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Automatically ingest flow logs",
						},
						"tenant_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Active Directory ID associated with Azure",
						},
						"service_principal_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique ID of the service principle object associated with the Prisma Cloud application that you create",
						},
						"account_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Account type - tenant or account",
							ValidateFunc: validation.StringInSlice(
								[]string{
									"tenant",
									"account",
								},
								false,
							),
						},
						"protection_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Protection mode",
						},
						"cloud_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Cloud type",
						},
						"hierarchy_selection": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "List of hierarchy selection. Each item has resource id, display name, node type and selection type",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"resource_id": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Management group ID or subscription ID.\nNote you must escape any double quotes in the resource ID with a backslash.",
									},
									"display_name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Display name for management group or subscription",
									},
									"node_type": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Valid values: SUBSCRIPTION, TENANT, MANAGEMENT_GROUP",
										ValidateFunc: validation.StringInSlice(
											[]string{
												"SUBSCRIPTION",
												"TENANT",
												"MANAGEMENT_GROUP",
											},
											false,
										),
									},
									"selection_type": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Selection type.Valid values: INCLUDE to include the specified resource to onboard, EXCLUDE to exclude the specified resource and onboard the rest, ALL to onboard all resources in the tenant.",
										ValidateFunc: validation.StringInSlice(
											[]string{
												"INCLUDE",
												"EXCLUDE",
												"ALL",
											},
											false,
										),
									},
								},
							},
						},
						"default_account_group_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Account group id to which you are assigning this account",
						},
						"features": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Features applicable for azure account",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Feature name",
									},
									"state": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Feature state, one of enabled and disabled",
										ValidateFunc: validation.StringInSlice(
											[]string{
												"enabled",
												"disabled",
											},
											false,
										),
									},
								},
							},
						},
						"root_sync_enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Azure tenant has children. Must be set to true when azure tenant is onboarded with children",
						},
						"environment_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Environment type",
							Default:     "azure",
							ValidateFunc: validation.StringInSlice(
								[]string{
									"azure",
									"azure_gov",
								},
								false,
							),
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
			org.TypeGcpOrg: {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Azure account type",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "GCP project ID",
						},
						"account_type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Account type - organization or account",
							ValidateFunc: validation.StringInSlice(
								[]string{
									"organization",
									"account",
								},
								true,
							),
						},
						"enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether or not the account is enabled",
							Default:     false,
						},
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Name to be used for the account on the Prisma Cloud platform (must be unique)",
						},
						"compression_enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Enable or disable compressed network flow log generation. Default value: `false`",
						},
						"credentials": {
							Type:             schema.TypeString,
							Required:         true,
							Description:      "Content of the JSON credentials file",
							Sensitive:        true,
							DiffSuppressFunc: gcpOrgv2CredentialsMatch,
						},
						"dataflow_enabled_project": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Project ID where the Dataflow API is enabled. Required if `compressionEnabled` is set to `true` and if the `accountType` is `organization`. Optional if the `accountType` is `account` or `masterServiceAccount`",
						},
						"features": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Features applicable for aws account",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Feature name",
									},
									"state": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Feature state, one of enabled and disabled",
										ValidateFunc: validation.StringInSlice(
											[]string{
												"enabled",
												"disabled",
											},
											false,
										),
									},
								},
							},
						},
						"flow_log_storage_bucket": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Cloud Storage Bucket name that is used store the flow logs.",
						},
						"protection_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Protection mode",
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
						"customer_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Customer name",
						},
						"created_epoch_millis": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Account created epoch time",
						},
						"last_modified_epoch_millis": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Last modified at epoch millis",
						},
						"last_modified_by": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Last modified by",
						},
						"deleted": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the account is deleted or not",
						},
						"storage_scan_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the storage scan is enabled",
						},
						"added_on_ts": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Added on time stamp",
						},
						"deployment_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Deployment type",
						},
						"deployment_type_description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Deployment type description",
						},
						"project_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "GCP project ID",
						},
						"service_account_email": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Service account email",
						},
						"authentication_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Authentication type",
						},
						"account_group_creation_mode": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "MANUAL",
							Description: "Cloud account group creation mode. Valid values - MANUAL, AUTO or RECURSIVE",
							ValidateFunc: validation.StringInSlice(
								[]string{
									"MANUAL",
									"AUTO",
									"RECURSIVE",
								},
								true,
							),
						},
						"hierarchy_selection": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "List of hierarchy selection. Each item has resource id, display name, node type and selection type",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"resource_id": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Resource ID. For folders, format is folders/{folder ID}. For projects, format is {project number}. For orgs, format is organizations/{org ID}",
									},
									"display_name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Display name for folder, project, or organization",
									},
									"node_type": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Node type. Valid values - FOLDER, PROJECT, ORG",
										ValidateFunc: validation.StringInSlice(
											[]string{
												"FOLDER",
												"PROJECT",
												"ORG",
											},
											false,
										),
									},
									"selection_type": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Selection type. Valid values: INCLUDE, EXCLUDE, ALL. If hierarchySelection.nodeType is PROJECT or FOLDER, then a valid value is either INCLUDE or EXCLUDE",
										ValidateFunc: validation.StringInSlice(
											[]string{
												"INCLUDE",
												"EXCLUDE",
												"ALL",
											},
											false,
										),
									},
								},
							},
						},
						"organization_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "GCP organization name",
						},
						"default_account_group_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Account group id to which you are assigning this account",
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
		},
	}
}
func gcpOrgv2CredentialsMatch(k, old, new string, d *schema.ResourceData) bool {
	var (
		err       error
		prev, cur accountv2.GcpCredentials
	)

	if err = json.Unmarshal([]byte(old), &prev); err != nil {
		return false
	}

	if err = json.Unmarshal([]byte(new), &cur); err != nil {
		return false
	}

	return (prev.Type == cur.Type &&
		prev.ProjectId == cur.ProjectId &&
		prev.PrivateKeyId == cur.PrivateKeyId &&
		//prev.PrivateKey == cur.PrivateKey &&       //Commenting this comparison of privateKey to avoid diff on every terraform plan
		prev.ClientEmail == cur.ClientEmail &&
		prev.ClientId == cur.ClientId &&
		prev.AuthUri == cur.AuthUri &&
		prev.TokenUri == cur.TokenUri &&
		prev.ProviderCertUrl == cur.ProviderCertUrl &&
		prev.ClientCertUrl == cur.ClientCertUrl)
}
func parseOrgV2CloudAccount(d *schema.ResourceData) (string, string, string, interface{}) {
	if x := ResourceDataInterfaceMap(d, org.TypeAwsOrg); len(x) != 0 {
		ans := org.AwsOrg{
			AccountId:             x["account_id"].(string),
			Enabled:               x["enabled"].(bool),
			DefaultAccountGroupId: x["default_account_group_id"].(string),
			Name:                  x["name"].(string),
			RoleArn:               x["role_arn"].(string),
			AccountType:           x["account_type"].(string),
			GroupIds:              SetToStringSlice(x["group_ids"].(*schema.Set)),
		}
		hsl := x["hierarchy_selection"].(*schema.Set).List()
		ans.HierarchySelection = make([]org.HierarchySelection, 0, len(hsl))
		for _, hsi := range hsl {
			hs := hsi.(map[string]interface{})
			ans.HierarchySelection = append(ans.HierarchySelection, org.HierarchySelection{
				ResourceId:    hs["resource_id"].(string),
				DisplayName:   hs["display_name"].(string),
				SelectionType: hs["selection_type"].(string),
				NodeType:      hs["node_type"].(string),
			})
		}
		features := x["features"].(*schema.Set).List()
		ans.Features = make([]org.Features, 0, len(features))
		for _, featuresi := range features {
			ftr := featuresi.(map[string]interface{})
			ans.Features = append(ans.Features, org.Features{
				Name:  ftr["name"].(string),
				State: ftr["state"].(string),
			})
		}
		return org.TypeAwsOrg, x["name"].(string), x["account_id"].(string), ans
	} else if x := ResourceDataInterfaceMap(d, org.TypeAzureOrg); len(x) != 0 {
		ans := org.AzureOrg{
			ClientId:              x["client_id"].(string),
			DefaultAccountGroupId: x["default_account_group_id"].(string),
			Key:                   x["key"].(string),
			MonitorFlowLogs:       x["monitor_flow_logs"].(bool),
			TenantId:              x["tenant_id"].(string),
			ServicePrincipalId:    x["service_principal_id"].(string),
			RootSyncEnabled:       x["root_sync_enabled"].(bool),
		}
		account := org.OrgAccountAzure{
			AccountId:   x["account_id"].(string),
			Enabled:     x["enabled"].(bool),
			AccountType: x["account_type"].(string),
			Name:        x["name"].(string),
			GroupIds:    SetToStringSlice(x["group_ids"].(*schema.Set)),
		}
		hsl := x["hierarchy_selection"].(*schema.Set).List()
		ans.HierarchySelection = make([]org.HierarchySelection, 0, len(hsl))
		for _, hsi := range hsl {
			hs := hsi.(map[string]interface{})
			ans.HierarchySelection = append(ans.HierarchySelection, org.HierarchySelection{
				ResourceId:    hs["resource_id"].(string),
				DisplayName:   hs["display_name"].(string),
				SelectionType: hs["selection_type"].(string),
				NodeType:      hs["node_type"].(string),
			})
		}
		ans.OrgAccountAzure = account
		features := x["features"].(*schema.Set).List()
		ans.Features = make([]org.Features, 0, len(features))
		for _, featuresi := range features {
			ftr := featuresi.(map[string]interface{})
			ans.Features = append(ans.Features, org.Features{
				Name:  ftr["name"].(string),
				State: ftr["state"].(string),
			})
		}
		return org.TypeAzureOrg, x["name"].(string), x["account_id"].(string), ans
	} else if x := ResourceDataInterfaceMap(d, org.TypeGcpOrg); len(x) != 0 {
		var creds org.GcpCredentials
		_ = json.Unmarshal([]byte(x["credentials"].(string)), &creds)
		ans := org.GcpOrg{
			CompressionEnabled:       x["compression_enabled"].(bool),
			DataflowEnabledProject:   x["dataflow_enabled_project"].(string),
			FlowLogStorageBucket:     x["flow_log_storage_bucket"].(string),
			Credentials:              creds,
			DefaultAccountGroupId:    x["default_account_group_id"].(string),
			OrganizationName:         x["organization_name"].(string),
			AccountGroupCreationMode: x["account_group_creation_mode"].(string),
		}
		account := org.OrgAccountGcp{
			AccountId:   x["account_id"].(string),
			AccountType: x["account_type"].(string),
			Enabled:     x["enabled"].(bool),
			Name:        x["name"].(string),
			ProjectId:   x["project_id"].(string),
		}
		hsl := x["hierarchy_selection"].(*schema.Set).List()
		ans.HierarchySelection = make([]org.HierarchySelection, 0, len(hsl))
		for _, hsi := range hsl {
			hs := hsi.(map[string]interface{})
			ans.HierarchySelection = append(ans.HierarchySelection, org.HierarchySelection{
				ResourceId:    hs["resource_id"].(string),
				DisplayName:   hs["display_name"].(string),
				SelectionType: hs["selection_type"].(string),
				NodeType:      hs["node_type"].(string),
			})
		}
		ans.OrgAccountGcp = account
		features := x["features"].(*schema.Set).List()
		ans.Features = make([]org.Features, 0, len(features))
		for _, featuresi := range features {
			ftr := featuresi.(map[string]interface{})
			ans.Features = append(ans.Features, org.Features{
				Name:  ftr["name"].(string),
				State: ftr["state"].(string),
			})
		}
		return org.TypeGcpOrg, x["name"].(string), x["account_id"].(string), ans
	}
	return "", "", "", nil
}

func saveOrgV2CloudAccount(d *schema.ResourceData, dest string, obj interface{}) {
	var val map[string]interface{}

	switch v := obj.(type) {
	case org.AwsOrgV2:
		val = map[string]interface{}{
			"account_id":                   v.CloudAccountResp.AccountId,
			"enabled":                      v.CloudAccountResp.Enabled,
			"default_account_group_id":     v.DefaultAccountGroupId,
			"name":                         v.CloudAccountResp.Name,
			"role_arn":                     v.RoleArn,
			"account_type":                 v.CloudAccountResp.AccountType,
			"group_ids":                    v.GroupIds,
			"cloud_type":                   v.CloudAccountResp.CloudType,
			"parent_id":                    v.CloudAccountResp.ParentId,
			"deleted":                      v.CloudAccountResp.Deleted,
			"protection_mode":              v.CloudAccountResp.ProtectionMode,
			"deployment_type":              v.CloudAccountResp.DeploymentType,
			"customer_name":                v.CloudAccountResp.CustomerName,
			"created_epoch_millis":         v.CloudAccountResp.CreatedEpochMillis,
			"last_modified_epoch_millis":   v.CloudAccountResp.LastModifiedEpochMillis,
			"last_modified_by":             v.CloudAccountResp.LastModifiedBy,
			"external_id":                  v.ExternalId,
			"has_member_role":              v.HasMemberRole,
			"template_url":                 v.TemplateUrl,
			"eventbridge_rule_name_prefix": v.EventbridgeRuleNamePrefix,
		}
		if len(v.HierarchySelection) == 0 {
			val["hierarchy_selection"] = nil
		} else {
			hsList := make([]interface{}, 0, len(v.HierarchySelection))
			for _, hs := range v.HierarchySelection {
				hsList = append(hsList, map[string]interface{}{
					"resource_id":    hs.ResourceId,
					"display_name":   hs.DisplayName,
					"node_type":      hs.NodeType,
					"selection_type": hs.SelectionType,
				})
			}
			val["hierarchy_selection"] = hsList
		}

		if len(v.CloudAccountResp.Features) == 0 {
			val["features"] = nil
		} else {
			ftrList := make([]interface{}, 0, len(v.CloudAccountResp.Features))
			for _, fti := range v.CloudAccountResp.Features {
				ftrList = append(ftrList, map[string]interface{}{
					"name":  fti.Name,
					"state": fti.State,
				})
			}
			val["features"] = ftrList
		}
	case org.AzureOrgV2:
		x := ResourceDataInterfaceMap(d, org.TypeAzureOrg)
		var key string
		if x["key"] == nil {
			key = v.Key
		} else {
			key = x["key"].(string)
		}

		val = map[string]interface{}{
			"account_id":                  v.CloudAccountAzureResp.AccountId,
			"enabled":                     v.CloudAccountAzureResp.Enabled,
			"group_ids":                   v.GroupIds,
			"name":                        v.CloudAccountAzureResp.Name,
			"account_type":                v.CloudAccountAzureResp.AccountType,
			"cloud_type":                  v.CloudAccountAzureResp.CloudType,
			"default_account_group_id":    v.DefaultAccountGroupId,
			"environment_type":            v.EnvironmentType,
			"client_id":                   v.ClientId,
			"key":                         key,
			"monitor_flow_logs":           v.MonitorFlowLogs,
			"tenant_id":                   v.TenantId,
			"service_principal_id":        v.ServicePrincipalId,
			"parent_id":                   v.CloudAccountAzureResp.ParentId,
			"deleted":                     v.CloudAccountAzureResp.Deleted,
			"customer_name":               v.CloudAccountAzureResp.CustomerName,
			"created_epoch_millis":        v.CloudAccountAzureResp.CreatedEpochMillis,
			"last_modified_epoch_millis":  v.CloudAccountAzureResp.LastModifiedEpochMillis,
			"last_modified_by":            v.CloudAccountAzureResp.LastModifiedBy,
			"template_url":                v.TemplateUrl,
			"protection_mode":             v.CloudAccountAzureResp.ProtectionMode,
			"deployment_type":             v.CloudAccountAzureResp.DeploymentType,
			"deployment_type_description": v.CloudAccountAzureResp.DeploymentTypeDescription,
			"member_sync_enabled":         v.MemberSyncEnabled,
			"root_sync_enabled":           v.MemberSyncEnabled,
		}
		if len(v.HierarchySelection) == 0 {
			val["hierarchy_selection"] = nil
		} else {
			hsList := make([]interface{}, 0, len(v.HierarchySelection))
			for _, hs := range v.HierarchySelection {
				hsList = append(hsList, map[string]interface{}{
					"resource_id":    hs.ResourceId,
					"display_name":   hs.DisplayName,
					"node_type":      hs.NodeType,
					"selection_type": hs.SelectionType,
				})
			}
			val["hierarchy_selection"] = hsList
		}

		if len(v.CloudAccountAzureResp.Features) == 0 {
			val["features"] = nil
		} else {
			ftrList := make([]interface{}, 0, len(v.CloudAccountAzureResp.Features))
			for _, fti := range v.CloudAccountAzureResp.Features {
				ftrList = append(ftrList, map[string]interface{}{
					"name":  fti.Name,
					"state": fti.State,
				})
			}
			val["features"] = ftrList
		}
	case org.GcpOrgV2:
		b, _ := json.Marshal(v.Credentials)
		val = map[string]interface{}{
			"account_id":                  v.CloudAccountGcpResp.AccountId,
			"enabled":                     v.CloudAccountGcpResp.Enabled,
			"name":                        v.CloudAccountGcpResp.Name,
			"account_type":                v.CloudAccountGcpResp.AccountType,
			"cloud_type":                  v.CloudAccountGcpResp.CloudType,
			"parent_id":                   v.CloudAccountGcpResp.ParentId,
			"deleted":                     v.CloudAccountGcpResp.Deleted,
			"customer_name":               v.CloudAccountGcpResp.CustomerName,
			"created_epoch_millis":        v.CloudAccountGcpResp.CreatedEpochMillis,
			"last_modified_epoch_millis":  v.CloudAccountGcpResp.LastModifiedEpochMillis,
			"last_modified_by":            v.CloudAccountGcpResp.LastModifiedBy,
			"protection_mode":             v.CloudAccountGcpResp.ProtectionMode,
			"credentials":                 string(b),
			"compression_enabled":         v.CompressionEnabled,
			"dataflow_enabled_project":    v.DataflowEnabledProject,
			"flow_log_storage_bucket":     v.FlowLogStorageBucket,
			"storage_scan_enabled":        v.CloudAccountGcpResp.StorageScanEnabled,
			"added_on_ts":                 v.CloudAccountGcpResp.AddedOnTs,
			"deployment_type":             v.CloudAccountGcpResp.DeploymentType,
			"deployment_type_description": v.CloudAccountGcpResp.DeploymentTypeDescription,
			"project_id":                  v.ProjectId,
			"service_account_email":       v.ServiceAccountEmail,
			"authentication_type":         v.AuthenticationType,
			"account_group_creation_mode": v.AccountGroupCreationMode,
			"default_account_group_id":    v.DefaultAccountGroupId,
			"group_ids":                   v.GroupIds,
			"organization_name":           v.OrganizationName,
		}
		if len(v.HierarchySelection) == 0 {
			val["hierarchy_selection"] = nil
		} else {
			hsList := make([]interface{}, 0, len(v.HierarchySelection))
			for _, hs := range v.HierarchySelection {
				hsList = append(hsList, map[string]interface{}{
					"resource_id":    hs.ResourceId,
					"display_name":   hs.DisplayName,
					"node_type":      hs.NodeType,
					"selection_type": hs.SelectionType,
				})
			}
			val["hierarchy_selection"] = hsList
		}
		if len(v.CloudAccountGcpResp.Features) == 0 {
			val["features"] = nil
		} else {
			ftrList := make([]interface{}, 0, len(v.CloudAccountGcpResp.Features))
			for _, fti := range v.CloudAccountGcpResp.Features {
				ftrList = append(ftrList, map[string]interface{}{
					"name":  fti.Name,
					"state": fti.State,
				})
			}
			val["features"] = ftrList
		}
	}
	for _, key := range []string{org.TypeAwsOrg, org.TypeAzureOrg, org.TypeGcpOrg} {
		if key != dest {
			d.Set(key, nil)
			continue
		}

		if err := d.Set(key, []interface{}{val}); err != nil {
			log.Printf("[WARN] Error setting %q field for %q: %s", key, d.Id(), err)
		}
	}
}
func createOrgV2CloudAccount(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	cloudType, name, _, obj := parseOrgV2CloudAccount(d)

	if err := org.Create(client, obj); err != nil {
		if strings.Contains(err.Error(), "duplicate_cloud_account") {
			if err := org.Update(client, obj); err != nil {
				return diag.FromErr(err)
			}
		} else {
			return diag.FromErr(err)
		}
	}
	PollApiUntilSuccess(func() error {
		_, err := org.Identify(client, cloudType, name)
		return err
	})

	accId, err := org.Identify(client, cloudType, name)
	if err != nil {
		return diag.FromErr(err)
	}

	var resp1 interface{}
	PollApiUntilSuccess(func() error {
		resp, err := org.Get(client, cloudType, accId)
		resp1 = resp
		return err
	})

	d.SetId(TwoStringsToId(cloudType, accId))
	saveOrgV2CloudAccount(d, cloudType, resp1)
	return readOrgV2CloudAccount(ctx, d, meta)
}

func readOrgV2CloudAccount(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	cloudType, id := IdToTwoStrings(d.Id())

	_, err := org.Get(client, cloudType, id)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	return nil
}

func updateOrgV2CloudAccount(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	cloudType, _, accId, obj := parseOrgV2CloudAccount(d)
	var resp1 interface{}

	if err := org.Update(client, obj); err != nil {
		return diag.FromErr(err)
	}

	PollApiUntilSuccess(func() error {
		resp, err := org.Get(client, cloudType, accId)
		resp1 = resp
		return err
	})

	d.SetId(TwoStringsToId(cloudType, accId))
	saveOrgV2CloudAccount(d, cloudType, resp1)
	return readOrgV2CloudAccount(ctx, d, meta)
}

func deleteOrgV2CloudAccount(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	cloudType, id := IdToTwoStrings(d.Id())
	disable := d.Get("disable_on_destroy").(bool)
	if disable {
		switch cloudType {
		case org.TypeAwsOrg:
			cloudAccount, _ := org.Get(client, cloudType, id)
			cloudAccountAws := cloudAccount.(org.AwsOrgV2)
			cloudAccountAws.CloudAccountResp.Enabled = false
			if err := org.DisableCloudAccount(client, cloudAccountAws.CloudAccountResp.AccountId); err != nil {
				return diag.FromErr(err)
			}
			return nil
		case org.TypeAzureOrg:
			cloudAccount, _ := org.Get(client, cloudType, id)
			orgAccountAzure := cloudAccount.(org.AzureOrgV2)
			orgAccountAzure.CloudAccountAzureResp.Enabled = false
			if err := org.DisableCloudAccount(client, orgAccountAzure.CloudAccountAzureResp.AccountId); err != nil {
				return diag.FromErr(err)
			}
			return nil
		case org.TypeGcpOrg:
			cloudAccount, _ := org.Get(client, cloudType, id)
			orgAccountGcp := cloudAccount.(org.GcpOrgV2)
			orgAccountGcp.CloudAccountGcpResp.Enabled = false
			if err := org.DisableCloudAccount(client, orgAccountGcp.CloudAccountGcpResp.AccountId); err != nil {
				return diag.FromErr(err)
			}
			return nil
		}
	}

	err := org.Delete(client, cloudType, id)
	if err != nil {
		if err != pc.ObjectNotFoundError {
			return diag.FromErr(err)
		}
	}

	d.SetId("")
	return nil
}
