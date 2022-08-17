package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"golang.org/x/net/context"
	"log"
	"strconv"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/rql/history"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceRqlHistoricSearches() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRqlHistoricSearchesRead,

		Schema: map[string]*schema.Schema{
			// Input.
			"filter": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter for historic RQL searches",
				Default:     history.Saved,
				ValidateFunc: validation.StringInSlice(
					[]string{
						history.Recent,
						history.Saved,
					},
					false,
				),
			},
			"limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Max number of historic RQL searches to return",
				Default:     1000,
			},

			// Output.
			"total": totalSchema("RQL saved searches"),
			"listing": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of historic RQL searches",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"created_by": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Created by",
						},
						"last_modified_by": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Last modified by",
						},
						"search_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Historic RQL search ID",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name",
						},
						"search_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Search type",
						},
						"saved": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "If this is a saved search",
						},
					},
				},
			},
		},
	}
}

func dataSourceRqlHistoricSearchesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)

	filter := d.Get("filter").(string)
	limit := d.Get("limit").(int)

	items, err := history.List(client, filter, limit)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(TwoStringsToId(filter, strconv.Itoa(limit)))

	d.Set("total", len(items))

	list := make([]interface{}, 0, len(items))
	for _, i := range items {
		list = append(list, map[string]interface{}{
			"created_by":       i.CreatedBy,
			"last_modified_by": i.LastModifiedBy,
			"search_id":        i.Model.Id,
			"name":             i.Model.Name,
			"search_type":      i.Model.SearchType,
			"saved":            i.Model.Saved,
		})
	}

	if err := d.Set("listing", list); err != nil {
		log.Printf("[WARN] Error setting 'listing' field for %q: %s", d.Id(), err)
	}

	return nil
}
