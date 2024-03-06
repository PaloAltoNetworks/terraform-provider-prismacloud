package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	collection "github.com/paloaltonetworks/prisma-cloud-go/collection"
	"golang.org/x/net/context"
)

func dataSourceCollection() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCollectionRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Required:    true,
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
			"asset_groups": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_group_ids": {
							Type:        schema.TypeList,
							Computed:    true,
							Optional:    true,
							Description: "List of account group ids",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"account_ids": {
							Type:        schema.TypeSet,
							Computed:    true,
							Optional:    true,
							Description: "List of account group ids",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"repository_ids": {
							Type:        schema.TypeSet,
							Computed:    true,
							Optional:    true,
							Description: "List of account group ids",
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

func dataSourceCollectionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var err error
	client := meta.(*pc.Client)
	id := d.Get("id").(string)
	if id == "" {
		return nil
	}
	obj, err := collection.Get(client, id)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}
	d.SetId(id)
	saveCollection(d, obj)
	return nil
}
