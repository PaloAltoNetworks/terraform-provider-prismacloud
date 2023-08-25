package prismacloud

import (
	"errors"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	collection "github.com/paloaltonetworks/prisma-cloud-go/collection"
	"golang.org/x/net/context"
	"log"
)

func resourceCollection() *schema.Resource {
	return &schema.Resource{
		CreateContext: createCollection,
		ReadContext:   readCollection,
		UpdateContext: updateCollection,
		DeleteContext: deleteCollection,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Collection ID",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Collection name",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
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
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_group_ids": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of account group ids",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"account_ids": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of account group ids",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"repository_ids": {
							Type:        schema.TypeList,
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

func deleteCollection(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	id := d.Id()
	log.Printf("[INFO]: Deleting Collection, Id:%+v\n", id)
	if err := collection.Delete(client, id); err != nil {
		if !errors.Is(err, pc.ObjectNotFoundError) {
			return diag.FromErr(err)
		}
	}
	d.SetId("")
	return nil
}

func updateCollection(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var err error
	client := meta.(*pc.Client)
	o, err := parseCollection(d)
	if err != nil {
		return diag.FromErr(err)
	}
	log.Printf("[INFO]: Updating Collection, Id:%+v\n", d.Get("id"))
	if _, err = collection.Update(client, o, d.Get("id").(string)); err != nil {
		if errors.Is(err, pc.ObjectNotFoundError) {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}
	return readCollection(ctx, d, meta)
}

func createCollection(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var err error
	client := meta.(*pc.Client)
	o, err := parseCollection(d)
	if err != nil {
		return diag.FromErr(err)
	}
	var collectionRes collection.Collection
	if collectionRes, err = collection.Create(client, o); err != nil {
		return diag.FromErr(err)
	}
	log.Printf("[INFO]: Collection Created Successfully, Id:%+v\n", collectionRes.Id)
	d.SetId(collectionRes.Id)
	return readCollection(ctx, d, meta)
}

func parseCollection(d *schema.ResourceData) (collection.CollectionRequest, error) {

	colReq := collection.CollectionRequest{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		AssetGroups: collection.AssetGroups{
			AccountGroupIds: []string{},
			AccountIds:      []string{},
			RepositoryIds:   []string{},
		},
	}
	//assetGroups := d.Get("asset_groups").([]interface{})[0]
	assetGroups := ResourceDataInterfaceMap(d, "asset_groups")
	if val, ok := assetGroups["account_group_ids"]; ok {
		list := make([]string, 0, len(val.([]interface{})))
		for _, v := range val.([]interface{}) {
			list = append(list, v.(string))
		}
		colReq.AssetGroups.AccountGroupIds = list
	}
	if val, ok := assetGroups["account_ids"]; ok {
		list := make([]string, 0, len(val.([]interface{})))
		for _, v := range val.([]interface{}) {
			list = append(list, v.(string))
		}
		colReq.AssetGroups.AccountIds = list
	}
	if val, ok := assetGroups["repository_ids"]; ok {
		list := make([]string, 0, len(val.([]interface{})))
		for _, v := range val.([]interface{}) {
			list = append(list, v.(string))
		}
		colReq.AssetGroups.RepositoryIds = list
	}

	return colReq, nil
}

func readCollection(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	id := d.Id()
	obj, err := collection.Get(client, id)
	if err != nil {
		if errors.Is(err, pc.ObjectNotFoundError) {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}
	saveCollection(d, obj)
	return nil
}
func saveCollection(d *schema.ResourceData, o collection.Collection) {
	d.Set("name", o.Name)
	d.Set("description", o.Description)
	d.Set("created_by", o.CreatedBy)
	d.Set("created_ts", o.CreatedTs)
	d.Set("last_modified_ts", o.LastModifiedTs)
	d.Set("last_modified_by", o.LastModifiedBy)

	assetGroups := map[string]interface{}{
		"account_group_ids": o.AssetGroups.AccountGroupIds,
		"account_ids":       o.AssetGroups.AccountIds,
		"repository_ids":    o.AssetGroups.RepositoryIds,
	}

	if err := d.Set("asset_groups", []map[string]interface{}{assetGroups}); err != nil {
		log.Printf("[WARN] Error setting 'asset_groups' for %s: %s", d.Id(), err)
	}
}
