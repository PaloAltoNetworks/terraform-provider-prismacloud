package prismacloud

import (
	"log"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/user/role"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceUserRole() *schema.Resource {
	return &schema.Resource{
		Create: createUserRole,
		Read:   readUserRole,
		Update: updateUserRole,
		Delete: deleteUserRole,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the role",
			},
			"role_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Role UUID",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description",
			},
			"role_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "User role type",
				ValidateFunc: validation.StringInSlice(
					[]string{
						role.TypeSystemAdmin,
						role.TypeAccountGroupAdmin,
						role.TypeAccountGroupReadOnly,
						role.TypeCloudProvisioningAdmin,
						role.TypeAccountAndCloudProvisioningAdmin,
					},
					false,
				),
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
				// TODO: Is this ordered or unordered?
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Accessible account group IDs",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"associated_users": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Associated application users which cannot exist in the system without the user role",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"restrict_dismissal_access": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Restrict dismissal access",
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
		},
	}
}

func parseUserRole(d *schema.ResourceData) *role.Role {
	return &role.Role{
		Name:                    d.Get("name").(string),
		Description:             d.Get("description").(string),
		RoleType:                d.Get("role_type").(string),
		AccountGroupIds:         SetToStringSlice(d.Get("account_group_ids").(*schema.Set)),
		AssociatedUsers:         SetToStringSlice(d.Get("associated_users").(*schema.Set)),
		RestrictDismissalAccess: d.Get("restrict_dismissal_access").(bool),
	}
}

func saveUserRole(d *schema.ResourceData, obj role.Role) {
	var err error

	d.Set("name", obj.Name)
	d.Set("role_id", obj.Id)
	d.Set("description", obj.Description)
	d.Set("role_type", obj.RoleType)
	d.Set("last_modified_by", obj.LastModifiedBy)
	d.Set("last_modified_ts", obj.LastModifiedTs)
	if err = d.Set("account_group_ids", StringSliceToSet(obj.AccountGroupIds)); err != nil {
		log.Printf("[WARN] Error setting 'account_group_ids' field for %q: %s", d.Id(), err)
	}
	if err = d.Set("associated_users", StringSliceToSet(obj.AssociatedUsers)); err != nil {
		log.Printf("[WARN] Error setting 'associated_users' field for %q: %s", d.Id(), err)
	}
	d.Set("restrict_dismissal_access", obj.RestrictDismissalAccess)

	agInfo := make([]interface{}, 0, len(obj.AccountGroups))
	for _, ag := range obj.AccountGroups {
		agInfo = append(agInfo, map[string]interface{}{
			"group_id": ag.Id,
			"name":     ag.Name,
		})
	}
	if err = d.Set("account_groups", agInfo); err != nil {
		log.Printf("[WARN] Error setting 'account_groups' field for %q: %s", d.Id(), err)
	}
}

func createUserRole(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	obj := parseUserRole(d)

	if err := role.Create(client, *obj); err != nil {
		return err
	}

	PollApiUntilSuccess(func() error {
		_, err := role.Identify(client, obj.Name)
		return err
	})

	id, err := role.Identify(client, obj.Name)
	if err != nil {
		return err
	}

	PollApiUntilSuccess(func() error {
		_, err := role.Get(client, id)
		return err
	})

	d.SetId(id)
	return readUserRole(d, meta)
}

func readUserRole(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	id := d.Id()

	obj, err := role.Get(client, id)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return err
	}

	saveUserRole(d, obj)

	return nil
}

func updateUserRole(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	obj := parseUserRole(d)
	obj.Id = d.Id()

	if err := role.Update(client, *obj); err != nil {
		return err
	}

	return readUserRole(d, meta)
}

func deleteUserRole(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	id := d.Id()

	err := role.Delete(client, id)
	if err != nil {
		if err != pc.ObjectNotFoundError {
			return err
		}
	}

	d.SetId("")
	return nil
}
