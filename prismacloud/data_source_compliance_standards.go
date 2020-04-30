package prismacloud

import (
	"log"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/compliance/standard"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceComplianceStandards() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceComplianceStandardsRead,

		Schema: map[string]*schema.Schema{
			// Output.
			"standard_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Number of system supported and custom compliance standards",
			},
			"standards": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of system supported and custom compliance standards",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cs_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Compliance standard ID",
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
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Compliance standard name",
						},
						"cloud_types": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Cloud type (determined based on policies assigned)",
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

func dataSourceComplianceStandardsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)

	items, err := standard.List(client)
	if err != nil {
		return err
	}

	d.SetId("compliance_standards")
	d.Set("standard_count", len(items))

	list := make([]interface{}, 0, len(items))
	for _, o := range items {
		list = append(list, map[string]interface{}{
			"cs_id":                   o.Id,
			"description":             o.Description,
			"created_by":              o.CreatedBy,
			"created_on":              o.CreatedOn,
			"last_modified_by":        o.LastModifiedBy,
			"last_modified_on":        o.LastModifiedOn,
			"system_default":          o.SystemDefault,
			"policies_assigned_count": o.PoliciesAssignedCount,
			"name":                    o.Name,
			"cloud_types":             o.CloudTypes,
		})
	}

	if err := d.Set("standards", list); err != nil {
		log.Printf("[WARN] Error setting 'standards' field for %q: %s", d.Id(), err)
	}

	return nil
}
