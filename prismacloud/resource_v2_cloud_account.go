package prismacloud

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/paloaltonetworks/prisma-cloud-go/cloud/account"
	"github.com/paloaltonetworks/prisma-cloud-go/cloud/account-v2"
	"golang.org/x/net/context"
	"log"
	"strings"
	"time"

	pc "github.com/paloaltonetworks/prisma-cloud-go"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceV2CloudAccount() *schema.Resource {
	return &schema.Resource{
		CreateContext: createV2CloudAccount,
		ReadContext:   readV2CloudAccount,
		UpdateContext: updateV2CloudAccount,
		DeleteContext: deleteV2CloudAccount,

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

			// AWS type.
			accountv2.TypeAws: {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "AWS account type",
				MaxItems:    1,
				ConflictsWith: []string{
					accountv2.TypeAzure,
					accountv2.TypeGcp,
					accountv2.TypeIbm,
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
						"group_ids": {
							Type:        schema.TypeSet,
							Required:    true,
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
							Default:     "account",
							Description: "Account type - organization or account",
							ValidateFunc: validation.StringInSlice(
								[]string{
									"account",
									"organization",
								},
								false,
							),
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
						"account_type_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Account type id",
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
					},
				},
			},
			//Azure type
			accountv2.TypeAzure: {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Azure account type",
				MaxItems:    1,
				ConflictsWith: []string{
					accountv2.TypeAws,
					accountv2.TypeGcp,
					accountv2.TypeIbm,
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
							Optional:    true,
							Description: "Whether or not the account is enabled",
							Default:     true,
						},
						"group_ids": {
							Type:        schema.TypeSet,
							Required:    true,
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
							Default:     "account",
							Description: "Account type - tenant or account",
							ValidateFunc: validation.StringInSlice(
								[]string{
									"account",
									"tenant",
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
						"environment_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Environment type",
							Default:     "azure",
							ValidateFunc: validation.StringInSlice(
								[]string{
									"azure",
									"azure_gov",
									"azure_china",
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
					},
				},
			},
			//Gcp type
			accountv2.TypeGcp: {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "GCP account type",
				MaxItems:    1,
				ConflictsWith: []string{
					accountv2.TypeAws,
					accountv2.TypeAzure,
					accountv2.TypeIbm,
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "GCP project ID",
						},
						"account_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "account",
							Description: "Account type - account, masterServiceAccount or organization",
							ValidateFunc: validation.StringInSlice(
								[]string{
									"account",
									"masterServiceAccount",
									"organization",
								},
								false,
							),
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
						"compression_enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable or disable compressed network flow log generation. Default value: `false`",
						},
						"credentials": {
							Type:             schema.TypeString,
							Required:         true,
							Description:      "Content of the JSON credentials file",
							Sensitive:        true,
							DiffSuppressFunc: gcpv2CredentialsMatch,
						},
						"dataflow_enabled_project": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Project ID where the Dataflow API is enabled. Required if `compressionEnabled` is set to `true` and if the `accountType` is `organization`. Optional if the `accountType` is `account` or `masterServiceAccount`",
						},
						"features": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Features applicable for gcp account",
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
							Description: "Cloud Storage Bucket name that is used store the flow logs",
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
							Description: "Parent Id",
						},
						"customer_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Prisma customer name",
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
							Description: "Storage scan enabled",
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
							Description: "GCP project ID",
						},
						"service_account_email": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Service account email",
						},
						"authentication_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Authentication type",
							Default:     "service_account",
							ValidateFunc: validation.StringInSlice(
								[]string{
									"service_account",
								},
								false,
							),
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
								false,
							),
						},
					},
				},
			},
			//IBM type
			accountv2.TypeIbm: {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "IBM account type",
				MaxItems:    1,
				ConflictsWith: []string{
					accountv2.TypeAws,
					accountv2.TypeAzure,
					accountv2.TypeGcp,
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "IBM account ID",
						},
						"account_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "account",
							Description: "Account type -  account",
						},
						"api_key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "IBM service API key",
							Sensitive:   true,
						},
						"enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether or not the account is enabled",
							Default:     true,
						},
						"group_ids": {
							Type:        schema.TypeSet,
							Required:    true,
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
						"svc_id_iam_id": {
							Type:        schema.TypeString,
							Required:    true,
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
							Description: "Parent Id",
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

func gcpv2CredentialsMatch(k, old, new string, d *schema.ResourceData) bool {
	var (
		err       error
		prev, cur account.GcpCredentials
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

func parseV2CloudAccount(d *schema.ResourceData) (string, string, string, interface{}) {
	if x := ResourceDataInterfaceMap(d, accountv2.TypeAws); len(x) != 0 {
		ans := accountv2.Aws{
			AccountId:   x["account_id"].(string),
			Enabled:     x["enabled"].(bool),
			GroupIds:    SetToStringSlice(x["group_ids"].(*schema.Set)),
			Name:        x["name"].(string),
			RoleArn:     x["role_arn"].(string),
			AccountType: x["account_type"].(string),
		}
		features := x["features"].(*schema.Set).List()
		ans.Features = make([]accountv2.Features, 0, len(features))
		for _, featuresi := range features {
			ftr := featuresi.(map[string]interface{})
			ans.Features = append(ans.Features, accountv2.Features{
				Name:  ftr["name"].(string),
				State: ftr["state"].(string),
			})
		}
		return accountv2.TypeAws, x["name"].(string), x["account_id"].(string), ans
	} else if x := ResourceDataInterfaceMap(d, accountv2.TypeAzure); len(x) != 0 {
		ans := accountv2.Azure{
			EnvironmentType:    x["environment_type"].(string),
			ClientId:           x["client_id"].(string),
			Key:                x["key"].(string),
			MonitorFlowLogs:    x["monitor_flow_logs"].(bool),
			TenantId:           x["tenant_id"].(string),
			ServicePrincipalId: x["service_principal_id"].(string),
		}
		account := accountv2.CloudAccountAzure{
			AccountId:   x["account_id"].(string),
			AccountType: x["account_type"].(string),
			Enabled:     x["enabled"].(bool),
			Name:        x["name"].(string),
			GroupIds:    SetToStringSlice(x["group_ids"].(*schema.Set)),
		}
		ans.CloudAccountAzure = account
		features := x["features"].(*schema.Set).List()
		ans.Features = make([]accountv2.Features, 0, len(features))
		for _, featuresi := range features {
			ftr := featuresi.(map[string]interface{})
			ans.Features = append(ans.Features, accountv2.Features{
				Name:  ftr["name"].(string),
				State: ftr["state"].(string),
			})
		}
		return accountv2.TypeAzure, x["name"].(string), x["account_id"].(string), ans
	} else if x := ResourceDataInterfaceMap(d, accountv2.TypeGcp); len(x) != 0 {
		var creds accountv2.GcpCredentials
		_ = json.Unmarshal([]byte(x["credentials"].(string)), &creds)
		ans := accountv2.Gcp{
			CompressionEnabled:     x["compression_enabled"].(bool),
			DataflowEnabledProject: x["dataflow_enabled_project"].(string),
			FlowLogStorageBucket:   x["flow_log_storage_bucket"].(string),
			DefaultAccountGroupId:  x["default_account_group_id"].(string),
			Credentials:            creds,
		}
		account := accountv2.CloudAccountGcp{
			AccountId:   x["account_id"].(string),
			AccountType: x["account_type"].(string),
			Enabled:     x["enabled"].(bool),
			Name:        x["name"].(string),
			GroupIds:    SetToStringSlice(x["group_ids"].(*schema.Set)),
		}
		ans.CloudAccountGcp = account
		features := x["features"].(*schema.Set).List()
		ans.Features = make([]accountv2.Features, 0, len(features))
		for _, featuresi := range features {
			ftr := featuresi.(map[string]interface{})
			ans.Features = append(ans.Features, accountv2.Features{
				Name:  ftr["name"].(string),
				State: ftr["state"].(string),
			})
		}
		return accountv2.TypeGcp, x["name"].(string), x["account_id"].(string), ans
	} else if x := ResourceDataInterfaceMap(d, accountv2.TypeIbm); len(x) != 0 {
		ans := accountv2.Ibm{
			AccountId:   x["account_id"].(string),
			AccountType: x["account_type"].(string),
			ApiKey:      x["api_key"].(string),
			Enabled:     x["enabled"].(bool),
			GroupIds:    SetToStringSlice(x["group_ids"].(*schema.Set)),
			Name:        x["name"].(string),
			SvcIdIamId:  x["svc_id_iam_id"].(string),
		}
		return accountv2.TypeIbm, x["name"].(string), x["account_id"].(string), ans
	}
	return "", "", "", nil
}

func saveV2CloudAccount(d *schema.ResourceData, dest string, obj interface{}) {
	var val map[string]interface{}

	switch v := obj.(type) {
	case accountv2.AwsV2:
		val = map[string]interface{}{
			"account_id":                   v.CloudAccountResp.AccountId,
			"enabled":                      v.CloudAccountResp.Enabled,
			"group_ids":                    v.GroupIds,
			"name":                         v.CloudAccountResp.Name,
			"role_arn":                     v.RoleArn,
			"account_type":                 v.CloudAccountResp.AccountType,
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
	case accountv2.AzureV2:
		x := ResourceDataInterfaceMap(d, accountv2.TypeAzure)
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
	case accountv2.GcpV2:
		b, _ := json.Marshal(v.Credentials)
		val = map[string]interface{}{
			"account_id":                  v.CloudAccountGcpResp.AccountId,
			"enabled":                     v.CloudAccountGcpResp.Enabled,
			"group_ids":                   v.GroupIds,
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
	case accountv2.IbmV2:
		x := ResourceDataInterfaceMap(d, accountv2.TypeIbm)
		var apiKey string
		if x["api_key"] == nil {
			apiKey = v.ApiKey
		} else {
			apiKey = x["api_key"].(string)
		}
		val = map[string]interface{}{
			"account_id":                  v.CloudAccountIbmResp.AccountId,
			"account_type":                v.CloudAccountIbmResp.AccountType,
			"added_on_ts":                 v.CloudAccountIbmResp.AddedOnTs,
			"cloud_type":                  v.CloudAccountIbmResp.CloudType,
			"created_epoch_millis":        v.CloudAccountIbmResp.CreatedEpochMillis,
			"customer_name":               v.CloudAccountIbmResp.CustomerName,
			"deleted":                     v.CloudAccountIbmResp.Deleted,
			"deployment_type":             v.CloudAccountIbmResp.DeploymentType,
			"deployment_type_description": v.CloudAccountIbmResp.DeploymentTypeDescription,
			"enabled":                     v.CloudAccountIbmResp.Enabled,
			"last_modified_epoch_millis":  v.CloudAccountIbmResp.LastModifiedEpochMillis,
			"last_modified_by":            v.CloudAccountIbmResp.LastModifiedBy,
			"name":                        v.CloudAccountIbmResp.Name,
			"parent_id":                   v.CloudAccountIbmResp.ParentId,
			"protection_mode":             v.CloudAccountIbmResp.ProtectionMode,
			"storage_scan_enabled":        v.CloudAccountIbmResp.StorageScanEnabled,
			"group_ids":                   v.GroupIds,
			"svc_id_iam_id":               v.SvcIdIamId,
			"api_key":                     apiKey,
		}
		if len(v.CloudAccountIbmResp.Features) == 0 {
			val["features"] = nil
		} else {
			ftrList := make([]interface{}, 0, len(v.CloudAccountIbmResp.Features))
			for _, fti := range v.CloudAccountIbmResp.Features {
				ftrList = append(ftrList, map[string]interface{}{
					"name":  fti.Name,
					"state": fti.State,
				})
			}
			val["features"] = ftrList
		}
	}
	for _, key := range []string{accountv2.TypeAws, accountv2.TypeAzure, accountv2.TypeGcp, accountv2.TypeIbm} {
		if key != dest {
			d.Set(key, nil)
			continue
		}

		if err := d.Set(key, []interface{}{val}); err != nil {
			log.Printf("[WARN] Error setting %q field for %q: %s", key, d.Id(), err)
		}
	}
}

func createV2CloudAccount(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	cloudType, _, accId, obj := parseV2CloudAccount(d)

	if err := accountv2.Create(client, obj); err != nil {
		if strings.Contains(err.Error(), "duplicate_cloud_account") {
			if err := accountv2.Update(client, obj); err != nil {
				return diag.FromErr(err)
			}
		} else {
			return diag.FromErr(err)
		}
	}

	PollApiUntilSuccess(func() error {
		_, err := accountv2.Get(client, cloudType, accId)
		return err
	})

	d.SetId(TwoStringsToId(cloudType, accId))
	return readV2CloudAccount(ctx, d, meta)
}

func readV2CloudAccount(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	cloudType, id := IdToTwoStrings(d.Id())

	obj, err := accountv2.Get(client, cloudType, id)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	saveV2CloudAccount(d, cloudType, obj)

	return nil
}

func updateV2CloudAccount(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)

	_, _, _, obj := parseV2CloudAccount(d)

	if err := accountv2.Update(client, obj); err != nil {
		return diag.FromErr(err)
	}

	return readV2CloudAccount(ctx, d, meta)
}

func deleteV2CloudAccount(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	cloudType, id := IdToTwoStrings(d.Id())
	disable := d.Get("disable_on_destroy").(bool)

	if disable {
		switch cloudType {
		case accountv2.TypeAws:
			cloudAccount, _ := accountv2.Get(client, cloudType, id)
			cloudAccountAws := cloudAccount.(accountv2.AwsV2)
			cloudAccountAws.CloudAccountResp.Enabled = false
			if err := accountv2.DisableCloudAccount(client, cloudAccountAws.CloudAccountResp.AccountId); err != nil {
				return diag.FromErr(err)
			}
			return nil
		case accountv2.TypeAzure:
			cloudAccount, _ := accountv2.Get(client, cloudType, id)
			cloudAccountAzure := cloudAccount.(accountv2.AzureV2)
			cloudAccountAzure.CloudAccountAzureResp.Enabled = false
			if err := accountv2.DisableCloudAccount(client, cloudAccountAzure.CloudAccountAzureResp.AccountId); err != nil {
				return diag.FromErr(err)
			}
			return nil
		case accountv2.TypeGcp:
			cloudAccount, _ := accountv2.Get(client, cloudType, id)
			cloudAccountGcp := cloudAccount.(accountv2.GcpV2)
			cloudAccountGcp.CloudAccountGcpResp.Enabled = false
			if err := accountv2.DisableCloudAccount(client, cloudAccountGcp.CloudAccountGcpResp.AccountId); err != nil {
				return diag.FromErr(err)
			}
			return nil
		case accountv2.TypeIbm:
			cloudAccount, _ := accountv2.Get(client, cloudType, id)
			cloudAccountIbm := cloudAccount.(accountv2.IbmV2)
			cloudAccountIbm.CloudAccountIbmResp.Enabled = false
			if err := accountv2.DisableCloudAccount(client, cloudAccountIbm.CloudAccountIbmResp.AccountId); err != nil {
				return diag.FromErr(err)
			}
			return nil
		}
	}

	err := accountv2.Delete(client, cloudType, id)
	if err != nil {
		if err != pc.ObjectNotFoundError {
			return diag.FromErr(err)
		}
	}

	d.SetId("")
	return nil
}
