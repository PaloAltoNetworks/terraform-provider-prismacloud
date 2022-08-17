package prismacloud

import (
	"log"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/data-security/dataprofile"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDataProfiles() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDataProfilesRead,

		Schema: map[string]*schema.Schema{
			// Output.
			"total": totalSchema("data profiles"),
			"listing": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of data profiles",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"profile_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Profile ID",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Profile Name",
						},
						"profile_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Profile status (active or disabled)",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Updated at (unix time)",
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
					},
				},
			},
		},
	}
}

func dataSourceDataProfilesRead(d *schema.ResourceData, meta interface{}) error {
	var err error
	client := meta.(*pc.Client)

	items, err := dataprofile.List(client)
	if err != nil {
		return err
	}

	ans := make([]interface{}, 0, len(items))
	for _, v := range items {
		ans = append(ans, map[string]interface{}{
			"profile_id":     v.Id,
			"name":           v.Name,
			"profile_status": v.Status,
			"updated_at":     v.UpdatedAt,
			"updated_by":     v.UpdatedBy,
			"created_by":     v.CreatedBy,
		})
	}

	d.SetId("data profiles")
	d.Set("total", len(items))
	if err = d.Set("listing", ans); err != nil {
		log.Printf("[WARN] Error setting 'listing' field for %q: %s", d.Id(), err)
	}

	return nil
}
