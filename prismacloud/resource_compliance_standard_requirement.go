package prismacloud

import (
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/compliance/standard/requirement"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceComplianceStandardRequirement() *schema.Resource {
	return &schema.Resource{
		Create: createComplianceStandardRequirement,
		Read:   readComplianceStandardRequirement,
		Update: updateComplianceStandardRequirement,
		Delete: deleteComplianceStandardRequirement,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"cs_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Compliance standard ID",
			},
			"csr_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Compliance standard requirement ID",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Compliance requirement name",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
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
				Optional:    true,
				Description: "Compliance requirement number",
			},
			"view_order": {
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
				Description: "View order",
			},
		},
	}
}

func parseComplianceStandardRequirement(d *schema.ResourceData, csrId string) requirement.Requirement {
	return requirement.Requirement{
		ComplianceId:  d.Get("cs_id").(string),
		Id:            csrId,
		Name:          d.Get("name").(string),
		Description:   d.Get("description").(string),
		RequirementId: d.Get("requirement_id").(string),
		ViewOrder:     d.Get("view_order").(int),
	}
}

func saveComplianceStandardRequirement(d *schema.ResourceData, csId string, o requirement.Requirement) {
	d.Set("cs_id", csId)
	d.Set("csr_id", o.Id)
	d.Set("name", o.Name)
	d.Set("description", o.Description)
	d.Set("created_by", o.CreatedBy)
	d.Set("created_on", o.CreatedOn)
	d.Set("last_modified_by", o.LastModifiedBy)
	d.Set("last_modified_on", o.LastModifiedOn)
	d.Set("system_default", o.SystemDefault)
	d.Set("policies_assigned_count", o.PoliciesAssignedCount)
	d.Set("standard_name", o.StandardName)
	d.Set("requirement_id", o.RequirementId)
	d.Set("view_order", o.ViewOrder)
}

func createComplianceStandardRequirement(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	o := parseComplianceStandardRequirement(d, "")

	if err := requirement.Create(client, o); err != nil {
		return err
	}

	PollApiUntilSuccess(func() error {
		_, err := requirement.Identify(client, o.ComplianceId, o.Name)
		return err
	})

	csrId, err := requirement.Identify(client, o.ComplianceId, o.Name)
	if err != nil {
		return err
	}

	PollApiUntilSuccess(func() error {
		_, err := requirement.Get(client, csrId)
		return err
	})

	d.SetId(TwoStringsToId(o.ComplianceId, csrId))
	return readComplianceStandardRequirement(d, meta)
}

func readComplianceStandardRequirement(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	csId, csrId := IdToTwoStrings(d.Id())

	o, err := requirement.Get(client, csrId)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return err
	}

	saveComplianceStandardRequirement(d, csId, o)

	return nil
}

func updateComplianceStandardRequirement(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	csrId := d.Get("csr_id").(string)
	o := parseComplianceStandardRequirement(d, csrId)

	if err := requirement.Update(client, o); err != nil {
		return err
	}

	return readComplianceStandardRequirement(d, meta)
}

func deleteComplianceStandardRequirement(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	_, csrId := IdToTwoStrings(d.Id())

	err := requirement.Delete(client, csrId)
	if err != nil {
		if err != pc.ObjectNotFoundError {
			return err
		}
	}

	d.SetId("")
	return nil
}
