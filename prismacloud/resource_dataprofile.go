package prismacloud

import (
	"log"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/data-security/dataprofile"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDataProfile() *schema.Resource {
	return &schema.Resource{
		Create: createDataProfile,
		Read:   readDataProfile,
		Update: updateDataProfile,
		Delete: deleteDataProfile,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"profile_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Profile ID",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Profile Name",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Profile description",
			},
			"types": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Type (basic or advance)",
				Default:     "basic",
			},
			"profile_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Profile Type (custom or system)",
				Default:     "custom",
			},
			"tenant_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Tenant ID",
			},
			"status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Status (hidden or non_hidden)",
				Default:     "non_hidden",
			},
			"profile_status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Profile status (active or disabled)",
				Default:     "active",
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
				Required:    true,
				MinItems:    1,
				MaxItems:    1,
				Description: "Model for DataProfile Rules",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"operator_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Pattern operator type",
							Default:     "or",
						},
						"data_pattern_rules": {
							Type:        schema.TypeSet,
							Required:    true,
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
										Required:    true,
										Description: "Pattern name",
									},
									"detection_technique": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Detection technique",
									},
									"match_type": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Match type",
										ValidateFunc: validation.StringInSlice(
											[]string{
												"include",
												"exclude",
											},
											false,
										),
									},
									"occurrence_operator_type": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Occurrence operator type",
										ValidateFunc: validation.StringInSlice(
											[]string{
												"any",
												"more_than_equal_to",
												"less_than_equal_to",
												"between",
											},
											false,
										),
									},
									"occurrence_count": {
										Type:         schema.TypeInt,
										Optional:     true,
										Description:  "Occurrence count",
										ValidateFunc: validation.IntBetween(1, 250),
									},
									"confidence_level": {
										Type:        schema.TypeString,
										Required:    true,
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
										Type:         schema.TypeInt,
										Optional:     true,
										Description:  "High Occurrence value",
										ValidateFunc: validation.IntBetween(1, 250),
									},
									"occurrence_low": {
										Type:         schema.TypeInt,
										Optional:     true,
										Description:  "Low Occurrence value",
										ValidateFunc: validation.IntBetween(1, 250),
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

func parseDataProfile(d *schema.ResourceData, id string) dataprofile.Profile {
	dpRule1 := ResourceDataInterfaceMap(d, "data_patterns_rule_1")
	var patternRules []dataprofile.DataPatternRule1

	if dpRule1["data_pattern_rules"] != nil && len(dpRule1["data_pattern_rules"].(*schema.Set).List()) > 0 {
		dpRules := dpRule1["data_pattern_rules"].(*schema.Set).List()
		patternRules = make([]dataprofile.DataPatternRule1, 0, len(dpRules))

		for i := range dpRules {
			dpRule := dpRules[i].(map[string]interface{})
			patternRules = append(patternRules, dataprofile.DataPatternRule1{
				Name:                   dpRule["name"].(string),
				MatchType:              dpRule["match_type"].(string),
				OccurrenceOperatorType: dpRule["occurrence_operator_type"].(string),
				OccurrenceCount:        dpRule["occurrence_count"].(int),
				ConfidenceLevel:        dpRule["confidence_level"].(string),
				OccurrenceHigh:         dpRule["occurrence_high"].(int),
				OccurrenceLow:          dpRule["occurrence_low"].(int),
			})
		}
	}

	return dataprofile.Profile{
		Id:            id,
		Name:          d.Get("name").(string),
		Description:   d.Get("description").(string),
		Types:         d.Get("types").(string),
		ProfileType:   d.Get("profile_type").(string),
		ProfileStatus: d.Get("profile_status").(string),
		Status:        d.Get("status").(string),
		DataPatternsRule1: dataprofile.DataPatternsRule1{
			OperatorType:     dpRule1["operator_type"].(string),
			DataPatternRules: patternRules,
		},
	}
}

func saveDataProfile(d *schema.ResourceData, o dataprofile.Profile) {
	var err error

	d.Set("profile_id", o.Id)
	d.Set("name", o.Name)
	d.Set("description", o.Description)
	d.Set("types", o.Types)
	d.Set("profile_type", o.ProfileType)
	d.Set("tenant_id", o.TenantId)
	d.Set("status", o.Status)
	d.Set("profile_status", o.ProfileStatus)
	d.Set("created_by", o.CreatedBy)
	d.Set("updated_by", o.UpdatedBy)
	d.Set("created_at", o.CreatedAt)
	d.Set("updated_at", o.UpdatedAt)

	dpRuleOne := map[string]interface{}{
		"operator_type":      o.DataPatternsRulesOne.OperatorType,
		"data_pattern_rules": nil,
	}

	if len(o.DataPatternsRulesOne.DataPatternRules) != 0 {
		dpRules := make([]interface{}, 0, len(o.DataPatternsRulesOne.DataPatternRules))
		for _, dpRule := range o.DataPatternsRulesOne.DataPatternRules {
			dpRules = append(dpRules, map[string]interface{}{
				"pattern_id":                  dpRule.Id,
				"name":                        dpRule.Name,
				"detection_technique":         dpRule.DetectionTechnique,
				"match_type":                  dpRule.MatchType,
				"occurrence_operator_type":    dpRule.OccurrenceOperatorType,
				"occurrence_count":            dpRule.OccurrenceCount,
				"confidence_level":            dpRule.ConfidencLevel,
				"supported_confidence_levels": dpRule.SupportedConfidenceLevels,
				"occurrence_high":             dpRule.OccurrenceHigh,
				"occurrence_low":              dpRule.OccurrenceLow,
			})
		}
		dpRuleOne["data_pattern_rules"] = dpRules
	}

	if err = d.Set("data_patterns_rule_1", []interface{}{dpRuleOne}); err != nil {
		log.Printf("[WARN] Error setting 'data_patterns_rule_1' for %s: %s", d.Id(), err)
	}
}

func createDataProfile(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	obj := parseDataProfile(d, "")

	if err := dataprofile.Create(client, obj); err != nil {
		return err
	}

	PollApiUntilSuccess(func() error {
		_, err := dataprofile.Identify(client, obj.Name)
		return err
	})

	id, err := dataprofile.Identify(client, obj.Name)
	if err != nil {
		return err
	}

	PollApiUntilSuccess(func() error {
		_, err := dataprofile.Get(client, id)
		return err
	})

	d.SetId(id)
	return readDataProfile(d, meta)
}

func readDataProfile(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	id := d.Id()

	obj, err := dataprofile.Get(client, id)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return err
	}

	saveDataProfile(d, obj)

	return nil
}

func updateDataProfile(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	id := d.Id()
	obj := parseDataProfile(d, id)

	if err := dataprofile.Update(client, obj); err != nil {
		return err
	}

	return readDataProfile(d, meta)
}

func deleteDataProfile(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	id := d.Id()

	err := dataprofile.Delete(client, id)
	if err != nil {
		if err != pc.ObjectNotFoundError {
			return err
		}
	}

	d.SetId("")
	return nil
}
