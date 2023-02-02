package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
		},
	}
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
	}
	for _, key := range []string{accountv2.TypeAws} {
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
		if cloudType == accountv2.TypeAws {
			cloudAccount, _ := accountv2.Get(client, cloudType, id)
			cloudAccountAws := cloudAccount.(accountv2.AwsV2)
			cloudAccountAws.CloudAccountResp.Enabled = false
			if err := accountv2.DisableCloudAccount(client, cloudAccountAws.CloudAccountResp.AccountId); err != nil {
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
