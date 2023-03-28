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
						"associated_roles": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Associated permission roles",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"role_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The role ID",
									},
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Role name",
									},
								},
							},
						},
						"features": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Features",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"feature_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Feature name",
									},
									"operations": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Operations",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												//A mapping of operations and a boolean value representing whether the privilege to perform the operation needs to be granted.
												"create": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Create operation",
												},
												"read": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Read operation",
												},
												"update": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Update operation",
												},
												"delete": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Delete operation",
												},
											},
										},
									},
								},
							},
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
		aroles := map[string]interface{}{
			"role_id": v.AssociatedRoles.RoleId,
			"name":    v.AssociatedRoles.Name,
		}
		feat := make([]interface{}, 0, len(v.Features))
		for _, fe := range v.Features {
			feat = append(feat, map[string]interface{}{
				"feature_name": fe.FeatureName,
				"operations":   map[string]bool{"create": fe.Operations.CREATE, "read": fe.Operations.READ, "update": fe.Operations.UPDATE, "delete": fe.Operations.DELETE},
			})
		}
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
			"associated_roles":         []interface{}{aroles},
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
