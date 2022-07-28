package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"golang.org/x/net/context"
	"log"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/compliance/standard/requirement"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceComplianceStandardRequirements() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceComplianceStandardRequirementsRead,

		Schema: map[string]*schema.Schema{
			// Input.
			"cs_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Compliance standard ID",
			},

			// Output.
			"total": totalSchema("system supported and custom appliance standard requirements"),
			"listing": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of all compliance requirements",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"csr_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Compliance standard requirement ID",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Compliance requirement name",
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
				},
			},
		},
	}
}

func dataSourceComplianceStandardRequirementsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	csId := d.Get("cs_id").(string)

	items, err := requirement.List(client, csId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(csId)
	d.Set("cs_id", csId)
	d.Set("total", len(items))

	list := make([]interface{}, 0, len(items))
	for _, o := range items {
		list = append(list, map[string]interface{}{
			"csr_id":                  o.Id,
			"name":                    o.Name,
			"description":             o.Description,
			"created_by":              o.CreatedBy,
			"created_on":              o.CreatedOn,
			"last_modified_by":        o.LastModifiedBy,
			"last_modified_on":        o.LastModifiedOn,
			"system_default":          o.SystemDefault,
			"policies_assigned_count": o.PoliciesAssignedCount,
			"standard_name":           o.StandardName,
			"requirement_id":          o.RequirementId,
			"view_order":              o.ViewOrder,
		})
	}

	if err := d.Set("listing", list); err != nil {
		log.Printf("[WARN] Error setting 'listing' field for %q: %s", d.Id(), err)
	}

	return nil
}
