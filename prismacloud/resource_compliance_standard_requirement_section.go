package prismacloud

import (
	"log"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/compliance/standard/requirement/section"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceComplianceStandardRequirementSection() *schema.Resource {
	return &schema.Resource{
		Create: createComplianceStandardRequirementSection,
		Read:   readComplianceStandardRequirementSection,
		Update: updateComplianceStandardRequirementSection,
		Delete: deleteComplianceStandardRequirementSection,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"csr_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Compliance standard requirement ID",
			},
			"csrs_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Compliance standard requirement section ID",
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
			"requirement_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Compliance requirement name",
			},
			"section_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Compliance section ID",
			},
			"label": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Section label",
			},
			"view_order": {
				Type:        schema.TypeInt,
				Optional:    true,
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

func parseComplianceStandardRequirementSection(d *schema.ResourceData, csrsId string) section.Section {
	return section.Section{
		RequirementId: d.Get("csr_id").(string),
		Id:            csrsId,
		Description:   d.Get("description").(string),
		SectionId:     d.Get("section_id").(string),
		Label:         d.Get("label").(string),
		ViewOrder:     d.Get("view_order").(int),
	}
}

func saveComplianceStandardRequirementSection(d *schema.ResourceData, csrId string, o section.Section) {
	d.Set("csr_id", csrId)
	d.Set("csrs_id", o.Id)
	d.Set("description", o.Description)
	d.Set("created_by", o.CreatedBy)
	d.Set("created_on", o.CreatedOn)
	d.Set("last_modified_by", o.LastModifiedBy)
	d.Set("last_modified_on", o.LastModifiedOn)
	d.Set("system_default", o.SystemDefault)
	d.Set("policies_assigned_count", o.PoliciesAssignedCount)
	d.Set("standard_name", o.StandardName)
	d.Set("requirement_name", o.RequirementName)
	d.Set("section_id", o.SectionId)
	d.Set("label", o.Label)
	d.Set("view_order", o.ViewOrder)
	if err := d.Set("associated_policy_ids", o.AssociatedPolicyIds); err != nil {
		log.Printf("[WARN] Error setting 'associated_policy_ids' for %s: %s", d.Id(), err)
	}
}

func createComplianceStandardRequirementSection(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	o := parseComplianceStandardRequirementSection(d, "")

	if err := section.Create(client, o); err != nil {
		return err
	}

	PollApiUntilSuccess(func() error {
		_, err := section.Get(client, o.RequirementId, o.SectionId)
		return err
	})

	liveObj, err := section.Get(client, o.RequirementId, o.SectionId)
	if err != nil {
		return err
	}

	PollApiUntilSuccess(func() error {
		_, err := section.GetId(client, o.RequirementId, liveObj.Id)
		return err
	})

	d.SetId(TwoStringsToId(o.RequirementId, liveObj.Id))
	return readComplianceStandardRequirementSection(d, meta)
}

func readComplianceStandardRequirementSection(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	csrId, csrsId := IdToTwoStrings(d.Id())

	o, err := section.GetId(client, csrId, csrsId)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return err
	}

	saveComplianceStandardRequirementSection(d, csrId, o)

	return nil
}

func updateComplianceStandardRequirementSection(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	csrsId := d.Get("csrs_id").(string)
	o := parseComplianceStandardRequirementSection(d, csrsId)

	if err := section.Update(client, o); err != nil {
		return err
	}

	return readComplianceStandardRequirementSection(d, meta)
}

func deleteComplianceStandardRequirementSection(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	_, csrsId := IdToTwoStrings(d.Id())

	err := section.Delete(client, csrsId)
	if err != nil {
		if err != pc.ObjectNotFoundError {
			return err
		}
	}

	d.SetId("")
	return nil
}
