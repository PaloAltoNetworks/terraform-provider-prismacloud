package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	resource_list "github.com/paloaltonetworks/prisma-cloud-go/resource-list"
	"golang.org/x/net/context"
)

func dataSourceResourceList() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceResourceListRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Required:    true,
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
			"members": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Resource list members",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tags": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "List of key:value pairs of tag members",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"azure_resource_groups": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "List of Azure resource groups part of the resource list",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"compute_access_groups": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Members when resource list type = compute access group",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hosts": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Resource list ID",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"app_id": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Resource list ID",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"images": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Resource list ID",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"labels": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Resource list ID",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"clusters": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Resource list ID",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"code_repos": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Resource list ID",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"functions": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Resource list ID",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"containers": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Resource list ID",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"namespaces": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Resource list ID",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceResourceListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var err error
	client := meta.(*pc.Client)
	id := d.Get("id").(string)
	if id == "" {
		return nil
	}
	obj, err := resource_list.Get(client, id)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}
	d.SetId(id)
	saveResourceList(d, obj)
	return nil
}
