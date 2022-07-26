package prismacloud

import (
	"log"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/report"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceReports() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceReportsRead,

		Schema: map[string]*schema.Schema{
			// Output.
			"total": totalSchema("reports"),
			"listing": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of reports",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"report_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Report ID",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Report name",
						},
						"report_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Report type",
						},
						"cloud_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Cloud type",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Report status",
						},
					},
				},
			},
		},
	}
}

func dataSourceReportsRead(d *schema.ResourceData, meta interface{}) error {
	var err error
	client := meta.(*pc.Client)

	items, err := report.List(client)
	if err != nil {
		return err
	}

	ans := make([]interface{}, 0, len(items))
	for _, v := range items {
		ans = append(ans, map[string]interface{}{
			"report_id":   v.Id,
			"name":        v.Name,
			"report_type": v.Type,
			"cloud_type":  v.CloudType,
			"status":      v.Status,
		})
	}

	d.SetId("reports")
	d.Set("total", len(items))
	if err = d.Set("listing", ans); err != nil {
		log.Printf("[WARN] Error setting 'listing' field for %q: %s", d.Id(), err)
	}

	return nil
}
