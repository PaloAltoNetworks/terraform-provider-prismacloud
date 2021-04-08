package prismacloud

import (
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/rql/history"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSavedSearch() *schema.Resource {
	return &schema.Resource{
		Create: createSavedSearch,
		Read:   readSavedSearch,
		Delete: deleteSavedSearch,

		Schema: map[string]*schema.Schema{
			// Input.
			"query": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The RQL search to perform",
				ForceNew:    true,
			},
			"search_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The RQL UUID",
				ForceNew:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Saved search name",
				ForceNew:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description",
				ForceNew:    true,
			},
			"time_range": timeRangeSchema("resource_saved_search"),

			// Output.
			"saved": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Set to true when the saved search is created",
			},
		},
	}
}

func createSavedSearch(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)

	req := history.SavedSearch{
		Id:          d.Get("search_id").(string),
		Name:        d.Get("name").(string),
		Query:       d.Get("query").(string),
		Description: d.Get("description").(string),
		TimeRange:   ParseTimeRange(ResourceDataInterfaceMap(d, "time_range")),
	}

	if err := history.Save(client, req); err != nil {
		return err
	}

	PollApiUntilSuccess(func() error {
		_, err := history.Get(client, req.Id)
		return err
	})

	d.SetId(req.Id)

	return readSavedSearch(d, meta)
}

func readSavedSearch(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	id := d.Id()

	info, err := history.Get(client, id)
	if err != nil {
		return err
	}

	d.Set("query", info.Query)
	d.Set("search_id", id)
	d.Set("name", info.Name)
	d.Set("description", info.Description)
	d.Set("saved", info.Saved)

	return nil
}

func deleteSavedSearch(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	id := d.Id()

	_ = history.Delete(client, id)
	return nil
}
