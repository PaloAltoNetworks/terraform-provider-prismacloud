package prismacloud

import (
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/user/role"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceUserRole() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceUserRoleRead,

		Schema: map[string]*schema.Schema{
			// Input.
			"role_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "Role ID",
				AtLeastOneOf: []string{"name", "role_id"},
			},
			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "Role name",
				AtLeastOneOf: []string{"name", "role_id"},
			},

			// Output.
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description",
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
			"account_group_ids": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Accessible account group IDs",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"resource_list_ids": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Resource list IDs",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"code_repository_ids": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Code repository IDs",
				Elem: &schema.Schema{
					Type: schema.TypeString,
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
	}
}

func dataSourceUserRoleRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)

	var err error
	id := d.Get("role_id").(string)

	if id == "" {
		name := d.Get("name").(string)
		id, err = role.Identify(client, name)
		if err != nil {
			if err == pc.ObjectNotFoundError {
				d.SetId("")
				return nil
			}
			return err
		}
	}

	obj, err := role.Get(client, id)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return err
	}

	d.SetId(id)
	saveUserRole(d, obj)

	return nil
}
