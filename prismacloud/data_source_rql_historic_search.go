package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"golang.org/x/net/context"
	"log"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/rql/history"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRqlHistoricSearch() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRqlHistoricSearchRead,

		Schema: map[string]*schema.Schema{
			// Input.
			"search_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "Historic RQL search ID",
				AtLeastOneOf: []string{"name", "search_id"},
			},
			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "Historic RQL search name",
				AtLeastOneOf: []string{"name", "search_id"},
			},

			// Output.
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description",
			},
			"search_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Search type",
			},
			"cloud_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cloud type",
			},
			"query": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "RQL query",
			},
			"saved": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "If this is a saved search",
			},
			"time_range": timeRangeSchema("data_source_rql_historic_search"),
		},
	}
}

func dataSourceRqlHistoricSearchRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var err error
	client := meta.(*pc.Client)

	id := d.Get("search_id").(string)
	if id == "" {
		name := d.Get("name").(string)
		id, err = history.Identify(client, name)
		if err != nil {
			if err == pc.ObjectNotFoundError {
				d.SetId("")
				return nil
			}
			return diag.FromErr(err)
		}
	}

	o, err := history.Get(client, id)
	log.Printf("Got: %#v", o)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	d.SetId(id)
	d.Set("search_id", o.Id)
	d.Set("name", o.Name)
	d.Set("description", o.Description)
	d.Set("search_type", o.SearchType)
	d.Set("cloud_type", o.CloudType)
	d.Set("query", o.Query)
	d.Set("saved", o.Saved)
	if err = d.Set("time_range", flattenTimeRange(o.TimeRange)); err != nil {
		log.Printf("[WARN] Error setting 'time_range' for %q: %s", d.Id(), err)
	}

	return nil
}
