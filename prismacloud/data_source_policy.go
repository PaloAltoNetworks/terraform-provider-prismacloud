package prismacloud

import (
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/policy"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourcePolicy() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePolicyRead,

		Schema: map[string]*schema.Schema{
			// Input.
			"policy_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "Policy ID",
				AtLeastOneOf: []string{"name", "policy_id"},
			},
			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "Policy name",
				AtLeastOneOf: []string{"name", "policy_id"},
			},

			// Output.
			"policy_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Policy type",
			},
			"system_default": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "If policy is a system default policy or not",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description",
			},
			"severity": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Severity",
			},
			"recommendation": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Remediation recommendation",
			},
			"cloud_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cloud type (Required for config policies)",
			},
			"labels": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Labels",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enabled",
			},
			"created_on": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Created on",
			},
			"created_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Created by",
			},
			"last_modified_on": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Last modified on",
			},
			"last_modified_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Last modified by",
			},
			"rule_last_modified_on": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Rule last modified on",
			},
			"overridden": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Overridden",
			},
			"deleted": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Deleted",
			},
			"restrict_alert_dismissal": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Restrict alert dismissal",
			},
			"open_alerts_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Open alerts count",
			},
			"owner": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Owner",
			},
			"policy_mode": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Policy mode",
			},
			"remediable": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Is remediable or not",
			},
			"rule": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Model for rule",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name",
						},
						"cloud_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Cloud type",
						},
						"cloud_account": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Cloud account",
						},
						"resource_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Resource type",
						},
						"api_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "API name",
						},
						"resource_id_path": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Resource ID path",
						},
						"criteria": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Saved search ID that defines the rule criteria",
						},
						"parameters": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Parameters",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"rule_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of rule or RQL query",
						},
					},
				},
			},
			"remediation": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Model for remediation",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"template_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CLI script template",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description",
						},
						"cli_script_template": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CLI script template",
						},
						"cli_script_json_schema_string": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CLI script JSON schema",
						},
					},
				},
			},
			"compliance_metadata": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of compliance data. Each item has compliance standard, requirement, and/or section information",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"standard_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Compliance standard name",
						},
						"standard_description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Compliance standard description",
						},
						"requirement_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Requirement ID",
						},
						"requirement_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Requirement name",
						},
						"requirement_description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Requirement description",
						},
						"section_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Section ID",
						},
						"section_description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Section description",
						},
						"policy_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Policy ID",
						},
						"compliance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Compliance ID",
						},
						"section_label": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Section label",
						},
						"custom_assigned": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Custom assigned",
						},
					},
				},
			},
		},
	}
}

func dataSourcePolicyRead(d *schema.ResourceData, meta interface{}) error {
	var err error
	client := meta.(*pc.Client)

	id := d.Get("policy_id").(string)

	if id == "" {
		name := d.Get("name").(string)
		id, err = policy.Identify(client, name)
		if err != nil {
			if err == pc.ObjectNotFoundError {
				d.SetId("")
				return nil
			}
			return err
		}
	}

	obj, err := policy.Get(client, id)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return err
	}

	d.SetId(id)
	savePolicy(d, obj)

	return nil
}
