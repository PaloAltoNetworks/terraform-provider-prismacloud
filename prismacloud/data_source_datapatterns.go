package prismacloud

import (
	"log"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/data-security/datapattern"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceDataPatterns() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDataPatternsRead,

		Schema: map[string]*schema.Schema{
			// Output.
			"total": totalSchema("data patterns"),
			"listing": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of data patterns",
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
						"mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Mode - predefined or custom",
						},
						"detection_technique": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Detection technique",
						},
						"updated_at": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Last updated at",
						},
						"updated_by": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Updated by",
						},
						"created_by": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Created by",
						},
						"is_editable": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Is editable",
						},
					},
				},
			},
		},
	}
}

func dataSourceDataPatternsRead(d *schema.ResourceData, meta interface{}) error {
	var err error
	client := meta.(*pc.Client)

	items, err := datapattern.List(client)
	if err != nil {
		return err
	}

	ans := make([]interface{}, 0, len(items))
	for _, v := range items {
		ans = append(ans, map[string]interface{}{
			"pattern_id":          v.Id,
			"name":                v.Name,
			"mode":                v.Mode,
			"detection_technique": v.DetectionTechnique,
			"updated_at":          v.UpdatedAt,
			"updated_by":          v.UpdatedBy,
			"created_by":          v.CreatedBy,
			"is_editable":         v.IsEditable,
		})
	}

	d.SetId("data patterns")
	d.Set("total", len(items))
	if err = d.Set("listing", ans); err != nil {
		log.Printf("[WARN] Error setting 'listing' field for %q: %s", d.Id(), err)
	}

	return nil
}
