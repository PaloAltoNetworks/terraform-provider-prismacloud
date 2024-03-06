package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	collection "github.com/paloaltonetworks/prisma-cloud-go/collection"
	"golang.org/x/net/context"
	"log"
)

func dataSourceCollections() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCollectionsRead,

		Schema: map[string]*schema.Schema{
			// Output.
			"total": totalSchema("collections"),
			"listing": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of collections",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Collection ID",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Collection name",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Collection description",
						},
						"created_by": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Collection created by",
						},
						"created_ts": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Collection created timestamp",
						},
						"last_modified_by": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Collection last modified by",
						},
						"last_modified_ts": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Collection last modified timestamp",
						},
					},
				},
			},
		},
	}
}

func dataSourceCollectionsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)

	res, err := collection.List(client)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("collections_list")
	d.Set("total", len(res.Value))
	list := make([]interface{}, 0, len(res.Value))
	for _, i := range res.Value {
		list = append(list, map[string]interface{}{
			"id":               i.Id,
			"name":             i.Name,
			"description":      i.Description,
			"created_by":       i.CreatedBy,
			"created_ts":       i.CreatedTs,
			"last_modified_by": i.LastModifiedBy,
			"last_modified_ts": i.LastModifiedTs,
		})
	}
	if err := d.Set("listing", list); err != nil {
		log.Printf("[WARN] Error setting 'listing' field for %q: %s", d.Id(), err)
	}
	return nil
}
