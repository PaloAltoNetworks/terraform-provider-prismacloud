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
		}
	}

	d.SetId(TwoStringsToId(cloudType, id))
	d.Set("cloud_type", cloudType)
	d.Set("name", name)
	d.Set("account_id", id)

	saveV2CloudAccount(d, cloudType, obj)

	return nil
}
