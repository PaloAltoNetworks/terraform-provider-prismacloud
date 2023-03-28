package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/permission_group"
	"golang.org/x/net/context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePermissionGroups() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePermissionGroupsRead,

		Schema: map[string]*schema.Schema{
			// Output.
			"total": totalSchema("permission groups"),
			"listing": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of permission groups",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the permission group",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description",
						},
						"permission_group_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Permission groups type",
						},
						"last_modified_by": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Last modified by",
						},
						"last_modified_ts": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Last modified timestamp",
						},
						"accept_account_groups": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Accept account groups",
						},
						"accept_resource_lists": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Accept resource lists",
						},
						"accept_code_repositories": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Accept code repositories",
						},
						"custom": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Custom",
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Permission group id",
						},
					},
				},
			},
		},
	}
}
func dataSourcePermissionGroupsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var err error
	client := meta.(*pc.Client)

	items, err := permission_group.List(client)
	if err != nil {
		return diag.FromErr(err)
	}
	ans := make([]interface{}, 0, len(items))
	for _, v := range items {
		ans = append(ans, map[string]interface{}{
			"name":                     v.Name,
			"description":              v.Description,
			"permission_group_type":    v.Type,
			"last_modified_by":         v.LastModifiedBy,
			"last_modified_ts":         v.LastModifiedTs,
			"accept_account_groups":    v.AcceptAccountGroups,
			"accept_resource_lists":    v.AcceptResourceLists,
			"accept_code_repositories": v.AcceptCodeRepositories,
			"id":                       v.Id,
			"custom":                   v.Custom,
		})

	}
	d.SetId("permission groups")
	d.Set("total", len(items))
	err = d.Set("listing", ans)
	if err != nil {
		log.Printf("[WARN] Error setting 'listing' field for %q: %s", d.Id(), err)
	}

	return nil
}
