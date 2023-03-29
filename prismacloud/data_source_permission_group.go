package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/permission_group"
	"golang.org/x/net/context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePermissionGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePermissionGroupRead,

		Schema: map[string]*schema.Schema{
			//input
			"id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "Permission group id",
				AtLeastOneOf: []string{"name", "id"},
			},
			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "Name of the permission group",
				AtLeastOneOf: []string{"name", "id"},
			},
			//output
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description",
			},
			"permission_group_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Permission group type",
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
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Associated permission roles",
				Elem: &schema.Schema{
					Type: schema.TypeString,
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
			"custom": { // Boolean value signifying whether this is a custom (i.e. user-defined) permission group. Is set to true if the attribute value of permissionGroupType is set to CUSTOM
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Custom",
			},
		},
	}
}
func dataSourcePermissionGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)

	var err error
	id := d.Get("id").(string)

	if id == "" {
		name := d.Get("name").(string)
		id, err = permission_group.Identify(client, name)
		if err != nil {
			if err == pc.ObjectNotFoundError {
				d.SetId("")
				return nil
			}
			return diag.FromErr(err)
		}
	}

	obj, err := permission_group.Get(client, id)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	d.SetId(id)
	savePermissionGroup(d, obj)

	return nil
}
