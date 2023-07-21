package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/rql/history"
	"golang.org/x/net/context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSavedSearch() *schema.Resource {
	return &schema.Resource{
		CreateContext: createSavedSearch,
		UpdateContext: updateSavedSearch,
		ReadContext:   readSavedSearch,
		DeleteContext: deleteSavedSearch,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			// Input.
			"query": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The RQL search to perform",
			},
			"search_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The RQL UUID",
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
			},
			"time_range": timeRangeSchema("resource_saved_search"),
			"cloud_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Cloud Type",
				ValidateFunc: validation.StringInSlice(
					[]string{
						"aws",
						"azure",
						"gcp",
						"alibaba_cloud",
						"oci",
					},
					false,
				),
			},

			// Output.
			"saved": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Set to true when the saved search is created",
			},
		},
	}
}

func createSavedSearch(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)

	req := history.SavedSearch{
		Id:          d.Get("search_id").(string),
		Name:        d.Get("name").(string),
		Query:       d.Get("query").(string),
		Description: d.Get("description").(string),
		TimeRange:   ParseTimeRange(ResourceDataInterfaceMap(d, "time_range")),
		CloudType:   d.Get("cloud_type").(string),
	}

	if err := history.Save(client, req); err != nil {
		return diag.FromErr(err)
	}

	PollApiUntilSuccess(func() error {
		_, err := history.Get(client, req.Id)
		return err
	})

	d.SetId(req.Id)

	return readSavedSearch(ctx, d, meta)
}

func updateSavedSearch(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)

	req := history.SavedSearch{
		Id:          d.Get("search_id").(string),
		Name:        d.Get("name").(string),
		Query:       d.Get("query").(string),
		Description: d.Get("description").(string),
		TimeRange:   ParseTimeRange(ResourceDataInterfaceMap(d, "time_range")),
		CloudType:   d.Get("cloud_type").(string),
	}

	if err := history.Save(client, req); err != nil {
		return diag.FromErr(err)
	}

	PollApiUntilSuccess(func() error {
		_, err := history.Get(client, req.Id)
		return err
	})

	d.SetId(req.Id)

	return readSavedSearch(ctx, d, meta)
}

func readSavedSearch(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	id := d.Id()

	info, err := history.Get(client, id)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("query", info.Query)
	d.Set("search_id", id)
	d.Set("name", info.Name)
	d.Set("description", info.Description)
	d.Set("saved", info.Saved)
	d.Set("cloud_type", info.CloudType)

	return nil
}

func deleteSavedSearch(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	id := d.Id()

	_ = history.Delete(client, id)
	return nil
}
