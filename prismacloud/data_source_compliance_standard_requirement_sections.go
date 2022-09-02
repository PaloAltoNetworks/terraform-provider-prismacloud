package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"golang.org/x/net/context"
	"log"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/compliance/standard/requirement/section"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceComplianceStandardRequirementSections() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceComplianceStandardRequirementSectionsRead,

		Schema: map[string]*schema.Schema{
			// Input.
			"csr_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Compliance standard requirement ID",
			},

			// Output.
			"total": totalSchema("system supported and custom compliance standard requirement sections"),
			"listing": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of all compliance requirement sections",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"csrs_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Compliance standard requirement section ID",
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
						"section_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Compliance section ID",
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
				},
			},
		},
	}
}

func dataSourceComplianceStandardRequirementSectionsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	csrId := d.Get("csr_id").(string)

	items, err := section.List(client, csrId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(csrId)
	d.Set("csr_id", csrId)
	d.Set("total", len(items))

	list := make([]interface{}, 0, len(items))
	for _, o := range items {
		list = append(list, map[string]interface{}{
			"csrs_id":                 o.Id,
			"description":             o.Description,
			"created_by":              o.CreatedBy,
			"created_on":              o.CreatedOn,
			"last_modified_by":        o.LastModifiedBy,
			"last_modified_on":        o.LastModifiedOn,
			"system_default":          o.SystemDefault,
			"policies_assigned_count": o.PoliciesAssignedCount,
			"standard_name":           o.StandardName,
			"requirement_name":        o.RequirementName,
			"section_id":              o.SectionId,
			"label":                   o.Label,
			"view_order":              o.ViewOrder,
			"associated_policy_ids":   o.AssociatedPolicyIds,
		})
	}

	if err := d.Set("listing", list); err != nil {
		log.Printf("[WARN] Error setting 'listing' field for %q: %s", d.Id(), err)
	}

	return nil
}
