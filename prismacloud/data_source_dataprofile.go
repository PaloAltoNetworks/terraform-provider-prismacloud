package prismacloud

import (
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/data-security/dataprofile"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceDataProfile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDataProfileRead,

		Schema: map[string]*schema.Schema{
			// Input.
			"profile_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "Profile ID",
				AtLeastOneOf: []string{"name", "profile_id"},
			},
			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "Profile Name",
				AtLeastOneOf: []string{"name", "profile_id"},
			},

			// Output.
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Profile description",
			},
			"types": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type (basic or advance)",
			},
			"profile_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Profile type (custom or system)",
			},
			"tenant_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Tenant ID",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status (hidden or non_hidden)",
			},
			"profile_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Profile status (active or disabled)",
			},
			"created_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Created by",
			},
			"updated_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Updated by",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Created at (unix time)",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Updated at (unix time)",
			},
			"data_patterns_rule_1": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Model for DataProfile Rules",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"operator_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Pattern operator type",
						},
						"data_pattern_rules": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of DataPattern Rules",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pattern_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Pattern ID",
									},
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Pattern name",
									},
									"detection_technique": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Detection technique",
									},
									"match_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Match type",
									},
									"occurrence_operator_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Occurrence operator type",
									},
									"occurrence_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Occurrence count",
									},
									"confidence_level": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Confidence level",
									},
									"supported_confidence_levels": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "List of supported confidence levels",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"occurrence_high": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "High occurrence value",
									},
									"occurrence_low": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Low occurrence value",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceDataProfileRead(d *schema.ResourceData, meta interface{}) error {
	var err error
	client := meta.(*pc.Client)

	id := d.Get("profile_id").(string)

	if id == "" {
		name := d.Get("name").(string)
		id, err = dataprofile.Identify(client, name)
		if err != nil {
			if err == pc.ObjectNotFoundError {
				d.SetId("")
				return nil
			}
			return err
		}
	}

	obj, err := dataprofile.Get(client, id)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return err
	}

	d.SetId(id)
	saveDataProfile(d, obj)

	return nil
}
