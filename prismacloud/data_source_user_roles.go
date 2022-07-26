package prismacloud

import (
	"log"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/user/role"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceUserRoles() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceUserRolesRead,

		Schema: map[string]*schema.Schema{
			// Output.
			"total": totalSchema("user roles"),
			"listing": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of user roles",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"role_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Role UUID",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the role",
						},
						"role_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "User role type",
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
						"account_groups": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Associated account groups",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"group_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The group ID",
									},
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Group name",
									},
								},
							},
						},
						"associated_users": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Associated application users which cannot exist in the system without the user role",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"restrict_dismissal_access": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Restrict dismissal access",
						},
						"additional_attributes": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Additional Parameters",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"only_allow_ci_access": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Allows only CI Access",
									},
									"only_allow_compute_access": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Give access to only compute tab and access key tab",
									},
									"only_allow_read_access": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Allow only read access",
									},
									"has_defender_permissions": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Has defender Permissions",
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

func dataSourceUserRolesRead(d *schema.ResourceData, meta interface{}) error {
	var err error
	client := meta.(*pc.Client)

	items, err := role.List(client)
	if err != nil {
		return err
	}

	ans := make([]interface{}, 0, len(items))
	for _, v := range items {
		addAttr := map[string]interface{}{
			"only_allow_ci_access":      v.AdditionalAttributes.OnlyAllowCIAccess,
			"only_allow_compute_access": v.AdditionalAttributes.OnlyAllowComputeAccess,
			"only_allow_read_access":    v.AdditionalAttributes.OnlyAllowReadAccess,
			"has_defender_permissions":  v.AdditionalAttributes.HasDefenderPermissions,
		}

		agInfo := make([]interface{}, 0, len(v.AccountGroups))
		for _, ag := range v.AccountGroups {
			agInfo = append(agInfo, map[string]interface{}{
				"group_id": ag.Id,
				"name":     ag.Name,
			})
		}

		ans = append(ans, map[string]interface{}{
			"role_id":                   v.Id,
			"name":                      v.Name,
			"role_type":                 v.RoleType,
			"last_modified_by":          v.LastModifiedBy,
			"last_modified_ts":          v.LastModifiedTs,
			"account_groups":            agInfo,
			"associated_users":          v.AssociatedUsers,
			"restrict_dismissal_access": v.RestrictDismissalAccess,
			"additional_attributes":     []interface{}{addAttr},
		})
	}

	d.SetId("user roles")
	d.Set("total", len(items))
	if err = d.Set("listing", ans); err != nil {
		log.Printf("[WARN] Error setting 'listing' field for %q: %s", d.Id(), err)
	}

	return nil
}
