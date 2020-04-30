package prismacloud

import (
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/compliance/standard/requirement"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceComplianceStandardRequirement() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceComplianceStandardRequirementRead,

		Schema: map[string]*schema.Schema{
			// Input.
			"cs_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Compliance standard ID",
			},
			"csr_id": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				Description:  "Compliance standard requirement ID",
				AtLeastOneOf: []string{"csr_id", "name"},
			},
			"name": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				Description:  "Compliance requirement name",
				AtLeastOneOf: []string{"csr_id", "name"},
			},

			// Output.
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
			"requirement_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Compliance requirement number",
			},
			"view_order": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "View order",
			},
		},
	}
}

func dataSourceComplianceStandardRequirementRead(d *schema.ResourceData, meta interface{}) error {
	var err error
	client := meta.(*pc.Client)
	csId := d.Get("cs_id").(string)
	csrId := d.Get("csr_id").(string)

	if csrId == "" {
		name := d.Get("name").(string)
		csrId, err = requirement.Identify(client, csId, name)
		if err != nil {
			if err == pc.ObjectNotFoundError {
				d.SetId("")
				return nil
			}
			return err
		}
	}

	o, err := requirement.Get(client, csrId)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return err
	}

	d.SetId(TwoStringsToId(csId, csrId))
	saveComplianceStandardRequirement(d, csId, o)

	return nil
}
