package prismacloud

import (
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/compliance/standard/requirement/section"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceComplianceStandardRequirementSection() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceComplianceStandardRequirementSectionRead,

		Schema: map[string]*schema.Schema{
			// Input.
			"csr_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Compliance standard requirement ID",
			},
			"csrs_id": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				Description:  "Compliance standard requirement section ID",
				AtLeastOneOf: []string{"csrs_id", "section_id"},
			},
			"section_id": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				Description:  "Compliance section ID",
				AtLeastOneOf: []string{"csrs_id", "section_id"},
			},

			// Output.
			"section_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Number of system supported and custom compliance standard requirement sections",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description",
			},
			"created_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Created by",
			},
			"created_on": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Created on",
			},
			"last_modified_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Last modified by",
			},
			"last_modified_on": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Last modified on",
			},
			"system_default": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "System default",
			},
			"policies_assigned_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Number of assigned policies",
			},
			"standard_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Compliance standard name",
			},
			"requirement_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Compliance requirement name",
			},
			"label": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Section label",
			},
			"view_order": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "View order",
			},
			"associated_policy_ids": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of associated policy IDs",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceComplianceStandardRequirementSectionRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	var o section.Section
	var err error
	csrId := d.Get("csr_id").(string)
	csrsId := d.Get("csrs_id").(string)
	sectionId := d.Get("section_id").(string)

	if csrsId != "" {
		o, err = section.GetId(client, csrId, csrsId)
	} else {
		o, err = section.Get(client, csrId, sectionId)
	}

	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return err
	}

	d.SetId(TwoStringsToId(csrId, o.Id))
	saveComplianceStandardRequirementSection(d, csrId, o)

	return nil
}
