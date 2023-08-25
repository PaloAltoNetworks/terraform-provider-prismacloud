package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	resource_list "github.com/paloaltonetworks/prisma-cloud-go/resource-list"
	"golang.org/x/net/context"
	"log"
)

func dataSourceResourceLists() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceResourceListsRead,

		Schema: map[string]*schema.Schema{
			// Output.
			"total": totalSchema("resource lists"),
			"listing": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of resource lists",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Resource list ID",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Resource list description",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Resource list name",
						},
						"resource_list_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Resource list type",
						},
						"last_modified_ts": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Resource list last modified timestamp",
						},
						"last_modified_by": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Resource list last modified by",
						},
					},
				},
			},
		},
	}
}

func dataSourceResourceListsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)

	items, err := resource_list.List(client)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("resource_lists_list")
	d.Set("total", len(items))
	list := make([]interface{}, 0, len(items))
	for _, i := range items {
		list = append(list, map[string]interface{}{
			"id":                 i.Id,
			"description":        i.Description,
			"name":               i.Name,
			"resource_list_type": i.ResourceListType,
			"last_modified_ts":   i.LastModifiedTs,
			"last_modified_by":   i.LastModifiedBy,
		})
	}
	if err := d.Set("listing", list); err != nil {
		log.Printf("[WARN] Error setting 'listing' field for %q: %s", d.Id(), err)
	}
	return nil
}
